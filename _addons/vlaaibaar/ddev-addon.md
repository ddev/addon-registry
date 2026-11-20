---
title: "vlaaibaar/ddev-addon"
github_url: "https://github.com/vlaaibaar/ddev-addon"
description: "Shared DDEV add-on for Vlaaibaar projects"
user: "vlaaibaar"
repo: "ddev-addon"
repo_id: 1462461683
default_branch: "main"
tag_name: "v1.1.1"
ddev_version_constraint: ">= v1.24.0"
dependencies: []
type: "contrib"
created_at: "2026-11-18"
updated_at: "2026-11-19"
workflow_status: "unknown"
stars: 0
---

# vlaaibaar/ddev-addon

Shared DDEV features for Vlaaibaar projects. Install or update in a project with:

```
ddev add-on get vlaaibaar/ddev-addon
ddev restart
```

Updating overwrites only the add-on's own files (marked `#ddev-generated`). Project settings are never overwritten.

## Sync: pull/push database and uploads

Adds the DDEV providers `production` and `staging` for Craft CMS sites.

1. Fill in the shared settings in `.ddev/vlaaibaar/sync/production.env` and `staging.env` (created on install). Commit them.
2. Each developer creates their own `production.local.env` / `staging.local.env` next to them, with only their SSH user (see `example.local.env`). These files are gitignored.
3. Use it:

```
ddev auth ssh                       # once per session
ddev pull production                # DB + uploads -> local
ddev pull staging --skip-files      # DB only
ddev pull staging --skip-db         # uploads only
ddev push staging                   # local DB + uploads -> staging
ddev push staging --skip-files      # DB only
ddev push staging --skip-db         # uploads only
```

Add `-y` to skip the confirmation prompt.

| Setting | Meaning |
| --- | --- |
| `SSH_HOST` | Host or IP of the server, without a user. An older `user@host` value still works |
| `SSH_USER` | Your own SSH user. Set it in `<env>.local.env`, not in the shared file. SSH keys handle authentication |
| `REMOTE_PATH` | Folder with the `craft` executable. Relative paths start from the SSH user's home folder |
| `REMOTE_UPLOADS` | Uploads folder on the server |
| `ALLOW_PUSH` | `true` allows `ddev push`. Created as `false` for production, `true` for staging |

How it works:
- The database is exported and restored with Craft's `php craft db/backup` / `db/restore` on the server, so no database passwords are needed locally. After a push, `php craft up` runs.
- Pulling uploads mirrors the server (local-only files are removed). Pushing uploads never deletes files on the server.
- Local uploads go to the project's `upload_dirs` (set in `.ddev/config.yaml`).
