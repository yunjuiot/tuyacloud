package tuyacloud

import (
	"net/http"
	"time"
)

type MemoryStore struct {
	token        string
	refreshToken string
	expiresTime  time.Time
}

func (s *MemoryStore) Token() string {
	//if time.Now().Before(s.expiresTime) {
	//	return s.token
	//}
	return s.token
}

func (s *MemoryStore) IsExpiresAt(d time.Duration) bool {
	if s.expiresTime.After(time.Now().Add(d)) {
		return false
	}
	return true
}

func (s *MemoryStore) Refresh(c *Client) (err error) {
	r := &TokenRequest{}
	if s.refreshToken != "" {
		r.SetRefresh(s.refreshToken)
	}
	var req *http.Request
	req, err = c.Request(&TokenRequest{})
	if err != nil {
		return
	}
	sign := c.PlainSign(req.Header.Get("t"))
	req.Header.Del("access_token")
	req.Header.Set("sign", sign)
	var resp *http.Response
	resp, err = c.Do(req)
	if err != nil {
		return
	}
	var body TokenResponse
	err = c.Parse(resp, &body)
	if err != nil {
		return
	}
	s.token = body.AccessToken
	s.refreshToken = body.RefreshToken
	s.expiresTime = time.Now().Add(time.Duration(body.ExpireTime-600) * time.Second)
	return
}
