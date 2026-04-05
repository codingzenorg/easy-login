# Claim Impact Analysis

## Reason For Impact Note

After durable guest persistence, the strongest remaining model gap is the absence of ownership. The system still cannot distinguish a temporary guest identity from one that the user has intentionally claimed.

## Impacted Areas

- domain model
  player state now needs an ownership transition from guest to claimed
- persistence
  player storage must carry claim state and recovery proof fields
- backend API
  a new claim endpoint is needed
- browser client
  the current identity view needs a path to submit a claim request

## Refinement Judgment

Claim should come before cross-device recovery.

Why this slice now:

- recovery without a prior claim model would have no stable ownership target
- claim introduces the ownership state transition with limited surface area
- it preserves the current continuity behavior while extending the model carefully

## Boundaries To Preserve

- do not implement full recovery on another device yet
- do not add email, OAuth, or unrelated account features
- do not turn this slice into broad security hardening work outside the current product model

## Expected Follow-Through For Build

- extend the current player representation to support claimed status
- add one new backend use case and HTTP endpoint for claim
- keep resume semantics stable while surfacing the updated claim status
