package tuyacloud

import (
	"github.com/go-log/log"
)

type options struct {
	httpClient HTTPClient
	logger     log.Logger
	storage    TokenStorage
}

// Option settings.
type Option func(o *options)

// WithHTTPClient setup HTTPClient
func WithHTTPClient(c HTTPClient) Option {
	return func(o *options) {
		o.httpClient = c
	}
}

// WithLogger setup log.Logger interface.
func WithLogger(l log.Logger) Option {
	return func(o *options) {
		o.logger = l
	}
}

// WithTokenStore setup token storage interface.
func WithTokenStore(s TokenStorage) Option {
	return func(o *options) {
		o.storage = s
	}
}
