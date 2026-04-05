# Pack: <pack_name>

## Purpose

Use this pack when <the kind of model or system it fits>.

This pack exists to define reusable implementation defaults for MRL repositories that share the same language, architecture, and runtime shape.

---

## Shape

- language: <ruby | go | typescript | multiple>
- runtime topology: <single runtime | client/server | multiple runtimes>
- architecture mode: <ddd monolith | service | event sourcing | rails monolith | etc>

---

## Good Fit

- <situation 1>
- <situation 2>
- <situation 3>

---

## Less Suitable

- <situation where this pack is a bad fit>
- <another mismatch>
- <another mismatch>

---

## Layout

```text
src or app/
  <main folders>

tests or spec/
  <test folders>

scripts/
  <optional runners>
```

---

## Boundaries

- business rules live in <where>
- orchestration lives in <where>
- persistence concerns live in <where>
- interface or transport concerns live in <where>
- shared contracts live in <where, if applicable>

---

## Testing Defaults

- unit/spec focus: <what gets tested first>
- integration focus: <what gets tested together>
- determinism rules: <time, IDs, external responses, event order>
- preferred test command: `<command>`

---

## Scenario Runner Shape

- preferred runner location: `<path>`
- preferred command: `<command>`
- runner responsibility: <what the scenario runner is supposed to do>
- runner must not: <what it should avoid>

---

## Runtime Targeting Rules

- slices using this pack should declare runtime targets as:
  - `<runtime 1>`
  - `<runtime 2>` if needed
- cross-runtime contracts should live in <where>
- event/message boundaries should be expressed in <where>

---

## Naming Conventions

- use cases or commands: `<style>`
- domain objects: `<style>`
- repositories/gateways: `<style>`
- events: `<style>`

---

## Rules

- <rule 1>
- <rule 2>
- <rule 3>
- <rule 4>

---

## Starter Tooling

- package manager: `<bundler | go modules | npm | etc>`
- formatter/linter: `<tools>`
- test framework: `<tools>`
- persistence default: `<sqlite | postgres | in-memory | event store>`

---

## Notes

This pack is a reusable default, not a law. If a repository diverges meaningfully, record that in `decisions.md`.
