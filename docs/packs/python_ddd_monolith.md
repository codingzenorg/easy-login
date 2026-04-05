# Pack: python_ddd_monolith

## Purpose

This pack defines the current example implementation default for the starter repository.

Use it when the model fits a single Python codebase with layered boundaries and deterministic local execution.

---

## Shape

- language: Python
- runtime topology: single runtime
- architecture mode: DDD-inspired modular monolith

---

## Layout

```text
src/
  app/
    domain/
    application/
    infrastructure/
    interfaces/

tests/
  unit/
  integration/
  builders/
  fixtures/
```

---

## Good Fit

- one main runtime
- strong use-case orchestration
- rich domain models
- explicit ports and adapters
- deterministic local simulation

---

## Less Suitable

- separate client and server implementations with independent models
- event stores and projections as the primary persistence model
- repositories where Go or TypeScript are the main execution environments

---

## Notes

This pack is an example default, not the definition of MRL itself.
