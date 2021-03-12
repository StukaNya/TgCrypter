package store

import (
	"context"
	"database/sql"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type CryptoStorer struct {
	db *sql.DB
}

type CryptoFile struct {
	name string
	data []byte
}

func (s *CryptoStorer) InitTable(ctx context.Context) error {
	const query = `CREATE TABLE crypto (
			id 			int 	PRIMARY KEY DEFAULT uuid_generate_v4(),
			user_id		uuid	NOT NULL REFERENCES user(id),
			file_name 	text,
			data		bytea,
			encrypt_at 	time	NOT NULL DEFAULT(now() at time zone 'utc')
		);`
	_, err := s.db.ExecContext(ctx, query)
	if err != nil {
		return fmt.Errorf("unable to create crypto storer db: %v", err)
	}

	return nil
}

func (s *CryptoStorer) FilterEncryptData(ctx context.Context, userID int) (*CryptoFile, error) {
	const query = "SELECT file_name, data FROM crypto WHERE user_id =?"
	row, err := s.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch new crypto file: %v", err)
	}
	defer row.Close()

	//TODO: filter multiply files
	file := new(CryptoFile)
	err = row.Scan(file.name, file.data)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (s *CryptoStorer) StoreEncryptData(ctx context.Context, userID uuid.UUID, data []byte) error {
	const query = "INSERT INTO crypto (user_id, data) VALUES ($1, $2)"
	_, err := s.db.ExecContext(ctx, query, userID, data)
	if err != nil {
		return fmt.Errorf("failed to store new crypto file: %v", err)
	}

	return nil
}
