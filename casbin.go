package casbin

import (
	"io"

	"github.com/gofiber/fiber/v2"
)

// Config is the configuration for Casbin middleware
type Config struct {
	Model io.ReadWriteCloser
}

// New creates a new middleware for Casbin middleware
func New(config Config) fiber.Handler {
	return nil
}
