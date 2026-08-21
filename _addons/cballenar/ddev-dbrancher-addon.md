---
title: "cballenar/ddev-dbrancher-addon"
github_url: "https://github.com/cballenar/ddev-dbrancher-addon"
description: "A DDEV add-on that brings isolated, dynamic database routing to your Drupal feature branches. Useful for tricky setups that can't easily use worktrees."
user: "cballenar"
repo: "ddev-dbrancher-addon"
repo_id: 1246962531
default_branch: "main"
tag_name: "v1.1.2"
ddev_version_constraint: ""
dependencies: []
type: "contrib"
created_at: "2026-05-22"
updated_at: "2026-06-15"
workflow_status: "unknown"
stars: 0
---

# DBrancher (DDEV Add-on)

A DDEV add-on that brings isolated, dynamic database routing to your Drupal feature branches. Useful for tricky setups that can't easily use worktrees.

## Features
- **Clones Database**: When checking out a feature branch, it will make a clone for that branch.
- **Auto-Provisioning**: Prompts you `[y/N]` to clone your database when checking out new branches.
- **Private Hook Injection**: Injects the automation directly into your `.git/hooks/post-checkout` folder to avoid enforcing it upon teammates who don't want it.
- **PHP Safe Routing**: Automatically appends the necessary PDO check to your `settings.local.php` file during installation.

### Smart Garbage Collection
To ensure orphaned databases are deleted when their corresponding pull requests are merged, you must configure Git to automatically prune deleted remote branches:

```bash
git config --global fetch.prune true
```
Once configured, the add-on will automatically sweep MariaDB and delete orphaned databases.

## Installation
Run this command from the root of your DDEV project:

```bash
ddev add-on get cballenar/ddev-dbrancher-addon
```

## Manual Commands
- `ddev dbranch list`: Lists all isolated databases and their statuses (Active, Orphaned, etc).
- `ddev dbranch drop <branch_name>`: Drops the isolated database for a specific branch.
- `ddev dbranch drop current`: Drops the isolated database for your active branch.
- `ddev dbranch drop orphaned`: Forces a garbage collection scan for orphaned databases.
- `ddev dbranch init [file_path]`: Imports a database dump and runs the Drupal update pipeline (`updb`, `cim`, `cr`). If no `file_path` is provided, it automatically finds the newest `.sql.gz` dump in your customized search directory.

## Customization
By default, the addon protects `develop,main,master,HEAD` from being dropped and uses the current directory for database dumps. 

You can customize the search directory and the protected branches by modifying the `.ddev/.dbranch-config` file created during installation:

```ini
# .ddev/.dbranch-config
DUMP_DIR="tmp"
PROTECTED_BRANCHES="develop,main,master,HEAD,production,staging"
```
