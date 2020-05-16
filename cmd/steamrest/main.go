package main

import (
	"flag"
	"log"

	"github.com/StukaNya/SteamREST/internal/app/httpserver"
	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
)

var (
	configPath string
)

// TODO: update for Linux
func init() {
	flag.StringVar(&configPath, "config-path", "C:\\pr\\SteamREST\\configs\\steamrest.toml", "path to config file")
}

func main() {
	flag.Parse()

	// Load TOML file
	tomlTree, err := toml.LoadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	// Parse TOML into config struct
	config := NewConfig()
	tomlTree.Unmarshal(config)

	// Configure logger
	logger := logrus.New()
	level, err := logrus.ParseLevel(config.Logger.LogLevel)
	if err != nil {
		log.Fatal(err)
	}
	logger.SetLevel(level)

	// Server startup
	server := httpserver.New(logger, config.ServerConfig)
	err = server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
