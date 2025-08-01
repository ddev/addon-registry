---
title: OpenForgeProject/ddev-gum
github_url: https://github.com/OpenForgeProject/ddev-gum
description: "A tool for glamorous shell scripts 🎀 like DDEV web commands"
user: OpenForgeProject
repo: ddev-gum
repo_id: 691401674
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2023-09-14
updated_at: 2024-11-21
workflow_status: success
stars: 8
---

<div align="center">
    <a href="https://ddev.com/">
        <img src="https://raw.githubusercontent.com/ddev/ddev/master/images/ddev-logo.svg" alt="DDEV logo" height="80">
    </a>
    <a href="https://github.com/charmbracelet/gum">
        <img src="https://stuff.charm.sh/gum/gum.png" alt="gum Logo" height="80">
    </a>
    <h1 align="center">ddev-gum</h1>
</div>

[![tests](https://github.com/OpenForgeProject/ddev-gum/actions/workflows/tests.yml/badge.svg)](https://github.com/OpenForgeProject/ddev-gum/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2024.svg)

## What is Gum?

> A tool for glamorous shell scripts. Leverage the power of [Bubbles](https://github.com/charmbracelet/bubbles) and [Lip Gloss](https://github.com/charmbracelet/lipgloss) in your scripts and aliases without writing any Go code!
>
> ![gum](https://github.com/OpenForgeProject/ddev-gum/assets/7961978/426f48a5-f02e-423e-a894-ca54bd2cd0d1)

For more information, see the [official repository](https://github.com/charmbracelet/gum) or visit <https://charm.sh/>.

YouTube: [Let's build a conventional commit helper script with Gum!](https://youtube.com/watch?v=vtCwt-4tIto)

YouTube Shorts: [Gum: Write Glamorous Shell Scripts](https://www.youtube.com/watch?v=J7S-qastsqg)

## Installation

```shell
ddev add-on get OpenForgeProject/ddev-gum
ddev restart
```

> [!NOTE]
> For DDEV versions prior to v1.23.5, use `ddev get` instead of `ddev add-on get`.

## Usage

### Host shell

```shell
ddev gum
```

### Inside DDEV web commands

```shell
gum
```

Please refer to the documentation at <https://github.com/charmbracelet/gum#interaction>.
