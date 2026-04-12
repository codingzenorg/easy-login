# Session Metric Guidance

## Purpose

This document gives the repository a simple way to track session productivity during MRL work.

Use it as a lightweight operational hint, not as a scoring system. The goal is to notice whether sessions remain controlled, grounded, and productive over time.

---

## Primary Metric

* `output_efficiency = output_tokens / total_tokens`

### Interpretation

* ideal range: `5% - 10%`
* below `5%`: investigate context size, loop friction, or low progress per session
* above `15%`: investigate drift risk, weak grounding, or slice size

### Tracking Intent

* do not maximize output
* prefer input-heavy, output-controlled sessions
* optimize for small, controlled changes, consistent progress, and stable reasoning

### Operational Rule

* if `output_efficiency` stays between `5%` and `10%`, the session pattern is healthy
* if it goes outside that range, inspect context size or slice size before continuing in the same pattern

---

## Supporting Metrics

These metrics are optional but useful when the tooling exposes them.

* `reasoning_efficiency = reasoning_tokens / output_tokens`
* `cache_ratio = cached_tokens / input_tokens`

Interpret them conservatively:

* reasoning efficiency helps indicate whether the session is spending a reasonable amount of explicit reasoning relative to produced output
* cache ratio helps show whether the workflow is benefiting from stable repeated context

These numbers are hints, not targets.

---

## How To Use This In MRL

Track metrics per working session or per short delivery interval.

Review them when:

* a session feels expensive without much progress
* slice size may be too large
* context keeps growing across loops
* the team wants a lightweight historical view of delivery efficiency

Do not use these metrics to replace artifact review, commit review, or slice evaluation.
They are operational hints for the shape of the work, not truth about the value of the work.

---

## Session - 20260405

### Raw

* total_tokens: 875011
* input_tokens: 806238
* cached_tokens: 19539456
* output_tokens: 68773
* reasoning_tokens: 9130

### Derived

* output_efficiency: 7.86%
* reasoning_efficiency: 13.28%
* cache_ratio: 24.24x

### Productivity

* slices_applied: 4
* tokens_per_slice: 218752.75
* docs_commits: 10

### Notes

* historical measurement recorded from `easy-login`
* productivity is interpreted cumulatively from repository start through `2026-04-05`
* `git log --until 2026-04-05` shows four `feat:` commits in `easy-login`, so four delivered slices were counted for that cumulative window
* the same cumulative history shows ten `docs:` commits, so the measurement includes substantial repository-shaping and refinement work alongside slice delivery

---

## Session - 20260412

### Raw

* total_tokens: 1094613
* input_tokens: 1022960
* cached_tokens: 20830464
* output_tokens: 71653
* reasoning_tokens: 9575

### Derived

* output_efficiency: 6.55%
* reasoning_efficiency: 13.36%
* cache_ratio: 20.36x

### Productivity

* slices_applied: 4
* tokens_per_slice: 273653.25
* docs_commits: 11

### Notes

* historical measurement recorded from `easy-login`
* productivity is interpreted cumulatively from repository start through `2026-04-12`
* `git log --until 2026-04-12` still shows four `feat:` commits in `easy-login`, so the delivered slice count is unchanged from the previous snapshot
* the same cumulative history now shows eleven `docs:` commits, so this measurement reflects another documentation-oriented increment on top of the already delivered slices
