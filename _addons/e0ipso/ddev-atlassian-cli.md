---
title: "e0ipso/ddev-atlassian-cli"
github_url: "https://github.com/e0ipso/ddev-atlassian-cli"
description: "Connect the Atlassian CLI to your account inside the ddev container."
user: "e0ipso"
repo: "ddev-atlassian-cli"
repo_id: 1436923670
default_branch: "main"
tag_name: "v2.0.1"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-10-28"
updated_at: "2026-10-28"
workflow_status: "unknown"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/e0ipso/ddev-atlassian-cli/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/e0ipso/ddev-atlassian-cli/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/e0ipso/ddev-atlassian-cli)](https://github.com/e0ipso/ddev-atlassian-cli/commits)
[![release](https://img.shields.io/github/v/release/e0ipso/ddev-atlassian-cli)](https://github.com/e0ipso/ddev-atlassian-cli/releases/latest)

# DDEV Atlassian CLI

## Overview

This add-on installs the [Atlassian Teamwork Graph CLI](https://developer.atlassian.com/platform/teamwork-graph/twg-cli/) (`twg`) in the [DDEV](https://ddev.com/) web container and adds a `ddev twg` command to run it.

`twg` covers Jira, Confluence, Bitbucket, Jira Service Management, Loom, Trello, Rovo and the rest of Teamwork Graph. Its commands and JSON output are built for coding agents as well as people.

## Installation

```bash
ddev add-on get e0ipso/ddev-atlassian-cli
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

```bash
ddev twg login
ddev twg jira workitem get ENT-1234
ddev twg doctor
```

`ddev twg` forwards all arguments to `twg`, so everything in the [command catalog](https://developer.atlassian.com/platform/teamwork-graph/twg-cli/commands/commands-catalog/) works. `twg` is also on `PATH` inside `ddev ssh` and `ddev exec`.

### Logging in

`twg login` uses the OAuth device flow, so the container needs no browser.

1. Run `ddev twg login`.
2. Answer yes to `Open in browser?`. There is no browser in the container, so `twg` just repeats the URL.
3. Open the URL on your host, check that the code matches the one in the terminal, and approve.

The login is stored in `~/.config/twg` inside the web container, and `ddev restart` recreates the container. Run `ddev twg login` again after a restart, or use environment variables.

### Environment variables

`twg` reads `TWG_USER`, `TWG_SITE` and `TWG_TOKEN` on every command. Put them in `.ddev/.env.web.local`, which DDEV loads into the web container and its generated `.gitignore` excludes:

```bash
ddev dotenv set .ddev/.env.web.local --twg-user=you@example.com --twg-site=example.atlassian.net
ddev restart
```

Add `TWG_TOKEN` to that file with an editor rather than on the command line, so the token stays out of your shell history. These variables also cover agent sandboxes that can't answer an interactive login. Run `ddev twg env --list` to see every variable `twg` understands.

## What gets installed

| File | Purpose |
| ---- | ------- |
| `.ddev/web-build/Dockerfile.atlassian-cli` | Runs Atlassian's [installer](https://teamwork-graph.atlassian.com/cli/install) unattended to install the latest stable `twg` as `/usr/local/bin/twg` |
| `.ddev/commands/web/twg` | The `ddev twg` command |

The build skips the installer's login and agent skills, and the image has no upkeep helper. Agent skills belong where your agent runs, and the helper needs a user service manager that the container lacks.

The installer runs with a throwaway home directory, so the image records no consent to Atlassian's terms on your behalf. Run `ddev twg consent` to review and accept them.

DDEV caches the web image, so `twg` stays at the version it was built with. Run `ddev debug rebuild` to pick up a new release. Do this when `twg` reports that it is too old for the current backend, because Atlassian blocks versions below a published minimum.

## Removal

```bash
ddev add-on remove atlassian-cli
ddev restart
```

## Contributing

Commits follow [Conventional Commits](https://www.conventionalcommits.org/), checked by [commitlint](https://commitlint.js.org/) in a husky `commit-msg` hook. Run `npm install` once to enable the hook.

Releases are automatic. After the `tests` workflow passes on `main`, [semantic-release](https://semantic-release.gitbook.io/) reads the commits since the last tag. A `feat:` commit makes a minor release, a `fix:` or `perf:` commit makes a patch release, and a `BREAKING CHANGE:` footer makes a major one. Other types do not release.

## Credits

**Contributed and maintained by [@e0ipso](https://github.com/e0ipso)**
