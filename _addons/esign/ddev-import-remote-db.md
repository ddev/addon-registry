---
title: esign/ddev-import-remote-db
github_url: https://github.com/esign/ddev-import-remote-db
description: "Import a remote MySQL database into your local DDEV database."
user: esign
repo: ddev-import-remote-db
repo_id: 1046928858
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2025-08-29
updated_at: 2025-09-25
workflow_status: disabled
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/esign/ddev-import-remote-db/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/esign/ddev-import-remote-db/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/esign/ddev-import-remote-db)](https://github.com/esign/ddev-import-remote-db/commits)
[![release](https://img.shields.io/github/v/release/esign/ddev-import-remote-db)](https://github.com/esign/ddev-import-remote-db/releases/latest)

# DDEV Import Remote Db

## Overview

This add-on allows you to import a remote MySQL database into your local [DDEV](https://ddev.com/) database.
It supports direct connections as well as SSH tunneling for secure access to remote databases.

## Installation

```bash
ddev add-on get esign/ddev-import-remote-db
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Usage


### Commands

#### Import a remote MySQL database

```bash
ddev import-remote-db --host <host> --user <user> --database <database> [--port <port>] [--ssh-host <ssh_host>] [--ssh-user <ssh_user>] [--password <password>]
```

**Flags:**

| Flag | Description | Required |
|------|-------------|----------|
| `--host` | Remote MySQL host | Yes |
| `--user` | MySQL username | Yes |
| `--database` | Database name | Yes |
| `--port` | MySQL port (default: 3306) | No |
| `--ssh-host` | SSH host for tunneling | No |
| `--ssh-user` | SSH username | No |
| `--password` | MySQL password | No (will prompt if not provided) |

**Example:**

```bash
ddev import-remote-db --host db.example.com --user dbuser --database dbname --port 3307 --ssh-host ssh.example.com --ssh-user remoteuser --password mypass
```

#### Import a remote MySQL database using 1Password

```bash
ddev import-remote-db-1pw <1pw-item-uuid>
```

This command uses the [1Password CLI](https://developer.1password.com/docs/cli/get-started/) to securely fetch database credentials from your 1Password vault.

**Steps:**
1. You will be prompted for the 1Password item UUID (or you can pass it as an argument).
2. The script will fetch the database credentials and (optionally) SSH details from the 1Password item.
3. You will be asked if SSH is required. If so, you can provide SSH host/user if not present in the item.
4. The script will run `ddev import-remote-db` with the fetched credentials.

**1Password Item Fields:**

The script will look for the following fields in your 1Password item (it tries multiple field names for flexibility):

| Purpose | Field Names (tries in order) | Required |
|---------|------------------------------|----------|
| Database Host | `host`, `server`, `DB_HOST` | Yes |
| Database Username | `username`, `DB_USERNAME` | Yes |
| Database Password | `password`, `DB_PASSWORD` | Yes |
| Database Name | `database`, `DB_DATABASE` | Yes |
| SSH Host | `SSH_HOST` | No (prompted if SSH required) |
| SSH Username | `SSH_USER` | No (prompted if SSH required) |

**Example:**

```bash
ddev import-remote-db-1pw 0198f64f-5ac6-79a5-8238-750890b87ce5
```

**Note:** The 1Password CLI (`op`) must be installed and you must be signed in. The script will check for this and prompt you if not.

---

| Command | Description |
| ------- | ----------- |
| `ddev import-remote-db` | Import a remote MySQL database using direct or SSH connection |
| `ddev import-remote-db-1pw` | Import a remote MySQL database using credentials from 1Password |
| `ddev describe` | View service status and used ports for Import Remote Db |
| `ddev logs -s import-remote-db` | Check Import Remote Db logs |

## Advanced Customization

To change the Docker image:

```bash
ddev dotenv set .ddev/.env.import-remote-db --import-remote-db-docker-image="ddev/ddev-utilities:latest"
ddev add-on get esign/ddev-import-remote-db
ddev restart
```

Make sure to commit the `.ddev/.env.import-remote-db` file to version control.

All customization options (use with caution):

| Variable | Flag | Default |
| -------- | ---- | ------- |
| `IMPORT_REMOTE_DB_DOCKER_IMAGE` | `--import-remote-db-docker-image` | `ddev/ddev-utilities:latest` |

## Credits

**Contributed and maintained by [@esign](https://github.com/esign)**
