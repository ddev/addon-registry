---
title: kanopi/ddev-kanopi-drupal
github_url: https://github.com/kanopi/ddev-kanopi-drupal
description: "This repository provides a DDEV add-on that configures a WordPress development environment with Kanopi's standard tooling and workflows. "
user: kanopi
repo: ddev-kanopi-drupal
repo_id: 1034499133
ddev_version_constraint: ">= v1.22.0"
dependencies: []
type: contrib
created_at: 2025-08-08
updated_at: 2025-09-25
workflow_status: failure
stars: 1
---

# DDEV Kanopi Drupal Add-on

[![tests](https://github.com/kanopi/ddev-kanopi-drupal/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/kanopi/ddev-kanopi-drupal/actions/workflows/test.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/kanopi/ddev-kanopi-drupal)](https://github.com/kanopi/ddev-kanopi-drupal/commits)
[![release](https://img.shields.io/github/v/release/kanopi/ddev-kanopi-drupal)](https://github.com/kanopi/ddev-kanopi-drupal/releases/latest)
![project is maintained](https://img.shields.io/maintenance/yes/2025.svg)
[![Documentation](https://img.shields.io/badge/docs-mkdocs-blue.svg)](https://kanopi.github.io/ddev-kanopi-drupal/)

A comprehensive DDEV add-on that provides Kanopi's battle-tested workflow for Drupal development with multi-provider hosting support.

## 🚀 Quick Start

```bash
# Install the add-on
ddev add-on get kanopi/ddev-kanopi-drupal

# Configure your hosting provider
ddev project-configure

# Initialize your development environment
ddev project-init
```

## ✨ Features

- **27+ Custom Commands** - Complete Drupal development workflow
- **Multi-Provider Hosting** - Pantheon and Acquia support
- **Smart Database Refresh** - 12-hour backup age detection
- **Recipe Development** - Drupal 11 recipe creation and management
- **Theme Development** - Node.js/NPM integration with build tools
- **E2E Testing** - Cypress integration with user management
- **Performance Tools** - Critical CSS generation and optimization
- **Service Integration** - PHPMyAdmin, Redis/Memcached, and Solr support

## 📚 Documentation

**[📖 Complete Documentation](https://kanopi.github.io/ddev-kanopi-drupal/)**

### Quick Links

| Topic | Description |
|-------|-------------|
| **[🏁 Getting Started](https://kanopi.github.io/ddev-kanopi-drupal/installation/)** | Installation and setup guide |
| **[⚙️ Configuration](https://kanopi.github.io/ddev-kanopi-drupal/configuration/)** | Hosting provider setup |
| **[🛠 Commands](https://kanopi.github.io/ddev-kanopi-drupal/commands/)** | Complete command reference |
| **[🎨 Theme Development](https://kanopi.github.io/ddev-kanopi-drupal/theme-development/)** | Asset compilation and build tools |
| **[📦 Recipe Development](https://kanopi.github.io/ddev-kanopi-drupal/recipe-development/)** | Drupal 11 recipe workflows |
| **[🗄 Database Operations](https://kanopi.github.io/ddev-kanopi-drupal/database-operations/)** | Smart refresh and migrations |
| **[🧪 Testing](https://kanopi.github.io/ddev-kanopi-drupal/testing/)** | Cypress E2E and automated testing |
| **[☁️ Hosting Providers](https://kanopi.github.io/ddev-kanopi-drupal/hosting-providers/)** | Platform-specific guides |
| **[🔧 Troubleshooting](https://kanopi.github.io/ddev-kanopi-drupal/troubleshooting/)** | Common issues and solutions |

## 🏛️ Hosting Providers

### Pantheon Integration
- **Nginx Configuration**: Automatic proxy setup for missing assets
- **Terminus Integration**: Full Pantheon API access with machine token
- **Smart Backups**: 12-hour backup age detection with automatic refresh
- **Redis Caching**: Optimized object caching for Pantheon environments

### Acquia Integration
- **Apache-FPM Configuration**: Native Apache setup matching Acquia Cloud
- **Acquia CLI Integration**: Full Acquia Cloud API access
- **File Proxy**: Apache .htaccess-based proxy for missing files
- **Memcached Caching**: Optimized caching for Acquia environments

## 🛠 Command Overview

### Project Management (5 commands)
- `ddev project-init` - Complete project initialization
- `ddev project-configure` - Interactive configuration wizard
- `ddev project-auth` - SSH key authorization
- `ddev project-lefthook` - Git hooks setup
- `ddev project-nvm` - Node.js environment setup

### Theme Development (5 commands)
- `ddev theme-install` - Set up build tools
- `ddev theme-watch` - Development with file watching
- `ddev theme-build` - Production asset compilation
- `ddev theme-npm` - Run NPM commands
- `ddev theme-npx` - Run NPX commands

### Database Operations (3 commands)
- `ddev db-refresh` - Smart database refresh with backup detection
- `ddev db-rebuild` - Composer install + database refresh
- `ddev db-prep-migrate` - Set up migration database

### Recipe Development (3 commands)
- `ddev recipe-apply` - Apply Drupal recipes
- `ddev recipe-unpack` - Unpack packaged recipes
- `ddev recipe-uuid-rm` - Clean UUIDs for development

### Testing & Quality (3 commands)
- `ddev cypress-install` - Install E2E testing
- `ddev cypress-run` - Run Cypress tests
- `ddev cypress-users` - Create test users

### Performance (2 commands)
- `ddev critical-install` - Install Critical CSS tools
- `ddev critical-run` - Generate critical CSS

### Hosting Integration (4 commands)
- `ddev pantheon-terminus` - Run Terminus commands
- `ddev pantheon-testenv` - Create test environments
- `ddev pantheon-tickle` - Keep environments awake
- `ddev drupal-uli` - Generate login links

### Utilities (2 commands)
- `ddev phpmyadmin` - Launch database interface
- `ddev drupal-open` - Open site in browser

## 🚦 Quick Workflows

### Daily Development
```bash
ddev start               # Start your day
ddev theme-watch         # Begin theme development
```

### Database Refresh
```bash
ddev db-refresh          # Get latest from default environment
ddev db-refresh live -f  # Force refresh from live
```

### Recipe Development
```bash
ddev recipe-apply ../recipes/my-recipe  # Apply recipe
ddev recipe-uuid-rm config/sync         # Clean for development
```

### Testing
```bash
ddev cypress-users       # Create test users
ddev cypress-run open    # Open Cypress UI
```

## 📋 Requirements

- DDEV v1.22.0 or higher
- Drupal 8+ project
- Hosting provider account (Pantheon or Acquia)
- Node.js (managed by add-on via NVM)

## 🔧 Management

### Update
```bash
ddev add-on get kanopi/ddev-kanopi-drupal
```

### Remove
```bash
ddev add-on remove kanopi-drupal
```

## 🤝 Contributing

This add-on is maintained by [Kanopi Studios](https://kanopi.com). For issues, feature requests, or contributions, please visit our [GitHub repository](https://github.com/kanopi/ddev-kanopi-drupal).

## 📄 License

This project is licensed under the GNU General Public License v2 - see the [LICENSE](https://github.com/kanopi/ddev-kanopi-drupal/blob/main/LICENSE) file for details.
