package controller

import (
	"github.com/StukaNya/SteamREST/internal/app/store"
	"github.com/sirupsen/logrus"
)

//TODO: change name, mb Manager
// Dependencies: db -> ctrl -> server
// add HTTP requests to SteamAPI and save to store DB

// Controller struct object
type Controller struct {
	config  *ControllerConfig
	logger  *logrus.Logger
	dbStore *store.Store
}

// Return controller instance
func NewController(log *logrus.Logger, config *ControllerConfig, store *store.Store) *Controller {
	return &Controller{
		config:  config,
		logger:  log,
		dbStore: store,
	}
}

// Получить все игры
// https://api.steampowered.com/ISteamApps/GetAppList/v2/?format=json
// Получить инфо об игре 57690 = Tropico 4
// http://store.steampowered.com/api/appdetails?appids=57690&cc=us&l=en
