# Domain Background Knowledge

## Purpose

This document captures broad background knowledge about the target domain.

It is not the repository's extracted glossary. It is reference material distilled from books, articles, standards, external systems, or other domain sources.

Use it mainly during expectation-gap detection.

---

## Domain Frame

Lightweight multiplayer games and simple browser applications often need continuity before they need full authentication.
Users expect immediate entry, but product owners still need a stable internal identity for scores, progression, and recovery.

This creates a recurring domain split:

- visible identity for social presence
- internal identity for continuity and state
- ownership proof for recovery and dispute handling

## Common Expectations In This Domain

- users can start immediately with minimal friction
- the same browser usually resumes the same identity automatically
- rankings and progression do not break when a display name changes
- losing browser storage may lose an unclaimed identity
- once ownership is claimed, users expect some portable recovery path

## Useful Domain Language

- guest identity
- persistent guest
- claimed identity
- recovery passphrase
- display name
- identity continuity
- ownership proof
- browser-local device token

## Practical Background Constraints

- browser storage is convenient but fragile
- display names alone are weak identifiers because they are social labels, not proof
- email and OAuth increase recovery strength but add friction and infrastructure cost
- passphrase-based recovery keeps the system lightweight but shifts responsibility to the user
- games and lightweight apps often prefer graceful degradation over hard account gates

## Architecture Reference Background

The `contacts` repository provides architecture background for expected implementation shape, even though it is from a different business domain.

Relevant transferable observations:

- frontend and backend are kept as separate apps in one repository
- the frontend uses Mithril and client-side routing
- the backend is a Go HTTP service
- local development uses containerized dependencies through compose
- the backend exposes `/healthz` and `/readyz`, which is useful for future Docker Swarm exposure
- frontend and backend each have their own Docker build path

These observations support an expected architecture reference, not a commitment to reuse the same libraries, persistence model, or exact folder layout.

## Exposure And Living Background

Future `expose` and `living` phases should assume a real deployment target on Docker Swarm, following the same general operational style already present in `contacts`.

Implications worth preserving during later phases:

- services should expose clear health and readiness signals
- frontend and backend should remain independently buildable and deployable
- deployment assumptions should be explicit rather than hidden in local-only scripts
- container images should run with production-safe defaults where practical

## Likely Omissions To Watch During EGD

- conflating display name with persistent identity
- treating device continuity as sufficient proof of ownership
- forgetting recovery behavior when storage is cleared or device changes
- introducing heavy login requirements too early and breaking instant entry
- coupling implementation too tightly to a reference architecture that was meant only as guidance
