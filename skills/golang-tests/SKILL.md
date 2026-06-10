---
name: golang-tests
description: Add or improve focused Go tests that verify changed behavior, preserved behavior, edge cases, and cleanup or concurrency paths where relevant.
compatibility: opencode
metadata:
  domain: software-development
  language: go
  stage: testing
---

## What I do

- Write or strengthen Go tests with the standard `testing` package unless the repo already uses a different established helper.
- Cover the changed behavior, important preserved behavior, and relevant failure or edge paths.
- Prefer deterministic tests with clear assertions and low mocking.
- Surface remaining testability gaps instead of hiding them.

## When to use me

Use this after a Go code change or when existing Go tests are weak around failure handling, concurrency, cleanup, or boundary conditions.

## Testing rules

1. Match repository test style and helpers.
2. Test behavior, not implementation details.
3. Avoid sleeps when synchronization or deterministic seams are available.
4. Keep scope tight: one strong regression test is better than many decorative tests.

## Expected output

- Test files and scenarios added or changed.
- Covered behaviors and preserved behaviors verified.
- Remaining gaps or flaky-risk areas.
