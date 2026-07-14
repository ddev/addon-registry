---
title: omega8cc/ddev-boa
github_url: https://github.com/omega8cc/ddev-boa
description: "DDEV add-on: pull a BOA-hosted Drupal site's database and files into a local DDEV project"
user: omega8cc
repo: ddev-boa
repo_id: 1299304061
default_branch: main
tag_name: v1.0.1
ddev_version_constraint: ">= v1.24.0"
dependencies: []
type: contrib
created_at: 2026-07-13
updated_at: 2026-07-13
workflow_status: unknown
stars: 0
---

# ddev-boa

Pull your **BOA-hosted** site's database and files into a local
[DDEV](https://ddev.com) project — through your normal BOA limited-shell account, with no
extra credentials and no changes on the server.

This add-on makes an existing DDEV project able to sync **from** a BOA server. It does
**not** install BOA locally and is **not** a BOA server: there is no Ægir/Hostmaster panel,
no Octopus multi-tenancy, no csf, and the local nginx/PHP/DB are DDEV's, not BOA's own
compiled builds. For a full local BOA box, see *BOA Local* (a prebuilt VM/LXC image) —
this add-on is the lighter, far more common need: "make my DDEV behave like my hosted site
and pull from it."

## Requirements

- A site hosted on BOA and your tenant shell account (`oN.ftp`) with your SSH key already
  installed (BOA reuses your SSH identity — the same key you use to `ssh` in).
- DDEV v1.24+.

## Install

```bash
ddev add-on get omega8cc/ddev-boa
```

## Configure

Edit `.ddev/providers/boa.yaml` (or set these in `.ddev/.env`):

| Variable | What | Example |
|---|---|---|
| `BOA_SSH_USER` | your tenant shell account | `o1.ftp` |
| `BOA_HOST` | your BOA server hostname | `server.example.com` |
| `BOA_ALIAS` | the site's Drush alias, no `@` | `mysite-com` |
| `BOA_DRUSH` | `drush11` (default) or `drush`/`drush8` for Drupal 6/7 | `drush11` |
| `BOA_FILES_PATH` | optional; auto-discovered if empty | |

Find the exact alias name (drush8 and drush10/11 alias names differ — dots become hyphens
except the last extension):

```bash
ddev boa-aliases
```

## Match your project to the site (optional)

```bash
ddev boa-config    # writes .ddev/config.boa.yaml, then:
ddev restart
```

`boa-config` queries the site once and sets your project's **PHP version**, **Drupal
project type** and **docroot** to match what BOA actually reports for the site (read from
`drush @alias status`, so it stays in step with the site rather than being hand-copied). It
notes the site's database engine (Percona/MySQL) as a commented, opt-in suggestion — DDEV's
default MariaDB imports BOA database dumps fine, so the database type is left unchanged to
keep things working out of the box. Review the generated `.ddev/config.boa.yaml` before
restarting.

## Use

```bash
ddev auth ssh      # forward your BOA SSH key into the web container (once per session)
ddev pull boa      # pulls the database + files. Never pulls or pushes code.
```

There is intentionally **no `ddev push boa`** — pushing to a live hosted site from a local
box is not something this add-on will do.

### Running Drush locally after a pull

Locally you do **not** use BOA Drush aliases — those are a server-side Ægir concept. Run
your site's own **site-local** Drush, which needs only the docroot and URI, not an alias:

```bash
ddev drush cr                     # ddev drush runs site-local drush with the right root/uri
# equivalently, site-local drush directly:
ddev exec vendor/drush/drush/drush.php --root=/var/www/html/web --uri=$DDEV_PRIMARY_URL cr
```

(Adjust `--root` to your docroot — e.g. `web` for Composer-based sites.) This is the same
site-local Drush model BOA documents for `vdrush`, just on your machine.

## How it works (and why it stays inside the limited shell)

Everything runs as single, allow-listed commands over your existing SSH access:

- **Database**: `drush @alias sql-dump --gzip` streamed to the local project. `sql-dump` is
  permitted by the BOA shell; drush supplies the credentials, so nothing is stored here.
- **Files**: `rsync` over SSH from your site's `files/` directory (read-only).

No server-side changes, no elevated commands, no master/server Drush contexts — only your
own site alias. Every remote command (`drush @alias sql-dump`, `rsync`, `drush aliases`,
`drush @alias status`) is one the BOA limited shell already permits over SSH.

## Scope / honesty

`ddev-boa` reproduces the **runtime your site sees** — its database and files (`ddev pull
boa`) and its PHP version / database / Drupal-type / docroot (`ddev boa-config`), all read
from what BOA reports for the site. It does **not** reproduce the BOA **server** itself:
there is no Ægir/Hostmaster panel, no Octopus multi-tenancy, no csf, and DDEV's nginx/PHP
are stock, not BOA's own compiled builds. Per-site php.ini tuning and BOA's nginx
directives are not exported (they aren't readable through the limited shell); behaviours
that depend on BOA's compiled modules will differ. This is honest **config + data**
parity, not a certification of BOA's stack.

## License

Copyright (C) 2009-2026 Omega8.cc. Free software under the GNU GPL, version 2 or
later — the same license as BOA itself.
