package device

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

// QueryDeviceDetailRequest gets device details.
type QueryDeviceDetailRequest struct {
	DeviceID string `validate:"required"`
}

// Method for Request.Method()
func (q *QueryDeviceDetailRequest) Method() string {
	return http.MethodGet
}

// URI for Request.URI()
func (q *QueryDeviceDetailRequest) URI() string {
	return "/v1.0/devices/" + q.DeviceID
}

// QueryDeviceDetailResponse returns device details.
type QueryDeviceDetailResponse struct {
	ID        string                   `json:"id"`
	UID       string                   `json:"uid"`
	Category  string                   `json:"category"`
	Sub       bool                     `json:"sub"`
	UUID      string                   `json:"uuid"`
	Online    bool                     `json:"online"`
	Status    []map[string]interface{} `json:"status"`
	LocalKey  string                   `json:"local_key"`
	ProductID string                   `json:"product_id"`
	OwnerID   string                   `json:"owner_id"`
}

// QueryUserDevicesRequest gets the list of devices under the user.
type QueryUserDevicesRequest struct {
	UID string `validate:"required"`
}

// Method for Request.Method()
func (q *QueryUserDevicesRequest) Method() string {
	return http.MethodGet
}

// URI for Request.URI()
func (q *QueryUserDevicesRequest) URI() string {
	return "/v1.0/users/" + q.UID + "/devices"
}

// UserDevice for device of user.
type UserDevice struct {
	Sub         bool                   `json:"sub"`
	CreateTime  int                    `json:"create_time"`
	LocalKey    string                 `json:"local_key"`
	OwnerID     string                 `json:"owner_id"`
	IP          string                 `json:"ip"`
	BizType     int                    `json:"biz_type"`
	Icon        string                 `json:"icon"`
	TimeZone    string                 `json:"time_zone"`
	UUID        string                 `json:"uuid"`
	ProductName string                 `json:"product_name"`
	ActiveTime  int                    `json:"active_time"`
	UID         string                 `json:"uid"`
	UpdateTime  int                    `json:"update_time"`
	ProductID   string                 `json:"product_id"`
	Name        string                 `json:"name"`
	Online      bool                   `json:"online"`
	ID          string                 `json:"id"`
	Category    string                 `json:"category"`
	Status      map[string]interface{} `json:"status"`
}

// QueryUserDevicesResponse returns devices.
type QueryUserDevicesResponse []UserDevice

// QueryDevicesRequest gets the list of devices with(out) filters.
type QueryDevicesRequest struct {
	Schema    string `url:"schema,omitempty"`
	ProductID string `url:"product_id,omitempty"`
	DeviceIDs string `url:"device_ids,omitempty"`
	PageNo    int    `validate:"required" url:"page_no"`
	PageSize  int    `validate:"required" url:"page_size"`
}

// Method for Request.Method()
func (q *QueryDevicesRequest) Method() string {
	return http.MethodGet
}

// URI for Request.URI()
func (q *QueryDevicesRequest) URI() string {
	v, _ := query.Values(q)
	return "/v1.0/devices?" + v.Encode()
}

// QueryDevicesResponse returns devices list.
type QueryDevicesResponse struct {
	Total  int                         `json:"total"`
	Devices []QueryDeviceDetailResponse `json:"devices"`
}

// UpdateFunctionPointNameRequest modifies the function point name.
type UpdateFunctionPointNameRequest struct {
	DeviceID     string `json:"-" validate:"required"`
	FunctionCode string `json:"-" validate:"required"`

	Name string `json:"name" validate:"required"`
}

// Method for Request.Method()
func (u *UpdateFunctionPointNameRequest) Method() string {
	return http.MethodPut
}

// URI for Request.URI()
func (u *UpdateFunctionPointNameRequest) URI() string {
	return "/v1.0/devices/" + u.DeviceID + "/functions/" + u.FunctionCode
}

// UpdateFunctionPointNameResponse returns success or not.
type UpdateFunctionPointNameResponse bool

// QueryLogsRequest gets device logs.
type QueryLogsRequest struct {
	DeviceID string `validate:"required" url:"-"`

	Type          string `validate:"required" url:"type"`
	StartTime     int64  `validate:"required" url:"start_time"`
	EndTime       int64  `validate:"required" url:"end_time"`
	Codes         string `url:"codes,omitempty"`
	StartRowKey   string `url:"start_row_key,omitempty"`
	LastRowKey    string `url:"last_row_key,omitempty"`
	LastEventTime int64  `url:"last_event_time,omitempty"`
	Size          int    `url:"size,omitempty"`
	QueryType     int    `url:"query_type,omitempty"`
}

// Method for Request.Method()
func (q *QueryLogsRequest) Method() string {
	return http.MethodGet
}

// URI for Request.URI()
func (q *QueryLogsRequest) URI() string {
	v, _ := query.Values(q)
	return "/v1.0/devices/" + q.DeviceID + "/logs?" + v.Encode()
}

// Log item
type Log struct {
	Code      string      `json:"code"`
	Value     interface{} `json:"value"`
	EventTime int64       `json:"event_time"`
	EventFrom string      `json:"event_from"`
	EventID   int         `json:"event_id"`
}

// QueryLogsResponse returns device logs.
type QueryLogsResponse struct {
	Logs          []Log  `json:"logs"`
	DeviceID      string `json:"device_id"`
	HasNext       bool   `json:"has_next"`
	CurrentRowKey string `json:"current_row_key"`
	NextRowKey    string `json:"next_row_key"`
}

// RestoreFactorySettingsRequest restores the factory settings
type RestoreFactorySettingsRequest struct {
	DeviceID string `validate:"required"`
}

// Method for Request.Method()
func (r *RestoreFactorySettingsRequest) Method() string {
	return http.MethodPut
}

// URI for Request.URI()
func (r *RestoreFactorySettingsRequest) URI() string {
	return "/v1.0/devices/{device_id}/reset-factory"
}

// RestoreFactorySettingsResponse shows success or not.
type RestoreFactorySettingsResponse bool

// DeleteDeviceRequest to delete device.
type DeleteDeviceRequest struct {
	DeviceID string `validate:"required"`
}

// Method for Request.Method()
func (d *DeleteDeviceRequest) Method() string {
	return http.MethodDelete
}

// URI for Request.URI()
func (d *DeleteDeviceRequest) URI() string {
	return "/v1.0/devices/" + d.DeviceID
}

// DeleteDeviceResponse shows success or not.
type DeleteDeviceResponse bool

// QuerySubDevicesRequest gets device for gateway.
type QuerySubDevicesRequest struct {
	DeviceID string `validate:"required"`
}

// Method for Request.Method()
func (q *QuerySubDevicesRequest) Method() string {
	return http.MethodGet
}

// URI for Request.URI()
func (q *QuerySubDevicesRequest) URI() string {
	return "/v1.0/devices/" + q.DeviceID + "/sub-devices"
}

// SubDevice for device for the gateway.
type SubDevice struct {
	ActiveTime int    `json:"active_time"`
	Category   string `json:"category"`
	ID         string `json:"id"`
	Name       string `json:"name"`
	Online     bool   `json:"online"`
	OwnerID    string `json:"owner_id"`
	ProductID  string `json:"product_id"`
	UpdateTime int    `json:"update_time"`
}

// QuerySubDevicesResponse returns device list.
type QuerySubDevicesResponse []SubDevice

// QueryFactoryInfoRequest gets device factory information
type QueryFactoryInfoRequest struct {
	DeviceIds string `url:"device_ids" validate:"required"`
}

// Method for Request.Method()
func (q *QueryFactoryInfoRequest) Method() string {
	return http.MethodGet
}

// URI for Request.URI()
func (q *QueryFactoryInfoRequest) URI() string {
	v, _ := query.Values(q)
	return "/v1.0/devices/factory-infos?" + v.Encode()
}

// FactoryInfo for device.
type FactoryInfo struct {
	ID   string `json:"id"`
	UUID string `json:"uuid"`
	SN   string `json:"sn"`
	MAC  string `json:"mac"`
}

// QueryFactoryInfoResponse returns factory information.
type QueryFactoryInfoResponse []FactoryInfo
