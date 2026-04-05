# Implementation Notes

## Implemented Scope

The first build slice now includes:

- Go backend service for guest identity creation and resume
- in-memory repositories for players and device registrations
- health and readiness endpoints for future Docker Swarm exposure
- minimal Mithril browser client that stores the device token in local storage
- explicit shared contract examples for guest creation and resume

## Deliberate Limits

- persistence is in memory for this first slice
- claim and recovery flows are not implemented
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
- health and readiness endpoints
