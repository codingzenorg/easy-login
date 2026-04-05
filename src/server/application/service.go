package application

import (
	"context"
	"errors"

	"easy-login/server/domain"
)

var ErrIdentityNotFound = errors.New("identity not found")

type PlayerRepository interface {
	Save(ctx context.Context, player domain.Player) error
	GetByID(ctx context.Context, playerID string) (domain.Player, error)
}

type DeviceRegistrationRepository interface {
	SaveRegistration(ctx context.Context, registration domain.DeviceRegistration) error
	GetByDeviceToken(ctx context.Context, deviceToken string) (domain.DeviceRegistration, error)
}

type GuestIdentityPersistence interface {
	SaveGuestIdentity(ctx context.Context, player domain.Player, registration domain.DeviceRegistration) error
}

type IDGenerator interface {
	NewPlayerID(ctx context.Context) (string, error)
}

type DeviceTokenGenerator interface {
	NewDeviceToken(ctx context.Context) (string, error)
}

type Service struct {
	players      PlayerRepository
	devices      DeviceRegistrationRepository
	idGenerator  IDGenerator
	tokenFactory DeviceTokenGenerator
}

type CreateGuestIdentityInput struct {
	DisplayName string
}

type IdentityView struct {
	PlayerID    string `json:"player_id"`
	DisplayName string `json:"display_name"`
	ClaimStatus string `json:"claim_status"`
	DeviceToken string `json:"device_token,omitempty"`
}

func NewService(
	players PlayerRepository,
	devices DeviceRegistrationRepository,
	idGenerator IDGenerator,
	tokenFactory DeviceTokenGenerator,
) Service {
	return Service{
		players:      players,
		devices:      devices,
		idGenerator:  idGenerator,
		tokenFactory: tokenFactory,
	}
}

func (s Service) CreateGuestIdentity(ctx context.Context, input CreateGuestIdentityInput) (IdentityView, error) {
	playerID, err := s.idGenerator.NewPlayerID(ctx)
	if err != nil {
		return IdentityView{}, err
	}

	player, err := domain.NewGuestPlayer(playerID, input.DisplayName)
	if err != nil {
		return IdentityView{}, err
	}

	deviceToken, err := s.tokenFactory.NewDeviceToken(ctx)
	if err != nil {
		return IdentityView{}, err
	}

	registration := domain.DeviceRegistration{
		DeviceToken: deviceToken,
		PlayerID:    player.PlayerID,
	}

	if persistence, ok := s.players.(GuestIdentityPersistence); ok {
		if err := persistence.SaveGuestIdentity(ctx, player, registration); err != nil {
			return IdentityView{}, err
		}
	} else {
		if err := s.players.Save(ctx, player); err != nil {
			return IdentityView{}, err
		}

		if err := s.devices.SaveRegistration(ctx, registration); err != nil {
			return IdentityView{}, err
		}
	}

	return IdentityView{
		PlayerID:    player.PlayerID,
		DisplayName: player.DisplayName,
		ClaimStatus: player.ClaimStatus,
		DeviceToken: deviceToken,
	}, nil
}

func (s Service) ResumeIdentityFromDeviceToken(ctx context.Context, deviceToken string) (IdentityView, error) {
	registration, err := s.devices.GetByDeviceToken(ctx, deviceToken)
	if err != nil {
		return IdentityView{}, ErrIdentityNotFound
	}

	player, err := s.players.GetByID(ctx, registration.PlayerID)
	if err != nil {
		return IdentityView{}, ErrIdentityNotFound
	}

	return IdentityView{
		PlayerID:    player.PlayerID,
		DisplayName: player.DisplayName,
		ClaimStatus: player.ClaimStatus,
	}, nil
}
