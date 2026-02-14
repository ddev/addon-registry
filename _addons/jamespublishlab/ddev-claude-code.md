---
title: jamespublishlab/ddev-claude-code
github_url: https://github.com/jamespublishlab/ddev-claude-code
description: "DDEV addon for Claude Code CLI integration"
user: jamespublishlab
repo: ddev-claude-code
repo_id: 1157025315
default_branch: main
tag_name: v1.1.0
ddev_version_constraint: ">= v1.23.0"
dependencies: []
type: contrib
created_at: 2026-02-13
updated_at: 2026-02-13
workflow_status: unknown
stars: 0
---

# ddev-claude-code

A DDEV addon that integrates [Claude Code](https://claude.ai/code) CLI into your DDEV environment.

## What It Does

- Installs Claude Code CLI, GitHub CLI (`gh`), and [ccusage](https://github.com/ryoppippi/ccusage) in the DDEV web container
- Mounts your host `~/.claude` and `~/.config/gh` directories into the container
- Injects `GH_TOKEN` and `CLAUDE_CODE_OAUTH_TOKEN` via a pre-start hook
- Provides `ddev claude` and `ddev claude-auto` commands

## Installation

```bash
ddev add-on get jamespublishlab/ddev-claude-code
ddev restart
```

### Prerequisites

- DDEV >= v1.23.0
- Claude Code installed on your host (`npm install -g @anthropic-ai/claude-code` or via `claude.ai/install.sh`)
- Run `claude` on your host at least once to initialize `~/.claude`
- GitHub CLI (`gh`) authenticated on your host (`gh auth login`)

## Usage

```bash
# Run Claude Code interactively
ddev claude

# Run with arguments
ddev claude --help
ddev claude --version

# Run in autonomous mode (skips permission prompts)
ddev claude-auto "add unit tests for the auth module"
```

### Checking the setup

```bash
ddev claude --version         # Claude Code CLI version
ddev exec gh --version        # GitHub CLI version
ddev exec 'echo $GH_TOKEN'   # Confirm token injection
```

## Authentication

The addon supports two authentication methods. The pre-start hook automatically captures tokens from your host and injects them into the container — no manual env var setup required.

### Option 1: OAuth Token (recommended)

Uses your Claude Pro/Max subscription allowance — no additional API costs.

1. **Install Claude Code on your host:**

   ```bash
   npm install -g @anthropic-ai/claude-code
   ```

2. **Run `claude` once to authenticate:**

   ```bash
   claude
   ```

   This opens a browser window where you sign in with your Anthropic account. On completion, Claude saves your OAuth credentials (access token + refresh token) to `~/.claude/.credentials.json`.

3. **Install the addon and restart:**

   ```bash
   ddev add-on get jamespublishlab/ddev-claude-code
   ddev restart
   ```

   The pre-start hook reads the token from `~/.claude/.credentials.json` and sets `CLAUDE_CODE_OAUTH_TOKEN` in the container automatically.

> **Note:** OAuth access tokens expire after ~1 hour. Because `~/.claude` is volume-mounted into the container, Claude Code can auto-refresh using the refresh token stored in `~/.claude/.credentials.json`. If you experience auth failures during long sessions, run `ddev restart` to re-capture a fresh token.

### Option 2: API Key

Uses per-token API billing — simpler but costs more than a subscription.

1. **Get an API key** from [console.anthropic.com](https://console.anthropic.com/)

2. **Add it to `.ddev/.env.tokens`:**

   ```bash
   echo "ANTHROPIC_API_KEY=sk-ant-api03-..." >> .ddev/.env.tokens
   ```

   > Do **not** set both `ANTHROPIC_API_KEY` and `CLAUDE_CODE_OAUTH_TOKEN` — they conflict and Claude will throw an auth error.

3. **Restart:** `ddev restart`

### GitHub CLI

The addon also captures your host's `gh auth token` for GitHub integration. Authenticate on your host first:

```bash
gh auth login
```

### Verifying authentication

```bash
ddev exec 'echo $CLAUDE_CODE_OAUTH_TOKEN'   # Should show your OAuth token
ddev exec 'echo $GH_TOKEN'                  # Should show your GitHub token
ddev claude --version                        # Should print version without auth errors
```

## How It Works

1. **Pre-start hook** (`config.claude-code.yaml`): Captures your host's `gh auth token` and Claude OAuth token into `.ddev/.env.tokens`
2. **Docker Compose overlay** (`docker-compose.claude-code.yaml`): Mounts `~/.claude` and `~/.config/gh` into the container and loads `.env.tokens`
3. **Dockerfile** (`web-build/Dockerfile.claude-code`): Installs `gh`, `claude`, and `ccusage` in the web container

## Files

| File | Type | Purpose |
|------|------|---------|
| `docker-compose.claude-code.yaml` | Project | Volume mounts + env_file |
| `config.claude-code.yaml` | Project | Pre-start hook for token capture |
| `web-build/Dockerfile.claude-code` | Project | Installs CLI tools |
| `commands/web/claude` | Global | `ddev claude` command |
| `commands/web/claude-auto` | Global | `ddev claude-auto` command |

## Removal

```bash
ddev add-on remove claude-code
ddev restart
```

## License

MIT
