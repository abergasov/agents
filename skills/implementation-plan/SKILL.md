---
name: implementation-plan
description: Turn a development request into a bounded execution plan with objective, scope, constraints, non-goals, and acceptance criteria.
compatibility: opencode
metadata:
  domain: software-development
  stage: planning
---

## What I do

- Convert a raw engineering request into a clear task definition.
- Define in-scope changes, out-of-scope changes, and constraints.
- Write acceptance criteria that can be reviewed or tested.
- Prefer the smallest safe change that solves the task.

## When to use me

Use this after research is sufficient and before implementation starts, especially when the task could drift or affect multiple files or behaviors.

## Planning rules

1. Do not plan from guesswork. Use current repository evidence.
2. Separate facts, assumptions, and decisions.
3. Prevent scope creep. Exclude unrelated cleanup and redesign unless required for safety.
4. Make the plan executable. Name likely change points, preserved behaviors, and validation targets.

## Expected output

- Objective and problem statement.
- In-scope and out-of-scope work.
- Constraints and assumptions.
- Acceptance criteria.
- Ordered workstreams for implementation, testing, and review.
