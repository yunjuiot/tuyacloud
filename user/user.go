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
	NickName     string `url:"-" json:"nick_name,omitempty"`               // nickname
}

func (s *SyncUserInfoRequest) Body() interface{} {
	return s
}

// Method for Request.Method()
func (s *SyncUserInfoRequest) Method() string {
	return http.MethodPost
}

// URL for Request.URL()
func (s *SyncUserInfoRequest) URL() string {
	return "/v1.0/apps/" + s.Schema + "/user"
}

type SyncUserInfoResponse struct {
	UID string `json:"uid"` // tuya user id
}

// QueryUserListRequest query user list
type QueryUserListRequest struct {
	Schema   string `url:"-" validate:"required" json:"-"`         // App schema
	PageNo   int    `url:"page_no" validate:"required" json:"-"`   // Page no
	PageSize int    `url:"page_size" validate:"required" json:"-"` // Page size, value range (0, 100]
}

func (s *QueryUserListRequest) Body() interface{} {
	return s
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

type UserInfo struct {
	CountryCode string `json:"country_code"`
	UID         string `json:"uid"`
	Username    string `json:"username"`
	Mobile      string `json:"mobile"`
}

type QueryUserListResponse struct {
	HasMore bool       `json:"has_more"`
	List    []UserInfo `json:"list"`
}

// QueryUserInfoRequest for get user information
type QueryUserInfoRequest struct {
	UID string `url:"-" validate:"required" json:"-"` // user id
}

func (s *QueryUserInfoRequest) Body() interface{} {
	return s
}

// Method for Request.Method()
func (s *QueryUserInfoRequest) Method() string {
	return http.MethodGet
}

// URL for Request.URL()
func (s *QueryUserInfoRequest) URL() string {
	return "/v1.0/users/" + s.UID + "/infos"
}

type QueryUserInfoResponse struct {
	CountryCode string `json:"country_code"` // country code
	Avatar      string `json:"avatar"`       // avatar
	Mobile      string `json:"mobile"`       // mobile
	NickName    string `json:"nick_name"`    // nick name
	UID         string `json:"uid"`          // user id
	Username    string `json:"username"`     // username
	CreateTime  int64  `json:"create_time"`  // create_time
	UpdateTime  int64  `json:"update_time"`  // update_time
}
