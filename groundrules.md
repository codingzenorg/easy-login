# Ground Rules

## Purpose

This document defines implementation rules for AI coding agents and humans working on this project.

Its purpose is to preserve architectural coherence while allowing incremental progress.

These rules should be treated as default constraints unless a deliberate architectural decision explicitly overrides them.

---

## Primary Objective

The objective of this project is to **refine business models** and make business behavior executable, testable, and inspectable.

This is **not** a production rewrite, infrastructure exercise, or framework showcase.

Every implementation choice should be evaluated by asking:

- does this clarify business behavior?
- does this improve testability?
- does this reduce accidental complexity?
- does this preserve inspectability?

If not, it is probably unnecessary.

---

## General Rules

### 1. Favor explicitness over cleverness

Prefer:

- explicit names
- explicit types
- explicit use cases
- explicit dependencies
- explicit state transitions

Avoid magic, hidden conventions, and “smart” abstractions that make the model harder to read.

### 2. Favor simplicity over infrastructure realism

Do not reproduce production infrastructure details unless the business depends on them.

This project exists to understand and refine behavior, not to mimic infrastructure, network failures, or deployment topology in full detail.

### 3. Favor business meaning over technical symmetry

The code should be shaped primarily by:

- business concepts
- use cases
- invariants
- state transitions

Do not force uniform abstractions where the business is asymmetrical.

### 4. Favor local reasoning

A reader should be able to understand a use case or model by reading a small number of files.

Avoid designs that require jumping across many indirection layers to understand simple behavior.

### 5. Favor deterministic behavior

Whenever possible, time, IDs, external responses, and message processing should be deterministic and controllable.

---

## Architectural Rules

### 6. Respect layer boundaries

Dependencies should point inward:

- domain should not depend on infrastructure
- application should not depend on concrete adapters
- interfaces should remain thin
- infrastructure should implement ports, not define business behavior

### 7. Keep business rules inside the domain or use cases

Do not place business rules in:

- repository SQL
- controllers/facades
- serializers
- fake API implementations
- test setup code

Use repositories for persistence, use cases for orchestration, and domain models/services for rules.

### 8. Use use cases as the main unit of behavior

Each meaningful workflow should have a named use case.

Avoid generic “manager”, “processor”, or “service” classes that accumulate unrelated orchestration.

Good examples:

- `PlaceOrder`
- `AddItemToCart`
- `CancelOrder`
- `RequestReturn`

Bad examples:

- `CheckoutManager`
- `BusinessLogicService`
- `GeneralProcessor`

### 9. Treat the API facade as an adapter, not as the application core

The API facade may translate requests and responses, but it should not own the workflow logic.

Business orchestration belongs in use cases.

### 10. Keep repositories intention-revealing

Repository methods should express intent clearly.

Prefer:

- `get_by_id(...)`
- `get_open_cart_for_shopper(...)`
- `save(...)`

Avoid repositories becoming large generic query bags.

### 11. Prefer one generic message table over one table per topic

For message simulation, default to a generic SQLite-backed message store.

Do not create one table per topic unless a strong and explicit reason appears.

### 12. Depend on abstractions

The model should publish through a message bus port.

Do not leak queue-specific assumptions into the domain or use cases unless the business meaning truly depends on them.

---

## Modeling Rules

### 13. Prefer behavior-rich models over anemic records

Domain objects should protect and express valid state transitions.

Avoid turning models into passive data containers with all logic pushed elsewhere.

### 14. Use value objects when they clarify meaning

When a concept carries meaning, constraints, or formatting rules, prefer a value object over raw primitives.

Likely candidates include:

- money
- dates
- statuses
- identifiers
- quantities
- correlation IDs

### 15. Use domain services only when needed

A domain service is appropriate when logic does not fit naturally in one entity/value object.

Do not use services as a default dumping ground for domain behavior.

### 16. Protect invariants at the model boundary

If a state transition is invalid, the model should reject it clearly.

Do not rely on UI validation or test discipline alone to keep state valid.

### 17. Prefer domain events for meaningful internal facts

When something meaningful happens in the model, represent it explicitly when useful.

Examples:

- order placed
- order cancelled
- item added to cart

Do not emit events for noise.

---

## Testing Rules

### 18. Tests are specification, not just regression nets

Tests should describe intended behavior clearly.

A reader should learn the model by reading tests.

### 19. Prefer domain tests first

When refining rules, begin with the smallest testable unit that captures business behavior.

Usually this means domain tests before infrastructure tests.

### 20. Use integration tests to validate mappings and flows

Integration tests should focus on:

- repository persistence fidelity
- SQLite behavior
- message recording/dispatch
- end-to-end use-case flows

Do not overuse integration tests for logic that belongs in simpler domain tests.

### 21. Keep tests deterministic

Tests should not depend on real network access, current wall-clock time, or nondeterministic ordering unless such behavior is under deliberate test.

### 22. Avoid bloated fixtures

Test setup should remain readable.

Prefer small explicit builders/factories over giant opaque fixtures.

---

## External Dependency Rules

### 23. All external APIs must be behind ports

No domain or use case should call external APIs directly.

All such access must go through interfaces/ports with fake implementations available locally.

### 24. Fakes must be simple and inspectable

Fake adapters should not try to become mini production systems.

They should support controlled scenarios and easy assertions.

### 25. No real external access in normal test runs

By default, tests should run fully locally.

Any exception should be explicit and isolated.

---

## Persistence Rules

### 26. SQLite is the default persistence target

Use SQLite as the default local persistence mechanism unless a strong reason appears otherwise.

### 27. Keep SQL isolated

SQL should live inside infrastructure repositories or dedicated persistence helpers.

Do not spread SQL strings through use cases or domain objects.

### 28. Preserve inspectability of stored state

Schema and persistence choices should make it easy to inspect current state manually when debugging.

This project benefits from visibility.

### 29. Avoid premature schema fragmentation

Do not create many specialized tables, stores, or persistence mechanisms too early.

Prefer fewer, clearer structures first.

---

## Coding Style Rules

### 30. Choose boring, readable Python

Prefer straightforward Python over highly meta-programmed or framework-heavy styles.

Good qualities:

- readable
- typed where useful
- small functions
- explicit constructors
- limited magic

### 31. Avoid framework lock-in early

Do not introduce heavy frameworks unless they solve a demonstrated problem.

Start with plain Python modules and explicit composition.

### 32. Keep files focused

A file should usually have one clear responsibility.

Avoid giant files mixing DTOs, use cases, repositories, and domain logic together.

### 33. Name things by business intent

Names should reflect domain meaning or use-case intent.

Avoid vague technical names such as:

- utils
- helpers
- common
- misc
- manager
- processor

### 34. Comments should explain why, not restate what

Use comments to explain intent, constraints, trade-offs, or non-obvious reasoning.

Do not clutter code with comments that merely narrate the syntax.

---

## Refactoring Rules

### 35. Refactor toward clarity, not abstraction density

Refactoring should make the model easier to understand.

Do not introduce abstraction layers only to make the code look more “architectural.”

### 36. Let discovered concepts reshape the model

As the frontend workflows and business rules become clearer, it is acceptable to change model boundaries.

Preserve behavior, not accidental historical structure.

---

## Message Bus Rules

### 37. Message processing should be explicit

Prefer explicit dispatch/polling in tests and local runners.

Avoid hidden background behavior unless required for a specific scenario.

### 38. Start with minimal semantics

The initial message bus should support only what is necessary:

- publish
- fetch pending
- mark processed
- mark failed
- retry count
- optional delay

Do not add dead-letter queues, visibility timeout simulation, or duplicate delivery logic until a concrete need appears.

### 39. Messages should remain inspectable

Payloads, headers, timestamps, and status should be easy to inspect in SQLite.

---

## Decision Rules for AI Agents

When implementing something ambiguous, prefer the option that is:

1. more explicit
2. easier to test
3. easier to inspect manually
4. less coupled to infrastructure
5. closer to business meaning
6. smaller in surface area
7. easier to replace later

If two options are functionally equivalent, choose the simpler one.

---

## When To Pause and Reconsider

Pause and reassess if a change introduces any of the following:

- domain importing infrastructure
- use case containing SQL or HTTP logic
- repository containing core business decisions
- facade/controller becoming orchestration-heavy
- fake adapter becoming complex and stateful without need
- excessive inheritance
- unexplained indirection
- framework setup larger than the business logic itself

These are signs of architectural drift.

---

## Default Delivery Mindset

Every increment should try to leave the project with:

- one clearer business concept
- one better-named use case
- one more trustworthy test
- one less infrastructure leak
- one more inspectable flow

Small coherent steps are preferred over large speculative rewrites.

---

## Summary

The default posture for this project is:

- modular monolith
- DDD-inspired
- use-case centered
- behavior-first
- SQLite-backed
- fake-integrations locally
- deterministic by default
- simple message simulation
- explicit code over clever code

When unsure, optimize for clarity and model quality, not for production mimicry.

