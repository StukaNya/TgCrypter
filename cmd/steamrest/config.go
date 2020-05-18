package main

import (
	"github.com/StukaNya/SteamREST/internal/app/controller"
	"github.com/StukaNya/SteamREST/internal/app/httpserver"
	"github.com/StukaNya/SteamREST/internal/app/store"
)

// Logger config
type Logger struct {
	LogLevel string
}

// Global config
type Config struct {
	Logger           Logger
	ServerConfig     httpserver.ServerConfig
	StoreConfig      store.StoreConfig
	ControllerConfig controller.ControllerConfig
}

// Return config instance
func NewConfig() *Config {
	return &Config{
		Logger: Logger{
			LogLevel: "debug",
		},
		ServerConfig:     *httpserver.NewConfig(),
		StoreConfig:      *store.NewConfig(),
		ControllerConfig: *controller.NewConfig(),
	}
}
