package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"

	"easy-login/server/domain"
)

var ErrNotFound = errors.New("not found")

type Store struct {
	db *sql.DB
}

func OpenDatabase(path string) (*sql.DB, error) {
	if path == "" {
		path = "data/easy-login.db"
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, fmt.Errorf("create sqlite directory: %w", err)
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("open sqlite database: %w", err)
	}

	return db, nil
}

func NewStore(db *sql.DB) (*Store, error) {
	store := &Store{db: db}
	if err := store.Initialize(context.Background()); err != nil {
		return nil, err
	}
	return store, nil
}

func (s *Store) Initialize(ctx context.Context) error {
	statements := []string{
		`PRAGMA foreign_keys = ON;`,
		`CREATE TABLE IF NOT EXISTS players (
			player_id TEXT PRIMARY KEY,
			display_name TEXT NOT NULL,
			claim_status TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS device_registrations (
			device_token TEXT PRIMARY KEY,
			player_id TEXT NOT NULL,
			FOREIGN KEY(player_id) REFERENCES players(player_id)
		);`,
	}

	for _, statement := range statements {
		if _, err := s.db.ExecContext(ctx, statement); err != nil {
			return fmt.Errorf("initialize sqlite schema: %w", err)
		}
	}

	return nil
}

func (s *Store) Save(ctx context.Context, player domain.Player) error {
	_, err := s.db.ExecContext(
		ctx,
		`INSERT OR REPLACE INTO players (player_id, display_name, claim_status)
		 VALUES (?, ?, ?)`,
		player.PlayerID,
		player.DisplayName,
		player.ClaimStatus,
	)
	return err
}

func (s *Store) SaveGuestIdentity(ctx context.Context, player domain.Player, registration domain.DeviceRegistration) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	if _, err = tx.ExecContext(
		ctx,
		`INSERT OR REPLACE INTO players (player_id, display_name, claim_status)
		 VALUES (?, ?, ?)`,
		player.PlayerID,
		player.DisplayName,
		player.ClaimStatus,
	); err != nil {
		return err
	}

	if _, err = tx.ExecContext(
		ctx,
		`INSERT OR REPLACE INTO device_registrations (device_token, player_id)
		 VALUES (?, ?)`,
		registration.DeviceToken,
		registration.PlayerID,
	); err != nil {
		return err
	}

	return tx.Commit()
}

func (s *Store) GetByID(ctx context.Context, playerID string) (domain.Player, error) {
	row := s.db.QueryRowContext(
		ctx,
		`SELECT player_id, display_name, claim_status
		 FROM players
		 WHERE player_id = ?`,
		playerID,
	)

	var player domain.Player
	if err := row.Scan(&player.PlayerID, &player.DisplayName, &player.ClaimStatus); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Player{}, ErrNotFound
		}
		return domain.Player{}, err
	}

	return player, nil
}

func (s *Store) SaveDeviceRegistration(ctx context.Context, registration domain.DeviceRegistration) error {
	_, err := s.db.ExecContext(
		ctx,
		`INSERT OR REPLACE INTO device_registrations (device_token, player_id)
		 VALUES (?, ?)`,
		registration.DeviceToken,
		registration.PlayerID,
	)
	return err
}

func (s *Store) SaveRegistration(ctx context.Context, registration domain.DeviceRegistration) error {
	return s.SaveDeviceRegistration(ctx, registration)
}

func (s *Store) GetByDeviceToken(ctx context.Context, deviceToken string) (domain.DeviceRegistration, error) {
	row := s.db.QueryRowContext(
		ctx,
		`SELECT device_token, player_id
		 FROM device_registrations
		 WHERE device_token = ?`,
		deviceToken,
	)

	var registration domain.DeviceRegistration
	if err := row.Scan(&registration.DeviceToken, &registration.PlayerID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.DeviceRegistration{}, ErrNotFound
		}
		return domain.DeviceRegistration{}, err
	}

	return registration, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}
