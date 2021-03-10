package store

import (
	"context"
	"database/sql"
	"fmt"
)

type PinCodeStorage struct {
	db *sql.DB
}

func (s *PinCodeStorage) InitTable(ctx context.Context) error {
	const query = `CREATE TABLE pin_code (
			id			uuid 	PRIMARY KEY DEFAULT uuid_generate_v4(),
			session_id  int 	NOT NULL REFERENCES session(id),
			hash 		text 	NOT NULL,
		);`
	_, err := s.db.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to create pin code table: %v", err)
	}

	return nil
}

func (s *PinCodeStorage) StorePinHash(ctx context.Context, sessionID int, pinHash string) error {
	const query = "INSERT INTO pin_code (session_id, hash) VALUES ($1, $2)"
	_, err := s.db.ExecContext(ctx, query, sessionID, pinHash)
	if err != nil {
		return fmt.Errorf("failed to store new pin code", err)
	}

	return nil
}

func (s *PinCodeStorage) FetchPinHash(ctx context.Context, sessionID int) (string, error) {
	const query = "SELECT hash FROM pin_code WHERE session_id=?"
	row, err := s.db.QueryContext(ctx, query, sessionID)
	if err != nil {
		return "", fmt.Errorf("unable to select pin code hash from db: %v", err)
	}

	var pinHash string
	if err := row.Scan(pinHash); err != nil {
		return pinHash, fmt.Errorf("unable to scan code hash: %v", err)
	}

	return pinHash, nil
}
