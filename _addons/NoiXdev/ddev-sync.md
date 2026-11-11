---
title: "NoiXdev/ddev-sync"
github_url: "https://github.com/NoiXdev/ddev-sync"
description: "Pull your DDEV project's remote database, files and setup into your local environment."
user: "NoiXdev"
repo: "ddev-sync"
repo_id: 1397686712
default_branch: "main"
tag_name: "v0.1.0-rc.7"
ddev_version_constraint: ">= v1.25.0"
dependencies: []
type: "contrib"
created_at: "2026-09-30"
updated_at: "2026-11-05"
workflow_status: "failure"
stars: 0
---

<img src="assets/logo-256.png" alt="" width="96" align="right">

# ddev-sync

[![CI](https://github.com/NoiXdev/ddev-sync/actions/workflows/ci.yaml/badge.svg)](https://github.com/NoiXdev/ddev-sync/actions/workflows/ci.yaml)
[![Release](https://img.shields.io/github/v/release/NoiXdev/ddev-sync?include_prereleases)](https://github.com/NoiXdev/ddev-sync/releases)
[![Docker Hub: ddev-sync](https://img.shields.io/docker/v/noixdev/ddev-sync?sort=semver&label=ddev-sync)](https://hub.docker.com/r/noixdev/ddev-sync)
[![Docker Hub: ddev-sync-endpoint](https://img.shields.io/docker/v/noixdev/ddev-sync-endpoint?sort=semver&label=ddev-sync-endpoint)](https://hub.docker.com/r/noixdev/ddev-sync-endpoint)
[![License](https://img.shields.io/github/license/NoiXdev/ddev-sync)](LICENSE)
[![DDEV add-on](https://img.shields.io/badge/DDEV-add--on-02a8e2)](https://addons.ddev.com)
[![Discord](https://img.shields.io/discord/1550965846594289727?label=Discord&logo=discord&logoColor=white)](https://discord.gg/sHcvJfXhq)

Pull a project's production database and files into a local DDEV environment,
anonymise the data on the way in, and run the framework's own setup commands
afterwards, all with one command.

## Quick start

```bash
cd /your/project
ddev add-on get NoiXdev/ddev-sync --version v0.1.0-rc.7
ddev auth ssh && ddev restart
ddev sync init
```

## Usage

```bash
ddev sync             # pick the target (if several) and the steps; db, files and sanitize are pre-ticked
ddev sync --dry-run   # print the plan without executing anything
ddev sync doctor      # check that a sync will work before it moves a byte
```

Supported project types: TYPO3, Laravel, Symfony, WordPress, Drupal and
Shopware 6, plus a generic fallback for any other DDEV project type.

## Documentation

- [Introduction](https://docs.noix.dev/ddev-sync/getting-started/introduction/)
- [Installation](https://docs.noix.dev/ddev-sync/getting-started/installation/)
- [Your first sync](https://docs.noix.dev/ddev-sync/getting-started/first-sync/)
- [Docker endpoint](https://docs.noix.dev/ddev-sync/guide/docker-endpoint/)
- [Sanitizing](https://docs.noix.dev/ddev-sync/guide/sanitizing/)
- [Configuration reference](https://docs.noix.dev/ddev-sync/reference/configuration/)
- [Troubleshooting](https://docs.noix.dev/ddev-sync/reference/troubleshooting/)

## Help

Questions and chat on [Discord](https://discord.gg/sHcvJfXhq), bugs and
requests in the [issue tracker](https://github.com/NoiXdev/ddev-sync/issues).
Report vulnerabilities as described in [SECURITY.md](https://github.com/NoiXdev/ddev-sync/blob/main/SECURITY.md).

## Contributing

See [CONTRIBUTING.md](https://github.com/NoiXdev/ddev-sync/blob/main/CONTRIBUTING.md) for working from a checkout, the test
suites and the conventions.

## License

[MIT](https://github.com/NoiXdev/ddev-sync/blob/main/LICENSE) © noix.dev
