---
name: scoped-implementation
description: Implement a bounded code change with minimal surface area while preserving existing invariants and repository conventions.
compatibility: opencode
metadata:
  domain: software-development
  stage: implementation
---

## What I do

- Implement the smallest correct change that satisfies the accepted plan.
- Reuse existing helpers, naming, and patterns.
- Preserve current behavior outside the scoped change.
- Record changed files, decisions, and residual risks.

## When to use me

Use this when the task definition is clear and likely change points are already known.

## Implementation rules

1. Stay inside scope. No side refactors, broad renames, or speculative abstractions.
2. Prefer explicit logic over cleverness.
3. Preserve lifecycle, cleanup, locking, API, and error-handling invariants.
4. If the task cannot be completed safely within scope, stop and surface the blocker.

## Expected output

- Changed files and symbols.
- Key behavior changes.
- Preserved behaviors.
- Deviations from plan, if any.
- Risks that tests and review should inspect.
