package controller

import (
	"testing"

	"github.com/StukaNya/SteamREST/internal/app/store"
	"github.com/sirupsen/logrus"
)

// TestControllerLoad...
func TestControllerLoad(t *testing.T) {

	// Init config and log instance
	logger := logrus.New()
	logger.SetLevel(logrus.WarnLevel)
	config := ControllerConfig{}

	// Init mock store object with test data
	testApp := store.AppInfo{1, "Stellaris"}
	st := store.NewStoreMock(5)
	st.InsertAppInfo(&testApp)

	// Init controller
	c := NewController(st, logger, &config)

	// Test load App info func
	appName, _ := c.AppInfo(testApp.AppID)
	expectedName := "1.Stellaris"
	if appName != expectedName {
		t.Error(
			"For", testApp.AppID,
			"expected", expectedName,
			"got", appName,
		)
	}

}
