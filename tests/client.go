package tests

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/go-log/log/log"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/yunjuiot/tuyacloud"
	"github.com/yunjuiot/tuyacloud/tests/helpers"
)

// NewClient return client.
func NewClient(t *testing.T) (client *tuyacloud.Client) {
	id := os.Getenv("ACCESSID")
	key := os.Getenv("ACCESSKEY")
	if id == "" || key == "" {
		t.SkipNow()
		return
	}
	var store tuyacloud.TokenStorage
	store = &tuyacloud.MemoryStore{}
	token := os.Getenv("TOKEN")
	if token != "" {
		store = helpers.NewStaticStore(token)
	}
	require.NotEmpty(t, token, "Access Token is empty.")
	client = tuyacloud.NewClient(
		tuyacloud.APIEndpointCN, id, key, tuyacloud.WithTokenStore(store),
		tuyacloud.WithLogger(log.New()),
	)
	return
}

// NewMockClient returns mock client
func NewMockClient(t *testing.T, ctrl *gomock.Controller) (client *tuyacloud.Client, httpClient *helpers.MockHTTPClient) {
	httpClient = helpers.NewMockHTTPClient(ctrl)
	fakeID := "123"
	fakeKey := "456"
	client = tuyacloud.NewClient(
		tuyacloud.APIEndpointEU, fakeID, fakeKey,
		tuyacloud.WithHTTPClient(httpClient),
		tuyacloud.WithTokenStore(helpers.NewStaticStore("789")),
	)
	return
}

type Mocked struct {
	ResponseBody string
	MockedError  error
}

func (m Mocked) Response() (*http.Response, error) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(m.ResponseBody)))
	resp := &http.Response{Body: r}
	return resp, m.MockedError
}
