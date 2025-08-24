---
title: iljapolanskis/ddev-sftp
github_url: https://github.com/iljapolanskis/ddev-sftp
description: "DDEV SFTP Server Addon"
user: iljapolanskis
repo: ddev-sftp
repo_id: 1028210891
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2025-07-29
updated_at: 2025-08-22
workflow_status: success
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

You can customize the SFTP service by setting environment variables in your `.ddev/.env` file:

```bash
# SFTP port (host port mapping)
SFTP_PORT=2222

# SFTP credentials
SFTP_USERNAME=sftp
SFTP_PASSWORD=sftp

# SFTP directory and user/group IDs
SFTP_UID=1001
SFTP_GID=1001
SFTP_DIR=upload

# Host path to mount
SFTP_HOST_PATH=../upload/sftp
```

After making changes, restart DDEV:

```bash
ddev restart
```

## Configuration Options

| Variable | Default | Description |
| -------- | ------- | ----------- |
| `SFTP_PORT` | `2222` | Host port for SFTP access |
| `SFTP_USERNAME` | `sftp` | SFTP username |
| `SFTP_PASSWORD` | `sftp` | SFTP password |
| `SFTP_UID` | `1001` | User ID for SFTP user |
| `SFTP_GID` | `1001` | Group ID for SFTP user |
| `SFTP_DIR` | `upload` | Directory name inside SFTP home |
| `SFTP_HOST_PATH` | `../upload/sftp` | Host path to mount as SFTP directory |

## Connecting to SFTP

Once configured, you can connect to the SFTP service using:
- **Host**: `localhost` (or your DDEV hostname)
- **Port**: Value of `SFTP_PORT` (default: 2222)
- **Username**: Value of `SFTP_USERNAME` (default: sftp)
- **Password**: Value of `SFTP_PASSWORD` (default: sftp)

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
