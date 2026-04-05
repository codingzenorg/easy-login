# Scenario Evaluation

## Purpose

This document defines how the project should evaluate implemented behavior beyond classic unit and integration correctness.

Its goal is to support **expectation gap detection**:

- identify plausible domain expectations that are absent
- detect awkward or semantically weak behavior
- surface possible omissions that deterministic tests do not catch
- support refinement of models, use cases, and API contracts

This evaluation is **not** a replacement for tests.

It is a complementary review loop.

---

## What This Is

Scenario evaluation is a process where:

1. a deterministic runner executes a realistic workflow
2. the system records what happened
3. a local LLM reviews the evidence
4. the LLM reports possible expectation gaps and review questions
5. a human and/or Codex decides whether refinement is needed

This is not traditional pass/fail testing.

It is closer to:

- domain expectation review
- semantic completeness review
- behavior plausibility review
- expectation gap detection

---

## What This Is Not

This process is **not** intended to:

- replace unit tests
- replace integration tests
- blindly trust the LLM as a domain authority
- declare final truth about the business
- invent requirements without review

The LLM should act as a **reviewer that raises plausible hypotheses**, not as a final judge.

---

## Core Distinction

### Correctness tests ask

- did the model preserve invariants?
- did the use case return the expected result?
- was the state transition valid?
- were the required events/messages emitted?

### Scenario evaluation asks

- does this behavior feel complete for the domain?
- what would a user or business stakeholder plausibly expect here?
- what is notably absent?
- what is technically valid but semantically weak?
- what should be reviewed as a possible omission or design gap?

This distinction is essential.

---

## Example

A shopper places an order successfully.

Deterministic tests may confirm:

- the order was placed
- cart state changed correctly
- a message was emitted
- the API returned success

But an expectation reviewer may still observe:

- no order confirmation was produced
- no shopper-visible summary artifact exists
- no follow-up notification seems present

This does **not** automatically mean the implementation is wrong.

It means:

- a plausible expectation may be missing
- the omission may be intentional or accidental
- the design should be reviewed explicitly

---

## Evaluation Philosophy

Scenario probing is part of the normal refinement loop for this lab.

It may be used at every iteration, but it only becomes meaningful once at least one executable scenario exists.

The practical prerequisite is therefore simple:

- at least one implemented vertical slice
- at least one scenario runner flow that produces evidence for review

The LLM reviewer should behave as a:

- domain-informed critic
- expectation-gap detector
- semantic reviewer
- source of review questions

It should **not** behave as a:

- strict correctness oracle
- product manager with unilateral authority
- hallucinating requirements engine

Findings should be treated as:

- hypotheses
- suspicions
- prompts for refinement discussion

---

## Recommended Evaluation Architecture

Use a **two-step architecture**.

### Step 1 - Deterministic scenario execution
A Python runner executes a scenario against the local system.

This runner should:

- set up initial state
- invoke API/use-case calls
- capture requests and responses
- capture state snapshots
- capture emitted events/messages/artifacts
- store all evidence in machine-readable files

### Step 2 - LLM expectation review
A local Ollama model reviews the evidence and produces a structured report.

Recommended starting point:

- begin with one general-purpose local model such as `llama3`
- optimize for review quality, prompt-following, and stable summarization rather than coding ability
- treat the first model choice as provisional and replace it if review quality is weak

This reviewer should:

- infer plausible expectations
- identify likely gaps
- distinguish between certainty and suspicion
- propose review questions

This separation is strongly preferred over a freeform agent doing everything.

---

## Minimum Setup

The minimum setup for scenario evaluation is:

- one implemented scenario runner
- one scenario definition
- one machine-readable evidence packet
- one local Ollama-served review model

There is no stricter hardware requirement defined in this repository yet.

The current working assumption is a local machine capable of running Ollama comfortably, for example a Ryzen-class workstation with 32 GB RAM.

That is a practical starting point, not a hard requirement.

---

## Why This Architecture

This design gives:

- reproducibility
- inspectability
- lower noise
- clearer evidence trail
- easier debugging
- easier comparison across model revisions

It also avoids giving too much responsibility to an unconstrained LLM.

---

## Recommended Input Package For Review

Each evaluation should be based on a compact structured packet.

### 1. Domain context
A short explanation of the bounded context, concepts, and important expectations.

Examples:

- what this part of the business does
- who the main actors are
- what common expectations usually exist
- what constraints matter

### 2. Contract summary
A distilled summary of the relevant use case or API contract.

Include:

- operation name
- inputs
- outputs
- expected known side effects
- emitted messages/artifacts if already modeled

### 3. Scenario definition
The scenario being executed.

Include:

- scenario name
- business purpose
- initial assumptions
- requested action sequence

### 4. Execution transcript
What actually happened during execution.

Include:

- requests
- responses
- errors if any
- timestamps if useful

### 5. Final state snapshot
Observed relevant final business state.

### 6. Emitted artifacts
Anything produced by the system.

Examples:

- domain events
- integration messages
- receipts
- notifications
- audit records
- confirmation objects

### 7. Review instructions
Clear instructions telling the LLM what kind of review to perform.

---

## Folder Structure Suggestion

```text
scenarios/
  place_order/
    domain_context.md
    contract_summary.json
    scenario.md
    input.json
    expected_notes.md

runs/
  place_order/
    2026-03-21T10-30-00/
      transcript.json
      final_state.json
      artifacts.json
      evaluation_packet.json
      evaluation_report.json
      evaluation_report.md
```

### Notes

- `scenarios/` contains reusable definitions
- `runs/` contains evidence and reports for concrete executions
- timestamps or run IDs make comparisons easier

---

## Scenario Types

It is useful to classify scenarios.

### 1. Happy-path scenario
Normal business success flow.

Purpose:

- evaluate completeness of ordinary behavior
- detect missing confirmations, acknowledgments, artifacts

### 2. Edge-case scenario
Rare but plausible conditions.

Purpose:

- reveal awkward outputs
- surface domain gaps around unusual inputs

### 3. Failure scenario
Intentional invalid or rejected flow.

Purpose:

- evaluate clarity of failure semantics
- detect missing error explanation or audit behavior

### 4. Ambiguous scenario
A situation where the model may still be underspecified.

Purpose:

- provoke refinement discussion
- reveal unresolved business decisions

---

## What The Reviewer Should Look For

The reviewer should be asked to inspect the scenario along a small stable set of dimensions.

### A. Missing artifact
Was something plausibly expected but absent?

Examples:

- receipt
- acknowledgment
- confirmation message
- audit record
- follow-up status object

### B. Missing feedback
Did the system technically succeed or fail, but leave the caller with weak or incomplete information?

### C. Missing downstream consequence
Did the flow imply a consequence that did not appear?

Examples:

- status not updated
- no follow-up event
- no visible outcome for a meaningful action

### D. Naming or contract mismatch
Does the API contract or output language seem misaligned with domain meaning?

### E. Suspicious flow design
Does the flow appear awkward, incomplete, or unexpectedly silent?

---

## What The Reviewer Should Not Do

The reviewer should avoid:

- declaring unsupported business facts
- assuming every common industry behavior applies here
- inventing hidden implementation details
- treating absence as proof of defect
- generating long generic advice disconnected from the scenario

The reviewer must stay grounded in the supplied packet.

---

## Reviewer Output Style

The reviewer should return **structured JSON**.

This is strongly preferred over freeform prose because it allows:

- machine processing
- report comparison
- filtering by severity/type
- stable rendering into markdown

---

## Recommended JSON Schema

```json
{
  "type": "object",
  "properties": {
    "scenario_name": {"type": "string"},
    "overall_assessment": {
      "type": "string",
      "enum": ["aligned", "mostly_aligned", "unclear", "concerning"]
    },
    "gaps": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "title": {"type": "string"},
          "severity": {
            "type": "string",
            "enum": ["low", "medium", "high"]
          },
          "type": {
            "type": "string",
            "enum": [
              "missing_artifact",
              "missing_feedback",
              "missing_downstream_consequence",
              "naming_mismatch",
              "suspicious_flow",
              "possible_rule_gap"
            ]
          },
          "observation": {"type": "string"},
          "why_it_might_be_expected": {"type": "string"},
          "confidence": {"type": "number"},
          "review_question": {"type": "string"}
        },
        "required": [
          "title",
          "severity",
          "type",
          "observation",
          "why_it_might_be_expected",
          "confidence",
          "review_question"
        ]
      }
    },
    "surprising_presences": {
      "type": "array",
      "items": {
        "type": "object",
        "properties": {
          "observation": {"type": "string"},
          "why_surprising": {"type": "string"}
        },
        "required": ["observation", "why_surprising"]
      }
    },
    "intentional_omissions_to_confirm": {
      "type": "array",
      "items": {"type": "string"}
    },
    "recommended_human_review": {
      "type": "array",
      "items": {"type": "string"}
    }
  },
  "required": [
    "scenario_name",
    "overall_assessment",
    "gaps",
    "surprising_presences",
    "intentional_omissions_to_confirm",
    "recommended_human_review"
  ]
}
```

---

## Output Interpretation Rules

### A gap is not automatically a defect
A gap means:

- this may be missing
- this may be intentional
- this should be reviewed

### Confidence is not certainty
The model’s confidence expresses how plausible the expectation appears given the scenario packet.

It should not be treated as proof.

### Review questions are the most valuable field
The system should prefer review questions over aggressive conclusions.

A good question is often more useful than a strong claim.

---

## Reviewer Prompt Philosophy

The prompt should keep the LLM in a narrow role.

### Good role framing

- review this scenario as a domain expectation analyst
- identify plausible expectation gaps
- do not decide final truth
- stay grounded in the evidence
- return hypotheses, not verdicts

### Bad role framing

- test whether the system is correct
- invent all missing requirements
- judge the product as if you are the business owner

---

## Recommended System Prompt

```text
You are a domain expectation gap reviewer.

Your job is to review a completed scenario execution and identify plausible expectation gaps.

Do not act as a correctness verifier.
Do not assume every missing behavior is a defect.
Treat your findings as review hypotheses, not final truths.

A gap is something that:
- a user, operator, or business stakeholder would plausibly expect
- is commonly implied by this kind of workflow
- seems notably absent, under-expressed, or awkward

Focus on:
- missing confirmations
- missing artifacts
- missing downstream consequences
- confusing or weak outputs
- domain-language mismatch
- suspiciously incomplete flows

Stay grounded in the provided domain context, contract summary, transcript, final state, and emitted artifacts.

Return only valid JSON matching the provided schema.
```

---

## Recommended Review Packet Shape

The reviewer input may be assembled as a JSON object like this:

```json
{
  "domain_context": "...",
  "contract_summary": {
    "operation": "PlaceOrder",
    "inputs": ["cart_id", "placed_at"],
    "outputs": ["order_id", "status"],
    "known_side_effects": ["cart marked as checked_out"]
  },
  "scenario": {
    "name": "place_order_happy_path",
    "purpose": "shopper places a valid order from an open cart"
  },
  "transcript": [
    {
      "step": 1,
      "request": {"cart_id": "cart-1"},
      "response": {"status": "success", "order_id": "ord-1"}
    }
  ],
  "final_state": {
    "cart_status": "checked_out",
    "order_status": "placed"
  },
  "artifacts": {
    "events": ["OrderPlaced"],
    "messages": ["order.placed"],
    "documents": []
  }
}
```

---

## Runner Responsibilities

The deterministic Python runner should:

- load scenario inputs
- prepare initial local state
- execute actions against the use-case/API boundary
- capture transcript
- capture final state
- capture artifacts
- persist all run files
- optionally call the reviewer and persist the review report

This runner should remain explicit and inspectable.

---

## Suggested Python Harness Shape

A simple internal structure could be:

```text
src/app/interfaces/cli/run_scenario.py
src/app/application/services/scenario_runner.py
src/app/application/services/scenario_evaluator.py
src/app/infrastructure/fakes/...
src/app/infrastructure/sqlite/...
```

### Possible service responsibilities

#### `scenario_runner.py`
- execute scenario deterministically
- write transcript and final state

#### `scenario_evaluator.py`
- assemble evaluation packet
- call Ollama
- validate reviewer JSON
- save report
- optionally render markdown summary

---

## Suggested Scenario Runner Flow

1. load scenario definition
2. set up initial DB/fakes/state
3. invoke the use case or facade
4. capture request/response transcript
5. capture relevant final state snapshot
6. collect emitted events/messages/artifacts
7. build evaluation packet
8. call local Ollama reviewer
9. validate structured JSON response
10. store report
11. optionally render markdown review

---

## Suggested Report Rendering

A markdown rendering can make reports easier to inspect.

Example shape:

```text
# Evaluation Report - place_order_happy_path

Overall assessment: mostly_aligned

## Gaps
- Missing artifact: no order confirmation produced after successful checkout
  - Severity: medium
  - Confidence: 0.61
  - Review question: should this domain generate an order confirmation, acknowledgment, or intentionally neither?

## Surprising presences
- None

## Intentional omissions to confirm
- Order confirmation may belong to another subsystem

## Recommended human review
- Confirm whether successful checkout must produce a shopper-visible confirmation artifact
```

This markdown is secondary. The JSON report remains the source of truth.

---

## Triage Of Findings

After a report is produced, each finding should be classified.

### 1. Required missing behavior
Outcome:

- add use case support
- add event/message
- add artifact
- add state transition
- add deterministic tests

### 2. Optional improvement
Outcome:

- backlog item
- UX improvement
- future design enhancement

### 3. Intentional omission
Outcome:

- document explicitly in `decisions.md`
- keep implementation unchanged

### 4. False positive / weak suggestion
Outcome:

- ignore
- optionally refine prompt or domain context

---

## Calibration Strategy

The reviewer should be calibrated with a small scenario set.

Recommended benchmark categories:

- scenarios where a meaningful omission should be detected
- scenarios where omission is intentional and should only raise a question
- scenarios that are fully adequate and should produce low noise

Calibration goals:

- reduce false positives
- reduce vague generic commentary
- improve review-question quality
- improve alignment with the actual domain

---

## Local Ollama Guidance

For this use case, the local Ollama model should be configured as:

- instruction-following
- low-temperature
- structured-output constrained
- evidence-grounded
- narrow in role

Do not use it as a highly autonomous open-ended agent in the first version.

Prefer:

- deterministic runner
- evidence packet
- structured review

This is simpler and more trustworthy.

---

## Evolution Path

### Version 1
Deterministic runner + one-shot structured review.

### Version 2
Reviewer can request additional evidence if needed.

### Version 3
Reviewer participates in constrained exploratory probing.

Start with Version 1.

---

## Decision Rule

If a recurring finding appears across multiple scenarios, review whether it should become:

- a new use case
- a new domain event
- a new artifact/output
- a stronger invariant
- an explicit intentional omission recorded in `decisions.md`

This keeps the evaluation loop connected to actual model refinement.

---

## Summary

Scenario evaluation exists to detect **expectation gaps**, not only correctness failures.

The healthy split is:

- tests answer: **is the model internally correct?**
- scenario evaluation answers: **does the model seem behaviorally complete and domain-plausible?**

The local Ollama reviewer should therefore be treated as:

- a hypothesis generator
- a semantic reviewer
- a detector of possible omissions

Its findings must be reviewed by humans and/or Codex before changing the model.

