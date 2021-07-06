package crypter

import (
	"context"
	"crypto/aes"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type EncryptedFile struct {
	ID     uuid.UUID
	Name   string
	Data   []byte
	UserID uuid.UUID
}

type Crypter struct {
	cryptRepo CryptoStorer
	pinRepo   PinCodeFetcher
}

func NewCrypter(crypt CryptoStorer, pin PinCodeFetcher) *Crypter {
	return &Crypter{
		cryptRepo: crypt,
		pinRepo:   pin,
	}
}

func (c *Crypter) EncryptFile(ctx context.Context, file *EncryptedFile) error {
	if len(file.Data) == 0 {
		return fmt.Errorf("empty encrypted data")
	}

	pin, err := c.pinRepo.FetchPinHash(ctx, file.UserID)
	if err != nil {
		return fmt.Errorf("unable to fetch pin code: %v", err)
	}

	cipher, err := aes.NewCipher(pin)
	if err != nil {
		return fmt.Errorf("unable to create AES cipher: %v", err)
	}
	cipher.Encrypt(file.Data, file.Data)

	if err := c.cryptRepo.StoreEncryptData(ctx, file); err != nil {
		return fmt.Errorf("unable to store encrypt data: %v", err)
	}

	return nil
}

type CryptoStorer interface {
	StoreEncryptData(ctx context.Context, file *EncryptedFile) error
	FetchEncryptData(ctx context.Context, fileID uuid.UUID) (*EncryptedFile, error)
}

type PinCodeFetcher interface {
	FetchPinHash(ctx context.Context, userID uuid.UUID) ([]byte, error)
}
