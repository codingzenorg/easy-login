package sqlite_test

import (
	"context"
	"path/filepath"
	"testing"

	"easy-login/server/application"
	"easy-login/server/domain"
	"easy-login/server/infrastructure/sqlite"
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

func openTestStore(t *testing.T, filename string) *sqlite.Store {
	t.Helper()

	db, err := sqlite.OpenDatabase(filepath.Join(t.TempDir(), filename))
	if err != nil {
		t.Fatalf("open sqlite database: %v", err)
	}

	store, err := sqlite.NewStore(db)
	if err != nil {
		t.Fatalf("create sqlite store: %v", err)
	}

	t.Cleanup(func() {
		_ = store.Close()
	})

	return store
}

func TestStorePersistsGuestIdentityAcrossStoreRecreation(t *testing.T) {
	databasePath := filepath.Join(t.TempDir(), "easy-login.db")

	db, err := sqlite.OpenDatabase(databasePath)
	if err != nil {
		t.Fatalf("open sqlite database: %v", err)
	}

	store, err := sqlite.NewStore(db)
	if err != nil {
		t.Fatalf("create sqlite store: %v", err)
	}

	service := application.NewService(
		store,
		store,
		fixedPlayerIDGenerator{value: "player-001"},
		fixedDeviceTokenGenerator{value: "device-001"},
	)

	created, err := service.CreateGuestIdentity(context.Background(), application.CreateGuestIdentityInput{
		DisplayName: "henrique",
	})
	if err != nil {
		t.Fatalf("create guest identity: %v", err)
	}

	if err := store.Close(); err != nil {
		t.Fatalf("close sqlite store: %v", err)
	}

	db, err = sqlite.OpenDatabase(databasePath)
	if err != nil {
		t.Fatalf("reopen sqlite database: %v", err)
	}

	reopenedStore, err := sqlite.NewStore(db)
	if err != nil {
		t.Fatalf("recreate sqlite store: %v", err)
	}
	defer reopenedStore.Close()

	resumeService := application.NewService(
		reopenedStore,
		reopenedStore,
		fixedPlayerIDGenerator{value: "unused"},
		fixedDeviceTokenGenerator{value: "unused"},
	)

	resumed, err := resumeService.ResumeIdentityFromDeviceToken(context.Background(), created.DeviceToken)
	if err != nil {
		t.Fatalf("resume guest identity: %v", err)
	}

	if resumed.PlayerID != created.PlayerID {
		t.Fatalf("expected player id %q after restart, got %q", created.PlayerID, resumed.PlayerID)
	}
}

func TestStoreReturnsNotFoundForUnknownDeviceToken(t *testing.T) {
	store := openTestStore(t, "missing.db")

	_, err := store.GetByDeviceToken(context.Background(), "missing-device-token")
	if err == nil {
		t.Fatal("expected missing device token to fail")
	}
}

func TestSchemaInitializationIsIdempotent(t *testing.T) {
	store := openTestStore(t, "schema.db")

	if err := store.Initialize(context.Background()); err != nil {
		t.Fatalf("expected repeated initialization to succeed: %v", err)
	}
}

func TestStoreReadsPersistedPlayerByID(t *testing.T) {
	store := openTestStore(t, "players.db")

	player := domain.Player{
		PlayerID:    "player-001",
		DisplayName: "henrique",
		ClaimStatus: domain.ClaimStatusGuest,
	}

	if err := store.Save(context.Background(), player); err != nil {
		t.Fatalf("save player: %v", err)
	}

	saved, err := store.GetByID(context.Background(), "player-001")
	if err != nil {
		t.Fatalf("get player by id: %v", err)
	}

	if saved.DisplayName != "henrique" {
		t.Fatalf("expected persisted display name, got %q", saved.DisplayName)
	}
}

func TestClaimedStatusPersistsAcrossStoreRecreation(t *testing.T) {
	databasePath := filepath.Join(t.TempDir(), "claimed.db")

	db, err := sqlite.OpenDatabase(databasePath)
	if err != nil {
		t.Fatalf("open sqlite database: %v", err)
	}

	store, err := sqlite.NewStore(db)
	if err != nil {
		t.Fatalf("create sqlite store: %v", err)
	}

	service := application.NewService(
		store,
		store,
		fixedPlayerIDGenerator{value: "player-001"},
		fixedDeviceTokenGenerator{value: "device-001"},
	)

	created, err := service.CreateGuestIdentity(context.Background(), application.CreateGuestIdentityInput{
		DisplayName: "henrique",
	})
	if err != nil {
		t.Fatalf("create guest identity: %v", err)
	}

	_, err = service.ClaimGuestIdentity(context.Background(), application.ClaimGuestIdentityInput{
		DeviceToken:        created.DeviceToken,
		RecoveryPassphrase: "moon-river-42",
	})
	if err != nil {
		t.Fatalf("claim guest identity: %v", err)
	}

	if err := store.Close(); err != nil {
		t.Fatalf("close sqlite store: %v", err)
	}

	db, err = sqlite.OpenDatabase(databasePath)
	if err != nil {
		t.Fatalf("reopen sqlite database: %v", err)
	}

	reopenedStore, err := sqlite.NewStore(db)
	if err != nil {
		t.Fatalf("recreate sqlite store: %v", err)
	}
	defer reopenedStore.Close()

	resumeService := application.NewService(
		reopenedStore,
		reopenedStore,
		fixedPlayerIDGenerator{value: "unused"},
		fixedDeviceTokenGenerator{value: "unused"},
	)

	resumed, err := resumeService.ResumeIdentityFromDeviceToken(context.Background(), created.DeviceToken)
	if err != nil {
		t.Fatalf("resume claimed identity: %v", err)
	}

	if resumed.ClaimStatus != domain.ClaimStatusClaimed {
		t.Fatalf("expected claimed status after restart, got %q", resumed.ClaimStatus)
	}
}

func TestRecoveryPersistsDeviceRegistrationUsableByResume(t *testing.T) {
	databasePath := filepath.Join(t.TempDir(), "recovery.db")

	db, err := sqlite.OpenDatabase(databasePath)
	if err != nil {
		t.Fatalf("open sqlite database: %v", err)
	}

	store, err := sqlite.NewStore(db)
	if err != nil {
		t.Fatalf("create sqlite store: %v", err)
	}

	service := application.NewService(
		store,
		store,
		fixedPlayerIDGenerator{value: "player-001"},
		fixedDeviceTokenGenerator{value: "device-001"},
	)

	created, err := service.CreateGuestIdentity(context.Background(), application.CreateGuestIdentityInput{
		DisplayName: "henrique",
	})
	if err != nil {
		t.Fatalf("create guest identity: %v", err)
	}

	_, err = service.ClaimGuestIdentity(context.Background(), application.ClaimGuestIdentityInput{
		DeviceToken:        created.DeviceToken,
		RecoveryPassphrase: "moon-river-42",
	})
	if err != nil {
		t.Fatalf("claim guest identity: %v", err)
	}

	recovered, err := service.RecoverClaimedIdentity(context.Background(), application.RecoverClaimedIdentityInput{
		RecoveryPassphrase: "moon-river-42",
	})
	if err != nil {
		t.Fatalf("recover claimed identity: %v", err)
	}

	resumed, err := service.ResumeIdentityFromDeviceToken(context.Background(), recovered.DeviceToken)
	if err != nil {
		t.Fatalf("resume recovered identity: %v", err)
	}

	if resumed.PlayerID != created.PlayerID {
		t.Fatalf("expected resumed recovery to preserve player id, got %q", resumed.PlayerID)
	}
}

func TestUnclaimedGuestIdentityCannotBeRecovered(t *testing.T) {
	store := openTestStore(t, "unclaimed.db")

	service := application.NewService(
		store,
		store,
		fixedPlayerIDGenerator{value: "player-001"},
		fixedDeviceTokenGenerator{value: "device-001"},
	)

	_, err := service.CreateGuestIdentity(context.Background(), application.CreateGuestIdentityInput{
		DisplayName: "henrique",
	})
	if err != nil {
		t.Fatalf("create guest identity: %v", err)
	}

	_, err = service.RecoverClaimedIdentity(context.Background(), application.RecoverClaimedIdentityInput{
		RecoveryPassphrase: "moon-river-42",
	})
	if err == nil {
		t.Fatal("expected unclaimed identity recovery to fail")
	}
}
