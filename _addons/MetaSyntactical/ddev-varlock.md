---
title: MetaSyntactical/ddev-varlock
github_url: https://github.com/MetaSyntactical/ddev-varlock
description: "DDEV add-on that installs the Varlock CLI into the web container"
user: MetaSyntactical
repo: ddev-varlock
repo_id: 1295829552
default_branch: main
tag_name: v0.1.0
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2026-07-09
updated_at: 2026-07-09
workflow_status: success
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![release](https://img.shields.io/github/v/release/MetaSyntactical/ddev-varlock)](https://github.com/MetaSyntactical/ddev-varlock/releases/latest)
[![license](https://img.shields.io/github/license/MetaSyntactical/ddev-varlock)](LICENSE)

# DDEV Varlock

## Overview

[Varlock](https://varlock.dev) provides AI-safe `.env` files: schemas for agents, secrets for humans. It validates and type-checks environment variables, prevents secret leaks, and can resolve secrets from providers like 1Password, Infisical, and Doppler at runtime.

This add-on installs the [Varlock CLI](https://varlock.dev/reference/cli-commands/) into the DDEV web container and wires it up to automatically load and validate your project's environment variables on every `ddev start`.

## Installation

```bash
ddev add-on get MetaSyntactical/ddev-varlock
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

### Prerequisites

* Your project needs an existing [`.env.schema`](https://varlock.dev/getting-started/) — this add-on doesn't create one for you.

## Usage

* `ddev varlock [command]` runs any [Varlock CLI](https://varlock.dev/reference/cli-commands/) command inside the web container, e.g. `ddev varlock load` or `ddev varlock run -- <command>`.
* On every `ddev start`, a `pre-start` hook runs `varlock load` on the host to print a readable validation summary before the containers even come up — useful for catching a broken schema or missing credentials early, without waiting for a container boot. This is optional: if the [Varlock CLI](https://varlock.dev/getting-started/installation/) isn't installed on your host, you'll just see a notice and `ddev start` continues normally — only the early feedback is skipped, not the actual secret injection below.
* The secrets actually reach your app via `.ddev/web-entrypoint.d/varlock.sh`. DDEV sources this script inside the web container right before starting nginx and php-fpm, so `eval "$(varlock load --format=shell)"` there injects resolved values straight into the environment every request sees via `getenv()`/`$_ENV` — no need to wrap your app's own start command.
* To resolve secrets, export the credentials for your provider in your **host** shell environment before running `ddev start`. These are only forwarded to the web container at runtime and are never written into the project:
  * `OP_SERVICE_ACCOUNT_TOKEN` — 1Password
  * `INFISICAL_CLIENT_ID` + `INFISICAL_CLIENT_SECRET` — Infisical (Universal Auth)
  * `DOPPLER_TOKEN` — Doppler
  * `AWS_ACCESS_KEY_ID` + `AWS_SECRET_ACCESS_KEY` — AWS Secrets Manager
  * `VAULT_TOKEN` — HashiCorp Vault

### Obtaining provider credentials

* **1Password** — create a [service account](https://developer.1password.com/docs/service-accounts/get-started/) via the web UI, or from the CLI:
  ```bash
  op service-account create "ddev-varlock" --vault "<vault-name>:read_items" --expires-in 90d
  ```
  The token is only printed once — capture it as `OP_SERVICE_ACCOUNT_TOKEN`. See the [Varlock 1Password plugin docs](https://varlock.dev/plugins/1password/) for `.env.schema` setup (its default config var is `OP_TOKEN`, so reference `$OP_SERVICE_ACCOUNT_TOKEN` explicitly in `@initOp()` to match this add-on).
* **HashiCorp Vault** — authenticate and mint a token:
  ```bash
  export VAULT_ADDR="https://vault.example.com"
  vault token create -policy="ddev-varlock" -ttl=24h -format=json | jq -r '.auth.client_token'
  ```
  Export the result as `VAULT_TOKEN`. See the [Varlock Vault plugin docs](https://varlock.dev/plugins/hashicorp-vault/) for `.env.schema` setup.
* **Infisical, Doppler, AWS Secrets Manager** — issue credentials from each provider's own console or CLI; see the [full plugin list](https://varlock.dev/guides/plugins/) for `.env.schema` setup per provider.

Scope credentials as narrowly as possible (read-only, single vault/policy, short TTL) rather than reusing a personal or admin token — they'll sit in your host shell environment for the life of the session.

## Configuration

### Pinning the Varlock version

By default, the add-on installs the latest Varlock release every time the web image is built. To pin a specific version instead:

```bash
ddev dotenv set .ddev/.env.varlock --varlock-version=<version>
ddev restart
```

For example:

```bash
ddev dotenv set .ddev/.env.varlock --varlock-version=1.9.0
ddev restart
```

Use a plain semantic version without a `v` prefix (e.g. `1.9.0`), matching the tags in the [Varlock releases](https://github.com/dmno-dev/varlock/releases). Leave `.ddev/.env.varlock` absent, or its value empty, to keep installing the latest release on every build.

Make sure to commit `.ddev/.env.varlock` to version control once you've pinned a version.

### Using a different secrets provider

Varlock supports many more providers than the ones listed above — for example Azure Key Vault, Google Secret Manager, Akeyless, Bitwarden, Keeper, Passbolt, Proton Pass, and Dashlane (see the [full plugin list](https://varlock.dev/guides/plugins/)). Each has its own auth variable(s), e.g. Azure Key Vault needs `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and `AZURE_CLIENT_SECRET`.

This add-on only pre-declares `web_environment` passthroughs for the most common providers. `web_environment` entries listed as a bare name (no `=value`) tell DDEV to forward that variable from your **host** shell into the web container at start time — the value itself is never written to `.ddev/config.yaml`, so it's safe to add even for secrets.

To add support for another provider's variable(s) in your project, run:

```bash
ddev config --web-environment-add="AZURE_TENANT_ID"
ddev config --web-environment-add="AZURE_CLIENT_ID"
ddev config --web-environment-add="AZURE_CLIENT_SECRET"
ddev restart
```

This appends the passthrough entries to your project's own `.ddev/config.yaml` (under `web_environment`), on top of the ones this add-on already registers globally — no need to fork or modify the add-on itself. Reference the resulting variables in your `.env.schema` the same way as the built-in providers (e.g. `@initAzureKeyVault(tenantId=$AZURE_TENANT_ID, ...)`), export the values in your host shell, then `ddev restart`.

## Credits

**Maintained by [MetaSyntactical](https://github.com/MetaSyntactical)**
