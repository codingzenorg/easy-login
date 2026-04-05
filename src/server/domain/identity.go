package domain

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"strings"
)

const ClaimStatusGuest = "guest"
const ClaimStatusClaimed = "claimed"

var ErrEmptyDisplayName = errors.New("display name is required")
var ErrInvalidRecoveryPassphrase = errors.New("recovery passphrase is required")
var ErrIdentityAlreadyClaimed = errors.New("identity is already claimed")

type Player struct {
	PlayerID               string
	DisplayName            string
	ClaimStatus            string
	RecoveryPassphraseHash string
}

type DeviceRegistration struct {
	DeviceToken string
	PlayerID    string
}

func NewGuestPlayer(playerID string, displayName string) (Player, error) {
	normalized := strings.TrimSpace(displayName)
	if normalized == "" {
		return Player{}, ErrEmptyDisplayName
	}

	return Player{
		PlayerID:    playerID,
		DisplayName: normalized,
		ClaimStatus: ClaimStatusGuest,
	}, nil
}

func (p Player) Claim(recoveryPassphrase string) (Player, error) {
	if p.ClaimStatus == ClaimStatusClaimed {
		return Player{}, ErrIdentityAlreadyClaimed
	}

	normalized := strings.TrimSpace(recoveryPassphrase)
	if normalized == "" {
		return Player{}, ErrInvalidRecoveryPassphrase
	}

	p.ClaimStatus = ClaimStatusClaimed
	p.RecoveryPassphraseHash = hashRecoveryPassphrase(normalized)
	return p, nil
}

func hashRecoveryPassphrase(passphrase string) string {
	sum := sha256.Sum256([]byte(passphrase))
	return hex.EncodeToString(sum[:])
}
