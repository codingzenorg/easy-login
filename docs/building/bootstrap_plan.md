# Bootstrap Plan

## Purpose

This document defines the first implementation path for a new repository using MRL.

---

## Sequence

1. use `extract` to create the first semantic baseline
2. use `refine` to define the first slice in `docs/slices/`
3. use `build` to implement one executable vertical slice
4. use `egd` to review the built behavior against the model and background knowledge
5. make a `release` decision before any real exposure

---

## Rules

- start with one thin vertical slice
- avoid broad infrastructure setup before behavior exists
- keep tooling lightweight
- defer frameworks and integrations until a slice requires them
