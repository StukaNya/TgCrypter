package main

import (
	"flag"
	"os"

	"github.com/StukaNya/TgCrypter/api"
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
	Logger       Logger           `yaml:"logger"`
	DbConfig     DatabaseConfig   `yaml:"database"`
	ServerConfig api.ServerConfig `yaml:"api_server"`
}

func NewConfig() *Config {
	return &Config{
		ServerConfig: api.ServerConfig{},
	}
}

func (cfg *Config) Parse() error {
	flag.Parse()
	cfgFile, err := os.Open(configPath)
	if err != nil {
		cfgFile, err = os.Open(localConfigPath)
		if err != nil {
			return err
		}
	}

	if err = yaml.NewDecoder(cfgFile).Decode(cfg); err != nil {
		return err
	}

	return nil
}
