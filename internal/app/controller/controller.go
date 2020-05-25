package controller

import (
	"strconv"

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
	dbStore store.Model
}

// NewController return controller instance
func NewController(store store.Model, log *logrus.Logger, config *ControllerConfig) *Controller {
	return &Controller{
		config:  config,
		logger:  log,
		dbStore: store,
	}
}

// AppInfo return serialize string info of App
func (c *Controller) AppInfo(appID int) (string, error) {
	appInfo, err := c.dbStore.GetAppInfo(appID)
	if err != nil {
		return "", err
	}

	serialInfo := strconv.Itoa(appID) + "." + appInfo.Name
	c.logger.Info("Get app info from DB: ", serialInfo)
	return serialInfo, nil
}

// LoadApps ...
func (c *Controller) LoadApps() error {
	return nil
}

// Получить все игры
// https://api.steampowered.com/ISteamApps/GetAppList/v2/?format=json
// Получить инфо об игре 57690 = Tropico 4
// http://store.steampowered.com/api/appdetails?appids=57690&cc=us&l=en
