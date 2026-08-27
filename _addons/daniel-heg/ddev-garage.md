---
title: "daniel-heg/ddev-garage"
github_url: "https://github.com/daniel-heg/ddev-garage"
description: "DDEV add-on for Garage S3-compatible object storage"
user: "daniel-heg"
repo: "ddev-garage"
repo_id: 1346222110
default_branch: "main"
tag_name: "v0.1.1"
ddev_version_constraint: ">= v1.25.2"
dependencies: []
type: "contrib"
created_at: "2026-08-25"
updated_at: "2026-08-26"
workflow_status: "success"
stars: 0
---

[![tests](https://github.com/daniel-heg/ddev-garage/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/daniel-heg/ddev-garage/actions/workflows/tests.yml?query=branch%3Amain)
[![release](https://img.shields.io/github/v/release/daniel-heg/ddev-garage)](https://github.com/daniel-heg/ddev-garage/releases/latest)

# DDEV Garage

A DDEV add-on for running a persistent, single-node
[Garage](https://garagehq.deuxfleurs.fr/) S3-compatible object store in local development.

The add-on uses the official `dxflrs/garage` image. It creates a default access key and bucket,
enables anonymous reads through Garage's website endpoint, and provides the Garage administration
CLI as `ddev garage`. It does not include a separate Web UI.

## Requirements

- DDEV v1.25.2 or newer

## Installation

```bash
ddev add-on get daniel-heg/ddev-garage
ddev restart
```

Commit the generated `.ddev` files to the consuming project's repository.

## Endpoints and credentials

| Setting | Default |
| --- | --- |
| Internal S3 API | `http://garage:3900` |
| Host S3 API | `https://<project>.ddev.site:3900` |
| Public bucket URL | `https://garage.<project>.ddev.site:3902` |
| Region | `garage` |
| Bucket | `garage` |
| Access key | `garage-access-key` |
| Secret key | `garage-secret-key` |

The credentials are deterministic and intended only for local development.

Laravel/Flysystem can use:

```dotenv
AWS_ACCESS_KEY_ID=garage-access-key
AWS_SECRET_ACCESS_KEY=garage-secret-key
AWS_DEFAULT_REGION=garage
AWS_BUCKET=garage
AWS_ENDPOINT=http://garage:3900
AWS_USE_PATH_STYLE_ENDPOINT=true
AWS_URL=https://garage.<project>.ddev.site:3902
```

`AWS_URL` deliberately has no bucket path. Garage's website endpoint identifies the bucket from the
hostname and provides the anonymous reads that S3 bucket policies would normally provide.

## Garage CLI

```bash
ddev garage status
ddev garage bucket list
ddev garage bucket info garage
ddev garage key list
```

## Configuration

The add-on uses these defaults when `.ddev/.env.garage` is absent:

| Variable | Default |
| --- | --- |
| `GARAGE_DOCKER_IMAGE` | `dxflrs/garage:v2.3.0` pinned by digest |
| `GARAGE_DEFAULT_ACCESS_KEY` | Development access key shown above |
| `GARAGE_DEFAULT_SECRET_KEY` | Development secret shown above |
| `GARAGE_DEFAULT_BUCKET` | `garage` |
| `GARAGE_DEFAULT_BUCKET_PUBLIC` | `true` |
| `GARAGE_LOG_LEVEL` | `garage=info` |

Create the user-managed `.ddev/.env.garage` with `ddev dotenv set` to override a value. The add-on
does not overwrite this file during installation or upgrades. A bucket-name change also changes the
generated DDEV hostname, so reinstall the add-on after changing it:

```bash
ddev dotenv set .ddev/.env.garage --garage-default-bucket=assets
ddev add-on get daniel-heg/ddev-garage
ddev restart
```

Bucket names must be lowercase DNS labels between 3 and 63 characters. Set
`GARAGE_DEFAULT_BUCKET_PUBLIC=false` to keep the default bucket private.

## Persistence and removal

Garage metadata and objects are stored in the named volumes
`ddev-<project>-garage-metadata` and `ddev-<project>-garage-data`. They survive `ddev restart` and
add-on removal.

```bash
ddev add-on remove garage
ddev restart
```

Use `ddev delete -O` only when the project's Garage objects may also be discarded.

## Development

Run the directory-install tests from the add-on root:

```bash
bats ./tests/test.bats --filter-tags '!release'
```

The GitHub Actions workflow tests DDEV stable and HEAD. The Garage image is pinned to its verified
multi-platform digest and supports both amd64 and arm64.

## Credits

Contributed and maintained by [@daniel-heg](https://github.com/daniel-heg).
