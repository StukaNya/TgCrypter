package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/StukaNya/SteamREST/model/crypter"
	"github.com/StukaNya/SteamREST/model/session"
	uuid "github.com/satori/go.uuid"
)

var (
	_ session.PinCodeStorer  = (*PinCodeStorage)(nil)
	_ crypter.PinCodeFetcher = (*PinCodeStorage)(nil)
)

type PinCodeStorage struct {
	db *sql.DB
}

func (s *PinCodeStorage) InitTable(ctx context.Context) error {
	const query = `CREATE TABLE pin_code (
			id			uuid 	PRIMARY KEY DEFAULT uuid_generate_v4(),
			user_id  	uuid 	NOT NULL REFERENCES user(id),
			hash 		bytea 	NOT NULL,
		);`
	_, err := s.db.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to create pin code table: %v", err)
	}

	return nil
}

func (s *PinCodeStorage) StorePinHash(ctx context.Context, userID uuid.UUID, pinHash []byte) error {
	const query = "INSERT INTO pin_code (session_id, hash) VALUES ($1, $2)"
	_, err := s.db.ExecContext(ctx, query, userID, pinHash)
	if err != nil {
		return fmt.Errorf("failed to store new pin code: %v", err)
	}

	return nil
}

func (s *PinCodeStorage) FetchPinHash(ctx context.Context, userID uuid.UUID) ([]byte, error) {
	const query = "SELECT hash FROM pin_code WHERE session_id=?"
	row, err := s.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("unable to select pin code hash from db: %v", err)
	}

	var pinHash []byte
	if err := row.Scan(&pinHash); err != nil {
		return pinHash, fmt.Errorf("unable to scan code hash: %v", err)
	}

	return pinHash, nil
}
