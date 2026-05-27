---
title: s-ayers/ddev-drupal-contrib-mkdocs
github_url: https://github.com/s-ayers/ddev-drupal-contrib-mkdocs
description: "ddev addon to provide a mkdocs container for Drupal project documentation (GitLab Pages)"
user: s-ayers
repo: ddev-drupal-contrib-mkdocs
repo_id: 1250471075
default_branch: main
tag_name: 1.0.0
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2026-05-26
updated_at: 2026-05-26
workflow_status: unknown
stars: 0
---

# DDEV Drupal contrib MkDocs add-on

[![tests](https://github.com/Metadrop/ddev-drupal-contrib-mkdocs/actions/workflows/tests.yml/badge.svg)](https://github.com/Metadrop/ddev-drupal-contrib-mkdocs/actions/workflows/tests.yml)

A [DDEV](https://ddev.readthedocs.io) add-on for building and previewing [MkDocs](https://www.mkdocs.org/) documentation in **Drupal contrib** module and theme projects.

Run DDEV from the extension root (where the `.info.yml` file lives). MkDocs uses the standard layout: `mkdocs.yml` at the project root and Markdown sources in `docs/`.

## Getting started

For DDEV v1.23.5 or above:

```shell
ddev add-on get s-ayers/ddev-drupal-contrib-mkdocs
```

For earlier versions of DDEV:

```shell
ddev get s-ayers/ddev-drupal-contrib-mkdocs
```

Restart the project:

```shell
ddev restart
```

## Project layout

After install (only missing files are created):

```
my_module/
├── mkdocs.yml          # MkDocs config (not overwritten if present)
├── docs/
│   └── index.md        # Home page (not overwritten if present)
└── .ddev/
    └── mkdocs/
        └── mkdocs.yml  # Add-on template copy
```

The add-on does **not** create a nested `docs/docs/` directory. It never overwrites an existing `mkdocs.yml` or `docs/index.md`.

If you previously used a layout with `docs/mkdocs.yml` and content under `docs/docs/`, move `mkdocs.yml` to the project root and your Markdown files into `docs/` manually.

## Using MkDocs

The service uses [MkDocs Material](https://squidfunk.github.io/mkdocs-material/) with PHP syntax highlighting suitable for Drupal code samples.

```shell
ddev mkdocs build
```

View the site at `https://docs.<project_name>.ddev.site` (HTTPS) or `http://docs.<project_name>.ddev.site` (HTTP), where `<project_name>` is your DDEV sitename. If DDEV uses alternate router ports because 80/443 are busy (see [port conflict](https://ddev.com/s/port-conflict)), append the port the same way as the main site, e.g. `https://docs.<project_name>.ddev.site:33001`. Run `ddev describe` for the authoritative URLs. The add-on registers `docs.<sitename>` in `additional_hostnames` on install.

The `mkdocs` service runs `mkdocs serve` with live reload. After you change Markdown under `docs/`, the site should rebuild and the browser should refresh. If it does not, run `ddev logs -s mkdocs` while saving a file; you should see a rebuild. On Docker Desktop for Mac, `WATCHDOG_FORCE_POLLING=true` is set to improve file watching across the bind mount.

### Configuration

The MkDocs container mounts the **project root** (`$DDEV_APPROOT`), so `mkdocs.yml` and the `docs/` folder resolve as MkDocs expects. Change the mount or `working_dir` in `docker-compose.mkdocs.yaml` only if you use a non-standard layout.

### Further reading

- [MkDocs getting started](https://www.mkdocs.org/getting-started/)
- [MkDocs user guide](https://www.mkdocs.org/user-guide/)
- [Material for MkDocs](https://squidfunk.github.io/mkdocs-material/)
