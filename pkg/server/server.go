package server

import (
	"github.com/AnkushJadhav/kamaji-node/pkg/store"
)

// Config is the config for the server
type Config struct {
	PopulatePool  bool
	EnableTLS     bool
	BindIP        string
	Port          int
	StorageDriver store.Driver
	JWTSecret     string
}

// Server is the interface to be fulfilled by any web server running on this application
type Server interface {
	Bootstrap(*Config) error
	Start() error
	Stop() error
}
