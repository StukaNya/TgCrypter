package store

// Database access interface
type StoreAccess interface {
	GetAppInfo(appId int) (AppInfo, error)
	InsertAppInfo(AppInfo) error
}

// Steam app info struct
type AppInfo struct {
	Id   int
	Name string
}

//TODO: add SQL requests
func (s *Store) GetAppInfo(appId int) (AppInfo, error) {
	return AppInfo{}, nil
}

func (s *Store) InsertAppInfo(AppInfo) error {
	return nil
}
