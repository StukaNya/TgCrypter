package crypter

import (
	"context"
	"crypto/aes"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type Encrypter struct {
	log     *logrus.Logger
	crypt   CryptoStorer
	pinCode PinCodeFetcher
}

func (c *Encrypter) RegisterData(ctx context.Context, sessionID uuid.UUID, rawData []byte) error {
	if len(rawData) == 0 {
		return fmt.Errorf("empty encrypted data")
	}

	pin, err := c.pinCode.FetchPinHash(ctx, sessionID)
	if err != nil {
		return fmt.Errorf("unable to fetch pin code: %v", err)
	}

	cipher, err := aes.NewCipher(pin)
	if err != nil {
		return fmt.Errorf("unable to create AES cipher: %v", err)
	}

	var cryptData = make([]byte, len(rawData))
	cipher.Encrypt(cryptData, rawData)
	c.log.Info("Encrypt new data: ", cryptData)

	if err := c.crypt.StoreEncryptData(ctx, sessionID, cryptData); err != nil {
		return fmt.Errorf("unable to store encrypt data: %v", err)
	}

	return nil
}

type CryptoStorer interface {
	StoreEncryptData(ctx context.Context, sessionID uuid.UUID, data []byte) error
}

type PinCodeFetcher interface {
	FetchPinHash(ctx context.Context, sessionID uuid.UUID) ([]byte, error)
}
