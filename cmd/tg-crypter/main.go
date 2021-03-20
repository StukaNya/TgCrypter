package main

import (
	"context"
	"database/sql"
	"flag"
	"log"

	"github.com/StukaNya/TgCrypter/api"
	store "github.com/StukaNya/TgCrypter/storage/user"

	"github.com/sirupsen/logrus"
)

var (
	configPath string
)

// TODO: update for Linux
func init() {
	flag.StringVar(&configPath, "config-path", "./config.yaml", "path to config file")
}

func main() {
	// Init context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Load config from YAML file
	config, err := NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Configure logger
	logger := logrus.New()
	level, err := logrus.ParseLevel(config.Logger.LogLevel)
	if err != nil {
		log.Fatal(err)
	}
	logger.SetLevel(level)

	// Configure DB and store layer
	dbURL := config.DbConfig.DatabaseURL
	db, err := newDB(dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// User storage
	userSt := store.NewUserStorage(db)
	err = userSt.InitTable(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Server startup
	server := api.NewAPIServer(logger, &config.ServerConfig, userSt)
	err = server.Start()
	if err != nil {
		log.Fatal(err)
	}
}

// Open database
func newDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return db, err
	}

	return db, nil
}
