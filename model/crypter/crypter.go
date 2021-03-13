package crypter

import (
	"context"
	"crypto/aes"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type Crypter struct {
	cryptRepo CryptoStorer
	pinRepo   PinCodeFetcher
}

type EncryptFile struct {
	Name   string
	Data   []byte
	UserID uuid.UUID
}

func (c *Crypter) RegisterFile(ctx context.Context, userID uuid.UUID, fileName string, rawData []byte) error {
	if len(rawData) == 0 {
		return fmt.Errorf("empty encrypted data")
	}

	pin, err := c.pinRepo.FetchPinHash(ctx, userID)
	if err != nil {
		return fmt.Errorf("unable to fetch pin code: %v", err)
	}

	cipher, err := aes.NewCipher(pin)
	if err != nil {
		return fmt.Errorf("unable to create AES cipher: %v", err)
	}

	encryptFile := EncryptFile{
		Name:   fileName,
		Data:   make([]byte, len(rawData)),
		UserID: userID,
	}
	cipher.Encrypt(encryptFile.Data, rawData)

	if err := c.cryptRepo.StoreEncryptData(ctx, &encryptFile); err != nil {
		return fmt.Errorf("unable to store encrypt data: %v", err)
	}

	return nil
}

type CryptoStorer interface {
	StoreEncryptData(ctx context.Context, file *EncryptFile) error
}

type PinCodeFetcher interface {
	FetchPinHash(ctx context.Context, sessionID uuid.UUID) ([]byte, error)
}
