# Packs

## Purpose

This document defines how implementation packs relate to the MRL core.

Use it when an adopting repository needs to choose a language, architecture, runtime topology, or test shape without rewriting the MRL loop itself.

---

## Definition

A pack is a repository-local set of implementation defaults.

A pack may define:

- programming language or languages
- architectural style
- source and test layout
- scenario runner shape
- integration and persistence defaults
- naming conventions tied to the selected architecture
- multi-runtime rules when one slice spans more than one executable surface

A pack must not redefine the MRL loop itself.

---

## MRL Core Versus Pack

MRL core owns:

- phase sequence
- artifact chain
- semantic docs
- slice docs
- evaluation and release discipline
- skill isolation model

Packs own:

- code layout
- toolchain files
- testing layout
- architectural boundaries
- runtime topology defaults
- implementation conventions

---

## Selection Rule

Each repository should make one pack selection explicit in:

- `architecture.md`
- `docs/building/project_structure.md`
- `decisions.md`

The selected pack may still allow per-slice variation, especially when one model spans multiple runtimes.

---

## Slice Metadata

Each slice document should declare at least:

- `pack`
- `runtime_targets`
- `architecture_mode`

Example:

```md
- pack: polyglot_client_server
- runtime_targets:
  - client/typescript
  - server/go
- architecture_mode: event_sourced
```

This metadata keeps the semantic model shared while making implementation placement explicit.

---

## Suggested Initial Packs

- `python_ddd_monolith`
- `typescript_application`
- `go_service`
- `event_sourced_domain`
- `polyglot_client_server`

Repositories do not need all packs. They should define only the packs they actually intend to use.

---

## Creating A New Pack

When promoting a new pack from a real project:

1. build a few real slices first
2. capture repeated language or architecture decisions
3. separate reusable pack guidance from project-specific choices
4. create a new file under `docs/packs/`
5. update this document to list the new pack
6. record the pack decision in `decisions.md`

Use `docs/packs/_template.md` as the canonical example when creating a new pack document.
