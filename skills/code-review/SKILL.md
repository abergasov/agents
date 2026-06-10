---
name: code-review
description: Review a scoped code change for correctness, hidden risk, and test sufficiency without drifting into style-only feedback.
compatibility: opencode
metadata:
  domain: software-development
  stage: review
---

## What I do

- Review changed code against the task objective, constraints, and acceptance criteria.
- Prioritize correctness, safety, cleanup, concurrency, and maintainability.
- Evaluate whether tests prove the changed and preserved behavior.
- Produce exact findings with requested fixes and explicit residual risks.

## When to use me

Use this after implementation and focused tests are ready.

## Review rules

1. Review against the task, not personal taste.
2. Ground findings in changed code, affected code, tests, and validation results.
3. Distinguish blockers from non-blocking notes.
4. Do not approve code whose correctness depends on weak assumptions or missing high-signal tests.

## Expected output

- Verdict: approved or changes required.
- Findings with severity, location, why it matters, and requested change.
- Test sufficiency assessment.
- Acceptance criteria check.
- Residual risks that remain acceptable.
