package store

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"
)

// Store app info object
type Store struct {
	logger *logrus.Logger
	db     *sql.DB
}

// Model database access interface
type Model interface {
	InitTable(ctx context.Context) error
	GetAppInfo(ctx context.Context, appID int) (*AppInfo, error)
	InsertAppInfo(ctx context.Context, info *AppInfo) error
}

type AppInfo struct {
	AppID int    `json:"appid"`
	Name  string `json:"name"`
}

func NewStore(log *logrus.Logger, db *sql.DB) *Store {
	return &Store{
		logger: log,
		db:     db,
	}
}

func (s *Store) InitTable(ctx context.Context) error {
	const query = "CREATE TABLE apps (app_id int PRIMARY KEY, app_name text NOT NULL);"
	_, err := s.db.ExecContext(ctx, query)
	if err != nil {
		s.logger.Info("Failed to create table on DB")
		return err
	}

	return nil
}

func (s *Store) GetAppInfo(ctx context.Context, appID int) (*AppInfo, error) {
	row, err := s.db.QueryContext(ctx, "SELECT app_id, app_name FROM apps WHERE app_id =?", appID)
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

func (s *Store) InsertAppInfo(ctx context.Context, info *AppInfo) error {
	_, err := s.db.ExecContext(ctx, "INSERT INTO apps (app_id, app_name) VALUES ($1, $2)", info.AppID, info.Name)
	if err != nil {
		s.logger.Info("Failed to INSERT info to DB")
		return err
	}

	return nil
}
