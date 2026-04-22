---
title: dkd-dobberkau/ddev-typo3-password
github_url: https://github.com/dkd-dobberkau/ddev-typo3-password
description: "DDEV add-on: Set TYPO3 backend user passwords without shell escaping issues"
user: dkd-dobberkau
repo: ddev-typo3-password
repo_id: 1216804617
default_branch: main
tag_name: v1.0.0
ddev_version_constraint: ">= v1.24.0"
dependencies: []
type: contrib
created_at: 2026-04-21
updated_at: 2026-04-21
workflow_status: unknown
stars: 0
---

# DDEV TYPO3 Password

A DDEV add-on that provides a command to set TYPO3 backend user passwords.

## Why?

TYPO3's CLI lacks a command to change an existing backend user's password. The available options all have limitations:

| Option | Limitation |
|--------|-----------|
| `ddev typo3 backend:resetpassword` | Requires email delivery to be configured |
| `ddev typo3 backend:createadmin` | Creates new user, can't update existing |
| Manual DB update via `ddev mysql` | Error-prone, requires knowledge of password hashing |

Additionally, `ddev exec` passes commands through two shells (host + container), which mangles special characters like `!`, `$`, `` ` `` in passwords. This command avoids the issue by passing the password via PHP's `$argv` instead of string interpolation.

## Installation

```bash
ddev add-on get dkd-dobberkau/ddev-typo3-password
ddev restart
```

## Usage

```bash
# Interactive (recommended - prompts for password with confirmation)
ddev typo3-password admin

# With argument
ddev typo3-password admin 'NewPass2024#'

# Non-existent user shows helpful error with available users
ddev typo3-password unknown
```

### Examples

Set password interactively:
```
$ ddev typo3-password admin
New password:
Confirm password:
Password for 'admin' updated successfully.
```

User not found:
```
$ ddev typo3-password unknown
Error: Backend user 'unknown' not found.

Available users:
+-----+----------+-------------------+-------+--------+
| uid | username | email             | admin | status |
+-----+----------+-------------------+-------+--------+
|   1 | admin    | admin@example.com | yes   | active |
+-----+----------+-------------------+-------+--------+
```

## How It Works

- Generates the password hash using PHP's `password_hash()` with `PASSWORD_ARGON2ID` (TYPO3's default algorithm)
- Passes the password via `$argv` to avoid shell escaping issues
- Validates user existence, password length (min 8 characters), and password confirmation in interactive mode
- Only available in TYPO3 project types

## Requirements

- DDEV >= 1.24.0
- TYPO3 project type

## License

Apache License 2.0
