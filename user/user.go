package user

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

// SyncUserInfoRequest sync user info
type SyncUserInfoRequest struct {
	Schema       string `url:"-" validate:"required" json:"-"`             // application unique identifier
	CountryCode  string `url:"-" validate:"required" json:"country_code"`  // Country code
	Username     string `url:"-" validate:"required" json:"username"`      // Username
	Password     string `url:"-" validate:"required" json:"password"`      // Password.It is recommended to use the MD5 hash original password.
	UsernameType int    `url:"-" validate:"required" json:"username_type"` // User name type,1:mobile，2:email，3:open id, default:3
	NickName     string `url:"-" json:"nick_name,omitempty"`                         // nickname
}

// Method for Request.Method()
func (s *SyncUserInfoRequest) Method() string {
	return http.MethodPost
}

// URL for Request.URL()
func (s *SyncUserInfoRequest) URL() string {
	return "/v1.0/apps/" + s.Schema + "/user"
}

// QueryUserListRequest query user list
type QueryUserListRequest struct {
	Schema   string      `url:"-" validate:"required" json:"-"`         // App schema
	PageNo   interface{} `url:"page_no" validate:"required" json:"-"`   // Page no
	PageSize interface{} `url:"page_size" validate:"required" json:"-"` // Page size, value range (0, 100]
}

// Method for Request.Method()
func (s *QueryUserListRequest) Method() string {
	return http.MethodGet
}

// URL for Request.URL()
func (s *QueryUserListRequest) URL() string {
	v, _ := query.Values(s)
	return "/v1.0/apps/" + s.Schema + "/users?" + v.Encode()
}

// QueryUserInfoRequest for get user information
type QueryUserInfoRequest struct {
	UID string `url:"-" validate:"required" json:"-"` // user id
}

// Method for Request.Method()
func (s *QueryUserInfoRequest) Method() string {
	return http.MethodGet
}

// URL for Request.URL()
func (s *QueryUserInfoRequest) URL() string {
	return "/v1.0/users/" + s.UID + "/infos"
}
