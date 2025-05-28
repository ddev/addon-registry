---
title: MurzNN/ddev-kubernetes
github_url: https://github.com/MurzNN/ddev-kubernetes
description: "Adds Kubernetes tools (kubectl, krew plugins) to DDEV projects"
user: MurzNN
repo: ddev-kubernetes
repo_id: 951743461
ddev_version_constraint: ">= v1.24.2"
dependencies: []
type: contrib
created_at: 2025-03-20
updated_at: 2025-05-27
stars: 1
---

[![tests](https://github.com/MurzNN/ddev-kubernetes/actions/workflows/tests.yml/badge.svg)](https://github.com/MurzNN/ddev-kubernetes/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2025.svg)

# DDEV Kubernetes add-on <!-- omit in toc -->

This is an add-on for the [DDEV](https://ddev.readthedocs.io) project that adds utilities to work with Kubernetes inside the DDEV Web container.

After adding, these commands will be available inside the `web` container:
- `kubectl` (and a shortcut as just `k`).
- `helm`.
- `k krew` - [details »](https://krew.sigs.k8s.io/)
- `k ns` and `k ctx` - [details »](https://github.com/ahmetb/kubectx)

## Installation

Run a command:

```sh
ddev add-on get MurzNN/ddev-kubernetes
```

Then, restart your project

```sh
ddev restart
```

**Contributed and maintained by [@MurzNN](https://github.com/MurzNN) (Alexey Murz Korepov).**
