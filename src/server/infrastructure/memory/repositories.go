package memory

import (
	"context"
	"errors"

	"easy-login/server/domain"
)

var ErrNotFound = errors.New("not found")

type PlayerRepository struct {
	players map[string]domain.Player
}

func NewPlayerRepository() *PlayerRepository {
	return &PlayerRepository{players: map[string]domain.Player{}}
}

func (r *PlayerRepository) Save(_ context.Context, player domain.Player) error {
	r.players[player.PlayerID] = player
	return nil
}

func (r *PlayerRepository) GetByID(_ context.Context, playerID string) (domain.Player, error) {
	player, ok := r.players[playerID]
	if !ok {
		return domain.Player{}, ErrNotFound
	}

	return player, nil
}

func (r *PlayerRepository) GetByRecoveryPassphraseHash(_ context.Context, recoveryPassphraseHash string) (domain.Player, error) {
	for _, player := range r.players {
		if player.RecoveryPassphraseHash == recoveryPassphraseHash && player.ClaimStatus == domain.ClaimStatusClaimed {
			return player, nil
		}
	}

	return domain.Player{}, ErrNotFound
}

type DeviceRegistrationRepository struct {
	registrations map[string]domain.DeviceRegistration
}

func NewDeviceRegistrationRepository() *DeviceRegistrationRepository {
	return &DeviceRegistrationRepository{registrations: map[string]domain.DeviceRegistration{}}
}

func (r *DeviceRegistrationRepository) Save(_ context.Context, registration domain.DeviceRegistration) error {
	r.registrations[registration.DeviceToken] = registration
	return nil
}

func (r *DeviceRegistrationRepository) SaveRegistration(_ context.Context, registration domain.DeviceRegistration) error {
	r.registrations[registration.DeviceToken] = registration
	return nil
}

func (r *DeviceRegistrationRepository) GetByDeviceToken(_ context.Context, deviceToken string) (domain.DeviceRegistration, error) {
	registration, ok := r.registrations[deviceToken]
	if !ok {
		return domain.DeviceRegistration{}, ErrNotFound
	}

	return registration, nil
}
