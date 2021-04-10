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
	userData map[uuid.UUID]*crypter.EncryptFile
}

func NewStoreMock(size int) *CryptoStoreMock {
	return &CryptoStoreMock{userData: make(map[uuid.UUID]*crypter.EncryptFile)}
}

func (s *CryptoStoreMock) FetchEncryptData(ctx context.Context, fileID uuid.UUID) (*crypter.EncryptFile, error) {
	user, ok := s.userData[fileID]
	if !ok {
		return nil, fmt.Errorf("file not found")
	}
	return user, nil
}

func (s *CryptoStoreMock) StoreEncryptData(ctx context.Context, file *crypter.EncryptFile) (uuid.UUID, error) {
	fileID := uuid.NewV4()
	s.userData[fileID] = file
	return fileID, nil
}
