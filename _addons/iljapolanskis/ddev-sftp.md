---
title: "iljapolanskis/ddev-sftp"
github_url: "https://github.com/iljapolanskis/ddev-sftp"
description: "DDEV SFTP Server Addon"
user: "iljapolanskis"
repo: "ddev-sftp"
repo_id: 1028210891
default_branch: "main"
tag_name: "1.0.2"
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: "contrib"
created_at: "2025-07-29"
updated_at: "2026-10-30"
workflow_status: "disabled"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/iljapolanskis/ddev-sftp/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/iljapolanskis/ddev-sftp/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/iljapolanskis/ddev-sftp)](https://github.com/iljapolanskis/ddev-sftp/commits)
[![release](https://img.shields.io/github/v/release/iljapolanskis/ddev-sftp)](https://github.com/iljapolanskis/ddev-sftp/releases/latest)

# DDEV SFTP

## Overview

This add-on integrates SFTP into your [DDEV](https://ddev.com/) project.

## Installation

```bash
ddev add-on get iljapolanskis/ddev-sftp
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev describe` | View service status and used ports for SFTP |
| `ddev logs -s sftp` | Check SFTP logs |

## Configuration

You can customize the SFTP service with `ddev dotenv set`, which writes to `.ddev/.env.sftp`. Keeping the settings there (instead of `.ddev/.env`) stops the SFTP credentials from leaking into every container of the project:

```bash
ddev dotenv set .ddev/.env.sftp --sftp-port=2223 --sftp-password=secret
ddev restart
```

## Configuration Options

| Variable | Default | Description |
| -------- | ------- | ----------- |
| `SFTP_PORT` | `2222` | Host port for SFTP access (bound to `127.0.0.1` only) |
| `SFTP_USERNAME` | `sftp` | SFTP username |
| `SFTP_PASSWORD` | `sftp` | SFTP password |
| `SFTP_UID` | your host UID (`DDEV_UID`) | User ID for SFTP user, so uploaded files are owned by you on the host |
| `SFTP_GID` | your host GID (`DDEV_GID`) | Group ID for SFTP user |
| `SFTP_DIR` | `upload` | Directory name inside SFTP home |
| `SFTP_HOST_PATH` | `../upload/sftp` | Host path to mount as SFTP directory (relative to `.ddev`) |

If several DDEV projects use this add-on at the same time, give each one its own `SFTP_PORT`, otherwise only the first project will start.

## Connecting to SFTP

From your host machine (FileZilla, `sftp`, etc.):
- **Host**: `127.0.0.1` (or your DDEV hostname, e.g. `myproject.ddev.site`)
- **Port**: Value of `SFTP_PORT` (default: 2222)
- **Username**: Value of `SFTP_USERNAME` (default: sftp)
- **Password**: Value of `SFTP_PASSWORD` (default: sftp)

From inside the DDEV containers, use host `sftp` on port `22`.

```bash
sftp -P 2222 sftp@127.0.0.1
```

The user is chrooted to its home directory, so files go into `/upload` (or whatever `SFTP_DIR` is set to); the home directory itself is not writable.

> [!NOTE]
> The SSH host key is regenerated every time the container is recreated (e.g. on `ddev restart`), so clients that remember host keys will warn about a changed key. That is expected for this local-only service.

> [!NOTE]
> The [atmoz/sftp](https://github.com/atmoz/sftp) image is only published for `linux/amd64`. On Apple Silicon it runs under emulation (Rosetta/QEMU), which works fine for local development.

## PHP Usage Example

Here's an example of how to use the SFTP service from within your PHP application:

First, install the required PHP library:

```bash
composer require phpseclib/phpseclib
```

Then use it in your code (playground.php file):

```php
<?php

require 'vendor/autoload.php';

use phpseclib3\Net\SFTP;

$sftp = new SFTP('sftp');
if (!$sftp->login('sftp', 'sftp')) {
    die('Login Failed');
}

$content = "This is a test file.\nLine 2 of the file.\n";
$remoteFilePath = '/upload/test_file_' . time() . '.txt';
if ($sftp->put($remoteFilePath, $content)) {
    echo "File uploaded successfully to $remoteFilePath\n";
} else {
    echo "Failed to upload file to $remoteFilePath\n";
}
```

Note: This example uses the default SFTP service hostname (`sftp`) which is accessible from within the DDEV environment.
```shell
ddev exec php playground.php
```

## Credits

**Contributed and maintained by [@iljapolanskis](https://github.com/iljapolanskis)**

This add-on uses the [atmoz/sftp](https://github.com/atmoz/sftp) Docker image as the base for the SFTP service.
