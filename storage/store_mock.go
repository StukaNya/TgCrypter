package store

import (
	_ "github.com/lib/pq"
)

// StoreMock object
type StoreMock struct {
	appsData map[int]string
}

// NewStoreMock return new store mock instance
func NewStoreMock(size int) *StoreMock {
	s := StoreMock{appsData: make(map[int]string)}

	return &s
}

// GetAppInfo data from mock object
func (s *StoreMock) GetAppInfo(appID int) (*AppInfo, error) {
	app := AppInfo{appID, s.appsData[appID]}
	return &app, nil
}

// InsertAppInfo data to mock object
func (s *StoreMock) InsertAppInfo(info *AppInfo) error {
	s.appsData[info.AppID] = info.Name
	return nil
}
