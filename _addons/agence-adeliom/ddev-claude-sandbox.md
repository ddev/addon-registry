---
title: agence-adeliom/ddev-claude-sandbox
github_url: https://github.com/agence-adeliom/ddev-claude-sandbox
description: "DDEV addon for sandboxing Claude Code in professional team environments with URL whitelist and secrets protection."
user: agence-adeliom
repo: ddev-claude-sandbox
repo_id: 1149650710
default_branch: main
tag_name: 
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2026-02-04
updated_at: 2026-02-25
workflow_status: failure
stars: 0
---

# DDEV Claude Sandbox

A DDEV addon that runs Claude Code inside a native sandbox (bubblewrap) with kernel-level isolation, network control, and file protection.

## Installation

```bash
ddev get agence-adeliom/ddev-claude-sandbox
ddev restart
```

## Usage

### Claude Code

```bash
ddev claude
```

Claude runs in `acceptEdits` mode with full sandbox isolation. Bash commands are auto-allowed when sandboxed.

### Execute commands with secrets

Protected files (`.env.local`, etc.) are blocked from Claude's Read/Edit tools. When you need to run commands that require those secrets:

```bash
ddev agent-env php bin/console app:call-api
ddev agent-env printenv | grep API_KEY
```

`ddev agent-env` loads the full environment (including `.env.local`) and executes the command, keeping secrets out of Claude's context.

## Configuration

Override defaults in `.ddev/config.local.yaml`:

```yaml
web_environment:
  # Customize protected files (comma-separated patterns)
  - CLAUDE_PROTECTED_FILES=.env.local,.env.*.local,credentials.json

  # Add project-specific allowed domains (comma-separated)
  - CLAUDE_ALLOWED_DOMAINS=api.example.com,*.internal.dev
```

Then restart: `ddev restart`

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `CLAUDE_PROTECTED_FILES` | `.env.local,.env.*.local` | File patterns denied in Read/Edit |
| `CLAUDE_ALLOWED_DOMAINS` | _(empty)_ | Additional domains to allow through sandbox network |

## Security

### Native Sandbox (bubblewrap)

Claude Code runs inside a bubblewrap sandbox providing:
- **Kernel-level process isolation** - filesystem and network restrictions enforced by the OS
- **Network allowlist** - only pre-approved domains are reachable (Anthropic API, GitHub, npm, Packagist, Composer, Symfony)
- **Auto-allowed bash** - shell commands run freely inside the sandbox without prompts (`autoAllowBashIfSandboxed`)
- **Docker excluded** - `docker` commands are excluded from the sandbox to avoid conflicts with DDEV

### File Protection

Protected files are enforced via `permissions.deny` rules in `settings.json`:
- `Read(<pattern>)` and `Edit(<pattern>)` deny rules block Claude from accessing sensitive files
- Default: `.env.local` and `.env.*.local` (with recursive `**/` variants)
- Use `ddev agent-env <command>` when you need to run commands with secrets loaded

### Persistent Configuration

- Claude credentials (`.credentials.json`) survive container restarts via symlink to project volume
- Claude config (`.claude.json`) also persisted
- Settings are regenerated on each `ddev restart` from environment variables

## Contributing

### Project Structure

```
ddev-claude-sandbox/
├── install.yaml                    # Addon manifest
├── config.claude-sandbox.yaml      # DDEV hooks and environment
├── web-build/
│   └── Dockerfile.claude-sandbox   # Installs claude, jq, bubblewrap, socat
├── claude/
│   └── .gitignore
├── commands/web/
│   ├── claude                      # ddev claude entrypoint
│   └── agent-env                   # ddev agent-env entrypoint
├── scripts/
│   ├── setup-claude.sh             # Post-start setup (symlinks, persistence)
│   └── generate-claude-settings.sh # Generates settings.json with jq
└── tests/
    └── test.bats
```

## Testing

This addon uses [BATS](https://bats-core.readthedocs.io/) (Bash Automated Testing System).

### Install Dependencies

```bash
# macOS
brew install bats-core
brew tap kaos/shell
brew install bats-file bats-support

# Linux (apt)
apt install bats bats-assert bats-file bats-support
```

### Run Tests

```bash
# All tests
bats ./tests/test.bats

# Exclude release tests (local development)
bats ./tests/test.bats --filter-tags '!release'

# Verbose output
bats ./tests/test.bats --show-output-of-passing-tests --verbose-run
```

### Test Coverage

| Test | Description |
|------|-------------|
| `install from directory` | Basic addon installation |
| `claude command is available` | Claude binary works |
| `agent-env command works` | Secrets wrapper executes |
| `claude config directory exists` | Config directory created |
| `credentials are persisted and symlinked` | Credentials survive restarts |
| `claude config is persisted and symlinked` | Config survives restarts |
| `settings.json has sandbox enabled` | Sandbox + autoAllowBashIfSandboxed |
| `bubblewrap is installed` | bwrap binary available |
| `settings.json has deny rules for protected files` | Read/Edit deny rules |
| `settings.json has allowed domains` | Network allowlist present |
| `protected files pattern is configurable` | Custom deny rules from env var |
| `custom allowed domains can be added` | Extra domains via env var |
| `install from release` | Install from GitHub release |

### CI/CD

Tests run automatically via GitHub Actions on:
- Pull requests
- Pushes to `main`
- Daily schedule (08:25 UTC)

## License

Apache License 2.0
