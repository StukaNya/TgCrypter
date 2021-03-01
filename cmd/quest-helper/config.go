package main

import (
	"flag"
	"os"

	httpserver "github.com/StukaNya/QuestHelper/http-server"
	controller "github.com/StukaNya/QuestHelper/model/controller"
	"gopkg.in/yaml.v3"
)

const (
	localConfigPath = "./config.yml"
)

type Logger struct {
	Format   string `yaml:"format"`
	LogLevel string `yaml:"log_level"`
}

type DatabaseConfig struct {
	DatabaseURL string `yaml:"database_url"`
}

type Config struct {
	Logger           Logger
	DbConfig         DatabaseConfig
	ServerConfig     httpserver.ServerConfig
	ControllerConfig controller.ControllerConfig
}

func NewConfig() (*Config, error) {
	flag.Parse()
	cfg, err := os.Open(configPath)
	if err != nil {
		cfg, err = os.Open(localConfigPath)
		if err != nil {
			return nil, err
		}
	}

	config := &Config{}
	if err = yaml.NewDecoder(cfg).Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}
