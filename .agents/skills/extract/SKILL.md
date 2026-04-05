---
name: extract
description: Extract change signals, external evidence, and domain material into explicit repository artifacts for the MRL loop. Use when starting a new loop or when new external code, docs, stakeholder input, runtime evidence, or reference material must be converted into updated semantic artifacts without designing implementation or writing production code.
---

# Mission

Capture relevant external or newly observed signals and turn them into explicit repository artifacts.

# Read First

- `readme.md`
- `docs/operating/skills_workflow.md`
- `docs/semantics/model_hypothesis.md`
- `docs/semantics/domain_background_knowledge.md`
- relevant external artifacts provided for the change

# Inputs

- stakeholder requests
- external code
- external docs
- runtime evidence
- broad domain references

# Model Guidance

- prefer a strong general reasoning model for this skill
- optimize for synthesis accuracy and traceability over raw coding speed
- use a coding-focused model only when the extraction task is mostly about interpreting source code rather than stakeholder or domain material

# Must Do

- identify which external signals matter to the loop
- update or create `request.md` when the change needs a bounded request artifact
- update `docs/semantics/model_hypothesis.md` when the target model changes
- update `docs/semantics/domain_background_knowledge.md` when the reviewer needs broader domain context
- keep source evidence explicit and traceable

# Must Not Do

- do not write production code
- do not design slice implementation details prematurely
- do not silently redefine business rules without updating artifacts
- do not rely on prior conversation as memory; use repository artifacts and provided source material only

# Outputs

- updated `docs/semantics/model_hypothesis.md`
- updated `docs/semantics/domain_background_knowledge.md`
- optional `work/changes/<id>/request.md`
