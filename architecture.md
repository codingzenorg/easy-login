# Architecture

## Purpose

This document defines the architectural shape selected for this repository. Its goal is to give AI coding agents and humans a stable base for implementation decisions, naming, boundaries, and trade-offs without implying that one architecture is mandatory for all MRL repositories.

This project is **not** intended to be a full production rewrite of the current system. It is a **model laboratory** used to:

- refine business models currently represented on model hypothesis
- isolate business rules from infrastructure concerns
- simulate external dependencies locally
- derive and validate use cases from frontend behavior
- make rules executable, testable, inspectable, and easier to evolve

---

## MRL Core Versus Pack

MRL itself is architecture-agnostic.

The MRL core defines:

- the refinement loop
- the artifact chain
- isolated phase execution
- semantic and slice documents
- evaluation and release discipline

Implementation shape is selected through a pack.

This repository currently adopts the `python_ddd_monolith` pack. Other repositories may instead adopt:

- `typescript_application`
- `go_service`
- `event_sourced_domain`
- `polyglot_client_server`

If a repository changes pack, record the decision in `decisions.md` and update this document so it describes the selected shape rather than pretending to be a universal MRL rule.

---

## Core Intent

Within this repository, the system should behave like a **DDD-inspired modular monolith**.

It should prefer:

- explicit use cases over generic service blobs
- domain models with behavior over anemic data containers
- ports and adapters over direct framework coupling
- deterministic local simulation over distributed complexity
- testability and inspectability over premature realism

This project is a **refinement environment**, not a microservices platform.

That statement is local to this selected pack. MRL as a workflow does not require a modular monolith and can support multi-process or multi-runtime systems when the model requires them.

---

## Architectural Style For This Pack

The project follows a layered approach:

```text
Tests -----------------> Use Cases -----------------> Domain Models
Interfaces/API Facade -> Use Cases -----------------> Repositories (ports)
                                          ---------> Message Bus (port)
                                          ---------> External APIs (ports)
```

A more explicit view:

```text
src/app/
  domain/
    models/
    services/
    events/
    value_objects/

  application/
    use_cases/
    ports/
    dto/

  interfaces/
    api_facade/

  infrastructure/
    sqlite/
    repositories/
    message_bus/
    fakes/
    clock/
    ids/

tests/
  unit/
  integration/
```

This is a pack-specific example, not a required layout for every MRL repository.

### Layer responsibilities

#### 1. Tests
Tests are the main executable specification.

They should validate:

- domain invariants
- use-case behavior
- persistence boundaries
- event/message emission
- frontend-derived workflows

#### 2. Application layer
The application layer contains **use cases**.

A use case:

- receives an intention/request
- loads domain objects through repositories
- coordinates domain behavior
- calls services when needed
- persists resulting state
- publishes domain/integration messages when needed
- returns a result DTO

A use case should **not** contain raw SQL, HTTP details, or framework-specific logic.

#### 3. Domain layer
The domain layer contains the business model.

It should include:

- entities
- aggregates
- value objects
- domain services
- domain events

The domain layer is where invariants and rules live.

Examples:

- order status transitions
- inventory constraints
- checkout eligibility
- cart-to-order conversion rules
- return eligibility constraints

The domain should avoid direct dependency on SQLite, HTTP clients, real queues mechanisms, or UI concerns.

#### 4. Interfaces layer
The interface layer exposes use cases in a convenient shape.

For this project, an **API facade** is acceptable as a boundary adapter. It can:

- translate frontend-like requests into use-case inputs
- compose response DTOs
- simulate endpoints if needed

This layer should stay thin.

#### 5. Infrastructure layer
Infrastructure implements ports required by the application/domain.

Examples:

- SQLite repositories
- SQLite-backed message bus
- in-memory message bus
- local fake external APIs
- test clocks
- ID generators

Infrastructure exists to support the model, not define it.

---

## Multi-Runtime Guidance

Some models span more than one runtime.

Examples:

- browser client plus Go server
- simulation engine plus UI shell
- event producer plus projection consumer

In those cases:

- keep one shared semantic model under `docs/semantics/`
- let slice documents declare more than one runtime target
- keep runtime boundaries explicit in contracts, messages, and events
- do not force one code tree to pretend the system is single-runtime when it is not

The pack may still define local boundaries for each runtime, but the semantic model stays shared unless the business itself diverges.

---

## Design Principles

### 1. Behavior over direct state manipulation
The main objective is to move away from direct state handling as the dominant design mode.

Instead of allowing logic to emerge from ad hoc database access, the system should expose behavior through:

- methods on domain objects
- explicit use cases
- constrained repositories

State exists, but behavior should control how state changes.

### 2. Use cases before endpoints
The frontend may inspire workflows, but implementation should be organized around **use cases**, not around controllers or pages.

Examples:

- `CreateCart`
- `AddItemToCart`
- `PlaceOrder`
- `RequestReturn`
- `CancelOrder`

### 3. Ports and adapters
External dependencies must be behind ports.

Typical ports:

- `OrderRepository`
- `CartRepository`
- `MessageBus`
- `Clock`
- `IdGenerator`
- `InventoryGateway`
- `PricingGateway`

This allows infrastructure replacement without changing the model.

### 4. Simple simulation over infrastructure fidelity
Do not reproduce infrastrcture or production integrations unless the business depends on their precise behavior.

The project should simulate only what is necessary to:

- validate use cases
- preserve behavior expectations
- inspect flows
- replay scenarios

### 5. Deterministic execution
The project should prefer deterministic behavior so scenarios can be replayed and reasoned about.

This includes controlling:

- time
- IDs
- external responses
- message processing order when relevant

---

## Dependency Rule

Dependencies should point inward.

- domain depends on nothing external
- application depends on domain and abstract ports
- interfaces depend on application
- infrastructure depends on application/domain contracts
- tests may depend on all layers

In practical terms:

- domain does not import infrastructure modules
- use cases do not know concrete SQLite implementations
- repositories implement interfaces defined closer to the application layer

---

## Domain Modeling Guidance

### Entities and aggregates
Prefer aggregates when a consistency boundary is needed.

A repository should generally persist/load aggregates, not arbitrary internal fragments.

### Value objects
Use value objects to clarify important concepts such as:

- money
- quantity
- sku
- order status
- shopper identifier
- correlation ID

### Domain services
Use domain services when a rule does not belong naturally to a single entity.

Examples:

- calculating shipping options
- deciding return eligibility rules
- cross-model validation

### Domain events
Domain events are useful when something meaningful happened inside the model.

Examples:

- `CartCreated`
- `ItemAddedToCart`
- `OrderPlaced`
- `OrderCancelled`
- `ReturnRequested`

These events may later be published through the message bus.

---

## Application Layer Guidance

Each use case should be explicit and named by intent.

Suggested pattern:

- input DTO / command
- use-case handler
- output DTO / result

Example shape:

```text
application/use_cases/place_order.py
application/dto/place_order_input.py
application/dto/place_order_result.py
```

A use case typically performs:

1. validate input shape
2. load required aggregates from repositories
3. invoke domain behavior
4. persist updated aggregates
5. publish resulting messages/events
6. return output

Keep orchestration in use cases, not inside controllers/facades.

---

## Repository Strategy

Repositories are abstractions over persistence.

They should:

- load/save aggregates
- hide raw SQL from the rest of the system
- provide intention-revealing methods

They should avoid becoming query dumping grounds.

Examples:

- `get_by_id(...)`
- `get_open_cart_for_shopper(...)`
- `save(order)`

For reporting-style or exploratory reads, it is acceptable to create specialized read repositories if needed.

---

## External API Simulation

All external APIs should initially be replaced with **local fakes**.

These fakes should be:

- deterministic
- scriptable in tests
- inspectable
- simple to swap

Possible strategies:

- in-memory fake objects
- fixture-backed fake clients
- SQLite-backed integration state when useful

The point is not realism for its own sake. The point is stable experimentation.

---

## Message Bus Simulation

In this refinement environment, message sharing should be simulated in a simple and inspectable way.

### Recommendation

Use a **generic SQLite-backed message table** plus an **in-memory variant** for fast unit tests.

Avoid creating one table per topic initially.

### Why

A generic message table provides:

- persistence
- replayability
- inspectability
- simple debugging
- low schema friction
- closer alignment with outbox/event-log thinking

A table per topic is usually too structural for this stage.

### Suggested message bus implementation

```python
class MessageBus:
    def publish(self, topic: str, message: dict, headers: dict | None = None) -> None:
        ...
```

Consumption may be modeled by explicit polling/dispatching in tests or runners rather than hidden background processes.

### Suggested SQLite schema

```sql
create table messages (
    id text primary key,
    topic text not null,
    message_type text not null,
    payload_json text not null,
    headers_json text,
    status text not null,
    available_at text not null,
    processed_at text,
    retries integer not null default 0,
    error text,
    correlation_id text,
    causation_id text,
    created_at text not null
);

create index idx_messages_pending
on messages(topic, status, available_at);
```

### Initial semantics to support

Start simple. Support only:

- publish message
- fetch pending messages
- mark processed
- mark failed
- retry count
- optional delayed availability

Only add more if needed:

- duplicate delivery simulation
- dead-letter behavior
- visibility timeout approximation
- ordering guarantees

### Rule

The model should depend on the **message bus abstraction**.

---

## Testing Strategy

The project should use at least three testing levels.

### 1. Domain tests
Fast tests over domain objects and value objects.

Validate:

- invariants
- transitions
- calculations
- invalid state changes

These tests should avoid database and API access.

### 2. Use-case tests
Test use-case orchestration with fakes or in-memory adapters.

Validate:

- repository interaction at the contract level
- event/message emission
- cross-aggregate orchestration
- response/output behavior

### 3. Integration tests
Use SQLite and local fakes.

Validate:

- mapping between repository and model
- persistence/reload fidelity
- end-to-end workflow behavior
- message recording and dispatch

---

## Frontend-Derived Use Case Extraction

Frontend usage should be treated as a source of **intention discovery**, not as direct architecture.

Suggested extraction flow:

1. observe a frontend action
2. identify user intent
3. define a use case
4. identify domain rules involved
5. model entities/value objects/events
6. implement tests
7. implement use case and adapters

Example:

Frontend action:
- shopper submits a return request

Use case:
- `RequestReturn`

Questions to answer:

- which order states allow returns?
- which line items are eligible?
- what must remain immutable?
- what events/messages are emitted?
- what data is returned to the UI?

This keeps the system centered on domain meaning rather than screen structure.

---

## Error Handling

Prefer explicit failure modes.

Use:

- domain exceptions for invariant violations
- application-level result mapping when returning user-facing outcomes
- structured error records for message processing failures

Avoid silent failures and overly generic exception swallowing.

---

## Observability Inside the Lab

Even as a local project, the system should be inspectable.

Useful inspection points:

- persisted aggregates
- published messages
- fake API interactions
- domain events produced
- scenario logs for test runs

This project benefits more from **clarity of state transitions** than from production-style telemetry.

---

## What Not To Do

To preserve the project’s purpose, avoid the following by default:

- do not recreate production infrastructure in detail
- do not introduce microservices unless a strong reason emerges
- do not let repositories absorb business rules
- do not turn use cases into generic CRUD wrappers
- do not couple domain logic to frontend payload shapes

---

## Decision Heuristics

When unsure, prefer the option that is:

1. easier to inspect
2. easier to test
3. more explicit in naming and intent
4. less coupled to framework/infrastructure details
5. closer to business meaning than transport detail

---

## Initial Build Order

Recommended implementation order:

1. define core domain concepts and use-case names
2. define ports/interfaces
3. implement in-memory fakes
4. write domain and use-case tests
5. implement domain models
6. implement use cases
7. implement SQLite repositories
8. implement SQLite message bus
9. add API facade adapters
10. expand scenarios from frontend usage

---

## Summary

This repository currently uses a **DDD-inspired modular monolith for model refinement**.

Its main characteristics are:

- behavior-centered domain model
- explicit use cases
- ports and adapters
- SQLite-backed persistence
- local fake external integrations
- simple message bus simulation
- tests as executable specification

The goal is not to mirror the current system literally.

The goal is to create a cleaner executable model that captures the real business behavior with lower accidental complexity.
