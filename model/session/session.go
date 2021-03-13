package session

import (
	"context"
	"crypto/sha256"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type UserInfo struct {
	Name         string
	ChatID       int64
	RegisteredAt time.Time
}

type SessionRegistry struct {
	log      *logrus.Logger
	userRepo UserRepository
	pinCode  PinCodeStorer
}

func (s *SessionRegistry) RegisterSession(ctx context.Context, userName string, chatID int64) (uuid.UUID, error) {
	userInfo := &UserInfo{
		Name:         userName,
		ChatID:       chatID,
		RegisteredAt: time.Now(),
	}

	userID, err := s.userRepo.RegisterUser(ctx, userInfo)
	if err != nil {
		return uuid.Nil, fmt.Errorf("unable to store session info: %v", err)
	}
	return userID, nil
}

func (s *SessionRegistry) RegisterPinCode(ctx context.Context, userID uuid.UUID, pin string) error {
	shaHash := sha256.New()
	shaHash.Write([]byte(pin))
	s.log.Info("Register new pin code hash: ", shaHash.Sum(nil))

	if err := s.pinCode.StorePinHash(ctx, userID, shaHash.Sum(nil)); err != nil {
		return fmt.Errorf("unable to store pin code: %v", err)
	}
	return nil
}

type UserRepository interface {
	UserRegistrar
	UserFetcher
}

type UserRegistrar interface {
	RegisterUser(ctx context.Context, info *UserInfo) (uuid.UUID, error)
}

type UserFetcher interface {
	FetchUser(ctx context.Context, userID uuid.UUID) (*UserInfo, error)
}

type PinCodeStorer interface {
	StorePinHash(ctx context.Context, userID uuid.UUID, pin []byte) error
}
