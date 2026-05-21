---
title: nicolabeghin/claude-code-dbhub-playwright
github_url: https://github.com/nicolabeghin/claude-code-dbhub-playwright
description: "Integrate Claude Code, Playwright and DBHub into your ddev setup to run AI assistant in the web container"
user: nicolabeghin
repo: claude-code-dbhub-playwright
repo_id: 1243951871
default_branch: main
tag_name: v0.0.7
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2026-05-19
updated_at: 2026-05-20
workflow_status: unknown
stars: 1
---

# ddev-claude-code-dbhub-playwright <!-- omit in toc -->

> Opinionated fork of [ddev-claude-code](https://github.com/FreelyGive/ddev-claude-code),
> extending it with DBHub (database MCP) and Playwright (browser MCP) services.

## Installation

```shell
ddev add-on get nicolabeghin/claude-code-dbhub-playwright
ddev restart
```

## Claude Code

Claude Code runs inside the DDEV web container and is accessible using
`ddev claude`. Running inside the container gives Claude direct access to
all project files, services (database, cache, etc.), and `ddev` commands —
no host credentials or tunnelling needed.

Configuration is stored in `.ddev/.claude.json` and `.ddev/.claude/`, and
persisted across restarts. You can drop in an existing `.ddev/.claude.json`
to reuse an API key or pre-approved tool list.

The host's `~/.claude/` and `~/.claude.json` are also mounted into the
container, so global Claude settings and credentials are available automatically.

To let Claude interact with GitLab, authenticate `glab` once with
`ddev glab auth login`. Configuration is stored in `.ddev/.glab-cli/` and
persisted across restarts.

## DBHub

[DBHub](https://github.com/bytebase/dbhub) is an MCP server that gives Claude
direct access to the project's MariaDB instance (`db:3306`). It also exposes
an optional database management UI at:

```
https://dbhub.<sitename>.<tld>
```

To enable the MCP integration, add the following to your project's `.mcp.json`:

```json
{
  "mcpServers": {
    "dbhub": {
      "type": "http",
      "url": "http://dbhub:22222/mcp"
    }
  }
}
```

To disable it, remove or rename `docker-compose.dbhub.yaml` in your `.ddev/`
directory.

## Playwright (MCP)

A Playwright service is included, running a headless Chromium browser with
both a WebSocket automation server (port 3000) and a
[Playwright MCP](https://github.com/microsoft/playwright-mcp) HTTP server
(port 3301). This lets Claude Code control a browser directly.

To enable the MCP integration, add the following to your project's `.mcp.json`
(create it at the project root if it doesn't exist):

```json
{
  "mcpServers": {
    "playwright": {
      "type": "http",
      "url": "http://playwright:3301/mcp"
    }
  }
}
```

The service is reachable from inside the DDEV web container only. If you need
host access, add a `ports:` entry to `docker-compose.playwright.yaml` — but
note this will conflict if multiple DDEV projects run the service simultaneously.

### Writing tests

Install `@playwright/test` in your project, then point `playwright.config.js`
at the in-container browser via its WebSocket endpoint:

```js
// e2e/playwright.config.js
const { defineConfig, devices } = require('@playwright/test');

module.exports = defineConfig({
  testDir: './tests',
  timeout: 120_000,
  use: {
    baseURL: process.env.BASE_URL ?? 'https://<sitename>.ddev.site',
    ignoreHTTPSErrors: true,
    screenshot: 'only-on-failure',
    trace: 'retain-on-failure',
  },
  projects: [
    {
      name: 'chromium',
      use: {
        ...devices['Desktop Chrome'],
        connectOptions: {
          wsEndpoint: process.env.PLAYWRIGHT_WS_ENDPOINT
            ?? 'ws://playwright:3000/ddev-pw-token',
          timeout: 30_000,
        },
      },
    },
  ],
});
```

A minimal smoke test:

```js
// e2e/tests/smoke.spec.js
const { test, expect } = require('@playwright/test');

test('browser is reachable', async ({ page }) => {
  await page.goto('https://www.google.com', { waitUntil: 'domcontentloaded' });
  await expect(page).toHaveTitle(/Google/);
});
```

Run tests from inside the web container:

```shell
ddev exec npx playwright test --config e2e/playwright.config.js
```
