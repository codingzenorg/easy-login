# Durable Guest Identity Persistence

## Selected Pack

`polyglot_client_server`

Current runtime direction remains:

- Mithril browser client
- Go backend service

This slice changes persistence behavior on the backend only. It does not introduce new client-side business behavior.

## Runtime Targets

- backend HTTP service
- SQLite persistence boundary

The browser client remains in scope only as the existing caller of the backend contracts. It should not need semantic changes to benefit from this slice.

## Architecture Mode

Infrastructure hardening in service of the existing business promise.

The slice exists to make the already-defined guest continuity behavior durable across backend process restarts by replacing the in-memory repositories with SQLite-backed repositories.

This is still a business slice, not generic infrastructure work, because the model promise is persistent identity rather than session-only identity.

## Discovery Scope

This slice covers:

- persistent storage for players
- persistent storage for device-token registrations
- backend startup with a configured SQLite database
- preserving the current guest-creation and resume contracts

It intentionally excludes:

- claim flow
- recovery passphrase flow
- display-name uniqueness beyond the current behavior
- frontend UX changes beyond any config needed to keep the client talking to the server

## Model Pressure

The first built slice proved identity continuity only while the server process stays alive.

That leaves a model gap:

- the domain says guest identity should be persistent on the same browser
- the implementation currently loses identity continuity when backend memory is lost

This slice should close that gap before ownership and recovery work continue.

## Use-Case Contract

The use cases remain the same:

### `CreateGuestIdentity`

#### Input

- `display_name`

#### Output

- `player_id`
- `display_name`
- `device_token`
- `claim_status`

#### Additional persistence guarantee

- after successful creation, the player and device registration survive backend restart when the same SQLite database is reused

### `ResumeIdentityFromDeviceToken`

#### Input

- `device_token`

#### Output on success

- `player_id`
- `display_name`
- `claim_status`

#### Additional persistence guarantee

- a device token created before a backend restart still resumes the same player identity after restart when the same SQLite database is reused

## Main Business Rules

- this slice must not change guest-identity semantics
- persistence durability is part of the user-facing continuity promise
- a successful guest creation must persist both the player and the device registration atomically enough that resume does not observe a half-created identity in normal operation
- an unknown device token still returns explicit not-found
- health and readiness endpoints stay available

## Required Ports And Boundaries

### Existing ports kept

- `PlayerRepository`
- `DeviceRegistrationRepository`
- `IdGenerator`
- `DeviceTokenGenerator`

### New or clarified infrastructure boundaries

- `SQLiteConnectionFactory` or equivalent database-opening boundary
- migration or schema initialization path for the required tables

### Expected tables

- `players`
- `device_registrations`

Exact schema names may vary if build documents the reason clearly, but the tables should preserve one player record and one device-token binding record per current model needs.

## Configuration Expectations

- backend should accept a database location through explicit configuration, such as `SQLITE_PATH`
- a stable local default is acceptable for manual runs, but tests must control the database location explicitly

## Initial Test Plan

### Integration tests

- creating a guest identity persists player and device registration in SQLite
- resuming with a known device token works when a new repository instance is created against the same SQLite database
- resuming with an unknown device token still returns explicit not-found
- schema initialization is idempotent enough for repeated local test runs

### Regression tests that must keep passing

- empty display name remains rejected
- health and readiness endpoints remain successful
- CORS behavior remains correct for the browser client

## Scenario Definition

1. backend starts with a configured SQLite file
2. user creates guest identity with display name `henrique`
3. backend persists player and device-token registration
4. backend process stops
5. backend process starts again using the same SQLite file
6. browser sends the previously stored device token
7. backend returns the same `player_id` and display name

## Done Criteria

- SQLite-backed repositories replace or sit alongside the in-memory ones as the default executable path
- the backend can restart and still resume a previously created guest identity from the same database
- integration tests prove durable continuity across repository or process recreation
- the slice does not expand into claim or recovery behavior
