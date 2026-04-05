# Implementation Notes

## Implemented Scope

The first four build slices now include:

- Go backend service for guest identity creation and resume
- SQLite-backed repositories for players and device registrations
- claim of an existing guest identity with a recovery passphrase
- recovery of a claimed identity on another browser using the recovery passphrase
- health and readiness endpoints for future Docker Swarm exposure
- minimal Mithril browser client that stores the device token in local storage
- explicit shared contract examples for guest creation and resume
- backend database configuration through `SQLITE_PATH`

## Deliberate Limits

- passphrase rotation is not implemented
- device-token revocation policy is not implemented
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
- successful recovery of a claimed identity with a usable new device token
- health and readiness endpoints

## Local Validation Loop

For manual local validation during future slices:

- run the backend from `src/server` with `air`
- run the frontend from `src/client` with `nvm use` first, then `npm run dev`

This keeps the current browser client and Go backend in a fast local feedback loop while preserving the same API contracts used by automated tests.

When running frontend install or build commands in automation or manual validation, prefer the Node version pinned in the repository `.nvmrc`. Running the client without `nvm use` may trigger avoidable engine-version warnings.
