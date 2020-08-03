package main

import (
	"flag"
	"log"

	"github.com/StukaNya/SteamREST/internal/app/controller"
	"github.com/StukaNya/SteamREST/internal/app/httpserver"
	"github.com/StukaNya/SteamREST/internal/app/store"
	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
)

var (
	configPath string
)

// TODO: update for Linux
func init() {
	flag.StringVar(&configPath, "config-path", "/home/stuka/go/SteamREST/configs/steamrest.toml", "path to config file")
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

	// Configure DB
	st := store.NewStore(logger, &config.StoreConfig)
	err = st.Open()
	if err != nil {
		log.Fatal(err)
	}

	// Configure manage controller
	ctrl := controller.NewController(st, logger, &config.ControllerConfig)
	err = ctrl.LoadAppList()
	if err != nil {
		log.Fatal(err)
	}

	// Server startup
	server := httpserver.New(logger, &config.ServerConfig, ctrl)
	err = server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
