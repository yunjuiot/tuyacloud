package tuyacloud

import "net/http"

// TokenRequest manages token refresh method.
type TokenRequest struct {
	refresh string
}

// Method for Request.Method()
func (t *TokenRequest) Method() string {
	return http.MethodGet
}

// SetRefresh refresh token.
func (t *TokenRequest) SetRefresh(token string) {
	t.refresh = token
}

// URL for Request.URL()
func (t *TokenRequest) URL() string {
	uri := "/v1.0/token"
	if t.refresh != "" {
		uri += "/" + t.refresh
	} else {
		uri += "?grant_type=1"
	}
	return uri
}

// TokenResponse for token response from tuya server.
type TokenResponse struct {
	ExpireTime   int    `json:"expire_time"`
	UID          string `json:"uid"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
