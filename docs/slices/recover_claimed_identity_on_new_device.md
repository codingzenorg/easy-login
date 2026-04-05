# Recover Claimed Identity On New Device

## Selected Pack

`polyglot_client_server`

Current runtime direction remains:

- Mithril browser client
- Go backend service

This slice adds cross-device recovery to the existing guest-plus-claim model.

## Runtime Targets

- browser client
- backend HTTP service
- SQLite persistence boundary

## Architecture Mode

Thin vertical extension focused on ownership proof becoming operational.

The browser is responsible for:

- presenting a recovery path when the current browser has no usable device token
- submitting a recovery passphrase
- storing the replacement device token returned after successful recovery

The backend is responsible for:

- locating the claimed identity by recovery proof
- issuing continuity for the current browser through a device token
- preserving stable `player_id` while moving to the new browser context

This slice should not redesign the claim model. It should make the already-claimed state useful across devices.

## Discovery Scope

This slice covers:

- recovering a claimed identity on another browser or after local storage loss
- issuing a valid device token for the recovering browser
- returning the recovered identity with claimed status

It intentionally excludes:

- passphrase rotation
- revoking old devices
- multi-device session management policies beyond what is necessary to recover
- email, OAuth, or stronger account systems

## Model Pressure

The current system now has:

- guest continuity on the same browser
- ownership claim through a recovery passphrase

But the main promise is still incomplete:

- ownership proof exists
- recovery across devices does not

This slice should close that gap so claimed identity becomes portable rather than only annotated.

## Use-Case Contract

### `RecoverClaimedIdentity`

#### Input

- `recovery_passphrase`

#### Output on success

- `player_id`
- `display_name`
- `claim_status` set to `claimed`
- `device_token` for the recovering browser

#### Output on failure

- explicit not-found or unauthorized-style result when the recovery proof does not match a claimed identity
- explicit validation failure when the recovery passphrase is invalid

#### Required behavior

- find a claimed player by recovery proof
- reject recovery when no claimed identity matches
- issue a device token for the recovering browser
- keep the same stable `player_id`
- return the recovered claimed identity plus the new device token

## Main Business Rules

- only claimed identities are recoverable in this slice
- recovery proof is the passphrase established during claim
- recovery does not create a new player identity
- a successful recovery returns the same `player_id` as before
- a successful recovery yields a usable device token for subsequent same-browser resume
- recovery must not silently downgrade claim state

## Deferred Questions

- whether recovery should invalidate previous device tokens
- whether multiple active device tokens are allowed long-term
- whether recovery attempts need throttling or audit trails
- whether passphrase recovery should reveal less information on failure

These should remain deferred unless the slice document is updated first.

## Required Ports And Boundaries

### Existing ports kept

- `PlayerRepository`
- `DeviceRegistrationRepository`
- `DeviceTokenGenerator`

### New or clarified persistence responsibility

- player lookup by recovery-proof derivative
- device registration creation for a recovered browser context

### Interface boundaries

- backend HTTP endpoint for recovery
- browser UI path to recover identity when no current device token is available

## Initial API Shape

- `POST /identities/guest`
- `POST /identities/resume`
- `POST /identities/claim`
- `POST /identities/recover`
  recovers a claimed identity from a recovery passphrase and returns a device token
- `GET /healthz`
- `GET /readyz`

## Initial Test Plan

### Domain or use-case tests

- recovery with a valid claimed passphrase returns the original `player_id`
- recovery with an invalid passphrase is rejected
- recovery returns `claim_status = claimed`
- recovery issues a device token for later resume

### Integration tests

- recovery finds a claimed identity in SQLite through persisted recovery-proof data
- recovery persists a device registration usable by the existing resume flow
- resume after recovery returns the same claimed identity
- unclaimed guest identities cannot be recovered

## Scenario Definition

1. browser A creates a guest identity
2. browser A claims the identity with a recovery passphrase
3. browser A is lost or local storage is cleared
4. browser B opens the app with no usable device token
5. browser B submits the recovery passphrase
6. backend returns the same `player_id`, `display_name`, `claim_status = claimed`, and a new `device_token`
7. browser B stores the returned device token and can resume normally afterward

## Done Criteria

- the repository supports recovering a claimed identity on another browser through a recovery passphrase
- recovery preserves stable identity and claimed status
- recovery returns a device token that works with the existing resume flow
- tests prove recovery behavior against SQLite-backed persistence
- the slice does not expand into token revocation or full account management
