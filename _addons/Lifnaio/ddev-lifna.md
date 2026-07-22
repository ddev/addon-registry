---
title: Lifnaio/ddev-lifna
github_url: https://github.com/Lifnaio/ddev-lifna
description: ""
user: Lifnaio
repo: ddev-lifna
repo_id: 1266494277
default_branch: main
tag_name: v0.1.7
ddev_version_constraint: ">= v1.24.0"
dependencies: []
type: contrib
created_at: 2026-06-11
updated_at: 2026-07-21
workflow_status: success
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![CI](https://github.com/Lifnaio/ddev-lifna/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/Lifnaio/ddev-lifna/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/Lifnaio/ddev-lifna)](https://github.com/Lifnaio/ddev-lifna/commits)
[![release](https://img.shields.io/github/v/release/Lifnaio/ddev-lifna)](https://github.com/Lifnaio/ddev-lifna/releases/latest)

# DDEV Lifna

Connect an existing DDEV Drupal project to a Lifna-hosted environment.

This add-on installs a native DDEV provider named `lifna`, so existing projects can use:

```bash
ddev pull lifna
ddev push lifna
ddev lifna status
ddev lifna spinup
ddev lifna pause
```

Code remains Git-first. The provider only syncs the database and Drupal public files, matching DDEV hosting-provider conventions.

## Install

Requires DDEV v1.24.0+ and an existing DDEV Drupal project.

From a DDEV project root:

```bash
ddev add-on get Lifnaio/ddev-lifna
ddev restart
```

To test a local checkout while developing this add-on, use the same command with a local path, for example `ddev add-on get /path/to/ddev-lifna`.

## Upgrade

Repeat the install command to update to the latest release:

```bash
ddev add-on get Lifnaio/ddev-lifna
ddev restart
```

Check the installed version with `ddev add-on list --installed`.

## Connect A Project

In Lifna, open the site, choose **Download for DDEV / local development**, then use the **DDEV access token** card. The direct URL is:

```text
https://app.lifna.io/sites/<site-slug>/export
```

Choose **Account-wide token** to reuse one token across every Lifna site you can access, or choose a specific environment for a narrower token. Select the allowed actions, click **Create DDEV token**, and copy the one-time token.

Then run:

```bash
ddev lifna connect \
  --site=my-site \
  --environment=main
```

Paste the token when prompted.

For local Lifna platform development only, set `LIFNA_DEV_MODE=1` or pass `--dev` with `--base-url` when connecting to a local HTTPS or localhost URL.

The command writes:

- `.lifna/environment.json` with the Lifna site/environment link.
- `.ddev/lifna/.env` with the local token.

Both are ignored locally by generated `.gitignore` files.

## Daily Workflow

```bash
ddev lifna status
ddev pull lifna
ddev push lifna
ddev lifna pause
ddev lifna spinup
ddev lifna open
```

`ddev pull lifna` downloads the Lifna database and public files into the DDEV project.

`ddev push lifna` uploads the local database and public files back to the scoped Lifna environment. Pushes to protected environments such as `main`, `live`, `prod`, or `production` require a typed confirmation.

## What Has Been Tested

Automated tests cover:

- Add-on installation into a disposable DDEV Drupal project.
- Generated DDEV provider and host command files.
- Scoped token storage without command-line token input.
- Production base URL protection and explicit local development mode.
- Token-safe diagnostics.
- Protected environment push confirmation.
- Files archive traversal protection before extraction.

## Limitations

- Database and public files are synced; code stays Git-first.
- Lifna API endpoints and token generation must be available in the Lifna app before pull, push, and lifecycle commands can complete.
- Local development URLs require explicit `--dev` or `LIFNA_DEV_MODE=1`.

## Security Model

The add-on uses a Lifna local access token created in the Lifna UI. Tokens can be account-wide or scoped to one site/environment, and always carry limited actions plus an expiry date. Lifna still checks your account permissions for the requested site on every API call.

Tokens are read from the secure prompt during `ddev lifna connect` or from `LIFNA_TOKEN`. Do not pass tokens as command-line arguments and do not commit `.ddev/lifna/.env`.

Production connections default to `https://app.lifna.io` unless explicit dev mode is enabled.

## Testing

The test suite is written with Bats and includes shell-level security tests plus a DDEV install smoke test.

```bash
bash -n commands/host/lifna lifna/client.sh
shellcheck commands/host/lifna lifna/client.sh
yamllint .
bats tests
```

GitHub Actions runs the same lint checks and uses DDEV's add-on test action to install the add-on into a disposable DDEV project.

## Publish

The recommended GitHub repo name is `ddev-lifna`, with the `ddev-get` topic added for DDEV add-on discovery.
