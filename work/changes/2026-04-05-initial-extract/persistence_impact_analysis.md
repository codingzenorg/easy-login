# Persistence Impact Analysis

## Reason For Impact Note

The current build proves guest identity continuity semantically, but not durably. The repositories are in memory, so continuity is lost when the backend process stops.

That is now the strongest mismatch between the model and the executable system.

## Impacted Areas

- backend infrastructure
  in-memory repositories need SQLite-backed equivalents
- backend startup
  the service needs database configuration and schema initialization
- tests
  integration coverage should verify continuity across repository or process recreation
- local development
  developers need a stable way to point the backend at a SQLite file

## Refinement Judgment

The next slice should harden persistence before claim and recovery are added.

Why this slice now:

- it preserves the current business boundary
- it closes a real gap in the continuity promise
- it makes later claim and recovery work build on durable identity records rather than process-local state

## Boundaries To Preserve

- do not introduce claim or recovery logic in this slice
- do not redesign the client flow unless backend configuration forces a minimal adjustment
- do not turn the slice into broad operational infrastructure work unrelated to durable identity continuity

## Expected Follow-Through For Build

- keep repository interfaces stable if possible
- add SQLite-backed implementations behind those interfaces
- make tests deterministic by controlling database paths and startup state
