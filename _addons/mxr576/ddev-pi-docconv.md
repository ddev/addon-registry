---
title: "mxr576/ddev-pi-docconv"
github_url: "https://github.com/mxr576/ddev-pi-docconv"
description: "Adds deterministic document conversion to PI Coding Agent via Pandoc and headless LibreOffice."
user: "mxr576"
repo: "ddev-pi-docconv"
repo_id: 1451243578
default_branch: "main"
tag_name: "1.0.0-alpha1"
ddev_version_constraint: ">= v1.25.5"
dependencies: ["mxr576/ddev-pi"]
type: "contrib"
created_at: "2026-11-09"
updated_at: "2026-11-09"
workflow_status: "failure"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/mxr576/ddev-pi-docconv/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/mxr576/ddev-pi-docconv/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/mxr576/ddev-pi-docconv)](https://github.com/mxr576/ddev-pi-docconv/commits)
[![release](https://img.shields.io/github/v/release/mxr576/ddev-pi-docconv)](https://github.com/mxr576/ddev-pi-docconv/releases/latest)

# DDEV Pi Docconv

## Overview

This add-on enhances the PI Coding Agent inside the [`ddev-pi`](https://github.com/mxr576/ddev-pi) DDEV add-on with deterministic document conversion and visual page verification capabilities.

It provisions Pandoc, headless LibreOffice, Poppler utilities, and drop-in Microsoft metric-compatible fonts (Carlito for Calibri, Caladea for Cambria, Liberation for Arial/Times/Courier) at image build time, delivering a sandboxed CLI tool (`docconv`) accessible inside the Pi sidecar container and from the host via `ddev docconv`.

## Installation

```bash
ddev add-on get mxr576/ddev-pi-docconv
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

You can invoke `docconv` from your host terminal or within the `pi` container:

| Command | Description |
| ------- | ----------- |
| `ddev docconv --help` | Display CLI usage information and options |
| `ddev docconv convert <file>` | Convert Markdown/DOCX to DOCX, PDF, and per-page PNG images |
| `ddev docconv verify <file> --exact-pages <n>` | Verify that a document meets an exact page count budget |
| `ddev docconv verify <file> --min-pages <n> --max-pages <m>` | Validate page boundary limits |

### Examples

#### Convert Markdown to Word, PDF, and PNG Images
```bash
ddev docconv convert document.md --to docx,pdf,png
```

#### Compile with a Custom Word Template
```bash
ddev docconv convert document.md --reference-doc custom-template.docx
```

#### Verify 2-Page Budget with JSON Output for Agents
```bash
ddev docconv verify CV.md --exact-pages 2 --to-png --json
```

Output:
```json
{
  "status": "success",
  "pages": 2,
  "target_pages": {
    "exact": 2,
    "min": null,
    "max": null
  },
  "artifacts": {
    "docx": "/var/www/html/CV.docx",
    "pdf": "/var/www/html/CV.pdf",
    "images": [
      "/var/www/html/CV-page-1.png",
      "/var/www/html/CV-page-2.png"
    ]
  },
  "errors": []
}
```

## Customization

The default version tag is defined in `docker-compose.pi-docconv.yaml` via `${PI_DOCCONV_VERSION:-main}`.

To override this version locally on your machine without altering tracked repository files:

1. Add the variable to `.ddev/.env.pi`:
   ```ini
   PI_DOCCONV_VERSION="custom-override"
   ```
2. Restart the project:
   ```bash
   ddev restart
   ```

All customization options:

| Variable | Default | Description |
| -------- | ------- | ----------- |
| `PI_DOCCONV_VERSION` | `main` | Version / git ref of Pi Docconv |

## Credits

**Contributed and maintained by [@mxr576](https://github.com/mxr576)**
