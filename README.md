# agents

Reusable agents and skills for software-development workflows.

## Install Context7 MCP

```bash
npx ctx7
npx ctx7 setup
export CONTEXT7_API_KEY=ctx7sk-abc
```

## Adopt generated outputs

```bash
make adopt_opencode
make adopt_copilot
```

`adopt_opencode` installs agents and skills into `~/.config/opencode`.

## Local OpenCode config

`opencode.json` enables the Context7 MCP server and exposes repo-local skills from `./skills`.