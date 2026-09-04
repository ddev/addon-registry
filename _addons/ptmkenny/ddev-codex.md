---
title: "ptmkenny/ddev-codex"
github_url: "https://github.com/ptmkenny/ddev-codex"
description: "An up-to-date ddev add-on for running codex in your project with bubblewrap."
user: "ptmkenny"
repo: "ddev-codex"
repo_id: 1354489766
default_branch: "main"
tag_name: "v1.0.0"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-09-02"
updated_at: "2026-09-02"
workflow_status: "success"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/ptmkenny/ddev-codex/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/ptmkenny/ddev-codex/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/ptmkenny/ddev-codex)](https://github.com/ptmkenny/ddev-codex/commits)
[![release](https://img.shields.io/github/v/release/ptmkenny/ddev-codex)](https://github.com/ptmkenny/ddev-codex/releases/latest)

# DDEV Codex

This add-on installs the [OpenAI Codex CLI](https://learn.chatgpt.com/docs/codex/cli)
in a dedicated DDEV sidecar. The sidecar uses the project's built web image, so
it has the same PHP, Composer, Node.js, and project-specific tools as the web
container. It supports AMD64 and ARM64 hosts.

## AI Usage Disclosure

This ddev addon has been written with codex assistance.

## Installation

```bash
ddev add-on get ptmkenny/ddev-codex
ddev restart
```

Commit the generated files under `.ddev`, including `.ddev/codex/.gitignore`.
Codex state and credentials within `.ddev/codex/` will be ignored.

## Authentication

If you have a ChatGPT subscription, enable device code authorization in your
ChatGPT security settings and run:

```bash
ddev codex login --device-auth
```

Open the displayed URL in a browser and enter the one-time code. Codex stores
the resulting credentials in `.ddev/codex/`, where they persist across DDEV
restarts.

To use an API key instead, pipe it to Codex from your host environment:

```bash
printf '%s' "$OPENAI_API_KEY" | ddev codex login --with-api-key
```

Treat `.ddev/codex/auth.json` as a secret. Do not commit or share it.

## Usage

```bash
# Start an interactive session in the project root.
ddev codex

# Show CLI help or the installed version.
ddev codex --help
ddev codex --version

# Run a prompt directly.
ddev codex "explain this project"
```

The sidecar starts with the project and mounts the same project filesystem as
the web container. This includes DDEV's Mutagen volume when Mutagen is enabled,
so Codex does not use a second, slower project bind mount. Commands such as
PHPCS, PHPStan, Composer, and Drush run directly in the sidecar:

```bash
ddev exec --service codex vendor/bin/phpcs
ddev exec --service codex vendor/bin/phpstan analyse
ddev exec --service codex drush status
```

The add-on also installs Bubblewrap and Python Pillow in the shared project
image. Pillow supports Codex image-generation helper scripts that need the
`PIL` Python package.

## Codex version

The latest Codex release is installed by default. To pin a release, create
`.ddev/.env.codex`:

```dotenv
CODEX_VERSION=0.152.1
```

Then rebuild the project image without its Docker cache:

```bash
ddev restart --no-cache
```

Use the same command to update a project that follows `latest`.

## Linux sandbox inside DDEV

Codex's Linux sandbox uses Bubblewrap and nested user and mount namespaces.
Docker's default seccomp, AppArmor, and protected-system-path settings block
operations Bubblewrap needs. Only the Codex sidecar disables those three outer
restrictions with `seccomp=unconfined`, `apparmor=unconfined`, and
`systempaths=unconfined`. The web and database containers retain their normal
Docker security profiles. This is the same three-option configuration
[Moby documents for a rootless sandbox running inside Docker](https://github.com/moby/buildkit/blob/master/docs/rootless.md#docker).

The sidecar also drops all Linux capabilities and enables
`no-new-privileges`. It is not privileged and does not mount the Docker socket
or DDEV SSH agent. Codex can modify the project and connect to DDEV services by
their Compose names, such as `web` and `db`, but it cannot use Docker to inspect
or execute commands inside those containers.

`systempaths=unconfined` removes Docker's masks from sensitive `/proc` and
`/sys` paths inside the sidecar. This is necessary for Bubblewrap to create its
inner mount layout, but it weakens the sidecar's outer Docker boundary. Keep the
sidecar non-root and do not add capabilities or host-level mounts.

DDEV injects variables from `.ddev/.env` and `.ddev/.env.local` into every
project service, including this sidecar. Put web-only secrets in a targeted
file such as `.ddev/.env.web.local` when Codex should not receive them.

The idle sidecar runs only a keepalive process. It reuses the already-built web
image and DDEV project mount, so the steady-state CPU impact is negligible and
the main additional cost is one small container process and its writable layer.

### Optional SSH agent access

SSH agent access is disabled by default. To opt in, add
`.ddev/docker-compose.codex-ssh-agent.yaml`:

```yaml
services:
  codex:
    environment:
      - SSH_AUTH_SOCK=/home/.ssh-agent/socket
    volumes:
      - ddev-ssh-agent_socket_dir:/home/.ssh-agent

volumes:
  ddev-ssh-agent_socket_dir:
    external: true
```

## Removal

```bash
ddev add-on remove codex
```

DDEV preserves `.ddev/codex/` when it contains Codex state. Remove that
directory manually if you also want to delete the stored credentials and
configuration.

## Credits

This add-on was inspired by
[`Gonzalo2683/ddev-codex`](https://github.com/Gonzalo2683/ddev-codex).
Patrick Kenny maintains this implementation at
[`ptmkenny/ddev-codex`](https://github.com/ptmkenny/ddev-codex).

## License

Apache License 2.0. See [LICENSE](https://github.com/ptmkenny/ddev-codex/blob/main/LICENSE).
