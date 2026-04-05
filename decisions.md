# Decisions Log

## Purpose

This document records relevant architectural and implementation decisions for the adopting repository.

Use it to preserve reasoning, avoid re-discussing settled trade-offs without context, and document deviations from `architecture.md` and `groundrules.md`.

---

## Entry Template

```md
## DEC-XXXX - Title

- Date: YYYY-MM-DD
- Status: proposed | accepted | superseded | rejected
- Owners: human | codex | both

### Context
What problem or tension led to this decision?

### Decision
What was decided?

### Consequences
What becomes easier, harder, or different because of this?

### Alternatives considered
What other options were considered and why were they not chosen?

### Notes
Any additional implementation guidance, migration note, or follow-up.
```

---

## Index

Add entries as the repository evolves.

## DEC-0001 - Separate MRL Core From Implementation Packs

- Date: 2026-03-29
- Status: accepted
- Owners: both

### Context
The starter was presenting Python plus a DDD-inspired modular monolith as if that were the default shape of MRL itself. That creates confusion when a repository needs another language, another architecture such as event sourcing, or more than one runtime.

### Decision
The repository now distinguishes between:

- MRL core, which stays artifact-driven and architecture-agnostic
- implementation packs, which define language, architecture, structure, and testing defaults

The current repository keeps `python_ddd_monolith` as the example selected pack.

### Consequences
It becomes easier to reuse the same refinement workflow across Python, JavaScript, Go, event-sourced, and polyglot client/server repositories. It also becomes necessary to make the selected pack explicit in architecture docs and slice docs.

### Alternatives considered
Keep one universal Python starter and treat every other shape as an undocumented deviation. This was rejected because it would keep conflating MRL with one implementation style.

### Notes
Future pack additions should live under `docs/packs/` and should be referenced by slice documents when the runtime topology matters.

## DEC-0002 - Treat Skill Model Guidance As Advisory

- Date: 2026-04-02
- Status: accepted
- Owners: both

### Context
The repository skills now include model guidance for tasks such as `build`, `refine`, and `extract`. That creates a potential ambiguity: a reader could assume that naming a preferred model in a skill will automatically switch the active Codex model or force sub-agent routing during execution.

### Decision
Model guidance inside repository skills is advisory only. It documents which model shape is usually a good fit for the task, but it does not by itself require automatic model switching, worker spawning, or hard routing behavior.

### Consequences
The skills remain durable even if model names, availability, or routing capabilities change. The repository gains clearer guidance for future operators and tooling, but predictable model selection still requires explicit runtime policy or orchestration outside the skill text.

### Alternatives considered
Encode specific model names in skills as if they were enforced execution rules. This was rejected because skill text alone does not guarantee runtime behavior and would overstate what the repository can currently control.

### Notes
If the repository later wants deterministic skill-to-model routing, document that as a separate operational decision and implement it in the calling workflow or agent orchestration layer.

## DEC-0003 - Avoid `.codex` Repository Artifacts For Now

- Date: 2026-04-02
- Status: accepted
- Owners: both

### Context
The repository has used `.codex`-specific artifacts while shaping the workflow. That creates a risk that the MRL process starts to look tool-defined rather than process-defined before it is clear whether MRL can stay tool-agnostic in practice.

### Decision
For now, the repository should avoid relying on `.codex` as part of the committed process shape. The workflow should stay centered on MRL artifacts and repository documents rather than tool-specific folders. This decision is explicitly reversible while the team validates whether MRL can remain tool-agnostic or whether a specific AI tool will prove operationally necessary.

### Consequences
The repository stays cleaner and more focused on portable process artifacts. It may require some extra translation when using Codex or another tool because fewer tool-native conventions are captured directly in versioned files. The decision also preserves room to adopt tool-specific support later if real usage shows that generic artifacts are not enough.

### Alternatives considered
Keep `.codex` as a first-class part of the repository workflow immediately. This was rejected for now because it would prematurely optimize for one tool before validating whether that coupling is necessary for MRL.

### Notes
If future evidence shows that MRL depends on stable capabilities from a specific AI tool, record a follow-up decision describing the required coupling, why generic artifacts were insufficient, and which tool-specific assets should become part of the repository.

## DEC-0004 - License MRL-Created Project Artifacts Under MPL-2.0

- Date: 2026-04-05
- Status: accepted
- Owners: both

### Context
This repository starts from the MRL starter, but it is now producing project-specific artifacts for `easy-login`. The repository needs to keep the reusable MRL framework and starter material distinct from the artifacts created by the MRL loop for this project.

### Decision
Use a split-license model:

- keep reusable MRL framework and starter material under `MIT`
- license project-specific artifacts created through the `easy-login` MRL process under `MPL-2.0`

The root `LICENSE` defines the scope. The full license texts live in `LICENSES/MIT.txt` and `LICENSES/MPL-2.0.txt`.

### Consequences
The repository preserves the reusable status of the MRL framework while making project-specific semantic artifacts, slices, code, tests, and other produced assets file-level copyleft under `MPL-2.0`. Future contributors need to keep the scope boundary clear when adding new files.

### Alternatives considered
Keep the whole repository under one permissive license. This was rejected because it would blur the distinction between reusable MRL process material and project artifacts produced for `easy-login`.

### Notes
When a future file is clearly part of the `easy-login` artifact set, prefer `MPL-2.0`. When it is clearly reusable MRL framework or starter material, keep it under `MIT` unless a later decision says otherwise.
