package store

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/StukaNya/TgCrypter/model/crypter"
	"github.com/StukaNya/TgCrypter/model/session"
	uuid "github.com/satori/go.uuid"
)

var (
	_ session.PinCodeStorer  = (*PinStoreMock)(nil)
	_ crypter.PinCodeFetcher = (*PinStoreMock)(nil)
)

type PinMock struct {
	id   uuid.UUID
	data []byte
}

type PinStoreMock struct {
	pinData map[uuid.UUID]*PinMock
}

func NewStoreMock(size int) *PinStoreMock {
	return &PinStoreMock{pinData: make(map[uuid.UUID]*PinMock)}
}

func (s *PinStoreMock) FetchPinHash(ctx context.Context, userID uuid.UUID) ([]byte, error) {
	pin, ok := s.pinData[userID]
	if !ok {
		return nil, fmt.Errorf("pin code not found")
	}
	return pin.data, nil
}

func (s *PinStoreMock) StorePinHash(ctx context.Context, userID uuid.UUID, pinHash []byte) error {
	s.pinData[userID] = &PinMock{id: uuid.NewV4(), data: pinHash}
	return nil
}
