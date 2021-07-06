package store

import (
	"context"
	"fmt"

	"github.com/StukaNya/TgCrypter/model/crypter"
	uuid "github.com/satori/go.uuid"
)

var (
	_ crypter.CryptoStorer = (*CryptoStoreMock)(nil)
)

type CryptoStoreMock struct {
	userData map[uuid.UUID]*crypter.EncryptedFile
}

func NewStoreMock(size int) *CryptoStoreMock {
	return &CryptoStoreMock{userData: make(map[uuid.UUID]*crypter.EncryptedFile)}
}

func (s *CryptoStoreMock) FetchEncryptData(ctx context.Context, fileID uuid.UUID) (*crypter.EncryptedFile, error) {
	user, ok := s.userData[fileID]
	if !ok {
		return nil, fmt.Errorf("file not found")
	}
	return user, nil
}

func (s *CryptoStoreMock) StoreEncryptData(ctx context.Context, file *crypter.EncryptedFile) error {
	file.ID = uuid.NewV4()
	s.userData[file.ID] = file
	return nil
}
