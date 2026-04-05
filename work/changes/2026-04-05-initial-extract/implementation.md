# Implementation Notes

## Implemented Scope

The first three build slices now include:

- Go backend service for guest identity creation and resume
- SQLite-backed repositories for players and device registrations
- claim of an existing guest identity with a recovery passphrase
- health and readiness endpoints for future Docker Swarm exposure
- minimal Mithril browser client that stores the device token in local storage
- explicit shared contract examples for guest creation and resume
- backend database configuration through `SQLITE_PATH`

## Deliberate Limits

- cross-device recovery is not implemented
- no ranking integration exists yet
- display-name validation is limited to non-empty trimmed input

## Architecture Follow-Through

The repository architecture baseline was updated to `polyglot_client_server` during this build so the implementation direction is no longer hidden drift from the starter.

## Validation Intent

Primary validation should come from Go tests covering:

- guest identity creation
- invalid empty display names
- resume by known device token
- explicit not-found on unknown device token
- durable resume from the same SQLite database after store recreation
- successful claim of a guest identity with claimed status persisted
- health and readiness endpoints

## Local Validation Loop

For manual local validation during future slices:

- run the backend from `src/server` with `air`
- run the frontend from `src/client` with `npm run dev`

This keeps the current browser client and Go backend in a fast local feedback loop while preserving the same API contracts used by automated tests.
