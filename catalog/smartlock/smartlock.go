package smartlock

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// CreateTempPwdRequest creates a temporary password with(out) name
type CreateTempPwdRequest struct {
	DeviceID      string `json:"-" validate:"required"`
	Name          string `json:"name,omitempty"`
	Password      string `json:"password" validate:"required"`
	EffectiveTime int64  `json:"effective_time" validate:"required"`
	InvalidTime   int64  `json:"invalid_time" validate:"required"`
	Phone         string `json:"phone,omitempty"`
}

// Method for Request.Method()
func (c *CreateTempPwdRequest) Method() string {
	return http.MethodPost
}

// URL for Request.URL()
func (c *CreateTempPwdRequest) URL() string {
	return "/v1.0/devices/" + c.DeviceID + "/door-lock/temp-password"
}

// CreateTempPwdResponse returns temp password id.
type CreateTempPwdResponse struct {
	ID int64 `json:"id"`
}

// QueryTempPwdRequest returns temp passwords info.
type QueryTempPwdRequest struct {
	DeviceID   string `json:"device_id" validate:"required"`
	PasswordID int64  `json:"password_id" validate:"required"`
}

// Method for Request.Method()
func (q *QueryTempPwdRequest) Method() string {
	return http.MethodGet
}

// URL for Request.URL()
func (q *QueryTempPwdRequest) URL() string {
	return fmt.Sprintf("/v1.0/devices/%s/door-lock/temp-passwords/%d", q.DeviceID, q.PasswordID)
}

// refer to https://developer.tuya.com/en/docs/iot/open-api/api-list/api/proprietary-category/smart-door-lock?temporary%20passwords
const (
	ZigbeeToBeCreate   = 1
	ZigbeeNormal       = 2
	ZigbeeFrozen       = 3
	ZigbeeDeleted      = 4
	ZigbeeCreateFailed = 5

	WiFiDeleted     = 0
	WiFiToBeIssued  = 1
	WiFiIssued      = 2
	WiFiToBeDeleted = 3
)

// QueryTempPwdResponse returns temp password information.
type QueryTempPwdResponse struct {
	PasswordID    int64  `json:"password_id"`
	Name          string `json:"name"`
	Phase         int    `json:"phase"`
	EffectiveTime int64  `json:"effective_time"`
	InvalidTime   int64  `json:"invalid_time"`
	Phone         string `json:"phone"`
}

// QueryTempPwdListRequest returns a list of temp passwords
type QueryTempPwdListRequest struct {
	DeviceID string `json:"device_id" validate:"required"`
}

// Method for Request.Method()
func (q *QueryTempPwdListRequest) Method() string {
	return http.MethodGet
}

// URL for Request.URL()
func (q *QueryTempPwdListRequest) URL() string {
	return "/v1.0/devices/" + q.DeviceID + "/door-lock/temp-passwords"
}

// QueryTempPwdListResponse return a list of QueryTempPwdResponse
type QueryTempPwdListResponse []QueryTempPwdResponse

// UpdateTempPwdRequest updates temp password.
type UpdateTempPwdRequest struct {
	DeviceID   string `json:"-" validate:"required"`
	PasswordID int64  `json:"-" validate:"required"`

	Phone         string `json:"phone,omitempty"`
	EffectiveTime int64  `json:"effective_time,omitempty"`
	InvalidTime   int64  `json:"invalid_time,omitempty"`
	Name          string `json:"name,omitempty"`
	Password      string `json:"password,omitempty"`
}

// Method for Request.Method()
func (u *UpdateTempPwdRequest) Method() string {
	return http.MethodPut
}

// URL for Request.URL()
func (u *UpdateTempPwdRequest) URL() string {
	return fmt.Sprintf("/v1.0/devices/%s/door-lock/temp-passwords/%d/modify-password", u.DeviceID, u.PasswordID)
}

// UpdateTempPwdResponse show success or not.
type UpdateTempPwdResponse bool

// DeleteTempPwdRequest delete temp password
type DeleteTempPwdRequest struct {
	DeviceID   string `json:"-" validate:"required"`
	PasswordID int64  `json:"-" validate:"required"`
}

// Method for Request.Method()
func (d *DeleteTempPwdRequest) Method() string {
	return http.MethodDelete
}

// URL for Request.URL()
func (d *DeleteTempPwdRequest) URL() string {
	return fmt.Sprintf("/v1.0/devices/%s/door-lock/temp-passwords/%d/modify-password", d.DeviceID, d.PasswordID)
}

// DeleteTempPwdResponse show success or not.
type DeleteTempPwdResponse bool

// DynamicPwdRequest request to a dynamic password
type DynamicPwdRequest struct {
	DeviceID string `json:"device_id" validate:"required"`
}

// Method for Request.Method()
func (d *DynamicPwdRequest) Method() string {
	return http.MethodGet
}

// URL for Request.URL()
func (d *DynamicPwdRequest) URL() string {
	return "/v1.0/devices/" + d.DeviceID + "/door-lock/dynamic-password"
}

// DynamicPwdResponse returns a dynamic password
type DynamicPwdResponse struct {
	DynamicPassword string `json:"dynamic_password"`
}

const (
	// OfflineTempPwdTypeMultiTime multiple uses.
	OfflineTempPwdTypeMultiTime = 0
	// OfflineTempPwdTypeOneTime single use.
	OfflineTempPwdTypeOneTime = 1
	// OfflineTempPwdTClean clear all passwords.
	OfflineTempPwdTClean = 9
)

// OfflineTempPwdRequest request to a offline temp password
type OfflineTempPwdRequest struct {
	DeviceID string `json:"-" validate:"required"`

	EffectiveTime int64  `json:"effective_time" validate:"required"`
	InvalidTime   int64  `json:"invalid_time" validate:"required"`
	Name          string `json:"name"`
	Type          int    `json:"type"`
	Lang          string `json:"lang" validate:"required"`
}

// Body struct.
func (d *OfflineTempPwdRequest) Body() interface{} {
	return d
}

// Method for Request.Method()
func (d *OfflineTempPwdRequest) Method() string {
	return http.MethodPost
}

// URL for Request.URL()
func (d *OfflineTempPwdRequest) URL() string {
	return "/v1.0/devices/" + d.DeviceID + "/door-lock/offline-temp-password"
}

// OfflineTempPwdResponse returns a offline temp password
type OfflineTempPwdResponse struct {
	OfflineTempPassword string `json:"offline_temp_password"`
}

// QueryOpenLogsRequest returns a list of logs.
type QueryOpenLogsRequest struct {
	DeviceID string `url:"-" validate:"required"`

	PageNo    int   `url:"page_no" validate:"required"`
	PageSize  int   `url:"page_size" validate:"required"`
	StartTime int64 `url:"start_time" validate:"required"`
	EndTime   int64 `url:"end_time"  validate:"required"`
}

// Method for Request.Method()
func (q *QueryOpenLogsRequest) Method() string {
	return http.MethodGet
}

// URL for Request.URL()
func (q *QueryOpenLogsRequest) URL() string {
	u, _ := query.Values(q)
	return "/v1.0/devices/" + q.DeviceID + "/door-lock/open-logs?" + u.Encode()
}

// OpenLog details of open action.
type OpenLog struct {
	Status struct {
		Code  string `json:"code"`
		Value string `json:"value"`
	} `json:"status"`
	UpdateTime int64 `json:"update_time"`
}

// QueryOpenLogsResponse returns a list if OpenLog
type QueryOpenLogsResponse struct {
	Total  int       `json:"total"`
	Status []OpenLog `json:"status"`
}

// IssuePasswordRequest sync password(s) to device.
type IssuePasswordRequest struct {
	DeviceID string `json:"-" validate:"required"`

	PasswordID int64 `json:"password_id,omitempty"`
}

// Method for Request.Method()
func (i *IssuePasswordRequest) Method() string {
	return http.MethodGet
}

// URL for Request.URL()
func (i *IssuePasswordRequest) URL() string {
	return "/v1.0/devices/" + i.DeviceID + "/door-lock/issue-password"
}

// IssuePasswordResponse shows success or not.
type IssuePasswordResponse bool
