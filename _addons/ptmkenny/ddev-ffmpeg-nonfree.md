---
title: ptmkenny/ddev-ffmpeg-nonfree
github_url: https://github.com/ptmkenny/ddev-ffmpeg-nonfree
description: "Build ffmpeg with libfdk-aac from source for DDEV web containers"
user: ptmkenny
repo: ddev-ffmpeg-nonfree
repo_id: 1155821061
default_branch: main
tag_name: 1.0.0
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2026-02-12
updated_at: 2026-02-12
workflow_status: success
stars: 1
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/ptmkenny/ddev-ffmpeg-nonfree/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/ptmkenny/ddev-ffmpeg-nonfree/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/ptmkenny/ddev-ffmpeg-nonfree)](https://github.com/ptmkenny/ddev-ffmpeg-nonfree/commits)
[![release](https://img.shields.io/github/v/release/ptmkenny/ddev-ffmpeg-nonfree)](https://github.com/ptmkenny/ddev-ffmpeg-nonfree/releases/latest)

# DDEV ffmpeg-nonfree

Builds ffmpeg from source inside the DDEV web container with `--enable-nonfree --enable-libfdk-aac`.

## Why?

Debian's packaged ffmpeg excludes `libfdk_aac` because the Fraunhofer FDK AAC license is incompatible with the GPL. This add-on compiles ffmpeg from source, linking against Debian's `libfdk-aac-dev` (from the `non-free` apt component), so you get a fully functional AAC encoder inside your DDEV project.

## Installation

```bash
ddev add-on get ptmkenny/ddev-ffmpeg-nonfree
ddev restart
```

The first `ddev restart` after installation will be slow because ffmpeg is compiled from source inside the web container.

After installation, commit the `.ddev` directory to version control.

## Usage

ffmpeg runs inside the web container:

```bash
ddev exec ffmpeg -version
ddev exec ffmpeg -i input.wav -c:a libfdk_aac -b:a 128k output.m4a
```

## Credits

**Contributed and maintained by [@ptmkenny](https://github.com/ptmkenny)**

Part of the code for this add-on was generated with [Claude Code](https://claude.ai/claude-code).
