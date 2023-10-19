// Package app implements the application layer
package app

import (
	"github.com/mirshahriar/marketplace/config"
	"github.com/mirshahriar/marketplace/internal/ports"
)

// Adapter implements the application layer
type Adapter struct {
	config config.AppConfig
	db     ports.DBPort
}

// This validates that Adapter implements the ports.APIPort interface
var _ ports.APIPort = Adapter{}

// NewApplication returns a new adapter with the given db port
func NewApplication(
	cfg config.AppConfig,
	db ports.DBPort,
) *Adapter {
	return &Adapter{
		config: cfg,
		db:     db,
	}
}
