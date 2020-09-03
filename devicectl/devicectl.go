package devicectl

import "net/http"

type Function struct {
	Code   string `json:"code"`   // code
	Type   string `json:"type"`   // type
	Values string `json:"values"` // value range
}

type QueryCategoryFunctionRequest struct {
	Category string `url:"-" validate:"required" json:"-"` // category name
}

// Method for Request.Method()
func (s *QueryCategoryFunctionRequest) Method() string {
	return http.MethodGet
}

// URL for Request.URL()
func (s *QueryCategoryFunctionRequest) URL() string {
	return "/v1.0/functions/" + s.Category
}

type QueryCategoryFunctionResponse Function

type QueryDeviceFunctionRequest struct {
	DeviceID string `url:"-" validate:"required" json:"-"` // Device id
}

// Method for Request.Method()
func (s *QueryDeviceFunctionRequest) Method() string {
	return http.MethodGet
}

// URL for Request.URL()
func (s *QueryDeviceFunctionRequest) URL() string {
	return "/v1.0/devices/" + s.DeviceID + "/functions"
}

type DeviceFunction struct {
	Code   string `json:"code"`   // code
	Type   string `json:"type"`   // type
	Values string `json:"values"` // value ranges
	Name   string `json:"name"`   // Instruction name
	Desc   string `json:"desc"`   // description
}

type QueryDeviceFunctionResponse struct {
	Category  string           `json:"category"`  // Category
	Functions []DeviceFunction `json:"functions"` // Functions
}

type QuerySpecificationsRequest struct {
	DeviceID string `url:"-" validate:"required" json:"-"` // device id
}

// Method for Request.Method()
func (s *QuerySpecificationsRequest) Method() string {
	return http.MethodGet
}

// URL for Request.URL()
func (s *QuerySpecificationsRequest) URL() string {
	return "/v1.0/devices/" + s.DeviceID + "/specifications"
}

type QuerySpecificationsResponse struct {
	Category  string   `json:"category"`  // Product category
	Functions Function `json:"functions"` // Instruction set
	Status    Function `json:"status"`    // State set
}

type DeviceCommand struct {
	Code  string      `json:"code"`
	Value interface{} `json:"value"`
}

type ControlDeviceRequest struct {
	DeviceID string          `url:"-" validate:"required" json:"-"`        // device id
	Commands []DeviceCommand `url:"-" validate:"required" json:"commands"` // commands
}

// Method for Request.Method()
func (s *ControlDeviceRequest) Method() string {
	return http.MethodPost
}

// URL for Request.URL()
func (s *ControlDeviceRequest) URL() string {
	return "/v1.0/devices/" + s.DeviceID + "/commands"
}

type ControlDeviceResponse bool

type QueryDeviceStatusRequest struct {
	DeviceID string `url:"-" validate:"required" json:"-"` // device id
}

// Method for Request.Method()
func (s *QueryDeviceStatusRequest) Method() string {
	return http.MethodGet
}

// URL for Request.URL()
func (s *QueryDeviceStatusRequest) URL() string {
	return "/v1.0/devices/" + s.DeviceID + "/status"
}

type QueryDeviceStatusResponse []DeviceCommand
