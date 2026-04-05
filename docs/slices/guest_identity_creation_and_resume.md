# Guest Identity Creation And Resume

## Selected Pack

Current repository baseline still documents `python_ddd_monolith`, but this slice is intentionally shaped for a future polyglot client/server implementation direction derived from refinement artifacts:

- browser frontend in Mithril
- Go backend service

The build phase should not silently ignore this tension. It should either:

- implement the slice in a minimal shape that preserves the documented runtime split, or
- record a decision that updates the repository architecture baseline before deeper implementation continues

## Runtime Targets

- browser client
- backend HTTP service

## Architecture Mode

Thin first vertical slice focused on identity continuity.

The browser is responsible for:

- collecting a display name from the user
- storing the returned device token locally
- presenting the stored device token on later resume requests

The backend is responsible for:

- creating a persistent guest identity
- associating a device token with that identity
- resuming the same identity when a valid device token is presented

No claim or recovery behavior is included in this slice.

## Discovery Scope

This slice exists to validate the core product promise:

- instant entry without forced login
- continuity on the same browser

It intentionally excludes:

- claimed identity
- recovery passphrase
- cross-device recovery
- ranking integration
- display-name uniqueness policy beyond what is needed to create and resume a guest identity

## Model Pressure

The strongest current model pressure is proving that display identity and persistent identity are separate from the start.

Without that separation:

- the system cannot support same-browser continuity safely
- future claim and recovery flows will be built on unstable assumptions
- later ranking continuity will likely attach to the wrong concept

This first slice should therefore prove:

- a display name can create a player record
- the player record has a stable internal identity
- the same browser can resume that identity using a device token

## Use-Case Contract

### Use case 1: `CreateGuestIdentity`

#### Input

- `display_name`

#### Output

- `player_id`
- `display_name`
- `device_token`
- `claim_status` set to `guest`

#### Required behavior

- create a new player with a stable internal identity
- create a device token associated with that player
- return both values to the caller

### Use case 2: `ResumeIdentityFromDeviceToken`

#### Input

- `device_token`

#### Output on success

- `player_id`
- `display_name`
- `claim_status`

#### Output on miss

- explicit not-found result that allows the client to prompt for guest creation again

#### Required behavior

- resolve the device token to the existing player identity
- return the same `player_id` that was created earlier

## Main Business Rules

- creating a guest identity requires a display name
- a guest identity receives a stable `player_id`
- a guest identity receives a device token for same-browser continuity
- resuming with a valid device token returns the same persistent identity
- display name is not the ownership proof
- claim status is `guest` for all identities in this slice
- an unknown device token must not create an identity implicitly during resume

## Open Questions Deferred

- exact validation rules for display name
- whether duplicate display names are allowed immediately or only until claim exists
- whether a new guest creation should replace a previous local identity on the same browser
- expiration or rotation rules for device tokens

These questions should not be invented during build unless the slice document is updated first.

## Required Ports And Boundaries

### Backend-side ports

- `PlayerRepository`
  responsibilities:
  persist and retrieve player identities
- `DeviceRegistrationRepository`
  responsibilities:
  persist and resolve device-token bindings
- `IdGenerator`
  responsibilities:
  generate deterministic testable player IDs
- `DeviceTokenGenerator`
  responsibilities:
  generate deterministic testable device tokens

### Interface boundaries

- backend API facade or HTTP adapter for guest creation
- backend API facade or HTTP adapter for resume lookup
- browser-side local storage adapter for device token persistence

## Initial API Shape

The exact transport details may change, but build should preserve this semantic shape:

- `POST /identities/guest`
  creates guest identity from display name
- `POST /identities/resume`
  resumes identity from device token
- `GET /healthz`
- `GET /readyz`

The health endpoints are included because future exposure is expected to target Docker Swarm.

## Initial Test Plan

### Domain or use-case tests

- creating a guest identity returns a player ID, device token, and guest claim status
- creating a guest identity with an empty display name is rejected
- resuming with a known device token returns the original player identity
- resuming with an unknown device token returns explicit not-found

### Integration tests

- guest creation persists both player and device registration
- resume flow reads persisted device registration and player identity correctly
- health and readiness endpoints return success

## Scenario Definition

1. new browser session starts with no stored device token
2. user enters display name `henrique`
3. system creates a guest identity and returns `player_id` plus `device_token`
4. browser stores the device token locally
5. later, the same browser sends the stored device token to resume
6. system returns the same `player_id` and display name
7. if local storage is cleared, resume fails explicitly and the user must create a new guest identity

## Done Criteria

- the repository contains executable behavior for guest creation and same-browser resume
- tests prove stable identity continuity through device token lookup
- build does not accidentally introduce claim or recovery behavior into this slice
- the runtime split and architecture tension are acknowledged in implementation notes or follow-up decisions
- health and readiness endpoints exist for the backend service
