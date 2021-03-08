package session

import (
	"context"
	"crypto/sha256"
	"fmt"
	"hash"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type SessionInfo struct {
	ID           uuid.UUID
	UserName     string
	ChatID       int64
	RegisterTime time.Time
}

type SessionRegistry struct {
	log  *logrus.Logger
	repo SessionRepository
}

func (s *SessionRegistry) RegisterSession(ctx context.Context, userName string, chatID int64) (uuid.UUID, error) {
	sessionInfo := &SessionInfo{
		ID:           uuid.NewV4(),
		UserName:     userName,
		ChatID:       chatID,
		RegisterTime: time.Now(),
	}

	if err := s.repo.StoreSession(ctx, sessionInfo); err != nil {
		return uuid.Nil, fmt.Errorf("unable to store session info: %v", err)
	}
	return sessionInfo.ID, nil
}

func (s *SessionRegistry) RegisterPinCode(ctx context.Context, sessionID uuid.UUID, pin string) error {
	shaHash := sha256.New()
	shaHash.Write([]byte(pin))
	s.log.Info("Register new pin code hash: ", shaHash.Sum(nil))

	if err := s.repo.StorePinHash(ctx, sessionID, shaHash); err != nil {
		return fmt.Errorf("unable to store pin code: %v", err)
	}
	return nil
}

type SessionRepository interface {
	StoreSession(ctx context.Context, info *SessionInfo) error
	StorePinHash(ctx context.Context, sessionID uuid.UUID, pin hash.Hash) error
}
