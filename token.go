package tuyacloud

import "net/http"

type TokenRequest struct {
	refresh string
}

func (t *TokenRequest) Method() string {
	return http.MethodGet
}

func (t *TokenRequest) SetRefresh(token string) {
	t.refresh = token
}

func (t *TokenRequest) URL() string {
	uri := "/v1.0/token"
	if t.refresh != "" {
		uri += "/" + t.refresh
	} else {
		uri += "?grant_type=1"
	}
	return uri
}

type TokenResponse struct {
	ExpireTime   int    `json:"expire_time"`
	UID          string `json:"uid"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
