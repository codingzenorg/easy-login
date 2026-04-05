# Request

## Goal

Extract the first semantic baseline for `easy-login` from the initial raw material and record an expected architecture reference for future refinement and build work.

## Requested Inputs

- initial raw material in `work/sources/easy-login-initial.md`
- architecture reference repository: `https://github.com/wastingnotime/contacts`

## Explicit Extraction Signals

- `easy-login` is a lightweight identity service for browser-first games and simple applications
- instant access without forced login is a primary product constraint
- persistent identity should exist without requiring full authentication
- ownership should be progressive rather than mandatory at first contact
- the architecture reference for future implementation should follow the broad split used in `contacts`:
  frontend app and backend app developed independently but kept in one repository
- the expected frontend direction is Mithril
- the expected backend direction is Go
- future `expose` and `living` phases should assume planned publication on a Docker Swarm cluster, following the operational style already used in `contacts`

## Source Evidence

- `work/sources/easy-login-initial.md`
- `https://github.com/wastingnotime/contacts`
- `https://github.com/wastingnotime/contacts/blob/main/README.md`
- `https://github.com/wastingnotime/contacts/blob/main/apps/web/package.json`
- `https://github.com/wastingnotime/contacts/blob/main/apps/web/src/index.js`
- `https://github.com/wastingnotime/contacts/blob/main/apps/api/main.go`
- `https://github.com/wastingnotime/contacts/blob/main/apps/api/dockerfile`
- `https://github.com/wastingnotime/contacts/blob/main/apps/web/dockerfile`
- `https://github.com/wastingnotime/contacts/blob/main/deploy/compose/docker-compose.dev.yml`

## Extraction Boundaries

- treat `contacts` as an architecture and deployment reference, not as a domain template
- keep implementation choices at the level of expected direction, not final slice design
- preserve future room to refine storage, claim flow details, and deployment rollout choices
