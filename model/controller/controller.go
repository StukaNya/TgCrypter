package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

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

// AppList is JSON format struct
type AppList struct {
	Apps []store.AppInfo `json:"apps"`
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

	serialInfo := strconv.Itoa(appInfo.AppID) + "." + appInfo.Name
	c.logger.Info("Get app info from DB: ", serialInfo)
	return serialInfo, nil
}

// LoadAppList from Steam API
func (c *Controller) LoadAppList() error {
	// Get responce from HTTP request
	resp, err := http.Get(c.config.AppList)
	if err != nil {
		c.logger.Info("Error during GET AppList responce from HTTP request")
		return err
	}
	defer resp.Body.Close()

	// Read byte data from responce
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.logger.Info("Error during read data from response")
		return err
	}

	// Unmarshal JSON data to AppList Struct
	list := AppList{}
	err = json.Unmarshal(data, &list)
	if err != nil {
		c.logger.Info("Error during unmarshal JSON data to AppList struct")
		return err
	}

	// Insert data to DB
	for _, app := range list.Apps {
		err = c.dbStore.InsertAppInfo(&app)
		if err != nil {
			c.logger.Info("Error during insert app to db: ", app.Name)
		}
	}

	return nil
}
