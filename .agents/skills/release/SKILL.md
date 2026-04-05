---
name: release
description: Evaluate whether the current built slice is ready to be accepted as the intended internal version in the MRL loop. Use when implementation, tests, and EGD artifacts exist and a release decision or regression summary is needed without exposing the slice yet.
---

# Mission

Decide whether the current slice state should be accepted as the internal released version.

# Read First

- `docs/operating/skills_workflow.md`
- `docs/slices/<specific_slice>.md`
- relevant `implementation.md`
- relevant `egd.md`
- regression outputs and test results
- `decisions.md` when prior accepted constraints matter

# Inputs

- `docs/slices/<specific_slice>.md`
- implementation evidence
- test results
- EGD outputs
- regression outputs

# Must Do

- evaluate acceptance against the intended slice
- summarize regressions or meaningful diffs when present
- write `release_decision.md`
- write `regression_diff.md` when regression evidence exists
- make accept, reject, or return-to-loop decisions explicit

# Must Not Do

- do not reduce the phase to raw test execution only
- do not expose the slice automatically
- do not change code directly in this phase
- do not rely on hidden conversational context

# Outputs

- `work/changes/<id>/release_decision.md`
- optional `work/changes/<id>/regression_diff.md`
