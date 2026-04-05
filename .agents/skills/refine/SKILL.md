---
name: refine
description: Refine extracted model signals into a concrete slice definition for the MRL loop. Use when a request and semantic artifacts exist and the next step is to shape one bounded slice, clarify model pressure, and write or update a slice document without editing implementation code.
---

# Mission

Turn extracted evidence and semantic hypotheses into one explicit slice definition.

# Read First

- `readme.md`
- `architecture.md`
- `groundrules.md`
- `decisions.md`
- `docs/operating/skills_workflow.md`
- `docs/semantics/model_hypothesis.md`
- relevant `work/changes/<id>/request.md` when available
- existing slice docs under `docs/slices/`

# Inputs

- `work/changes/<id>/request.md` when available
- `docs/semantics/model_hypothesis.md`
- semantic and architectural docs
- relevant code only when needed to understand current boundaries

# Model Guidance

- prefer a strong general reasoning model for this skill
- if model selection is available, use a frontier reasoning model rather than a speed-optimized coding model
- use a coding-focused model only when repository structure or nearby implementation details dominate the refinement work

# Must Do

- identify the bounded change to implement next
- explain model pressure and impacted boundaries when useful
- write or update `docs/slices/<specific_slice>.md`
- write or update `impact_analysis.md` when the change affects multiple areas
- keep the slice concrete enough for build to execute deterministically

# Must Not Do

- do not write production code
- do not silently invent business rules unsupported by available artifacts
- do not bypass architecture or ground rules without documenting the tension
- do not rely on conversational continuity; rebuild context from artifacts

# Outputs

- `docs/slices/<specific_slice>.md`
- optional `work/changes/<id>/impact_analysis.md`
