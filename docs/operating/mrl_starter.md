# MRL Starter

## Purpose

This document defines the minimal portable shape for adopting the Model Refinement Lab in a new repository.

Use it when creating a fresh project that should follow the same refinement discipline without copying domain-specific content from this repository. The practical instantiation flow is to clone the starter into a new repository name and then let `extract` create the first semantic baseline.

---

## Split the Repository Into Two Parts

### MRL core

These parts should be portable across domains:

- `readme.md`
- `groundrules.md`
- `decisions.md`
- `AGENTS.md`
- `.agents/skills/`
- `docs/operating/`
- evaluator pattern
- semantic docs
- slice docs

### Pack-specific

These parts define language and architecture defaults and may be replaced:

- `architecture.md`
- `docs/building/project_structure.md`
- `docs/packs/`
- scenario runner pattern
- test shape
- code under `src/`
- tests under `tests/`
- language toolchain files such as `pyproject.toml`, `package.json`, or `go.mod`

### Project-specific

These parts should be rewritten for each new domain:

- `docs/semantics/model_hypothesis.md`
- `docs/semantics/domain_background_knowledge.md`
- `docs/slices/`
- local evaluation artifacts under `runs/`

---

## Minimal Starter Layout

```text
project_root/
  readme.md
  architecture.md
  groundrules.md
  decisions.md
  AGENTS.md
  .gitignore
  pyproject.toml

  .agents/
    skills/
      extract/
      refine/
      build/
      egd/
      release/
      expose/
      living/

  docs/
    operating/
      mrl_reference.md
      mrl_starter.md
      packs.md
      skills_workflow.md
    building/
      project_structure.md
    packs/
      python_ddd_monolith.md
    evaluation/
      scenario_evaluation.md
    semantics/
      model_hypothesis.md
      domain_background_knowledge.md
    slices/

  src/                        # shaped by the selected pack
  tests/                      # shaped by the selected pack
```

---

## First Bootstrap Sequence

A new repository should start with this order:

1. define root strategic docs and `docs/operating/`
2. choose and document the repository licensing approach before adding substantial project-specific code
3. install repo-local skills under `.agents/skills/`
4. select or define a pack and record it in `decisions.md`
5. keep the semantic docs as placeholders under `docs/semantics/`
6. run `extract` to build the first model baseline
7. run `refine` to create the first slice
8. run `build` to implement the first executable vertical slice
9. run `egd` before deciding `release`

## Licensing Guidance For Adopters

Treat licensing as an explicit adoption decision, not as an accidental consequence of cloning the starter.

The simplest path is to keep one license for the whole adopting repository. That is a good default when you do not need different sharing terms for process artifacts and implementation code.

When a project needs different terms, a split-license model is also valid. A practical pattern is:

- keep MRL process artifacts and reusable template-style material under a permissive license
- place project-specific implementation code and runtime assets under the license that fits the shipped product

Typical process or template-style material includes:

- `readme.md`, `groundrules.md`, `decisions.md`, and `AGENTS.md`
- `docs/operating/`
- generic pack definitions under `docs/packs/`
- generic skill instructions under `.agents/skills/`

Typical project-specific material includes:

- code under `src/`
- tests under `tests/`
- domain-specific semantic artifacts under `docs/semantics/`
- domain-specific slice documents under `docs/slices/`
- runtime assets and game or product content

If you adopt a split-license approach:

- rewrite the root `LICENSE` to explain which parts of the repository are covered by which license
- include the full text of each applicable license in a dedicated location such as `LICENSES/`
- update package metadata and release metadata to match the license that applies to the distributed code
- add header templates or file-level notices when the chosen code license requires clearer per-file scope
- record the reasoning and scope in `decisions.md`

The important rule is clarity. An adopting repository should make it easy for a contributor, user, or downstream maintainer to tell which repository content is governed by which license and why.

---

## Portability Test

MRL is portable only if a new domain can reuse the same loop and skills by changing semantic artifacts and slices rather than rewriting the workflow itself.

A practical test is:

- keep the operating docs and skills
- select or replace the implementation pack intentionally
- replace semantic docs and slice docs
- implement a new first slice
- run the same loop end to end

If that works, the starter is sufficiently portable.
