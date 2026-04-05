# MRL Starter

## Purpose

This repository is a starter template for the Model Refinement Lab.

Use it to begin a new repository that follows the MRL loop:

```text
extract -> refine -> build -> egd -> release -> expose -> living -> extract
```

This starter separates:

- MRL core: the artifact-driven refinement loop and workflow
- implementation packs: selectable language and architecture defaults

The repository currently includes `python_ddd_monolith` as the example selected pack. If you want another shape such as Go, TypeScript, event sourcing, or a multi-runtime client/server system, keep the MRL core and replace the pack guidance intentionally.

One of the repository's main characteristics is that pre-defined packs can be used to scaffold new MRL projects. That lets adopters keep the same refinement workflow while starting from a language and architecture shape that already fits their context.

MRL itself is AI-agent-friendly rather than tool-specific, but this starter is currently optimized for Codex-style repository workflows. It can still be used well with other AI coding agents when they operate from explicit repository artifacts instead of conversational memory.

---

## Getting Started

Clone the starter into the new repository name you want:

```bash
git clone <starter-url> my-new-project
cd my-new-project
```

The preferred way to start a new project is to use this repository as a GitHub template, not as a fork. Forking is more appropriate when working on the starter itself or maintaining a derivative starter.

Then read these files first:

- `docs/operating/mrl_reference.md`
- `docs/operating/skills_workflow.md`
- `docs/operating/packs.md`
- `docs/operating/best_practices.md`
- `architecture.md`
- `groundrules.md`
- `docs/building/project_structure.md`

Before writing substantial project-specific code, decide whether the adopting repository will keep the starter's default license or intentionally replace it with another license or a split-license model. Record that choice clearly in the root `LICENSE` and `decisions.md`.

Then start the loop:

1. run `extract` to build the first semantic baseline in `docs/semantics/`
2. run `refine` to define the first slice in `docs/slices/`
3. run `build` to implement one vertical slice
4. run `egd` to review the built behavior

The semantic placeholders in `docs/semantics/` are intentionally empty. They are meant to be filled by the `extract` phase, not by copying domain content from this starter.
Preserve original evidence in `work/sources/` before extraction or refinement artifacts are produced elsewhere.

---

## Starter Layout

```text
.agents/skills/            # repo-local MRL skills
/docs/operating/           # MRL model and workflow docs
/docs/packs/               # implementation pack definitions
/docs/building/            # structure and bootstrap guidance
/docs/evaluation/          # expectation-gap evaluation guidance
/docs/semantics/           # domain-specific meaning created by extract
/docs/slices/              # one slice document per increment
/work/sources/             # canonical folder for curated raw evidence and original source material
/work/changes/             # request, impact, and implementation artifacts
/src/                      # implementation root shaped by the selected pack
/tests/                    # executable specification
```

---

## Notes

- Treat this repository as a template, not as a finished application.
- Keep domain specifics out of the starter and in the adopting repository.
- Keep the MRL loop generic and move language or architecture assumptions into packs.
- Make licensing an explicit adoption choice; the starter can be reused under one license or adapted into a split-license repository when process material and implementation code need different terms.
- Treat `work/` as repository memory, not scratch space; preserve original evidence in `work/sources/` before downstream artifacts are created.
- Prefer one small slice over broad scaffolding.
