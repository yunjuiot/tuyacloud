package smarthome

import (
	"fmt"
	"net/http"
)

type Home struct {
	Name      string  `url:"-" validate:"required" json:"name"` // Family name
	GeoName   string  `url:"-" json:"geo_name,omitempty"`       // Family location
	Longitude float32 `url:"-" json:"lon,omitempty"`            // Longitude
	Latitude  float32 `url:"-" json:"lat,omitempty"`            // Latitude
}

type AddHomeRequest struct {
	UID   string   `url:"-" validate:"required" json:"uid"`  // Tuya user id
	Home  Home     `url:"-" validate:"required" json:"home"` // Family
	Rooms []string `url:"-" json:"rooms,omitempty"`          // Room list
}

func (s *AddHomeRequest) Body() interface{} {
	return s
}

// Method for Request.Method()
func (s *AddHomeRequest) Method() string {
	return http.MethodPost
}

// URL for Request.URL()
func (s *AddHomeRequest) URL() string {
	return "/v1.0/home/create-home"
}

type AddHomeResponse int64

type UpdateHomeRequest struct {
	HomeID    int64   `url:"-" validate:"required" json:"-"` // Family ID
	Name      string  `url:"-" json:"name,omitempty"`        // Family name
	GeoName   string  `url:"-" json:"geo_name,omitempty"`    // Family location
	Longitude float32 `url:"-" json:"lon,omitempty"`         // Longitude
	Latitude  float32 `url:"-" json:"lat,omitempty"`         // Latitude
}

func (s *UpdateHomeRequest) Body() interface{} {
	return s
}

// Method for Request.Method()
func (s *UpdateHomeRequest) Method() string {
	return http.MethodPut
}

// URL for Request.URL()
func (s *UpdateHomeRequest) URL() string {
	return fmt.Sprintf("/v1.0/homes/%d", s.HomeID)
}

type UpdateHomeResponse bool

type QueryHomeRequest struct {
	HomeID int64 `url:"-" validate:"required" json:"-"` // Family ID
}

func (s *QueryHomeRequest) Body() interface{} {
	return s
}

// Method for Request.Method()
func (s *QueryHomeRequest) Method() string {
	return http.MethodGet
}

// URL for Request.URL()
func (s *QueryHomeRequest) URL() string {
	return fmt.Sprintf("/v1.0/homes/%d", s.HomeID)
}

type QueryHomeResponse struct {
	HomeID    int64   `json:"home_id"`  // Family id
	Name      string  `json:"name"`     // Family name
	GeoName   string  `json:"geo_name"` // Family location
	Longitude float32 `json:"lon"`      // Longitude
	Latitude  float32 `json:"lat"`      // Latitude
}

type DeleteHomeRequest struct {
	HomeID int64 `url:"-" validate:"required" json:"-"` // Family id
}

func (s *DeleteHomeRequest) Body() interface{} {
	return s
}

// Method for Request.Method()
func (s *DeleteHomeRequest) Method() string {
	return http.MethodDelete
}

// URL for Request.URL()
func (s *DeleteHomeRequest) URL() string {
	return fmt.Sprintf("/v1.0/homes/%d", s.HomeID)
}

type DeleteHomeResponse bool

type QueryHomeDeviceRequest struct {
	HomeID int64 `url:"-" validate:"required" json:"-"` // Family id
}

func (s *QueryHomeDeviceRequest) Body() interface{} {
	return s
}

// Method for Request.Method()
func (s *QueryHomeDeviceRequest) Method() string {
	return http.MethodGet
}

// URL for Request.URL()
func (s *QueryHomeDeviceRequest) URL() string {
	return fmt.Sprintf("/v1.0/homes/%d/devices", s.HomeID)
}

type Status struct {
	Code  string `json:"code"`  // Function point code
	Value string `json:"value"` // Function point value
}

type QueryHomeDeviceResponse struct {
	ID        string `json:"id"`         // Device ID
	UID       string `json:"uid"`        // User ID
	LocalKey  string `json:"local_key"`  // Key
	Category  string `json:"category"`   // Product Category
	ProductID string `json:"product_id"` // Product ID
	Sub       bool   `json:"sub"`        // Whether it is a sub-device (true: yes, false: no)
	UUID      string `json:"uuid"`       // Device unique identifier
	OwnerID   string `json:"owner_id"`   // Device owner ID
	Online    bool   `json:"online"`     // Device online status
	Ip        string `json:"ip"`         // Device IP
	Name      string `json:"name"`       // Device name
	TimeZone  string `json:"time_zone"`  // Time zone of the device
	Status    Status `json:"status"`     // Device function status
}

type AddRoomRequest struct {
	HomeID int64  `url:"-" validate:"required" json:"-"`    // Family id
	Name   string `url:"-" validate:"required" json:"name"` // Room name
}

func (s *AddRoomRequest) Body() interface{} {
	return s
}

// Method for Request.Method()
func (s *AddRoomRequest) Method() string {
	return http.MethodPost
}

// URL for Request.URL()
func (s *AddRoomRequest) URL() string {
	return fmt.Sprintf("/v1.0/homes/%d/room", s.HomeID)
}

type AddRoomResponse int64

type UpdateRoomRequest struct {
	HomeID int64  `url:"-" validate:"required" json:"-"`    // Family ID
	RoomID int64  `url:"-" validate:"required" json:"-"`    // Room ID
	Name   string `url:"-" validate:"required" json:"name"` // Room name
}

func (s *UpdateRoomRequest) Body() interface{} {
	return s
}

// Method for Request.Method()
func (s *UpdateRoomRequest) Method() string {
	return http.MethodPut
}

// URL for Request.URL()
func (s *UpdateRoomRequest) URL() string {
	return fmt.Sprintf("/v1.0/homes/%d/rooms/%d", s.HomeID, s.RoomID)
}

type UpdateRoomResponse bool

type DeleteRoomRequest struct {
	HomeID int64 `url:"-" validate:"required" json:"-"` // Family id
	RoomID int64 `url:"-" validate:"required" json:"-"` // Room id
}

func (s *DeleteRoomRequest) Body() interface{} {
	return s
}

// Method for Request.Method()
func (s *DeleteRoomRequest) Method() string {
	return http.MethodDelete
}

// URL for Request.URL()
func (s *DeleteRoomRequest) URL() string {
	return fmt.Sprintf("/v1.0/homes/%d/rooms/%d", s.HomeID, s.RoomID)
}

type DeleteRoomResponse bool

type QueryRoomInfoRequest struct {
	HomeID int64 `url:"-" validate:"required" json:"-"` // Family id
}

func (s *QueryRoomInfoRequest) Body() interface{} {
	return s
}

// Method for Request.Method()
func (s *QueryRoomInfoRequest) Method() string {
	return http.MethodGet
}

// URL for Request.URL()
func (s *QueryRoomInfoRequest) URL() string {
	return fmt.Sprintf("/v1.0/homes/%d/rooms", s.HomeID)
}

type Room struct {
	RoomID int64  `json:"room_id"` // Room id
	Name   string `json:"name"`    // Room name
}

type QueryRoomInfoResponse struct {
	HomeID    int64   `json:"home_id"`  // Family ID
	Name      string  `json:"name"`     // Family name
	GeoName   string  `json:"geo_name"` // Family location
	Longitude float32 `json:"lon"`      // Longitude
	Latitude  float32 `json:"lat"`      // Latitude
	Rooms     []Room  `json:"rooms"`    // Room list
}

type Member struct {
	CountryCode   string `url:"-" validate:"required" json:"country_code"`   // Country code
	MemberAccount string `url:"-" validate:"required" json:"member_account"` // Member account
	Admin         bool   `url:"-" validate:"required" json:"admin"`          // Whether an administrator
	Name          string `url:"-" validate:"required" json:"name"`           // Member name
}

type AddHomeMemberRequest struct {
	HomeID    int64  `url:"-" validate:"required" json:"-"`          // Family id
	AppSchema string `url:"-" validate:"required" json:"app_schema"` // Account mark
	Member    Member `url:"-" validate:"required" json:"member"`     // Member object
}

func (s *AddHomeMemberRequest) Body() interface{} {
	return s
}

// Method for Request.Method()
func (s *AddHomeMemberRequest) Method() string {
	return http.MethodPost
}

// URL for Request.URL()
func (s *AddHomeMemberRequest) URL() string {
	return fmt.Sprintf("/v1.0/homes/%d/members", s.HomeID)
}

type AddHomeMemberResponse bool

type UpdateHomeMemberPermissionRequest struct {
	HomeID int64  `url:"-" validate:"required" json:"-"`     // Family id
	UID    string `url:"-" validate:"required" json:"-"`     // User ID in tuya
	Admin  bool   `url:"-" validate:"required" json:"admin"` // The home administrator indicates, true is the administrator and false is not the administrator.
}

func (s *UpdateHomeMemberPermissionRequest) Body() interface{} {
	return s
}

// Method for Request.Method()
func (s *UpdateHomeMemberPermissionRequest) Method() string {
	return http.MethodPut
}

// URL for Request.URL()
func (s *UpdateHomeMemberPermissionRequest) URL() string {
	return fmt.Sprintf("/v1.0/homes/%d/members/%s", s.HomeID, s.UID)
}

type UpdateHomeMemberPermissionResponse bool

type DeleteHomeMemberRequest struct {
	HomeID int64  `url:"-" validate:"required" json:"-"` // family id
	UID    string `url:"-" validate:"required" json:"-"` // User ID in tuya
}

func (s *DeleteHomeMemberRequest) Body() interface{} {
	return s
}

// Method for Request.Method()
func (s *DeleteHomeMemberRequest) Method() string {
	return http.MethodDelete
}

// URL for Request.URL()
func (s *DeleteHomeMemberRequest) URL() string {
	return fmt.Sprintf("/v1.0/homes/%d/members/%s", s.HomeID, s.UID)
}

type DeleteHomeMemberResponse bool

type QueryHomeMembersRequest struct {
	HomeID int64 `url:"-" validate:"required" json:"-"` // family id
}

func (s *QueryHomeMembersRequest) Body() interface{} {
	return s
}

// Method for Request.Method()
func (s *QueryHomeMembersRequest) Method() string {
	return http.MethodGet
}

// URL for Request.URL()
func (s *QueryHomeMembersRequest) URL() string {
	return fmt.Sprintf("/v1.0/homes/%d/members", s.HomeID)
}

type QueryHomeMembersResponse []Member

type QueryHomeRoomDeviceRequest struct {
	HomeID int64 `url:"-" validate:"required" json:"-"` // Family ID
	RoomID int64 `url:"-" validate:"required" json:"-"` // Room ID
}

func (s *QueryHomeRoomDeviceRequest) Body() interface{} {
	return s
}

// Method for Request.Method()
func (s *QueryHomeRoomDeviceRequest) Method() string {
	return http.MethodGet
}

// URL for Request.URL()
func (s *QueryHomeRoomDeviceRequest) URL() string {
	return fmt.Sprintf("/v1.0/homes/%d/rooms/%d/devices", s.HomeID, s.RoomID)
}

type Device struct {
	ID        string `json:"id"`         // device ID
	UID       string `json:"uid"`        // User id
	LocalKey  string `json:"local_key"`  // Key
	Category  string `json:"category"`   // Product Category
	ProductID string `json:"product_id"` // Product id
	Sub       bool   `json:"sub"`        // Whether it is a sub-equipment (true: yes, false: no)
	UUID      string `json:"uuid"`       // Device unique identifier
	OwnerID   string `json:"owner_id"`   // Device owner id
	Online    bool   `json:"online"`     // Device online status
	IP        string `json:"ip"`         // Device IP
	Name      string `json:"name"`       // Device name
	TimeZone  string `json:"time_zone"`  // Device time zone
	Status    Status `json:"status"`     // Device function status
}

type QueryHomeRoomDeviceResponse []Device

type AddHomeRoomDeviceRequest struct {
	HomeID    int64    `url:"-" validate:"required" json:"-"`          // home ID
	RoomID    int64    `url:"-" validate:"required" json:"-"`          // room ID
	DeviceIDs []string `url:"-" validate:"required" json:"device_ids"` // device IDs
}

func (s *AddHomeRoomDeviceRequest) Body() interface{} {
	return s
}

// Method for Request.Method()
func (s *AddHomeRoomDeviceRequest) Method() string {
	return http.MethodPost
}

// URL for Request.URL()
func (s *AddHomeRoomDeviceRequest) URL() string {
	return fmt.Sprintf("/v1.0/homes/%d/rooms/%d/devices", s.HomeID, s.RoomID)
}

type AddHomeRoomDeviceResponse bool

type UpdateHomeRoomDeviceRequest struct {
	HomeID    int64    `url:"-" validate:"required" json:"-"`          // home ID
	RoomID    int64    `url:"-" validate:"required" json:"-"`          // room ID
	DeviceIDs []string `url:"-" validate:"required" json:"device_ids"` // device IDs
}

func (s *UpdateHomeRoomDeviceRequest) Body() interface{} {
	return s
}

// Method for Request.Method()
func (s *UpdateHomeRoomDeviceRequest) Method() string {
	return http.MethodPut
}

// URL for Request.URL()
func (s *UpdateHomeRoomDeviceRequest) URL() string {
	return fmt.Sprintf("/v1.0/homes/%d/rooms/%d/devices", s.HomeID, s.RoomID)
}

type UpdateHomeRoomDeviceResponse bool

type DeleteHomeRoomDeviceRequest struct {
	HomeID    int64    `url:"-" validate:"required" json:"-"`          // home ID
	RoomID    int64    `url:"-" validate:"required" json:"-"`          // room ID
	DeviceIDs []string `url:"-" validate:"required" json:"device_ids"` // devices ID
}

func (s *DeleteHomeRoomDeviceRequest) Body() interface{} {
	return s
}

// Method for Request.Method()
func (s *DeleteHomeRoomDeviceRequest) Method() string {
	return http.MethodDelete
}

// URL for Request.URL()
func (s *DeleteHomeRoomDeviceRequest) URL() string {
	return fmt.Sprintf("/v1.0/homes/%d/rooms/%d/devices", s.HomeID, s.RoomID)
}

type QueryUserHomeRequest struct {
	UID int64 `url:"-" validate:"required" json:"-"` // home ID
}

func (s *QueryUserHomeRequest) Body() interface{} {
	return s
}

// Method for Request.Method()
func (s *QueryUserHomeRequest) Method() string {
	return http.MethodGet
}

// URL for Request.URL()
func (s *QueryUserHomeRequest) URL() string {
	return fmt.Sprintf("/v1.0/users/%d/homes", s.UID)
}

type QueryUserHomeResponse []QueryHomeResponse
