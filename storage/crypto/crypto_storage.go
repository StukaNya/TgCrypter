package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/StukaNya/TgCrypter/model/crypter"
)

var (
	_ crypter.CryptoStorer = (*CryptoStorage)(nil)
)

type CryptoStorage struct {
	db *sql.DB
}

func NewCryptoStorage(db *sql.DB) *CryptoStorage {
	return &CryptoStorage{
		db: db,
	}
}

func (s *CryptoStorage) InitTable(ctx context.Context) error {
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

func (s *CryptoStorage) FetchEncryptData(ctx context.Context, fileID int) (*crypter.EncryptFile, error) {
	const query = "SELECT file_name, user_id, data FROM crypto WHERE id =?"
	row, err := s.db.QueryContext(ctx, query, fileID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch new crypto file: %v", err)
	}
	defer row.Close()

	file := new(crypter.EncryptFile)
	err = row.Scan(file.Name, file.UserID, file.Data)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (s *CryptoStorage) StoreEncryptData(ctx context.Context, file *crypter.EncryptFile) error {
	const query = "INSERT INTO crypto (user_id, file_name, data) VALUES ($1, $2)"
	_, err := s.db.ExecContext(ctx, query, file.UserID, file.Name, file.Data)
	if err != nil {
		return fmt.Errorf("failed to store new crypto file: %v", err)
	}

	return nil
}
