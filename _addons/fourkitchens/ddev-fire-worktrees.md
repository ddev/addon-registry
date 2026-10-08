---
title: "fourkitchens/ddev-fire-worktrees"
github_url: "https://github.com/fourkitchens/ddev-fire-worktrees"
description: "DDEV add-on that runs each Git worktree of a Drupal site as its own DDEV project, with a copy of the main checkout's database."
user: "fourkitchens"
repo: "ddev-fire-worktrees"
repo_id: 1407733226
default_branch: "main"
tag_name: "v0.1.1"
ddev_version_constraint: ">= v1.25.4"
dependencies: []
type: "contrib"
created_at: "2026-10-06"
updated_at: "2026-10-06"
workflow_status: "unknown"
stars: 1
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/fourkitchens/ddev-fire-worktrees/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/fourkitchens/ddev-fire-worktrees/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/fourkitchens/ddev-fire-worktrees)](https://github.com/fourkitchens/ddev-fire-worktrees/commits)
[![release](https://img.shields.io/github/v/release/fourkitchens/ddev-fire-worktrees)](https://github.com/fourkitchens/ddev-fire-worktrees/releases/latest)

# DDEV Fire Worktrees

## Overview

This add-on runs several branches of a Drupal site side by side. Each
[Git worktree](https://git-scm.com/docs/git-worktree) gets its own DDEV project, database and URLs.
Its database starts as a copy of your main checkout's, then Drupal's deploy steps run so the copy
matches the branch. Your committed `.ddev/config.yaml` never changes.

A standard single-site Drupal project needs no settings. Multi-domain and multisite projects need
a few. The add-on is part of the Fire family with [`ddev-fire`](https://github.com/fourkitchens/ddev-fire),
but neither add-on needs the other.

```bash
ddev fire-worktree-create ABC-123-hero   # new branch and checkout, its own site, a copy of your database
ddev fire-worktree-list                  # every worktree, its status and URL
ddev fire-worktree-remove                # in a worktree: delete its site and checkout, keep the branch
```

## Requirements

* [DDEV](https://ddev.com/) v1.25.4 or later.
* Git, and a project with `.ddev/config.yaml` committed.
* macOS or Linux.

## Installation

```bash
ddev add-on get fourkitchens/ddev-fire-worktrees
git add .ddev
git commit -m "Add ddev-fire-worktrees"
```

Commit the installed files, because a new worktree only has the commands its branch has. The commands
run on your computer rather than in a container, so no `ddev restart` is needed.

Install only one worktree add-on per project. Other public add-ons also define `ddev worktree-remove`,
which this add-on provides as an alias.

## Usage

| Command | Alias | What it does |
| --- | --- | --- |
| `ddev fire-worktree-create <branch> [start-point]` | `worktree-create` | Creates the branch in a directory beside your main checkout, such as `../mysite-abc-123-hero`, then runs startup there. |
| `ddev fire-worktree-start` | `worktree-start` | Gives the worktree you're in its own DDEV project, copies the database into it, and builds it. Safe to rerun. |
| `ddev fire-worktree-list` | | Lists every worktree with its branch, DDEV project, status, uncommitted changes and URL. |
| `ddev fire-worktree-remove [--yes] [--keep-checkout] [path]` | `worktree-remove` | Deletes the worktree's DDEV project and database, then removes the checkout. Keeps the branch. |
| `ddev fire-worktree-prune [--yes]` | | Finds DDEV projects whose checkout is gone, for example after a tool deletes a worktree. Deletes them with `--yes`. |

Run `create`, `list` and `prune` from any checkout of the repository. Run `start` inside a worktree.
Run `remove` inside a worktree, or pass its path. Every command accepts `--help`.

`create` starts from `WORKTREE_BASE_REF`, else your remote's default branch, as it stands locally.
Fetch first to start from the newest commit. `create` refuses a branch or directory that already exists.

`create` and `start` copy your main checkout's database by default. Pass `--db=fresh` for an empty
database, `--db=<file>` for a snapshot or SQL dump, or `--no-build` to skip the build. See
[Choosing the database](https://github.com/fourkitchens/ddev-fire-worktrees/blob/main/docs/databases.md).

`remove` refuses the main checkout and a worktree with uncommitted or untracked files. It asks before
deleting anything; pass `--yes` when no one is at the terminal.

## Documentation

| Guide | Read it to |
| --- | --- |
| [Drupal projects](https://github.com/fourkitchens/ddev-fire-worktrees/blob/main/docs/drupal.md) | Set up a standard site, a FIRE project, or a multi-domain or multisite project. |
| [Configuration](https://github.com/fourkitchens/ddev-fire-worktrees/blob/main/docs/configuration.md) | Change worktree names, copied files, the default database or the build command. |
| [Choosing the database](https://github.com/fourkitchens/ddev-fire-worktrees/blob/main/docs/databases.md) | Copy, reuse or import a database, or start empty. |
| [Lando projects](https://github.com/fourkitchens/ddev-fire-worktrees/blob/main/docs/lando.md) | Use worktrees in a project that also has, or only has, a `.lando.yml`. |
| [T3 and other agent tools](https://github.com/fourkitchens/ddev-fire-worktrees/blob/main/docs/agent-tools.md) | Start worktrees from T3, Claude Code and similar tools, and clean up after them. |
| [Migrating from project-local commands](https://github.com/fourkitchens/ddev-fire-worktrees/blob/main/docs/migrating.md) | Replace a project's own `worktree-*` commands and keep existing worktrees. |
| [How it works](https://github.com/fourkitchens/ddev-fire-worktrees/blob/main/docs/how-it-works.md) | See what startup does and how names, hostnames, locks and prune work. |
| [Troubleshooting](https://github.com/fourkitchens/ddev-fire-worktrees/blob/main/docs/troubleshooting.md) | Fix common errors and warnings. |
| [Contributing](https://github.com/fourkitchens/ddev-fire-worktrees/blob/main/docs/contributing.md) | Run the tests, try changes in a project, and release. |

## Removal

```bash
ddev add-on remove fire-worktrees
```

This removes the commands but not the worktrees' DDEV projects. Run `ddev fire-worktree-remove` in
each worktree first, or delete projects afterwards with `ddev delete <name>`.

## Credits

**Contributed and maintained by [Four Kitchens](https://www.fourkitchens.com/), started by [@mcortes19](https://github.com/mcortes19)**
