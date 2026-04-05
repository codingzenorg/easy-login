package domain

import (
	"errors"
	"strings"
)

const ClaimStatusGuest = "guest"

var ErrEmptyDisplayName = errors.New("display name is required")

type Player struct {
	PlayerID    string
	DisplayName string
	ClaimStatus string
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
