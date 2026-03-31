---
title: owenbush/decodie-ddev
github_url: https://github.com/owenbush/decodie-ddev
description: ""
user: owenbush
repo: decodie-ddev
repo_id: 1195669181
default_branch: main
tag_name: 
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2026-03-30
updated_at: 2026-03-30
workflow_status: unknown
stars: 0
---

<p align="center"><img src="assets/decodie-logo.png" alt="Decodie" width="200"></p>

# Decodie DDEV Add-on

A DDEV add-on that integrates the Decodie learning companion into your local development environment.

## Requirements

- DDEV v1.22+

## Installation

```bash
ddev add-on get owenbush/decodie-ddev
ddev restart
ddev decodie
```

This installs everything you need:
- The Decodie UI (via npm)
- The Decodie skill for Claude Code (into `~/.claude/skills/decodie/`)

The UI will be available at `https://decodie.SITENAME.ddev.site` (where SITENAME is your DDEV project name).

To start generating learning entries, activate the skill in Claude Code by running `/decodie` at the start of your session. Entries will appear automatically as you code.

## Usage

```bash
ddev decodie            # Open in your browser
ddev decodie status     # Show entry statistics
ddev decodie cleanup    # Open the archival review page
ddev decodie help       # Show available commands
```

## Q&A Configuration

The Q&A feature requires authentication with Claude. Edit `.ddev/decodie/.env` and set one of the following:

### Option A: OAuth token (recommended -- uses your Claude subscription)

1. On your **host machine** (not inside DDEV), run: `claude setup-token`
2. Copy the token and add it to `.ddev/decodie/.env`:
   ```
   CLAUDE_CODE_OAUTH_TOKEN=sk-ant-oat01-...
   ```
3. `ddev restart`

### Option B: API key

1. Get a key from [console.anthropic.com](https://console.anthropic.com/)
2. Add it to `.ddev/decodie/.env`:
   ```
   CLAUDE_API_KEY=sk-ant-...
   ```
3. `ddev restart`

## How It Works

The add-on installs [`@owenbush/decodie-ui`](https://www.npmjs.com/package/@owenbush/decodie-ui) from npm and runs it as a daemon inside the DDEV web container. Traefik handles subdomain routing. No separate container is needed.

To update to the latest version of the UI:

```bash
ddev exec "cd /var/www/html/.ddev/decodie && npm update @owenbush/decodie-ui"
ddev restart
```

## Related Repositories

- [owenbush/decodie](https://github.com/owenbush/decodie) -- Main Decodie site
- [owenbush/decodie-ui](https://github.com/owenbush/decodie-ui) -- The learning companion UI
- [owenbush/decodie-skill](https://github.com/owenbush/decodie-skill) -- Claude Code skill for generating learning entries
