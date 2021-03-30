package store

import (
	"context"
	"fmt"

	"github.com/StukaNya/TgCrypter/model/session"
	uuid "github.com/satori/go.uuid"
)

var (
	_ session.UserRepository = (*UserStoreMock)(nil)
)

type UserStoreMock struct {
	userData map[uuid.UUID]*session.UserInfo
}

func NewStoreMock(size int) *UserStoreMock {
	return &UserStoreMock{userData: make(map[uuid.UUID]*session.UserInfo)}
}

func (s *UserStoreMock) FetchUser(ctx context.Context, userID uuid.UUID) (*session.UserInfo, error) {
	user, ok := s.userData[userID]
	if !ok {
		return nil, fmt.Errorf("pin code not found")
	}
	return user, nil
}

func (s *UserStoreMock) RegisterUser(ctx context.Context, info *session.UserInfo) (uuid.UUID, error) {
	userID := uuid.NewV4()
	s.userData[userID] = info
	return userID, nil
}
