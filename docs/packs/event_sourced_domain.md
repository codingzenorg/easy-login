# Pack: event_sourced_domain

## Purpose

Use this pack when the primary model is represented through events, streams, projections, and replayable state transitions.

---

## Shape

- language: one or more
- runtime topology: single or multiple runtimes
- architecture mode: event sourcing

---

## Layout

```text
src/
  domain/
  application/
  event_store/
  projections/
  interfaces/

tests/
  domain/
  projections/
  integration/
  replay/
```

---

## Rules

- model changes as append-only domain facts when appropriate
- make stream boundaries explicit
- test replay, projection consistency, and command handling deterministically
- distinguish clearly between source-of-truth events and derived read models

---

## Good Fit

- audit-heavy domains
- replayable simulations
- projection-driven read models
- systems where temporal history is part of the business meaning
