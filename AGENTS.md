# AGENTS.md — Shared baseline

These rules apply to every agent. Role-specific method lives in each agent's skill.

## Engineering rules

1. **Stay inside scope.** Before editing, identify the requested behavior, the likely touched packages, and the non-goals. Change only what the task requires. No side refactors, unrelated renames, file moves "because it's cleaner," or architecture "improvements" the task didn't ask for. If a correct fix genuinely needs broader changes, say so before editing.

2. **Respect existing conventions.** Look at the nearest production code and tests in the same package first, and match the established patterns: naming, package layout, error handling, logging (`logger.With*`), context plumbing, tests, configuration access through `AppConfig` getters/runtime helpers, dependency injection, concurrency control. Don't introduce a new style when an existing one already solves the problem.

3. **Prefer minimal change.** Choose the smallest change that solves the problem — modify existing flow, reuse helpers, extend types carefully, delete complexity. Avoid speculative abstractions, new layers, generic helpers "for future use," and broad rewrites.

4. **Make code readable.** Clear names, short control flow, explicit error handling, narrow responsibilities. Avoid clever compactness, hidden side effects, unnecessary indirection, dense nesting.

5. **Preserve invariants.** Before changing behavior, identify invariants in the touched path — protocol assumptions, lifecycle and cleanup guarantees, locking discipline, context cancellation semantics, persistence assumptions, existing API behavior. If an invariant must change, state the old behavior, the new behavior, and why.

6. **Required vs. optional.** Problems you notice nearby are not part of this task. Report them as risks or follow-ups; don't expand scope on your own.

7. **Be honest about uncertainty.** If code behavior is unclear or conflicts with research, stop and report the conflict. Don't patch blindly, and don't treat compiling or passing lint as proof of correctness.

8. **Don't grow the surface.** No new module dependencies, top-level packages, or public APIs unless the task requires them. Prefer existing packages, config fields, and interfaces.

## Comments

Default to no comment. Code says *what*; a comment exists only to say *why* when the code can't. Trivial functions, obvious logic, and typical patterns get none.

When a comment earns its place, make it laconic, sharp, and valuable:

- Explain intent, a non-obvious constraint, a tradeoff, or a "why not the obvious way" — never restate the code.
- One line where one line does. No preamble, no hedging, no decoration.
- Delete on sight: comments that narrate the next line, banner/section dividers, commented-out code, redundant doc-comments that echo the signature, and TODOs without an owner or reason.
- A comment that's wrong or stale is worse than none. If you change code, fix or remove its comment.

If you feel the urge to comment to make a block understandable, first try a better name or a smaller function. Reach for the comment only when the *why* still won't fit in the code.

## Response style

Be exact, concise, and technical. Don't pad, flatter, moralize, sell the code, claim elegance, or hide weak points. Use plain engineering language, and report uncertainty plainly rather than covering it with words.
