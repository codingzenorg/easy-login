# Project Structure

## Purpose

This document defines the intended folder layout and file responsibilities for the project.

Its purpose is to give AI coding agents and humans a concrete structural map so implementation can evolve coherently without each contributor inventing a different organization.

The structure is split into:

- MRL core artifacts that every adopting repository should keep
- pack-specific implementation layouts that can vary by language and architecture

If a change becomes necessary, register the reasoning in `decisions.md`.

---

## Structural Philosophy

The structure should optimize for:

- business meaning first
- clear boundaries
- easy local navigation
- low accidental complexity
- small focused files
- predictable placement of new code

The project should avoid:

- framework-shaped organization too early
- giant shared `utils` or `common` folders
- mixing domain, persistence, and interface code in the same files
- excessive nesting without purpose

---

## Top-Level Layout

```text
project_root/
  readme.md
  architecture.md
  groundrules.md
  decisions.md
  AGENTS.md

  docs/
    building/
      bootstrap_plan.md
      project_structure.md
    packs/
      <selected_pack>.md
    evaluation/
      scenario_evaluation.md
    operating/
      mrl_reference.md
      mrl_starter.md
      packs.md
      skills_workflow.md
    semantics/
      model_hypothesis.md
      domain_background_knowledge.md
    slices/

  work/
    sources/
    changes/

  pyproject.toml               # only when a Python pack uses it
  package.json                 # only when a JavaScript/TypeScript pack uses it
  go.mod                       # only when a Go pack uses it
  .python-version              # optional and pack-specific
  .gitignore
  .env.example                 # only if needed later

  src/                         # shaped by the selected pack

  tests/                       # shaped by the selected pack

  scripts/

  data/                        # optional and pack-specific
```

---

## Top-Level Responsibility of Each Area

### Root documents
These files define project intent and coordination rules.

- `readme.md`: repository overview
- `architecture.md`: architectural direction
- `groundrules.md`: implementation constraints
- `decisions.md`: decision history
- `AGENTS.md`: contributor guidance

### `docs/building/`
Contains implementation planning and structure guidance.

- `docs/building/bootstrap_plan.md`: implementation order
- `docs/building/project_structure.md`: structural map

### `docs/operating/`
Contains execution guidance for isolated skill-based workflows and the portable MRL reference set.

- `docs/operating/mrl_reference.md`: compact MRL definition
- `docs/operating/mrl_starter.md`: portable starter layout for new repositories
- `docs/operating/packs.md`: pack model and selection rules
- `docs/operating/skills_workflow.md`: skill execution model and artifact-driven loop

### `docs/packs/`
Contains pack definitions.

Each pack file should define:

- intended language or languages
- architectural shape
- repository layout
- testing defaults
- runtime targeting rules
- pack-specific constraints

### `docs/evaluation/`
Contains scenario-review design and evaluator guidance.

- `docs/evaluation/scenario_evaluation.md`: expectation-gap evaluation loop

### `docs/semantics/`
Contains business-language material used during analysis and evaluation.

- `docs/semantics/model_hypothesis.md`: analysis and build guidance
- `docs/semantics/domain_background_knowledge.md`: expectation-review context

### `work/`
Contains tracked working artifacts that preserve loop traceability without promoting every intermediate artifact into `docs/`.

- `work/sources/`: canonical folder for curated raw evidence and original source material
- `work/changes/`: bounded request artifacts and phase-local handoff material tied to a specific change

`work/` is repository memory, not personal scratch space.
Preserve original evidence in `work/sources/` before extraction or refinement artifacts are produced elsewhere.

### `src/`
Contains production code shaped by the selected pack.

### `tests/`
Contains executable specification and integration validation shaped by the selected pack.

### `scripts/`
Contains small runners or maintenance scripts useful for local development.

### `data/`
Optional local persisted files for manual inspection or scenario runs.

This directory should not hold mandatory production-like assets. It is for local labs and experiments and its substructure is pack-specific.

### `docs/`
Contains implementation-adjacent documentation that should not live in the root strategic files.

### `docs/slices/`
Contains one document per vertical slice, including discovery scope, use-case contract, initial rules, test plan, scenario definition, and done criteria.

---

## Source Tree

```text
src/
  ... pack-specific code layout
```

### Example: `python_ddd_monolith` pack

```text
src/
  app/
    domain/
      models/
      value_objects/
      services/
      events/
      exceptions/

    application/
      use_cases/
      dto/
      ports/
      services/

    infrastructure/
      sqlite/
        repositories/
        message_bus/
        migrations/
      fakes/
      clocks/
      ids/
      gateways/
      mappers/

    interfaces/
      api_facade/
      cli/
```

### Example: `polyglot_client_server` pack

```text
src/
  client/
  server/
  shared_contracts/

tests/
  client/
  server/
  integration/
  contracts/
```

---

# 1. Pack-Specific Layers

## Purpose

The selected pack defines the concrete layers or modules inside `src/`.

For the `python_ddd_monolith` pack, the domain layer contains the business model.

It should represent:

- entities
- aggregates
- value objects
- domain services
- domain events
- domain-specific exceptions

It must not depend on SQLite, HTTP, or framework concerns.

---

## `src/app/domain/models/`

### Purpose
Contains entities and aggregates that own business behavior.

### Typical contents
Examples:

- `cart.py`
- `order.py`
- `order_item.py`
- `shopper.py`
- `return_request.py`

### Rules

- keep models behavior-rich
- protect invariants inside the model
- avoid turning models into passive records
- keep files focused on one main concept when possible

### Example responsibility

`order.py` should contain order-related state and transitions, not SQL and not facade logic.

---

## `src/app/domain/value_objects/`

### Purpose
Contains small immutable concepts with business meaning.

### Typical contents
Examples:

- `money.py`
- `quantity.py`
- `order_status.py`
- `cart_status.py`
- `shopper_id.py`
- `correlation_id.py`

### Rules

- prefer value objects when they clarify meaning or enforce constraints
- avoid primitive obsession for important concepts
- keep them small and explicit

---

## `src/app/domain/services/`

### Purpose
Contains domain services for rules that do not belong clearly inside a single entity or value object.

### Typical contents
Examples:

- `shipping_option_calculator.py`
- `return_policy.py`
- `pricing_service.py`

### Rules

- do not use this folder as a dumping ground
- add a domain service only when the model really needs cross-entity or free-standing business logic
- prefer model methods when ownership is clear

---

## `src/app/domain/events/`

### Purpose
Contains domain events representing meaningful facts that happened in the domain.

### Typical contents
Examples:

- `cart_created.py`
- `order_placed.py`
- `order_cancelled.py`
- `item_added_to_cart.py`

### Rules

- events should represent something meaningful, not noise
- keep event payloads explicit
- domain events should not contain transport concerns

---

## `src/app/domain/exceptions/`

### Purpose
Contains domain-specific exceptions for invariant violations or invalid state transitions.

### Typical contents
Examples:

- `invalid_order_transition.py`
- `cart_already_checked_out.py`
- `return_not_allowed.py`

### Rules

- prefer specific exception names
- do not centralize all errors into one generic exception type
- exceptions should signal domain meaning, not infrastructure failure

---

# 2. Application Layer

## Purpose

The application layer coordinates behavior.

It should contain:

- use cases
- input/output DTOs
- ports (abstractions)
- optional application services when orchestration helpers are needed

The application layer may depend on domain abstractions but should not depend on concrete infrastructure implementations.

---

## `src/app/application/use_cases/`

### Purpose
Contains explicit use cases named by business intent.

### Typical contents
Examples:

- `place_order.py`
- `add_item_to_cart.py`
- `cancel_order.py`
- `request_return.py`

### Rules

- one file per main use case
- orchestration belongs here
- use cases should load models via repositories, invoke domain behavior, persist changes, publish events/messages, and return results
- avoid raw SQL, HTTP calls, and framework coupling here

### Example responsibility

`place_order.py` may:

- validate input DTO shape
- load cart/order context
- convert the cart into an order
- persist updates
- publish a message/event
- return a result DTO

---

## `src/app/application/dto/`

### Purpose
Contains input and output data structures used by use cases and adapters.

### Typical contents
Examples:

- `place_order_input.py`
- `place_order_result.py`
- `request_return_input.py`

### Rules

- keep DTOs simple and explicit
- do not place domain behavior in DTOs
- DTOs exist to move structured data across boundaries

### Naming guidance
Prefer names based on the use case:

- `<use_case>_input.py`
- `<use_case>_result.py`

---

## `src/app/application/ports/`

### Purpose
Contains abstract contracts for dependencies needed by the application/domain.

### Typical contents
Examples:

- `order_repository.py`
- `cart_repository.py`
- `message_bus.py`
- `clock.py`
- `id_generator.py`
- `inventory_gateway.py`

### Rules

- ports should be small and intention-revealing
- define only methods that are currently needed
- do not create giant “god interfaces”

---

## `src/app/application/services/`

### Purpose
Contains optional application-level helper services for orchestration that does not belong in a single use case but is still not domain logic.

### Typical contents
Examples:

- `message_dispatcher.py`
- `scenario_runner.py`

### Rules

- do not create this folder unless a real need appears
- prefer keeping orchestration in use cases until reuse pressure becomes real
- application services should not hide the core business flow unnecessarily

---

# 3. Infrastructure Layer

## Purpose

The infrastructure layer implements ports and technical details.

It may contain:

- SQLite repositories
- SQLite message bus
- fake external integrations
- fake repositories
- clock implementations
- ID generator implementations
- persistence mappers
- migrations

Infrastructure should serve the model, not define it.

---

## `src/app/infrastructure/sqlite/`

### Purpose
Contains SQLite-specific persisted implementations.

### Subfolders

```text
sqlite/
  repositories/
  message_bus/
  migrations/
```

---

## `src/app/infrastructure/sqlite/repositories/`

### Purpose
Contains concrete SQLite repository implementations.

### Typical contents
Examples:

- `sqlite_order_repository.py`
- `sqlite_cart_repository.py`

### Rules

- keep SQL here or in nearby dedicated helpers
- map rows explicitly to domain models
- avoid leaking SQL outside infrastructure

---

## `src/app/infrastructure/sqlite/message_bus/`

### Purpose
Contains the SQLite-backed message bus implementation.

### Typical contents
Examples:

- `sqlite_message_bus.py`
- `sqlite_message_store.py`

### Rules

- start with minimal semantics
- keep messages inspectable in SQLite

---

## `src/app/infrastructure/sqlite/migrations/`

### Purpose
Contains schema creation scripts or migration helpers for the local SQLite database.

### Typical contents
Examples:

- `0001_create_orders.sql`
- `0002_create_messages.sql`

### Rules

- keep migrations explicit and readable
- create only what the implemented slices need
- the schema is a refinement schema

---

## `src/app/infrastructure/fakes/`

### Purpose
Contains in-memory or deterministic fake implementations of ports.

### Typical contents
Examples:

- `in_memory_order_repository.py`
- `in_memory_message_bus.py`
- `fake_inventory_gateway.py`
- `fake_pricing_gateway.py`

### Rules

- fakes must be simple and inspectable
- avoid hidden complexity or global state
- shape fakes around testing/scenario needs

---

## `src/app/infrastructure/clocks/`

### Purpose
Contains implementations of the `Clock` port.

### Typical contents
Examples:

- `system_clock.py`
- `fixed_clock.py`

### Rules

- tests should generally use fixed or controlled clocks
- do not call wall-clock time directly from the domain or use cases

---

## `src/app/infrastructure/ids/`

### Purpose
Contains implementations of ID generation.

### Typical contents
Examples:

- `uuid_id_generator.py`
- `incremental_id_generator.py`

### Rules

- tests may use deterministic generators
- production-like uniqueness is not the main goal during bootstrap

---

## `src/app/infrastructure/gateways/`

### Purpose
Contains concrete gateway implementations for external APIs if needed later.

### Typical contents
Examples:

- `http_inventory_gateway.py`
- `fixture_pricing_gateway.py`

### Rules

- keep real gateways optional and isolated
- do not let them bleed into the domain/application
- if only fakes exist initially, that is acceptable

---

## `src/app/infrastructure/mappers/`

### Purpose
Contains explicit mapping helpers when repository or integration mapping becomes non-trivial.

### Typical contents
Examples:

- `order_row_mapper.py`
- `order_event_mapper.py`

### Rules

- only add dedicated mappers when mapping complexity becomes real
- do not create a mapping layer for its own sake

---

# 4. Interface Layer

## Purpose

The interface layer adapts external input/output to the application layer.

It should stay thin.

Typical roles:

- API facade
- CLI commands
- scenario runners for manual execution

It should not become the place where business orchestration lives.

---

## `src/app/interfaces/api_facade/`

### Purpose
Contains request/response adapters that mimic or expose application behavior in a frontend-friendly way.

### Typical contents
Examples:

- `cart_facade.py`
- `order_facade.py`

### Rules

- keep facades thin
- translate request shapes into use-case DTOs
- translate results into response shapes
- avoid business decision logic here

---

## `src/app/interfaces/cli/`

### Purpose
Contains command-line entry points or local scenario execution commands.

### Typical contents
Examples:

- `run_scenario.py`
- `dispatch_messages.py`

### Rules

- use this for local convenience
- a CLI may be more appropriate than a web API in early phases
- keep command code thin and delegated to use cases/services

---

# 5. Tests

## Purpose

Tests are executable specification.

They should teach the model and protect behavior.

---

## Test Tree

```text
tests/
  unit/
  integration/
  builders/
  fixtures/
```

---

## `tests/unit/`

### Purpose
Contains fast tests for domain and application behavior without real persistence.

### Suggested sub-organization

```text
unit/
  domain/
  application/
```

### Typical contents
Examples:

- `tests/unit/domain/test_order_transitions.py`
- `tests/unit/application/test_place_order.py`

### Rules

- keep them fast
- prefer explicit setup
- use fakes rather than SQLite where possible

---

## `tests/integration/`

### Purpose
Contains tests that validate persistence, message storage, and end-to-end local flows.

### Typical contents
Examples:

- `test_sqlite_order_repository.py`
- `test_sqlite_message_bus.py`
- `test_place_order_with_sqlite.py`

### Rules

- focus on mappings and local infrastructure behavior
- keep each integration test centered on one clear concern

---

## `tests/builders/`

### Purpose
Contains test builders/factories that improve readability of setup.

### Typical contents
Examples:

- `cart_builder.py`
- `order_builder.py`

### Rules

- add builders only when setup repetition becomes noisy
- builders should improve clarity, not hide too much meaning

---

## `tests/fixtures/`

### Purpose
Contains shared pytest fixtures or fixture data that are broadly useful.

### Typical contents
Examples:

- temporary SQLite database fixture
- fixed clock fixture
- common fake message bus fixture

### Rules

- avoid giant fixture webs
- prefer local setup if a fixture makes the test harder to read

---

# 6. Scripts

## `scripts/`

### Purpose
Contains small developer-oriented scripts.

### Typical contents
Examples:

- initialize local db
- run a named scenario
- dump pending messages
- seed local test data

### Rules

- scripts should remain thin wrappers
- do not move business logic here
- use scripts for convenience, not architecture

---

# 7. Local Data

## `data/sqlite/`

### Purpose
Contains optional local SQLite files for experimentation and manual inspection.

### Typical contents
Examples:

- `lab.db`
- `scenario_return_request.db`

### Rules

- do not assume these files are committed unless deliberately chosen
- use this area for local reproducible labs
- tests should usually create temporary databases instead of relying on persistent files here

---

# 8. Naming Conventions

## Modules

Prefer names by business concept or use case intent.

Good examples:

- `order.py`
- `place_order.py`
- `sqlite_message_bus.py`

Avoid vague names:

- `utils.py`
- `helpers.py`
- `common.py`
- `manager.py`
- `processor.py`

## Tests

Use names that reveal behavior.

Good examples:

- `test_order_cannot_be_cancelled_twice.py`
- `test_place_order_publishes_order_placed_message.py`

---

# 9. What Should Not Exist Early

Avoid introducing these too early:

- `shared/` folders with mixed responsibilities
- framework-specific folders without real use
- generic `services/` full of unrelated code
- separate read/write package hierarchies before real pressure appears
- event sourcing infrastructure before the model actually needs it
- plugin systems or dependency injection frameworks before explicit composition becomes painful

---

# 10. Recommended Early File Set

A realistic early set of files could look like:

```text
src/app/domain/models/cart.py
src/app/domain/models/order.py
src/app/domain/value_objects/money.py
src/app/domain/value_objects/order_status.py
src/app/domain/events/order_placed.py
src/app/domain/exceptions/cart_already_checked_out.py

src/app/application/dto/place_order_input.py
src/app/application/dto/place_order_result.py
src/app/application/ports/order_repository.py
src/app/application/ports/message_bus.py
src/app/application/ports/clock.py
src/app/application/ports/id_generator.py
src/app/application/use_cases/place_order.py

src/app/infrastructure/fakes/in_memory_order_repository.py
src/app/infrastructure/fakes/in_memory_message_bus.py
src/app/infrastructure/clocks/fixed_clock.py
src/app/infrastructure/ids/incremental_id_generator.py
src/app/infrastructure/sqlite/repositories/sqlite_order_repository.py
src/app/infrastructure/sqlite/message_bus/sqlite_message_bus.py
src/app/infrastructure/sqlite/migrations/0001_create_orders.sql
src/app/infrastructure/sqlite/migrations/0002_create_messages.sql

src/app/interfaces/api_facade/order_facade.py
src/app/interfaces/cli/run_scenario.py

tests/unit/domain/test_order_checkout_rules.py
tests/unit/application/test_place_order.py
tests/integration/test_sqlite_order_repository.py
tests/integration/test_sqlite_message_bus.py
tests/integration/test_place_order_with_sqlite.py
```

This is only a starting pattern.

---

# 11. Structural Change Rule

If a contributor wants to change the structure substantially, they should answer:

- what problem in the current structure is causing friction?
- what becomes simpler with the new structure?
- what are the migration consequences?
- does the change preserve the project’s focus on explicit behavior and low accidental complexity?

If the answer is meaningful, register the change in `decisions.md`.

---

## Summary

For the `python_ddd_monolith` pack, the structure is organized by architectural responsibility:

- `domain` for business meaning
- `application` for use-case orchestration and ports
- `infrastructure` for concrete adapters
- `interfaces` for thin external-facing adapters
- `tests` for executable specification
- `scripts` for local convenience

This structure exists to keep the project legible while the model is being refined.
