package tuyacloud

import "encoding/json"

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

// Request for API call.
type Request interface {
	Method() string
	URI() string
}

type RequestBody interface {
	Body() interface{}
}

type Response struct {
	Success bool            `json:"success"`
	Code    int             `json:"code"`
	Msg     string          `json:"msg"`
	Result  json.RawMessage `json:"result"`
}

// TokenStorage stores token.
type TokenStorage interface {
	Token() string
	Refresh(c *Client) error
}
