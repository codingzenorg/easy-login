# Skills Workflow

## Purpose

This document defines how skills should be used inside the Model Refinement Lab. Read `docs/operating/mrl_reference.md` first when a compact restatement of the model is needed.

The goal is to preserve phase isolation, reduce hidden assumptions, and keep model refinement reproducible. See `docs/operating/best_practices.md` for practical phase guidance such as when architecture decisions should move from `extract` to `refine`.

---

## Core Principle

> Each phase is an isolated loop executed by a bounded agent.

Do not rely on conversational continuity.
Do not mix responsibilities.
Do not carry implicit state across phases.

Instead:

> State lives in artifacts. Execution is stateless.

---

## Why Strong Isolation

One long-running session creates avoidable problems:

- extraction contaminates refinement
- refinement contaminates build decisions
- validation contaminates hypothesis generation
- assumptions become implicit instead of documented
- reproducibility degrades over time

Strong isolation improves:

- role clarity
- deterministic execution
- easier debugging
- cleaner refinement loops

---

## Repository Skills

Repository-scoped skills should live under `.agents/skills/` with one folder per loop phase.

```text
.agents/skills/
  extract/
    SKILL.md
  refine/
    SKILL.md
  build/
    SKILL.md
  egd/
    SKILL.md
  release/
    SKILL.md
  expose/
    SKILL.md
  living/
    SKILL.md
```

Each skill should remain phase-bounded and artifact-driven.

---

## New Repository Use

When bootstrapping MRL in a fresh repository, use `docs/operating/mrl_starter.md` as the portable baseline and then install the same repository-scoped skills under `.agents/skills/`.

---

## Execution Model

### Rule

> One skill = one isolated execution context.

Preferred execution:

```bash
codex run .agents/skills/extract/SKILL.md
codex run .agents/skills/refine/SKILL.md
codex run .agents/skills/build/SKILL.md
codex run .agents/skills/egd/SKILL.md
```

Lighter but still acceptable execution:

```bash
/new
/skill extract

/new
/skill refine

/new
/skill build

/new
/skill egd
```

---

## Recommended Agent Mapping

The table below defines the current recommended agent, context, inputs, outputs, and boundaries for the main executable phases.

| Phase | Recommended agent | Role | Ideal context | Main inputs | Main outputs | Must not |
| --- | --- | --- | --- | --- | --- | --- |
| `extract` | `gpt-5.4` `medium` | external-domain analyst | current `docs/semantics/model_hypothesis.md`, current `docs/semantics/domain_background_knowledge.md`, relevant external material | external code, external docs, stakeholder requests, runtime evidence | updated `docs/semantics/model_hypothesis.md`, updated `docs/semantics/domain_background_knowledge.md`, optional `request.md` | design implementation prematurely or write production code |
| `refine` | `gpt-5.4` `medium` | model-to-slice designer | `docs/semantics/model_hypothesis.md`, `architecture.md`, `groundrules.md`, `decisions.md` | `request.md`, semantic docs, relevant code | `docs/slices/<specific_slice>.md`, optional `impact_analysis.md` | write code or silently invent business rules |
| `build` | `gpt-5.3-codex` `medium` | slice implementer | `docs/building/project_structure.md`, `architecture.md`, `groundrules.md`, `decisions.md` | `docs/slices/<specific_slice>.md` | code, tests, optional `implementation.md` | redefine semantic docs unless the slice explicitly requires it |
| `egd` | `gpt-5.4` `medium` plus optional local `llama3` reviewer | expectation-gap reviewer | `docs/semantics/model_hypothesis.md`, `docs/semantics/domain_background_knowledge.md`, `docs/evaluation/scenario_evaluation.md` | `docs/slices/<specific_slice>.md`, scenario outputs, test outputs, implementation artifacts | `runs/<slice>/<timestamp>/...`, optional `egd.md` | modify code or patch behavior directly |
| `release` | `gpt-5.4` `medium` | acceptance and regression judge | slice definition, implementation summary, test results, EGD outputs, decisions | `docs/slices/<specific_slice>.md`, regression results, EGD report | `regression_diff.md`, `release_decision.md` | reduce the phase to raw test execution only |

Notes:

- `extract` should usually produce or update both semantic reference docs, not only one.
- `refine` is where slice design becomes explicit; `extract` should stop earlier.
- `egd` may use a smaller local model such as `llama3` for first-pass review, but interpretation and loop decisions should remain explicit.
- `release` should end with a decision, not only a diff artifact.

---

## MRL Phase Structure

### 1. Extract

**Goal**

- capture relevant signals from requests, systems, documents, behavior, and prior evidence
- identify the material that should enter the refinement loop
- make source evidence explicit

**Input**

- request material
- relevant docs
- existing code
- runtime evidence when available

**Output**

- `request.md`
- extracted notes or source evidence references

**Must not**

- define the solution prematurely
- write implementation code

### 2. Refine

**Goal**

- transform extracted signals into clearer model hypotheses
- evaluate fit with current boundaries and language
- identify impacted slices and likely decisions

**Input**

- `request.md`
- semantic docs
- architecture docs
- relevant code

**Output**

- `hypothesis.md`
- `impact_analysis.md`

**Must not**

- write production code
- silently redefine business rules

### 3. Build

**Goal**

- implement or change one slice based on the refined model
- respect domain boundaries
- validate deterministic behavior with tests

**Input**

- `hypothesis.md`
- `impact_analysis.md`
- repository code

**Output**

- code changes
- `implementation.md`
- test results

**Must not**

- invent missing business rules
- bypass the refined model without documenting why

### 4. EGD

**Goal**

- detect mismatch between expected and actual behavior
- identify plausible omissions, weak outcomes, or semantically incomplete behavior

**Input**

- `request.md`
- `hypothesis.md`
- `implementation.md`
- runtime or test outputs
- expectation review artifacts

**Output**

- `egd.md`
- recommendation: return to extract, return to refine, return to build, or continue

**Must not**

- modify code
- patch behavior directly

### 5. Release

**Goal**

- declare that the current model state is the accepted internal version
- package the outcome of the loop as the current intended state

**Input**

- implementation state
- test evidence
- EGD findings
- decision records

**Output**

- release decision
- accepted version state

**Must not**

- imply external exposure automatically
- substitute for EGD or refinement decisions

### 6. Expose

**Goal**

- put the released state into contact with a real context
- expose it to users, stakeholders, workflows, or environments

**Input**

- released version
- exposure plan or target context

**Output**

- exposure event
- initial real-world feedback signals

**Must not**

- be treated as long-term operation
- be confused with release itself

### 7. Living

**Goal**

- treat the exposed state as something that now exists in reality
- collect new evidence from friction, surprises, drift, and feedback
- feed those signals back into `extract`

**Input**

- exposed version
- real-world signals

**Output**

- new evidence for the next loop

**Must not**

- be confused with general operations ownership
- bypass the loop by making undocumented corrections

---

## Artifact-Centered Memory

> Artifacts are the only source of truth.

Recommended structure:

```text
work/changes/<id>/
  request.md
  hypothesis.md
  impact_analysis.md
  implementation.md
  egd.md
  decision.md
```

Each phase should:

- read explicit artifacts
- write explicit artifacts
- avoid depending on conversational memory

---

## Skill Design Rules

Each skill should define:

- mission
- explicit inputs
- explicit outputs
- allowed actions
- forbidden actions

Critical rule:

> Every skill must state what it must not do.

This is the mechanism that preserves isolation.

---

## Operational Modes

### Fast loop

Use this for ordinary iteration speed:

```bash
/new
/skill extract

/new
/skill refine

/new
/skill build

/new
/skill egd
```

### Strict loop

Use this when you need higher confidence or easier debugging:

```bash
codex run .agents/skills/extract/SKILL.md
codex run .agents/skills/refine/SKILL.md
codex run .agents/skills/build/SKILL.md
codex run .agents/skills/egd/SKILL.md
```

---

## Mental Model

Not:

> one assistant evolving over time

But:

> a pipeline of isolated phase workers operating on shared artifacts

```text
extract
  -> refine
  -> build
  -> egd
  -> release
  -> expose
  -> living
  -> extract
```

---

## Summary

Strong isolation plus artifact-driven execution is the preferred operating model for MRL.

MRL defines the loop.
Skills execute the loop.
Artifacts preserve the truth.
