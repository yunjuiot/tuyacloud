package tuyacloud

import (
	"net/http"

	"github.com/go-log/log"
)

type options struct {
	httpClient *http.Client
	logger     log.Logger
	storage    TokenStorage
}

type Option func(o *options)

// HTTPClient setup HTTPClient
func HTTPClient(c *http.Client) Option {
	return func(o *options) {
		o.httpClient = c
	}
}

// Logger setup log.Logger interface.
func Logger(l log.Logger) Option {
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

