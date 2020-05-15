package main

import (
	"flag"
	"log"

	"github.com/StukaNya/SteamREST/internal/app/httpserver"
	"github.com/pelletier/go-toml"
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
	config := httpserver.NewConfig()
	tomlTree.Unmarshal(&config)

	server := httpserver.New(config)
	err = server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
