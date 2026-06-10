---
name: repo-research
description: Inspect a task-relevant code area, identify current behavior, and separate facts from unknowns before planning or implementation.
compatibility: opencode
metadata:
  domain: software-development
  stage: research
---

## What I do

- Identify entrypoints, nearby modules, callers, callees, tests, configs, and docs relevant to the task.
- Describe current behavior only from repository evidence.
- Separate facts, inferences, and unknowns.
- Call out invariants, risks, and weak test coverage that matter before code changes.

## When to use me

Use this when the request touches an unfamiliar part of the codebase, spans several files, or needs safe task-scoped understanding before planning or implementation.

## Working rules

1. Stay task-bounded. Inspect only the paths and dependencies needed for the request.
2. Stay grounded. Tie important claims to files, symbols, tests, configs, or docs.
3. Stop once the main flow, constraints, and risks are clear enough for planning.
4. Mark uncertainty plainly instead of guessing.

## Expected output

- Relevant files and why they matter.
- Relevant symbols and ownership boundaries.
- Current execution flow.
- Invariants that downstream work must preserve.
- Risks, unknowns, and missing coverage.
