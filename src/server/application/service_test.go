package application_test

import (
	"context"
	"testing"

	"easy-login/server/application"
	"easy-login/server/domain"
	"easy-login/server/infrastructure/memory"
)

type fixedPlayerIDGenerator struct {
	value string
}

func (g fixedPlayerIDGenerator) NewPlayerID(context.Context) (string, error) {
	return g.value, nil
}

type fixedDeviceTokenGenerator struct {
	value string
}

func (g fixedDeviceTokenGenerator) NewDeviceToken(context.Context) (string, error) {
	return g.value, nil
}

func TestCreateGuestIdentityReturnsStableIdentityAndDeviceToken(t *testing.T) {
	service := application.NewService(
		memory.NewPlayerRepository(),
		memory.NewDeviceRegistrationRepository(),
		fixedPlayerIDGenerator{value: "player-001"},
		fixedDeviceTokenGenerator{value: "device-001"},
	)

	identity, err := service.CreateGuestIdentity(context.Background(), application.CreateGuestIdentityInput{
		DisplayName: " henrique ",
	})
	if err != nil {
		t.Fatalf("expected guest identity creation to succeed: %v", err)
	}

	if identity.PlayerID != "player-001" {
		t.Fatalf("expected stable player id, got %q", identity.PlayerID)
	}

	if identity.DeviceToken != "device-001" {
		t.Fatalf("expected device token, got %q", identity.DeviceToken)
	}

	if identity.DisplayName != "henrique" {
		t.Fatalf("expected trimmed display name, got %q", identity.DisplayName)
	}

	if identity.ClaimStatus != domain.ClaimStatusGuest {
		t.Fatalf("expected guest claim status, got %q", identity.ClaimStatus)
	}
}

func TestCreateGuestIdentityRejectsEmptyDisplayName(t *testing.T) {
	service := application.NewService(
		memory.NewPlayerRepository(),
		memory.NewDeviceRegistrationRepository(),
		fixedPlayerIDGenerator{value: "player-001"},
		fixedDeviceTokenGenerator{value: "device-001"},
	)

	_, err := service.CreateGuestIdentity(context.Background(), application.CreateGuestIdentityInput{
		DisplayName: "   ",
	})
	if err == nil {
		t.Fatal("expected empty display name to be rejected")
	}
}

func TestResumeIdentityFromDeviceTokenReturnsOriginalPlayer(t *testing.T) {
	service := application.NewService(
		memory.NewPlayerRepository(),
		memory.NewDeviceRegistrationRepository(),
		fixedPlayerIDGenerator{value: "player-001"},
		fixedDeviceTokenGenerator{value: "device-001"},
	)

	created, err := service.CreateGuestIdentity(context.Background(), application.CreateGuestIdentityInput{
		DisplayName: "henrique",
	})
	if err != nil {
		t.Fatalf("expected guest identity creation to succeed: %v", err)
	}

	resumed, err := service.ResumeIdentityFromDeviceToken(context.Background(), created.DeviceToken)
	if err != nil {
		t.Fatalf("expected resume to succeed: %v", err)
	}

	if resumed.PlayerID != created.PlayerID {
		t.Fatalf("expected same player id on resume, got %q", resumed.PlayerID)
	}

	if resumed.DisplayName != "henrique" {
		t.Fatalf("expected display name to survive resume, got %q", resumed.DisplayName)
	}

	if resumed.DeviceToken != "" {
		t.Fatalf("expected resume response not to issue a new device token, got %q", resumed.DeviceToken)
	}
}

func TestResumeIdentityFromUnknownDeviceTokenReturnsNotFound(t *testing.T) {
	service := application.NewService(
		memory.NewPlayerRepository(),
		memory.NewDeviceRegistrationRepository(),
		fixedPlayerIDGenerator{value: "player-001"},
		fixedDeviceTokenGenerator{value: "device-001"},
	)

	_, err := service.ResumeIdentityFromDeviceToken(context.Background(), "missing-device-token")
	if err == nil {
		t.Fatal("expected missing device token to fail")
	}
}
