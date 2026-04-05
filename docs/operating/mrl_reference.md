# MRL Reference

## Purpose

This document is the compact reference for the Model Refinement Lab.

Use it to rebuild context quickly in a fresh session, a bounded skill run, or a new repository that adopts the same workflow.

---

## What MRL Is

MRL is a repository-centered way to evolve software through explicit model refinement.

It treats business understanding, slice definition, code, evaluation, and delivery as separate but connected phases.

The goal is not to preserve one long assistant conversation. The goal is to preserve truth in repository artifacts.

---

## What MRL Is Not

MRL is not:

- a chat-driven development process
- a hidden-state agent workflow
- a framework that assumes microservices or production infrastructure first
- an operations model

MRL is a refinement loop.

---

## Canonical Loop

```text
extract -> refine -> build -> egd -> release -> expose -> living -> extract
```

### `extract`
Capture requests, external systems, docs, runtime evidence, and reference material into explicit artifacts.

### `refine`
Turn extracted signals into bounded model hypotheses and slice definitions.

### `build`
Implement one slice in code with deterministic tests.

### `egd`
Run expectation-gap detection against realistic execution evidence.

### `release`
Accept the current state as the intended internal version.

### `expose`
Put the released state into contact with a real context.

### `living`
Capture the feedback, friction, surprises, and drift created by exposure.

---

## Artifact Rule

> Artifacts are the source of truth. Execution is stateless.

Each phase should read explicit files, write explicit files, and avoid depending on prior conversational memory.

Typical artifact chain:

```text
request.md
  -> docs/semantics/model_hypothesis.md
  -> docs/semantics/domain_background_knowledge.md
  -> docs/slices/<slice>.md
  -> code + tests + implementation notes
  -> runs/<slice>/<timestamp>/...
  -> release_decision.md
```

---

## Repository Layers

Use the repository in layers:

- root strategic docs: overall intent, architecture, constraints, and decisions
- `docs/operating/`: the MRL model and skill workflow
- `docs/semantics/`: extracted and reference business meaning
- `docs/slices/`: one bounded slice definition per increment
- `src/` and `tests/`: executable implementation
- `runs/`: untracked evaluation evidence

---

## Skills and Isolation

Repository-scoped skills under `.agents/skills/` execute the loop in isolated contexts.

Each skill should define:

- mission
- explicit inputs
- explicit outputs
- allowed actions
- forbidden actions

Use a fresh session or isolated run per phase whenever practical.

---

## Portable MRL

To make MRL portable to a fresh repository, keep the framework generic and move domain-specific material into semantic docs and slice docs.

Use `docs/operating/mrl_starter.md` as the starter layout for new repositories. Use `docs/operating/best_practices.md` for practical guidance such as architecture timing and phase discipline.

MRL portability does not require one implementation language or one architecture.

The portable part is:

- the loop
- the artifact rules
- the skill isolation model
- the semantic and slice documents

Language, runtime topology, and architectural style should be selected separately through repository pack guidance.
