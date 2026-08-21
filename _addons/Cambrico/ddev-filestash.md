---
title: "Cambrico/ddev-filestash"
github_url: "https://github.com/Cambrico/ddev-filestash"
description: "📂 DDEV add-on for Filestash — a web UI over your S3/SFTP/WebDAV storage."
user: "Cambrico"
repo: "ddev-filestash"
repo_id: 1300078219
default_branch: "main"
tag_name: ""
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-07-14"
updated_at: "2026-07-14"
workflow_status: "disabled"
stars: 0
---

# DDEV Filestash

[![tests](https://github.com/Cambrico/ddev-filestash/actions/workflows/tests.yml/badge.svg)](https://github.com/Cambrico/ddev-filestash/actions/workflows/tests.yml)
[![project is maintained](https://img.shields.io/maintenance/yes/2026.svg)](https://github.com/Cambrico/ddev-filestash/commits/main)

A [DDEV](https://ddev.com/) add-on that runs [**Filestash**](https://www.filestash.app/)
— a web file manager and UI over your storage backends (S3, SFTP, FTP, WebDAV,
Git, local disk, …) — and publishes it on a dedicated port of the DDEV router.

Filestash is a *frontend* to storage, so after installing you open it and connect
it to whichever backend you want to browse (for example an S3 bucket served by a
`ddev-minio`/`ddev-rustfs` add-on, or a remote SFTP server).

## Requirements

DDEV `v1.24.10` or newer.

## Installation

```bash
ddev add-on get Cambrico/ddev-filestash
ddev restart
```

## Accessing Filestash

| What | URL |
|------|-----|
| Filestash UI (HTTPS) | `https://<project>.ddev.site:8334` |
| Filestash UI (HTTP)  | `http://<project>.ddev.site:8333`  |
| Admin console        | `https://<project>.ddev.site:8334/admin` |

Or just run:

```bash
ddev filestash          # opens the UI in your browser
```

The add-on ships **turnkey** (minio-style): the admin console is pre-seeded with a
known password so you can log in right away.

- **Admin password:** `admin` — **change it** in the admin console after first login.

## Connecting to a storage backend

Two ways, both from the UI:

- **Ad-hoc (home page):** pick a backend type (S3, SFTP, FTP, WebDAV, …), enter the
  connection details, and browse. Nothing is saved.
- **Saved connections (admin console):** in `/admin`, add named connections so they
  appear as one-click buttons on the home page. Their credentials are stored
  encrypted with the per-project `secret_key` (see below).

Backends run on the same Docker network, so reach other DDEV services by their
container name, e.g.:

- **S3 (a `ddev-minio` / `ddev-rustfs` add-on):**
  - Endpoint: `http://ddev-<project>-minio:9000` (use the actual container name)
  - Access/secret key: the add-on's credentials · Region: `us-east-1` · Path style: on
- **SFTP / FTP / WebDAV:** host, port, and credentials of the remote server.
- **Local disk:** browses inside the Filestash container. Mount a host path into the
  service first (add a `volumes:` entry to `.ddev/docker-compose.filestash.yaml`).

See the [Filestash documentation](https://www.filestash.app/docs/) for every backend
type and its options.

## Configuration & data

- **Persistent secret key.** On install, a random 16-char `secret_key` is generated
  **once** and stored in `.ddev/.env.filestash` as `FILESTASH_SECRET_KEY`. DDEV
  injects it into the container; Filestash uses it to encrypt saved connection
  credentials. It is generated only if not already present, so re-installing or
  upgrading the add-on never rotates it (which would invalidate saved connections).
- **State volume.** `config.json` and Filestash's internal database live in the
  `filestash-state` Docker volume and survive `ddev restart`.
- **Change the admin password.** Easiest: log in to `/admin` and change it there.
  Advanced/reproducible: put a bcrypt (`$2a$`) hash in `.ddev/.env.filestash` and
  reseed:
  ```bash
  # generate a hash (needs the apache htpasswd tool):
  HASH=$(htpasswd -bnBC 5 "" 'my-password' | tr -d ':\n' | sed 's/^\$2y\$/$2a$/')
  ddev dotenv set .ddev/.env.filestash --filestash-admin-password-bcrypt="$HASH"
  # then drop the seeded config so it is regenerated:
  ddev exec -s filestash rm -f /app/data/state/config/config_initialized /app/data/state/config/config.json
  ddev restart
  ```
- **Pin the image tag:**
  ```bash
  ddev dotenv set .ddev/.env.filestash --filestash-docker-image=machines/filestash:<tag>
  ddev restart
  ```
- **Ports.** HTTPS `8334` and HTTP `8333` (external) both map to the container's
  `8334`. Change `HTTP_EXPOSE`/`HTTPS_EXPOSE` in `.ddev/docker-compose.filestash.yaml`
  if they clash with another service.

## Documentation

- [Filestash documentation](https://www.filestash.app/docs/) — storage backends,
  admin console, plugins, and API.
- [DDEV add-ons & additional services](https://docs.ddev.com/en/stable/users/extend/additional-services/)
  — how add-ons and custom services work, including the `x-ddev` options used here.

## Removing the add-on

```bash
ddev add-on remove filestash
```

This leaves `.ddev/.env.filestash` and the `filestash-state` volume in place so a
re-install keeps your data. To purge everything:

```bash
rm .ddev/.env.filestash
docker volume rm ddev-<project>-filestash   # the Filestash state volume
```

## Credits

- [Filestash](https://www.filestash.app/) by Mickael Kerjean.
- Built from the [DDEV add-on template](https://github.com/ddev/ddev-addon-template).

**Maintained by [Cambrico](https://github.com/Cambrico).**
