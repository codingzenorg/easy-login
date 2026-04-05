package system

import (
	"context"
	"crypto/rand"
	"encoding/hex"
)

type RandomHexGenerator struct{}

func (RandomHexGenerator) NewPlayerID(_ context.Context) (string, error) {
	return newHex(16)
}

func (RandomHexGenerator) NewDeviceToken(_ context.Context) (string, error) {
	return newHex(24)
}

func newHex(byteLength int) (string, error) {
	buffer := make([]byte, byteLength)
	if _, err := rand.Read(buffer); err != nil {
		return "", err
	}

	return hex.EncodeToString(buffer), nil
}
