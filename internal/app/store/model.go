package store

// Model database access interface
type Model interface {
	GetAppInfo(appID int) (*AppInfo, error)
	InsertAppInfo(info *AppInfo) error
}

// AppInfo stores DB row of this app
type AppInfo struct {
	AppID int    `json:"appid"`
	Name  string `json:"name"`
}

// GetAppInfo select info of app with current ID from DB
func (s *Store) GetAppInfo(appID int) (*AppInfo, error) {
	row, err := s.db.Query("SELECT app_id, app_name FROM apps WHERE app_id =?", appID)
	if err != nil {
		s.logger.Info("Failed to SELECT info from DB")
		return nil, err
	}
	defer row.Close()

	info := new(AppInfo)
	err = row.Scan(&info.AppID, &info.Name)
	if err != nil {
		s.logger.Info("Failed to scan info from DB row")
		return nil, err
	}

	return info, nil
}

// InsertAppInfo insert info of app to DB
func (s *Store) InsertAppInfo(info *AppInfo) error {
	_, err := s.db.Exec("INSERT INTO apps (app_id, app_name) VALUES ($1, $2)", info.AppID, info.Name)
	if err != nil {
		s.logger.Info("Failed to INSERT info to DB")
		return err
	}

	return nil
}
