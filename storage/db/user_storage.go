package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/StukaNya/SteamREST/model/session"
	uuid "github.com/satori/go.uuid"
)

var (
	_ session.UserRepository = (*UserStotage)(nil)
)

type UserStotage struct {
	db *sql.DB
}

func (s *UserStotage) InitTable(ctx context.Context) error {
	const query = `CREATE TABLE user (
			id			uuid 	PRIMARY KEY DEFAULT uuid_generate_v4(),
			chat_id		int,
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

func (s *UserStotage) RegisterUser(ctx context.Context, info *session.UserInfo) (uuid.UUID, error) {
	userID := uuid.NewV4()
	const query = "INSERT INTO user (id, chat_id, name, created_at) VALUES ($1, $2, $3)"
	_, err := s.db.ExecContext(ctx, query, userID, info.ChatID, info.Name, info.RegisteredAt)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to register new user: %v", err)
	}

	return userID, nil
}

func (s *UserStotage) FetchUser(ctx context.Context, userID uuid.UUID) (*session.UserInfo, error) {
	const query = "SELECT chat_id, name, created_at FROM user WHERE id=?"
	row, err := s.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("unable to select user info from db: %v", err)
	}

	userInfo := new(session.UserInfo)
	if err := row.Scan(userInfo.ChatID, userInfo.Name, userInfo.RegisteredAt); err != nil {
		return nil, fmt.Errorf("unable to scan user info: %v", err)
	}

	return userInfo, nil
}
