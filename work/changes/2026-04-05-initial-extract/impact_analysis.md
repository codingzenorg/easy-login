# Impact Analysis

## Reason For Impact Note

The extracted model and the requested architecture reference now point toward a polyglot client/server shape, but the repository baseline in `architecture.md` still names `python_ddd_monolith` as the selected pack.

This affects refinement because the first slice should be executable without hiding that mismatch.

## Impacted Areas

- `architecture.md`
  currently describes a Python modular-monolith pack as the active repository shape
- future `build`
  needs to know whether to preserve the documented pack or follow the extracted Mithril plus Go direction
- future `expose`
  should preserve Docker Swarm readiness expectations such as clear health and readiness probes

## Refinement Judgment

The first slice can still be defined now because the business behavior is clear:

- create guest identity
- resume identity from device token

But the slice must explicitly declare runtime targets and architecture pressure so build does not accidentally drift back into the starter's Python default.

## Required Follow-Through

Before or during build, the repository should make one of these moves explicit:

1. update `architecture.md` and record a decision that the active implementation direction is now polyglot client/server
2. deliberately choose a temporary implementation path for this slice while documenting why that does not invalidate the extracted target architecture

## Boundaries To Preserve

- do not let the reference repository define the business model
- do not introduce claim or recovery behavior into the first slice
- do not defer identity continuity itself; that is the core behavior this slice exists to prove
