# Claim Guest Identity With Passphrase

## Selected Pack

`polyglot_client_server`

Current runtime direction remains:

- Mithril browser client
- Go backend service

This slice adds ownership behavior to the existing guest identity flow without changing the overall runtime split.

## Runtime Targets

- browser client
- backend HTTP service
- SQLite persistence boundary

## Architecture Mode

Thin vertical extension of the current identity model.

The browser is responsible for:

- presenting a claim action for the current guest identity
- submitting the current device token plus a recovery passphrase
- rendering the resulting claimed status

The backend is responsible for:

- resolving the current guest identity from the device token
- attaching ownership proof to that identity
- persisting the new claimed state without breaking existing continuity

Cross-device recovery remains out of scope for this slice.

## Discovery Scope

This slice covers:

- claiming an existing guest identity
- persisting a recovery passphrase or its durable equivalent
- returning claimed status for the identity after successful claim
- preventing invalid claim transitions

It intentionally excludes:

- recovering an identity from another device
- rotating or changing a recovery passphrase
- display-name reservation semantics beyond what is required by claim state
- stronger authentication channels such as email or OAuth

## Model Pressure

The current system supports continuity, but not ownership.

That leaves the next major model gap:

- the user can keep using the same identity on one browser
- but the system cannot yet distinguish a disposable guest identity from one that the user has intentionally claimed

This slice should close that gap before cross-device recovery is attempted.

## Use-Case Contract

### `ClaimGuestIdentity`

#### Input

- `device_token`
- `recovery_passphrase`

#### Output on success

- `player_id`
- `display_name`
- `claim_status` set to `claimed`

#### Output on failure

- explicit not-found when the device token does not resolve an identity
- explicit validation failure when the recovery passphrase is invalid
- explicit conflict when the identity is already claimed and the operation is not allowed

#### Required behavior

- find the current player identity from the device token
- reject claim when the player cannot be found
- reject empty or clearly invalid recovery passphrases
- persist the player as claimed
- keep the existing device-token continuity intact

## Main Business Rules

- a guest identity may be claimed exactly once in this slice
- claim changes ownership state, not persistent identity
- `player_id` remains stable before and after claim
- display name is still not ownership proof
- the recovery passphrase becomes the ownership proof for later slices
- claiming an identity must not issue a new device token
- the existing browser should continue to resume the same identity after claim

## Deferred Questions

- whether the passphrase is stored directly or as a derived secret
- exact passphrase complexity rules
- whether the user supplies the passphrase or the system can generate one
- whether duplicate display names gain new rules after claim

These should not be invented during build unless the slice document is updated first.

## Required Ports And Boundaries

### Existing ports kept

- `PlayerRepository`
- `DeviceRegistrationRepository`
- `IdGenerator`
- `DeviceTokenGenerator`

### New or clarified persistence responsibility

- player persistence must carry ownership state and recovery proof fields required for claim

### Interface boundaries

- backend HTTP endpoint for claim
- browser UI path to submit a claim for the current identity

## Initial API Shape

- `POST /identities/guest`
- `POST /identities/resume`
- `POST /identities/claim`
  claims an identity from the current device token and recovery passphrase
- `GET /healthz`
- `GET /readyz`

## Initial Test Plan

### Domain or use-case tests

- claiming a guest identity changes claim status to `claimed`
- claiming with an empty passphrase is rejected
- claiming preserves the original `player_id`
- claiming an already claimed identity is rejected explicitly

### Integration tests

- claim persists claimed status in SQLite
- resume after claim still returns the same player identity with updated claim status
- claim with unknown device token returns explicit not-found

## Scenario Definition

1. browser creates a guest identity
2. backend returns `player_id`, `device_token`, and `claim_status = guest`
3. user decides to claim ownership
4. browser submits the current device token and recovery passphrase
5. backend marks the identity as claimed
6. a later resume on the same browser returns the same `player_id` with `claim_status = claimed`

## Done Criteria

- the repository supports claiming an existing guest identity with a recovery passphrase
- claiming preserves identity continuity and updates only ownership state
- tests prove that claim state survives SQLite-backed persistence
- the slice does not expand into cross-device recovery behavior
