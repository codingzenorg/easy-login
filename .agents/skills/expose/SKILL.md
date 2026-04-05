---
name: expose
description: Prepare or record exposure of a released slice to a real context in the MRL loop. Use when the internal released state is accepted and the next step is to define or capture how that state is exposed to users, stakeholders, workflows, or environments without treating this as long-term operations ownership.
---

# Mission

Define or record how a released slice is put into contact with reality.

# Read First

- `docs/operating/skills_workflow.md`
- `work/changes/<id>/release_decision.md`
- relevant slice and implementation artifacts

# Inputs

- released version state
- exposure target or context
- release decision

# Must Do

- describe the exposure target clearly
- record the exposure event or plan
- make the expected feedback channels explicit

# Must Not Do

- do not treat this as general operations ownership
- do not confuse exposure with release itself
- do not silently mutate the model during exposure planning

# Outputs

- optional `work/changes/<id>/exposure.md`
- exposure event or exposure plan
