# Repository Guidelines

## Project Structure & Module Organization
This repository is an MRL starter. Strategic docs live at the root and working design material lives under `docs/`. Read `docs/operating/mrl_reference.md` and `docs/operating/skills_workflow.md` before changing the workflow.

Use this structure as the default shape:

```text
src/app/{domain,application,infrastructure,interfaces}
tests/{unit,integration,builders,fixtures}
docs/{operating,building,evaluation,semantics,slices}
.agents/skills/
```

Record structural deviations in `decisions.md`.

## Build, Test, and Development Commands
Keep tooling lightweight until the first slice exists.

- `pip install -e .` installs the editable package.
- `pytest` runs the test suite.
- `python -m src.app.interfaces.cli.run_scenario` is the preferred shape for a local scenario runner when one exists.

## Coding Style & Naming Conventions
Prefer Python 3.12+, explicit types, 4-space indentation, and business-oriented names. Use verb-driven use cases such as `PlaceOrder` and intention-revealing repositories such as `get_by_id` and `save`.

## Testing Guidelines
Use tests as specification. Start with domain tests, add integration tests for mappings and end-to-end flows, and keep time, IDs, and external responses deterministic.

## Commit & Pull Request Guidelines
Use Conventional Commits for commit subjects, choosing an appropriate type such as `feat`, `fix`, `docs`, `refactor`, `test`, `build`, `ci`, or `chore`. Keep commits scoped to one slice or doc change, and include test evidence in pull requests.
