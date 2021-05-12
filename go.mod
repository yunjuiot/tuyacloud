// Deprecated: This repo is transfered to github.com/ekeynow/tuyacloud.
module github.com/yunjuiot/tuyacloud

go 1.17

require (
	github.com/go-log/log v0.2.0
	github.com/go-playground/validator/v10 v10.3.0
	github.com/golang/mock v1.4.4
	github.com/google/go-querystring v1.0.0
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.6.1
)

// Retract due to change the owner.
retract (
	v0.2.0
	v0.1.3
	v0.1.2
	v0.1.1
	v0.1.0
)
