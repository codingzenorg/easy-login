---
name: egd
description: Run expectation gap detection for a built slice in the MRL loop. Use when code and scenario outputs exist and the next step is to compare built behavior against the model hypothesis and broad domain background knowledge without modifying code.
---

# Mission

Detect expectation gaps between built behavior and the semantic understanding of the domain.

# Read First

- `docs/operating/skills_workflow.md`
- `docs/evaluation/scenario_evaluation.md`
- `docs/semantics/model_hypothesis.md`
- `docs/semantics/domain_background_knowledge.md`
- `docs/slices/<specific_slice>.md`
- relevant implementation and scenario artifacts

# Inputs

- `docs/slices/<specific_slice>.md`
- runtime outputs
- scenario transcripts
- implementation artifacts
- expectation-review packets and reports

# Must Do

- compare expected behavior with observed behavior
- identify plausible omissions, weak outcomes, and review questions
- write or update `egd.md` when a human-readable summary is needed
- persist run artifacts under `runs/<slice>/<timestamp>/` when the evaluator is used
- make return-to-loop recommendations explicit

# Must Not Do

- do not modify code
- do not patch behavior directly
- do not treat the reviewer as domain truth
- do not rely on prior conversation instead of artifacts

# Outputs

- `runs/<slice>/<timestamp>/...`
- optional `work/changes/<id>/egd.md`
- recommendation to return to `extract`, `refine`, `build`, or continue
