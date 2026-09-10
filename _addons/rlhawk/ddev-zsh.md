---
title: "rlhawk/ddev-zsh"
github_url: "https://github.com/rlhawk/ddev-zsh"
description: "A DDEV add-on that provides a customizable Zsh environment with Starship, Antidote, Oh My Zsh, Fuzzy Finder, and persistent project-local shell state."
user: "rlhawk"
repo: "ddev-zsh"
repo_id: 1360670062
default_branch: "main"
tag_name: "1.0.5"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-09-07"
updated_at: "2026-09-08"
workflow_status: "success"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/rlhawk/ddev-zsh/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/rlhawk/ddev-zsh/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/rlhawk/ddev-zsh)](https://github.com/rlhawk/ddev-zsh/commits)
[![release](https://img.shields.io/github/v/release/rlhawk/ddev-zsh)](https://github.com/rlhawk/ddev-zsh/releases/latest)

# ddev-zsh

A DDEV add-on that provides a customizable Zsh environment with Starship, Antidote, Oh My Zsh, Fuzzy Finder, and persistent project-local shell state.

## Installed tools

- Zsh
- fzf
- Antidote at `/usr/local/share/antidote`
- Starship at `/usr/local/bin/starship`
- Oh My Zsh at `/usr/local/share/oh-my-zsh`
- `ddev zsh` and `ddev zsh-doctor`
- Project-local persistent history at `.ddev/.zsh/.zsh_history`

## Install

From a configured DDEV project:

```bash
ddev add-on get /absolute/path/to/ddev-zsh
ddev restart
ddev zsh-doctor
ddev zsh
```

## Configuration layers

The add-on installs its generated entry point at:

```text
<project>/.ddev/homeadditions/.zshrc
```

DDEV copies it into the container as `~/.zshrc`. It conditionally sources these optional files in order:

1. `~/.zshrc.global`
2. `~/.zshrc.project`


### Global user configuration

A user may create the following file on their host computer to provide global Zsh configuration for all DDEV projects that have this add-on installed:

```text
~/.ddev/homeadditions/.zshrc.global
```

### Project or team configuration

A project may provide:

```text
<project>/.ddev/homeadditions/.zshrc.project
```

It loads after `.zshrc.global`, so project configuration can extend or override global choices.

Neither optional file is created by the add-on. Users choose whether to initialize Antidote, Starship, Oh My Zsh, fzf bindings, aliases, or other behavior.

## Example `.zshrc.global`

```zsh
source /usr/local/share/antidote/antidote.zsh
antidote load ~/.zsh_plugins.txt

source /usr/share/doc/fzf/examples/key-bindings.zsh 2>/dev/null || true
source /usr/share/doc/fzf/examples/completion.zsh 2>/dev/null || true

eval "$(starship init zsh)"
```

A user choosing Antidote must also provide any referenced bundle file, such as `~/.ddev/homeadditions/.zsh_plugins.txt`.

## Example `.zshrc.project`

```zsh
alias drush='vendor/bin/drush'
export PROJECT_ENV=local
```

## History

The add-on creates `.ddev/.zsh`, bind-mounts it at `/mnt/ddev-zsh`, and sets `HISTFILE=/mnt/ddev-zsh/.zsh_history`. Its nested `.gitignore` ignores itself and all generated state. Nothing in `.ddev/.zsh` needs to be committed.

Be aware that `git clean -xfd` can delete this ignored directory and its history.

## Keeping the Add-on Local to Your Machine

By default, the add-on installs several files into the project's `.ddev` directory. Teams may choose to commit these files so that everyone working on the project gets the same Zsh environment.

If you prefer to keep the add-on local to your workstation and do not want to commit its files, add the following paths to:

```text
.git/info/exclude
```

This file works like a repository-local `.gitignore`. Entries added here affect only your local checkout and are never shared with other developers.

Add the following paths:

```text
.ddev/config.zsh.yaml
.ddev/docker-compose.zsh.yaml
.ddev/commands/web/zsh
.ddev/commands/web/zsh-doctor
.ddev/homeadditions/.zshrc
.ddev/web-build/Dockerfile.zsh
```

You can append them automatically with:

```bash
cat >> .git/info/exclude <<'EOF'

# ddev-zsh
.ddev/config.zsh.yaml
.ddev/docker-compose.zsh.yaml
.ddev/commands/web/zsh
.ddev/commands/web/zsh-doctor
.ddev/homeadditions/.zshrc
.ddev/web-build/Dockerfile.zsh
EOF
```

After doing this:

- The add-on continues to function normally.
- The files remain in your working copy.
- Git will not report them as untracked files.
- Other developers are not affected.
- The repository's `.gitignore` file does not need to be modified.
- The add-on can be used on projects where committing local development tooling is undesirable.

### About `.ddev/.zsh`

The add-on stores project-specific shell history in:

```text
.ddev/.zsh/
```

This directory contains its own `.gitignore` file, which ignores all generated state within the directory, including the shell history file:

```text
.ddev/.zsh/
├── .gitignore
└── .zsh_history
```

Because `.ddev/.zsh` manages its own ignored contents, it does **not** need to be added to `.git/info/exclude`.

The history file is written directly by Zsh through a Docker bind mount and persists across container restarts. The directory is intended to remain local to each developer's workstation and should not be committed to the repository.

## Resetting State

All project-specific shell state is stored in:

```text
.ddev/.zsh/
```

You can safely remove individual directories or files to reset specific tools.

Reset Antidote plugin cache:

```bash
rm -rf .ddev/.zsh/antidote
```

Reset Oh My Zsh cache:

```bash
rm -rf .ddev/.zsh/oh-my-zsh-cache
```

Clear shell history:

```bash
rm .ddev/.zsh/.zsh_history
```

## Windows

The current design targets DDEV running in WSL2. Native Windows/PowerShell DDEV has not been validated.
