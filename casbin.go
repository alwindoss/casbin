package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

// Config is the configuration for Casbin middleware
type Config struct {
	Enforcer *casbin.Enforcer

	// Authorizer defines a function you can pass
	// to check the credentials however you want.
	// It will be called with a username and password
	// and is expected to return true or false to indicate
	// that the credentials were approved or not.
	//
	// Optional. Default: nil.
	Authorizer func(sub, obj, action string) bool

	// Unauthorized defines the response body for unauthorized responses.
	// By default it will return with a 401 Unauthorized and the correct WWW-Auth header
	//
	// Optional. Default: nil
	Unauthorized fiber.Handler
}

// DefaultConfig is the default config
var DefaultConfig = Config{}

// New creates a new middleware for Casbin middleware
func New(cfg Config) fiber.Handler {
	if cfg.Enforcer == nil {

	}
	// Return new handler
	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Next returns true
		// if cfg.Next != nil && cfg.Next(c) {
		// 	return c.Next()
		// }

		// Authentication failed
		return cfg.Unauthorized(c)
	}
}
