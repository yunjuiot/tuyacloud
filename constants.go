package tuyacloud

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Endpoint string

const (
	// APIEndpointCN prefer to China.
	APIEndpointCN Endpoint = "https://openapi.tuyacn.com"
	// APIEndpointUS prefer to America.
	APIEndpointUS Endpoint = "https://openapi.tuyaus.com"
	// APIEndpointEU prefer to Europe.
	APIEndpointEU Endpoint = "https://openapi.tuyaeu.com"
	// APIEndpointIN prefer to India.
	APIEndpointIN Endpoint = "https://openapi.tuyain.com"
)

// HTTPClient interface.
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Request for API call.
type Request interface {
	Method() string
	URL() string
}

type RequestBody interface {
	Body() interface{}
}

type Response struct {
	Success   bool            `json:"success"`
	Code      int             `json:"code"`
	Msg       string          `json:"msg"`
	Timestamp int64           `json:"t"`
	Result    json.RawMessage `json:"result"`
}

// TokenStorage stores token.
type TokenStorage interface {
	Token() string
	Refresh(c *Client) error
	IsExpiresAt(d time.Duration) bool
}

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Msg)
}
