package tuyacloud

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/go-log/log"
)

type Client struct {
	endpoint  string
	accessID  string
	accessKey string

	httpClient *http.Client
	logger     log.Logger
	storage    TokenStorage
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
	return
}

func (c *Client) Request(r Request) (req *http.Request, err error) {
	target := c.endpoint + "/" + r.URI()
	var buf io.Reader
	if r.Method() != http.MethodGet {
		i := r.(RequestBody).Body()
		var b []byte
		b, err = json.Marshal(i)
		if err != nil {
			return
		}
		buf = bytes.NewReader(b)
	}

	req, err = http.NewRequest(r.Method(), target, buf)
	if err != nil {
		return
	}
	var token string
	token, err = c.Token()
	if err != nil {
		return
	}
	timestamp := Timestamp()
	sign := c.PlainSign(timestamp)
	req.Header.Add("client_id", c.accessID)
	req.Header.Add("access_token", token)
	req.Header.Add("sign", sign)
	req.Header.Add("sign_method", "HMAC-SHA256")
	req.Header.Add("t", timestamp)
	if r.Method() != http.MethodGet {
		req.Header.Add("Content-Type", "application/json")
	}
	return
}

func (c *Client) Parse(res *http.Response, resp interface{}) error {
	defer res.Body.Close()
	var body Response
	err := json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body.Result, resp)
	if err != nil {
		return err
	}
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

func (c *Client) Token() (token string, err error) {
	token = c.storage.Token()
	if token == "" {
		err = c.storage.Refresh(c)
		if err != nil {
			return
		}
		token = c.storage.Token()
	}
	return
}
