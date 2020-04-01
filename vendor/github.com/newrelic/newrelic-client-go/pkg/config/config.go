// Package config provides cross-cutting configuration support for the newrelic-client-go project.
package config

import (
	"net/http"
	"time"

	"github.com/newrelic/newrelic-client-go/internal/logging"
	"github.com/newrelic/newrelic-client-go/internal/version"
	"github.com/newrelic/newrelic-client-go/pkg/region"
)

// Config contains all the configuration data for the API Client.
type Config struct {
	// PersonalAPIKey to authenticate API requests
	// see: https://docs.newrelic.com/docs/apis/get-started/intro-apis/types-new-relic-api-keys#personal-api-key
	PersonalAPIKey string

	// AdminAPIKey to authenticate API requests
	// Note this will be deprecated in the future!
	// see: https://docs.newrelic.com/docs/apis/get-started/intro-apis/types-new-relic-api-keys#admin
	AdminAPIKey string

	// region of the New Relic platform to use
	region *region.Region

	// Timeout is the client timeout for HTTP requests.
	Timeout *time.Duration

	// HTTPTransport allows customization of the client's underlying transport.
	HTTPTransport http.RoundTripper

	// UserAgent updates the default user agent string used by the client.
	UserAgent string

	// ServiceName is for New Relic internal use only.
	ServiceName string

	// LogLevel can be one of the following values:
	// "panic", "fatal", "error", "warn", "info", "debug", "trace"
	LogLevel string

	// LogJSON toggles formatting of log entries in JSON format.
	LogJSON bool

	// Logger allows customization of the client's underlying logger.
	Logger logging.Logger
}

// New creates a default configuration and returns it
func New() Config {
	regCopy := *region.Default

	return Config{
		region:    &regCopy,
		UserAgent: "newrelic/newrelic-client-go",
		LogLevel:  "info",
	}
}

// Region returns the region configuration struct
// if one has not been set, use the default region
func (c *Config) Region() *region.Region {
	if c.region == nil {
		regCopy := *region.Default
		c.region = &regCopy
	}

	return c.region
}

// SetRegion configures the region
func (c *Config) SetRegion(reg *region.Region) error {
	if reg == nil {
		return region.ErrorNil()
	}

	c.region = reg

	return nil
}

// GetLogger returns a logger instance based on the config values.
func (c *Config) GetLogger() logging.Logger {
	if c.Logger != nil {
		return c.Logger
	}

	l := logging.NewStructuredLogger().
		SetDefaultFields(map[string]string{"newrelic-client-go": version.Version}).
		LogJSON(c.LogJSON).
		SetLogLevel(c.LogLevel)

	c.Logger = l
	return l
}
