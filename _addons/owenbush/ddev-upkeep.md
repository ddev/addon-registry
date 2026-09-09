---
title: "owenbush/ddev-upkeep"
github_url: "https://github.com/owenbush/ddev-upkeep"
description: "Named database fixtures for Drupal projects on DDEV"
user: "owenbush"
repo: "ddev-upkeep"
repo_id: 1316338349
default_branch: "main"
tag_name: "1.0.0"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-07-29"
updated_at: "2026-09-08"
workflow_status: "failure"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/owenbush/ddev-upkeep/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/owenbush/ddev-upkeep/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/owenbush/ddev-upkeep)](https://github.com/owenbush/ddev-upkeep/commits)
[![release](https://img.shields.io/github/v/release/owenbush/ddev-upkeep)](https://github.com/owenbush/ddev-upkeep/releases/latest)

# DDEV Upkeep

## Overview

This add-on integrates Upkeep into your [DDEV](https://ddev.com/) project: it
provides named database **fixtures** — create a gzipped SQL dump of the
current database, reload it later in seconds, and share it via the module's
own repository or a local fixture library.

It is the fixture layer of the
[Upkeep](https://github.com/owenbush/upkeep) maintenance orchestrator (which
installs it into every environment it provisions), but it is equally usable
standalone in any Drupal ddev project — in particular a module checkout using
[ddev-drupal-contrib](https://github.com/ddev/ddev-drupal-contrib).

## Installation

```bash
ddev add-on get owenbush/ddev-upkeep
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

To pin a release, or to install from a checkout while developing the add-on:

```bash
ddev add-on get owenbush/ddev-upkeep --version 1.0.0
ddev add-on get /path/to/ddev-upkeep
```

[Upkeep](https://github.com/owenbush/upkeep) installs this add-on itself, at a
pinned release, into every environment it provisions — set
`UPKEEP_ADDON_SOURCE=/path/to/ddev-upkeep` to point it at a checkout instead.

### Upgrading from the un-namespaced commands

The fixture commands were `ddev fixture-create`, `-load`, `-list` and `-prune`
until they were namespaced. `ddev` puts every add-on's host commands in one
flat namespace per project, so a name as general as `fixture-load` claims
ground this add-on has no business claiming — the next add-on that wants it
has nowhere to go, and nothing in the name says where the command came from.

Re-running `ddev add-on get` performs the upgrade: the namespaced commands are
installed and the old ones are removed, so nothing is left behind under a name
this add-on no longer documents. Removal is guarded on the file being ours —
a `fixture-load` somebody else wrote is not this add-on's to delete.

Update any scripts or CI that call the old names; there is no alias.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev upkeep-fixture-create <name> [--dest=module\|library] [--no-sanitize]` | Dump the current DB to a portable fixture (`<name>.sql.gz`) and write its manifest (`<name>.yml`). Sanitizes via `drush sql:sanitize` by default when destined for the module repo. Warns when the dump exceeds 5 MB (`UPKEEP_FIXTURE_SIZE_WARN_MB`) |
| `ddev upkeep-fixture-load <name>` | Load a fixture: first use imports the dump and materializes a snapshot; later loads restore the snapshot (fast path). Module `tests/fixtures/` shadows the shared library |
| `ddev upkeep-fixture-list` | List fixtures in both scopes with size and snapshot state |
| `ddev upkeep-fixture-prune` | Delete this project's disposable materialized snapshots (never the `.sql.gz` dumps) |
| `ddev describe` | View service status and used ports for Upkeep |
| `ddev logs -s upkeep` | Check Upkeep logs |

Fixtures resolve module-first: `tests/fixtures/<name>.sql.gz` in the module checkout, then the shared library (`UPKEEP_FIXTURE_LIBRARY`, defaulting to `$UPKEEP_COCKPIT/fixtures`, defaulting to `~/.upkeep/fixtures`).

## What a fixture needs: the manifest

A dump is not a self-contained artifact. It encodes references to *code* —
enabled extensions, plugin IDs inside config entities, field types, schema
versions — and captures none of it. Import one into a codebase that does not
provide that code and Drupal cannot build its container; the failure lands in
whatever ran next, and reads as though that thing is broken.

So every dump gets a sidecar beside it:

```
tests/fixtures/smoke.sql.gz     # the dump, unchanged
tests/fixtures/smoke.yml        # what it needs in order to mean anything
```

```yaml
core: '11'
core_version: '11.4.6'
db_engine: 'mariadb:10.11'
created_at: '2026-09-08T18:00:00Z'
require:
  drupal/admin_toolbar: '^3.4'
extensions:
  - admin_toolbar
  - node
  - pathauto
```

**It is generated, never hand-written.** `upkeep-fixture-create` already has
the database open and the project's `composer.json` in front of it, so it
reads the enabled extensions from `core.extension` and the requirements from
the project's own direct dependencies, minus core, drush and the module the
project is a checkout of. Nothing to remember and nothing to keep in step.

`upkeep-fixture-load` then makes the codebase able to hold the dump, or
refuses before touching the database:

1. **Core major must match.** A dump captured on Drupal 11 is not loadable on
   Drupal 10, and finding that out through a schema error later helps nobody.
2. **Declared packages are installed** if the project does not already have
   them, with `ddev composer require`. What it installs is printed.
3. **Every declared extension must then be present.** If one still is not —
   a custom module, or a package the manifest does not name — the load is
   refused and the missing extensions are listed.

Refusing happens *before* the import, so a fixture that cannot work leaves the
database exactly as it was rather than half-replaced.

A dump with no sidecar behaves exactly as fixtures did before manifests
existed: it is imported, and nothing is checked. Nothing that already works
stops working.

The dump itself is untouched by any of this, so `gunzip -c <name>.sql.gz` and
import it with whatever you like remains true.

## The fixture model: dumps vs. snapshots

A fixture is a named gzipped SQL dump, `<name>.sql.gz`. **The dump is the
portable source of truth** — it is what you commit, share, and keep.

On first `ddev upkeep-fixture-load`, the dump is imported and *materialized* into a
fast-format snapshot artifact (`.ddev/upkeep/materialized/<name>.sql`,
streamed straight into the DB server on restore). Subsequent loads restore
that snapshot, which is much faster than re-importing the dump. A metadata
file records the DB engine identity and the dump's checksum at
materialization time; if either changes — you upgraded the database engine,
or the dump was updated — the snapshot is considered stale and is rebuilt
from the dump on the next load.

Materialized snapshots are **disposable, engine-tied local caches — never
authoritative**. `ddev upkeep-fixture-prune` deletes them (and only them); the next
load rebuilds from the dump. Don't commit them.

### Resolution order

`upkeep-fixture-load` resolves a name per-module first:

1. **Module scope:** `tests/fixtures/<name>.sql.gz` in the project root, when
   the project root is a module checkout (an `*.info.yml` at the root — the
   ddev-drupal-contrib layout).
2. **Library scope:** `$UPKEEP_FIXTURE_LIBRARY`, defaulting to
   `$UPKEEP_COCKPIT/fixtures`, defaulting to `~/.upkeep/fixtures`.

A module fixture always shadows a same-named library fixture. Configuration
can come from the caller's environment or from `.ddev/.env.upkeep` (the ddev
dotenv convention); real environment variables win over the dotenv file.

## The `tests/fixtures/` convention (for module maintainers)

This section stands alone: it applies to any Drupal contrib module, whether
or not you or your co-maintainers use Upkeep or this add-on.

- **Path.** Database fixtures live in `tests/fixtures/` in the module
  repository, alongside the module's other test resources.
- **Format and naming.** One gzipped SQL dump per fixture,
  `tests/fixtures/<name>.sql.gz`, and its generated manifest,
  `tests/fixtures/<name>.yml`. Commit both: the dump is what the fixture *is*,
  the manifest is what it *needs*, and a dump whose requirements nobody
  recorded is a dump that only works on the machine it was made on. Names are lowercase, filesystem-safe
  (letters, digits, dot, dash, underscore), and describe the state the
  fixture provides — e.g. `smoke.sql.gz`, `two-vocabularies.sql.gz`,
  `upgrade-from-2x.sql.gz`.
- **Sanitization is mandatory.** A committed fixture is public data. Never
  dump a database containing real user accounts, e-mail addresses, personal
  data, or secrets — sanitize first (e.g. `drush sql:sanitize`) or build the
  fixture from a scratch install that never held real data.
  (`ddev upkeep-fixture-create` runs `drush sql:sanitize` for you by default when
  the destination is the module repo; `--no-sanitize` opts out for databases
  that are already clean.)
- **Keep them lean.** A fixture should contain the *minimum* state that makes
  it useful: a minimal-profile install plus the entities your tests need, not
  a production copy. Gzipped SQL of a minimal Drupal install is well under
  5 MB; treat anything larger as a smell (this add-on warns at that
  threshold) and megabytes of it are usually cache and log tables — truncate
  them before dumping.

Anyone with this add-on can then load your fixture with
`ddev upkeep-fixture-load <name>`; anyone without it can simply
`gunzip -c tests/fixtures/<name>.sql.gz` and import it with the tool of
their choice.

## Advanced Customization

To change the Docker image:

```bash
ddev dotenv set .ddev/.env.upkeep --upkeep-docker-image="ddev/ddev-utilities:latest"
ddev restart
```

then re-run the `ddev add-on get` you installed with (see Installation) if
you want the change reflected in a fresh add-on install.

Make sure to commit the `.ddev/.env.upkeep` file to version control.

All customization options (use with caution):

| Variable | Flag | Default |
| -------- | ---- | ------- |
| `UPKEEP_DOCKER_IMAGE` | `--upkeep-docker-image` | `ddev/ddev-utilities:latest` |
| `UPKEEP_FIXTURE_LIBRARY` | `--upkeep-fixture-library` | `$UPKEEP_COCKPIT/fixtures`, else `~/.upkeep/fixtures` |
| `UPKEEP_FIXTURE_SIZE_WARN_MB` | `--upkeep-fixture-size-warn-mb` | `5` |

## Testing and verification

The bats test suite and how to run it (including the shim-HOME setup some
macOS configurations need) are documented in
[docs/testing.md](https://github.com/owenbush/ddev-upkeep/blob/main/docs/testing.md). A recorded end-to-end pass against a real
contrib module is in
[docs/manual-e2e-conditions-helper.md](https://github.com/owenbush/ddev-upkeep/blob/main/docs/manual-e2e-conditions-helper.md).

## Credits

**Contributed and maintained by [@owenbush](https://github.com/owenbush)**
