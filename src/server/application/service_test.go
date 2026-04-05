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

func TestClaimGuestIdentityChangesClaimStatusToClaimed(t *testing.T) {
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

	claimed, err := service.ClaimGuestIdentity(context.Background(), application.ClaimGuestIdentityInput{
		DeviceToken:        created.DeviceToken,
		RecoveryPassphrase: "moon-river-42",
	})
	if err != nil {
		t.Fatalf("expected claim to succeed: %v", err)
	}

	if claimed.PlayerID != created.PlayerID {
		t.Fatalf("expected claim to preserve player id, got %q", claimed.PlayerID)
	}

	if claimed.ClaimStatus != domain.ClaimStatusClaimed {
		t.Fatalf("expected claimed status, got %q", claimed.ClaimStatus)
	}
}

func TestClaimGuestIdentityRejectsEmptyPassphrase(t *testing.T) {
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

	_, err = service.ClaimGuestIdentity(context.Background(), application.ClaimGuestIdentityInput{
		DeviceToken:        created.DeviceToken,
		RecoveryPassphrase: "   ",
	})
	if err == nil {
		t.Fatal("expected empty passphrase to be rejected")
	}
}

func TestClaimGuestIdentityRejectsAlreadyClaimedIdentity(t *testing.T) {
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

	_, err = service.ClaimGuestIdentity(context.Background(), application.ClaimGuestIdentityInput{
		DeviceToken:        created.DeviceToken,
		RecoveryPassphrase: "moon-river-42",
	})
	if err != nil {
		t.Fatalf("expected first claim to succeed: %v", err)
	}

	_, err = service.ClaimGuestIdentity(context.Background(), application.ClaimGuestIdentityInput{
		DeviceToken:        created.DeviceToken,
		RecoveryPassphrase: "other-passphrase",
	})
	if err == nil {
		t.Fatal("expected second claim to fail")
	}
}
