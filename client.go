package tuyacloud

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/go-log/log"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

type Client struct {
	endpoint  string
	accessID  string
	accessKey string

	httpClient HTTPClient
	logger     log.Logger
	storage    TokenStorage
	validator  *validator.Validate
}

// NewClient returns API client.
func NewClient(endpoint Endpoint, accessID, accessKey string, opts ...Option) (c *Client) {
	c = &Client{endpoint: string(endpoint), accessID: accessID, accessKey: accessKey}
	conf := &options{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		logger:     log.DefaultLogger,
		storage:    &MemoryStore{},
	}
	for _, opt := range opts {
		opt(conf)
	}
	c.httpClient = conf.httpClient
	c.logger = conf.logger
	c.storage = conf.storage
	c.validator = validator.New()
	return
}

func (c *Client) isBody(r Request) bool {
	// WTF: Some cases uploads body with http.MethodDelete
	if _, ok := r.(RequestBody); ok {
		return true
	}
	if r.Method() != http.MethodGet && r.Method() != http.MethodDelete {
		return true
	}
	return false
}

func (c *Client) Request(r Request) (req *http.Request, err error) {
	// Check params by go-playground/validator
	err = c.validator.Struct(r)
	if err != nil {
		return
	}
	target := c.endpoint + r.URL()
	var buf io.Reader
	if c.isBody(r) {
		i := r.(RequestBody).Body()
		var b []byte
		b, err = json.Marshal(i)
		if err != nil {
			return
		}
		buf = bytes.NewReader(b)
	}

	c.logger.Logf("%s %s", r.Method(), target)
	req, err = http.NewRequest(r.Method(), target, buf)
	if err != nil {
		return
	}
	timestamp := Timestamp()
	// TODO: dirty hack for infinity loop
	if !strings.Contains(r.URL(), "/v1.0/token") {
		var token string
		token, err = c.Token()
		if err != nil {
			return
		}
		sign := c.TokenSign(token, timestamp)
		req.Header.Add("access_token", token)
		req.Header.Add("sign", sign)
	}
	req.Header.Add("client_id", c.accessID)
	req.Header.Add("sign_method", "HMAC-SHA256")
	req.Header.Add("t", timestamp)
	if c.isBody(r) {
		req.Header.Add("Content-Type", "application/json")
	}
	return
}

func (c *Client) Parse(res *http.Response, resp interface{}) error {
	defer res.Body.Close()
	var body Response
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	c.logger.Logf("Recv: %s", string(b))
	err = json.Unmarshal(b, &body)
	if err != nil {
		return err
	}
	if !body.Success {
		return errors.Wrap(&Error{
			Code: body.Code,
			Msg:  body.Msg,
		}, "call failed")
	}
	err = json.Unmarshal(body.Result, resp)
	return err
}

func (c *Client) Do(r *http.Request) (res *http.Response, err error) {
	res, err = c.httpClient.Do(r)
	return
}

func (c *Client) DoAndParse(r Request, resp interface{}) (err error) {
	var req *http.Request
	var res *http.Response
	req, err = c.Request(r)
	if err != nil {
		return
	}
	res, err = c.Do(req)
	if err != nil {
		return
	}
	err = c.Parse(res, resp)
	return
}

func (c *Client) PlainSign(timestamp string) string {
	return strings.ToUpper(HmacSha256(c.accessID+timestamp, c.accessKey))
}

func (c *Client) TokenSign(token, timestamp string) string {
	return strings.ToUpper(HmacSha256(c.accessID+token+timestamp, c.accessKey))
}

var (
	tokenLock sync.RWMutex
)

// add token lock
func (c *Client) Token() (token string, err error) {
	err = c.autoRefreshToken(time.Second * 30)
	if err != nil {
		c.logger.Logf("Auto Refresh Token Error: %s", err.Error())
		return
	}
	tokenLock.RLock()
	token = c.storage.Token()
	tokenLock.RUnlock()
	return
}

// token 快过期30秒之前自动刷新,防止token失效
func (c *Client) autoRefreshToken(d time.Duration) error {
	if !c.storage.IsExpiresAt(d) {
		return nil
	}
	tokenLock.Lock()
	defer tokenLock.Unlock()
	err := c.storage.Refresh(c)
	if err != nil {
		return err
	}
	return nil
}
