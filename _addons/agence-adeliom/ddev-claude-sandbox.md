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
updated_at: 2026-02-04
workflow_status: failure
stars: 0
---

# DDEV Claude Sandbox

A DDEV addon for sandboxing Claude Code in professional team environments with built-in security features: URL allow list and .env protection.

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

### Execute commands with secrets

Run commands that need environment variables from `.env.local`:

```bash
ddev agent-env php bin/console app:call-api
ddev agent-env printenv | grep API_KEY
```

## Configuration

Override defaults in `.ddev/config.local.yaml`:

```yaml
web_environment:
  # Disable URL allow list feature
  - CLAUDE_URL_ALLOWLIST_ENABLED=false

  # Disable .env protection
  - CLAUDE_ENV_PROTECTION_ENABLED=false

  # Customize protected files (comma-separated patterns)
  - CLAUDE_PROTECTED_FILES=.env.local,.env.*.local,credentials.json
```

Then restart: `ddev restart`

### Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `CLAUDE_URL_ALLOWLIST_ENABLED` | `true` | Auto-approve domains after first authorization |
| `CLAUDE_ENV_PROTECTION_ENABLED` | `true` | Block Claude from reading .env files |
| `CLAUDE_PROTECTED_FILES` | `.env.local,.env.*.local` | File patterns to protect |

## Benefits

### Security Features

**URL Allow list/Disallow list** - Control which domains Claude can access:
- First access to a new domain prompts for user approval
- Approved domains are automatically allowed for future requests
- Refused domains can be disallowed (Claude will call `~/.claude/hooks/url-disallowlist-add.sh <domain>`)
- Disallowed domains are auto-rejected, but Claude can still access other domains
- Lists persist in `.ddev/claude/url-list.json`

**Environment Protection** - Keep secrets safe:
- Claude cannot read `.env.local` or other protected files
- Use `ddev agent-env <command>` when you need secrets loaded
- Configurable file patterns for custom protection

### Developer Experience

- **Persistent configuration** - Claude settings survive container restarts
- **Native binary** - No Node.js overhead, faster startup
- **Team-ready** - Security controls for professional environments

## Contributing

### Project Structure

```
ddev-claude-sandbox/
├── install.yaml                 # Addon manifest
├── config.claude-sandbox.yaml   # DDEV hooks and environment
├── web-build/
│   └── Dockerfile.claude-sandbox
├── claude/
│   └── hooks/                   # Security hooks (tracked by git)
├── commands/web/
│   ├── claude
│   └── agent-env
├── scripts/
│   ├── setup-claude.sh
│   └── generate-claude-settings.php
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
| `setup script creates hook files` | Hooks generated correctly |
| `settings.json is generated` | Hook configuration created |
| `url allow list can be disabled` | Feature toggle works |
| `env protection can be disabled` | Feature toggle works |
| `claude config directory exists` | Config directory created |
| `protected files pattern is configurable` | Custom patterns work |

### CI/CD

Tests run automatically via GitHub Actions on:
- Pull requests
- Pushes to `main`
- Daily schedule (08:25 UTC)

## License

Apache License 2.0
