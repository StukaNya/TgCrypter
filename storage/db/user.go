package store

import (
	"context"
	"database/sql"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type UserStotage struct {
	db *sql.DB
}

func (s *UserStotage) InitTable(ctx context.Context) error {
	const query = `CREATE TABLE user (
			id			uuid 	PRIMARY KEY DEFAULT uuid_generate_v4(),
			session_id	int		NOT NULL REFERENCES session(id),
			name 		text,
			created_at 	time	NOT NULL DEFAULT(now() at time zone 'utc')
			is_active	boolean NOT NULL DEFAULT false
		);`
	_, err := s.db.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to create user table: %v", err)
	}

	return nil
}

func (s *UserStotage) StoreUser(ctx context.Context, sessionID int, userName string) error {
	const query = "INSERT INTO user (session_id, name) VALUES ($1, $2)"
	_, err := s.db.ExecContext(ctx, query, sessionID, userName)
	if err != nil {
		return fmt.Errorf("failed to store new user: %v", err)
	}

	return nil
}

func (s *UserStotage) FetchUser(ctx context.Context, userID uuid.UUID) (string, error) {
	const query = "SELECT name FROM user WHERE id=?"
	row, err := s.db.QueryContext(ctx, query, userID)
	if err != nil {
		return "", fmt.Errorf("unable to select user name from db: %v", err)
	}

	var userName string
	if err := row.Scan(userName); err != nil {
		return userName, fmt.Errorf("unable to scan user name: %v", err)
	}

	return userName, nil
}
