---
title: vanWittlaer/ddev-coolify-shopware
github_url: https://github.com/vanWittlaer/ddev-coolify-shopware
description: "Installs OpenTofu, the vanwittlaer/opentofu-coolify-shopware-stack and helper commands"
user: vanWittlaer
repo: ddev-coolify-shopware
repo_id: 1293359306
default_branch: main
tag_name: v1.2.0
ddev_version_constraint: ">= v1.24.0"
dependencies: []
type: contrib
created_at: 2026-07-08
updated_at: 2026-07-10
workflow_status: unknown
stars: 0
---

# ddev-coolify-bootstrap

[ddev](https://ddev.com) add-on that turns a Shopware 6 ddev project into a
**one-shot installer** for a production + staging stack on
[Coolify](https://coolify.io) v4 — powered by OpenTofu and the
[`terraform-coolify-shopware-stack`](https://github.com/vanWittlaer/terraform-coolify-shopware-stack)
module under the hood. You never run `tofu` yourself.

## Install

```bash
ddev add-on get vanWittlaer/ddev-coolify-bootstrap
ddev restart                      # bakes OpenTofu into the web container
```

## Use

```bash
ddev coolify-bootstrap init                 # scaffold infra/ (fresh projects only)
# fill in infra/secrets.auto.tfvars (from the .example) and infra/*.tfvars
ddev coolify-bootstrap up                   # provision production + staging on Coolify
ddev coolify-bootstrap up staging           # ... or one environment at a time
ddev coolify-bootstrap destroy [staging]    # tear the stack (or one env) down again
```

`up` runs prereq checks → plan → **one** confirmation → apply → a post-setup
checklist. What gets created, the required prerequisites (Coolify server, S3
buckets, registry image) and every knob are documented in the module repo's
[README](https://github.com/vanWittlaer/terraform-coolify-shopware-stack#readme),
[PREREQUISITES](https://github.com/vanWittlaer/terraform-coolify-shopware-stack/blob/main/PREREQUISITES.md)
and [STATE](https://github.com/vanWittlaer/terraform-coolify-shopware-stack/blob/main/STATE.md).
A complete reference project using this add-on: [swoofy](https://github.com/vanWittlaer/swoofy).

With an environment argument (`production`|`staging`), `up` and `destroy` operate
on that environment only (OpenTofu resource targeting under the hood — tofu's
"targeting is in effect" note is expected). Shared resources — the Coolify
project and the S3 CORS rule — are created by whichever per-env `up` runs first
and are only removed by a full, no-argument `destroy`: after destroying both
envs individually, an empty project (and the CORS rule) linger until then. The
one-shot contract is per environment: bootstrap each env once, then the Coolify
UI owns it.

## The one-shot contract

This is a **day-0 bootstrapper, not a management tool**. After `up` succeeds,
the **Coolify UI is the single source of truth** — maintain, tune and upgrade
the environment there, and never re-run the bootstrap against it (the Coolify
provider pushes env vars write-only, so a re-apply silently overwrites UI
changes). The command enforces this: it refuses to bootstrap when the Coolify
project already exists, and warns before any re-apply from existing local
state. Afterwards, archive `infra/secrets.auto.tfvars` + `infra/tofu.tfstate`
off-machine and delete them locally — they are recovery records.

## What the add-on installs

| File | Purpose |
|---|---|
| `.ddev/web-build/Dockerfile.opentofu` | installs OpenTofu into the web container (official installer, deb method) |
| `.ddev/commands/web/coolify-bootstrap` | the `init` / `up` / `destroy` command |
| `.ddev/coolify-bootstrap/templates/` | the `infra/` scaffold (consumer config for the tcss module, pinned to a released version) |

## Maintainer notes: template sync

The templates embed the tcss module version once, as `TCSS_VERSION` in
`commands/web/coolify-bootstrap` (`__TCSS_VERSION__` in `templates/main.tf` is
substituted at `init` time). **Release checklist when tcss releases:** bump
`TCSS_VERSION`, re-sync `templates/` against the module repo's
`examples/two-environment/`, run the bats tests, tag.

**Intentional deltas from the example — keep them when re-syncing:** the module
source is `?ref=__TCSS_VERSION__` (not `../../`), and `templates/versions.tf`
carries a `backend "local" { path = "tofu.tfstate" }` block the example omits —
dropping it would scaffold an un-gitignored `terraform.tfstate` full of
plaintext secrets (the bats suite asserts the block is present).

**Maintained by [@vanWittlaer](https://github.com/vanWittlaer)**
