---
title: "massiws/ddev-bearer"
github_url: "https://github.com/massiws/ddev-bearer"
description: "Integrate Bearer into DDEV to scan your projects against top security and privacy risks"
user: "massiws"
repo: "ddev-bearer"
repo_id: 1156317873
default_branch: "main"
tag_name: "v1.0.1"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-02-12"
updated_at: "2026-05-05"
workflow_status: "disabled"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/massiws/ddev-bearer/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/massiws/ddev-bearer/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/massiws/ddev-bearer)](https://github.com/massiws/ddev-bearer/commits)
[![release](https://img.shields.io/github/v/release/massiws/ddev-bearer)](https://github.com/massiws/ddev-bearer/releases/latest)

# DDEV Bearer

## Overview

This add-on integrates [Bearer](https://github.com/bearer/bearer), a powerful static application security testing (SAST) tool, directly into your [DDEV](https://ddev.com/) project.

It allows you to **scan your source code for security and privacy vulnerabilities** without leaving the DDEV workflow. 

### Key Benefits

- **Early Detection**: Discover security vulnerabilities in your codebase during development
- **Comprehensive Coverage**: Detects risks based on [OWASP Top 10](https://owasp.org/www-project-top-ten/) and [CWE Top 25](https://cwe.mitre.org/top25/archive/2023/2023_top25_list.html) standards
- **Data Flow Analysis**: Analyzes your source code's data flows to identify security and privacy risks
- **Flexible Scanning**: Customize scans by severity level, scanner type, and context
- **Built-in Rules**: Pre-configured rules for common security vulnerabilities

## Main Features

- **Multi-Scanner Support**: Scan for various vulnerability types including secrets, credentials, and OWASP violations
- **Severity Filtering**: Focus on critical and high-severity issues or customize to your needs
- **Configurable Rules**: Customize scanning behavior through `bearer.yml` configuration file
- **Fingerprint-based Ignore Management**: Mark specific findings as false positives and track them with author and comment metadata
- **Context-Aware Scanning**: Run scans with specific contexts (e.g., health, payment) for targeted analysis

## Installation

```bash
ddev add-on get massiws/ddev-bearer
ddev restart

# Generates a default config to `bearer.yml`
ddev bearer init
```

After installation, you may want to customize the default configuration settings by editing the `bearer.yml` file:
see [Bearer configuration docs](https://docs.bearer.com/reference/config/) for more information.

**Important**: Restart DDEV after making changes to the `bearer.yml` file.

Make sure to commit the `.ddev` directory and the `bearer.yml` file to version control.

## Usage

| Command | Description |
|---------|-------------|
| `ddev bearer` | Display available commands and usage information |
| `ddev bearer scan .` | Scan entire project using default configuration in `bearer.yml` |
| `ddev bearer scan . --severity critical,high` | Scan only for critical and high-severity issues |
| `ddev bearer scan . --scanner=secrets` | Scan specifically for hardcoded credentials and secrets |
| `ddev bearer scan <path/to/file> --context=health` | Scan specific file/folder with custom context (health, payment, finance, etc.) |
| `ddev bearer ignore add <fingerprint> --author "Your Name" --comment "Reason"` | Add fingerprint to ignore file and track with author/comment metadata |

### Tips & Tricks

**Auto-fill Author in Ignore Commands**: To avoid specifying the author repeatedly, configure your [Git username in DDEV globals](https://docs.ddev.com/en/stable/users/extend/in-container-configuration/#git-configuration):

```bash
ln -s $HOME/.gitconfig $HOME/.ddev/homeadditions/.gitconfig
```

This will automatically use your Git username for all `ddev bearer ignore` commands.

## Documentation

For comprehensive information, refer to the [Bearer CLI documentation](https://docs.bearer.com):
- [All available commands](https://docs.bearer.com/reference/commands)
- [Flags and options](https://docs.bearer.com/reference/commands/#bearer%20init-flags)
- [CLI configuration](https://docs.bearer.com/reference/config/)
- [Scanner types](https://docs.bearer.com/explanations/scanners/)

## Contributing & Support

**Contributed and maintained by [@massiws](https://github.com/massiws)**

For issues, feature requests, or contributions, please visit the [GitHub repository](https://github.com/massiws/ddev-bearer).
