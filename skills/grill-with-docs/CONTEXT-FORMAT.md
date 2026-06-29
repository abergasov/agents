# CONTEXT.md Format

`CONTEXT.md` is a **glossary** — the canonical, agreed meaning of the domain terms this codebase uses. It is not a spec, a scratch pad, or a home for implementation decisions (those go in ADRs). Keep it free of implementation detail.

Create it lazily — only when the first term is resolved during a grilling session.

## Template

```md
# Context: {Bounded context or project name}

One sentence on what this context is responsible for.

## Glossary

### {Term}
What it means here, in one or two sentences. Note what it is **not** when the
term is easily confused with another (e.g. "Customer — the billing entity; not
the User, which is a login identity").

### {Term}
...
```

## Rules

- **One canonical meaning per term.** If a word is overloaded ("account", "order"), split it into distinct terms and define each.
- **Capture inline.** Add or correct a term the moment it's resolved in conversation — don't batch.
- **Glossary only.** If you're tempted to write *how* something works, that's an ADR (see [ADR-FORMAT.md](./ADR-FORMAT.md)) or code, not this file.
- **Multiple contexts.** When a repo has a `CONTEXT-MAP.md`, each context owns its own `CONTEXT.md`; define a term in the context that owns it.
