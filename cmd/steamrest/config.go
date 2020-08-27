package main

import (
	"github.com/StukaNya/SteamREST/internal/app/controller"
	"github.com/StukaNya/SteamREST/internal/app/httpserver"
)

// Logger config
type Logger struct {
	LogLevel string
}

// Database config
type DatabaseConfig struct {
	DatabaseURL string
}

// Global config
type Config struct {
	Logger           Logger
	DbConfig         DatabaseConfig
	ServerConfig     httpserver.ServerConfig
	ControllerConfig controller.ControllerConfig
}

// Return config instance
func NewConfig() *Config {
	return &Config{
		Logger: Logger{
			LogLevel: "debug",
		},
		DbConfig: DatabaseConfig{
			DatabaseURL: "",
		},
		ServerConfig:     *httpserver.NewConfig(),
		ControllerConfig: *controller.NewConfig(),
	}
}
