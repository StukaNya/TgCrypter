package store

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// Database store object
type Store struct {
	config *StoreConfig
	logger *logrus.Logger
	db     *sql.DB
}

// Return db store instance
func NewStore(log *logrus.Logger, config *StoreConfig) *Store {
	return &Store{
		config: config,
		logger: log,
	}
}

// Open database
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	s.db = db

	return nil
}

// Close database
func (s *Store) Close() {
	s.db.Close()
}
