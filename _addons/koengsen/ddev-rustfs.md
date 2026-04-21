---
title: koengsen/ddev-rustfs
github_url: https://github.com/koengsen/ddev-rustfs
description: "`ddev-rustfs` is a DDEV add-on that provides a local, S3-compatible object storage using **RustFS**."
user: koengsen
repo: ddev-rustfs
repo_id: 1213419521
default_branch: main
tag_name: 
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: contrib
created_at: 2026-04-17
updated_at: 2026-04-20
workflow_status: failure
stars: 0
---

# ddev-rustfs

`ddev-rustfs` is a DDEV add-on that provides a local, S3-compatible object storage using **RustFS**.

It is intended for local development setups that expect `Storage::disk('s3')` to work (e.g. Laravel).

## What you get

- **S3 API** on port **9000**
- **Console/UI** on port **9001**
- Default credentials: **`rustfsadmin` / `rustfsadmin`**
- Automatically ensures a default bucket exists (configurable)

## Installation

From a local checkout:

```bash
ddev add-on get ./ddev-rustfs
ddev restart
```

From GitHub (once published):

```bash
ddev add-on get <org-or-user>/ddev-rustfs
ddev restart
```

## URLs

With a DDEV project name like `myproj`:

- API: `https://myproj.ddev.site:9000`
- Console: `https://myproj.ddev.site:9001`

## Credentials

- Access key: `rustfsadmin`
- Secret key: `rustfsadmin`

## Laravel `.env` example

Use the internal DDEV service hostname (`rustfs`) for the endpoint:

```dotenv
FILESYSTEM_DISK=s3

AWS_ACCESS_KEY_ID=rustfsadmin
AWS_SECRET_ACCESS_KEY=rustfsadmin
AWS_DEFAULT_REGION=us-east-1
AWS_BUCKET=rustfs

AWS_ENDPOINT=http://rustfs:9000
AWS_USE_PATH_STYLE_ENDPOINT=true
```

Notes:

- `AWS_DEFAULT_REGION` can be any non-empty value for local S3-compatible storage.
- `AWS_USE_PATH_STYLE_ENDPOINT=true` is recommended for local S3-compatible endpoints.
- If your app uses `Storage::url()` you can optionally set `AWS_URL` to a routable URL.

## Default bucket creation

On startup, the add-on runs a small init container that ensures a bucket exists.

- Default bucket name: `rustfs`
- Override with: `AWS__DEFAULT_BUCKET` (double underscore), e.g. in your project's `.env`:

```dotenv
AWS__DEFAULT_BUCKET=rustfs
```

Then set your application's `AWS_BUCKET` to the same value.

## Troubleshooting

- If the container can't write to `/data`, the add-on runs an init container that `chown`s the volume to RustFS' user id `10001`.
- If you change credentials, you may need to delete the `rustfs-data` volume.

## Tests

There is a small Bats smoke test in `tests/test.bats`
