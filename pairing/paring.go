package pairing

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type RequestExtension struct {
	UUID string `url:"-" json:"uuid"`
}

type PairingDeviceRequest struct {
	ParingType string           `url:"-" validate:"required" json:"paring_type"`  // paring type,support BLE, AP, EZ
	UID        string           `url:"-" validate:"required" json:"uid"`          // Tuya user id
	TimeZoneID string           `url:"-" validate:"required" json:"time_zone_id"` // User's time zone id, device needs for daylight saving time
	HomeID     string           `url:"-" json:"home_id"`                          // Home id . If empty ,default user home
	Extension  RequestExtension `url:"-" json:"extension"`                        // extension, the paring type is BLE and the device UUID is passed in
}

// Method for Request.Method()
func (s *PairingDeviceRequest) Method() string {
	return http.MethodPost
}

// URL for Request.URL()
func (s *PairingDeviceRequest) URL() string {
	return "/v1.0/device/paring/token"
}

type ResponseExtension struct {
	EncryptKey string `json:"encrypt_key"` // encrypt key
	Random     string `json:"random"`      // Encrypted string
}

type PairingDeviceResponse struct {
	ExpireTime int64             `json:"expire_time"` // Token expiration time
	Region     string            `json:"region"`      // Current region, support:AY EU US
	Token      string            `json:"token"`       // paring token
	Secret     string            `json:"secret"`      // secret
	Extension  ResponseExtension `json:"extension"`   // extension
}

type QueryPairingDeviceListRequest struct {
	Token string `url:"-" validate:"required" json:"-"` // device paring token
}

// Method for Request.Method()
func (s *QueryPairingDeviceListRequest) Method() string {
	return http.MethodGet
}

// URL for Request.URL()
func (s *QueryPairingDeviceListRequest) URL() string {
	return "/v1.0/device/paring/tokens/" + s.Token
}

type QueryPairingDeviceListResponse struct {
	Failed []struct { // Paring failed devices
		DeviceID string `json:"device_id"` // Device id
		Code     string `json:"code"`      // errorCode
		Msg      string `json:"msg"`       // errorMsg
		Name     string `json:"name"`      // Device name
	} `json:"failed"`
	Success []struct { // Paring success devices
		Category  string `json:"category"`   // Device category
		DeviceID  string `json:"device_id"`  // Device id
		Name      string `json:"name"`       // Device name
		ProductID string `json:"product_id"` // Product id
	} `json:"success"`
}

type ZigbeeGatewayPairingRequest struct {
	DeviceID string `url:"-" validate:"required" json:"-"` // device ID
	Duration int    `url:"duration" json:"-"`              // Gateway discovery time, default 100 seconds, maximum 3600 seconds , 0: stop discovery
}

// Method for Request.Method()
func (s *ZigbeeGatewayPairingRequest) Method() string {
	return http.MethodPut
}

// URL for Request.URL()
func (s *ZigbeeGatewayPairingRequest) URL() string {
	v, _ := query.Values(s)
	return "/v1.0/devices/" + s.DeviceID + "/enabled-sub-discovery?" + v.Encode()
}

type ZigbeeGatewayPairingResponse bool

type ZigbeeGatewayPairingListRequest struct {
	DeviceID      string `url:"-" validate:"required" json:"-"`              // gateway ID
	DiscoveryTime int64  `url:"discovery_time" validate:"required" json:"-"` // Gateway enabled time
}

// Method for Request.Method()
func (s *ZigbeeGatewayPairingListRequest) Method() string {
	return http.MethodGet
}

// URL for Request.URL()
func (s *ZigbeeGatewayPairingListRequest) URL() string {
	v, _ := query.Values(s)
	return "/v1.0/devices/" + s.DeviceID + "/list-sub?" + v.Encode()
}

type ZigbeeGatewayPairingListResponse struct {
	ID         string `json:"id"`          // Device id
	Category   string `json:"category"`    // Category
	ProductID  string `json:"product_id"`  // Product id
	UUID       string `json:"uuid"`        // Device chip id
	OwnerID    string `json:"owner_id"`    // Home id
	Online     bool   `json:"online"`      // Device online state
	ActiveTime int64  `json:"active_time"` // Active time
	UpdateTime int64  `json:"update_time"` // Device update time
}

type ZigbeeGatewayPairedListRequest struct {
	DeviceID string `url:"-" validate:"required" json:"-"` // Device id
}

// Method for Request.Method()
func (s *ZigbeeGatewayPairedListRequest) Method() string {
	return http.MethodGet
}

// URL for Request.URL()
func (s *ZigbeeGatewayPairedListRequest) URL() string {
	return "/v1.0/devices/" + s.DeviceID + "/sub-devices"
}

type ZigbeeGatewayPairedListResponse struct {
	ID         string `json:"id"`          // Device id
	ProductID  string `json:"product_id"`  // Product id
	OwnerID    string `json:"owner_id"`    // Home id
	Online     bool   `json:"online"`      // Device online state
	Name       string `json:"name"`        // Device name
	UpdateTime int64  `json:"update_time"` // Update time
	ActiveTime int64  `json:"active_time"` // Last actice time
	Category   string `json:"category"`    // Category
}
