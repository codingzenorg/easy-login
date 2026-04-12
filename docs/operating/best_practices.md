# MRL Best Practices

## Purpose

This document captures practical guidance for running MRL well.

Use it for habits that improve phase discipline, artifact quality, and architectural timing without changing the core loop definition.

---

## Architecture Timing

A common first-pass mistake is to define the implementation architecture during `extract`.

That usually compresses `extract`, `refine`, and part of `build` into one step.

A better default is:

- `extract` identifies domain signals, constraints, tensions, and architectural pressures
- `refine` turns those pressures into explicit architecture and slice decisions
- `build` implements only the architecture needed for the current slice

### What belongs in `extract`

Capture facts such as:

- fixed language or runtime constraints
- existing system boundaries
- latency or authority pressures
- external protocol requirements
- deployment or platform constraints that are already decided outside the lab

These are extracted constraints, not internal architecture choices.

### What belongs in `refine`

Decide only as far as needed for the next slice:

- authority model
- protocol shape
- pack selection
- runtime split
- persistence style
- client/server responsibility split
- event, command, or snapshot boundaries

Architecture is usually an interpretation of extracted evidence, so it should normally become explicit here.

### What must not happen in `build`

`build` should not silently redesign the architecture.

If implementation pressure forces a meaningful architectural change, push that back into:

- the current slice definition
- `decisions.md`
- or a return to `refine`

---

## Practical Rule

Use this test:

- if it is a discovered fact, keep it in `extract`
- if it is a chosen shape, decide it in `refine`
- if it is only needed to make the current slice executable, implement it in `build`

---

## Example

For a browser game with a Go backend:

### `extract`

Capture:

- browser client is required
- backend must be Go
- low-latency movement matters
- anti-cheat or authoritative validation matters
- real-time transport is required

### `refine`

Decide:

- authoritative server simulation
- WebSocket protocol
- client prediction rules
- reconciliation strategy
- which collisions are predicted locally and which are authoritative remotely

### `build`

Implement one thin slice such as:

- player moves
- player collides with one power-up
- server validates and applies the effect
- client reconciles rendered state

---

## Why This Matters

This separation keeps:

- `extract` evidence-oriented
- `refine` design-oriented
- `build` execution-oriented

That usually reduces accidental early lock-in and makes architecture easier to evolve.

---

## Session Productivity

If the repository wants lightweight operational feedback on how MRL work is progressing, track session metrics alongside normal artifact review.

Use `docs/session_metrics.md` as the repository pattern for:

- token-efficiency tracking
- historical session snapshots
- simple investigation thresholds when sessions become expensive or drift-prone

These metrics should guide inspection of session shape, not replace review of slice quality, model quality, or delivered behavior.
