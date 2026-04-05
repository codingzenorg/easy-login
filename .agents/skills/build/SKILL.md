---
name: build
description: Implement one refined slice in code for the MRL loop. Use when a slice document exists and the next step is to change code, add tests, and validate deterministic behavior while respecting repository structure, architecture, and decisions.
---

# Mission

Implement one bounded slice from the repository artifacts.

# Read First

- `readme.md`
- `architecture.md`
- `groundrules.md`
- `decisions.md`
- `docs/building/project_structure.md`
- `docs/operating/skills_workflow.md`
- `docs/slices/<specific_slice>.md`

# Inputs

- `docs/slices/<specific_slice>.md`
- current repository code
- architectural and structural guidance

# Model Guidance

- prefer a Codex-class coding model for this skill
- if model selection is available, `gpt-5.3-codex` is a good default for implementation-heavy work
- use a smaller or faster model only for narrow, low-risk changes

# Must Do

- implement the slice in code
- add or update deterministic tests
- keep boundaries aligned with architecture and project structure
- write `implementation.md` when the change needs an explicit implementation artifact
- run relevant validation commands when feasible

# Must Not Do

- do not silently rewrite semantic docs unless the slice explicitly requires it
- do not invent missing business rules without surfacing the gap
- do not expand the slice scope just because nearby code is convenient to change
- do not use prior conversation as hidden context; use repository artifacts

# Outputs

- code changes
- test changes
- optional `work/changes/<id>/implementation.md`
- validation results
