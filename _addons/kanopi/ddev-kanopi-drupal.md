---
title: kanopi/ddev-kanopi-drupal
github_url: https://github.com/kanopi/ddev-kanopi-drupal
description: "This repository provides a DDEV add-on that configures a Drupal development environment with Kanopi's standard tooling and workflows. "
user: kanopi
repo: ddev-kanopi-drupal
repo_id: 1034499133
ddev_version_constraint: ">= v1.22.0"
dependencies: []
type: contrib
created_at: 2025-08-08
updated_at: 2025-10-10
workflow_status: disabled
stars: 2
---

# DDEV Kanopi Drupal Add-on

[![tests](https://github.com/kanopi/ddev-kanopi-drupal/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/kanopi/ddev-kanopi-drupal/actions/workflows/test.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/kanopi/ddev-kanopi-drupal)](https://github.com/kanopi/ddev-kanopi-drupal/commits)
[![release](https://img.shields.io/github/v/release/kanopi/ddev-kanopi-drupal)](https://github.com/kanopi/ddev-kanopi-drupal/releases/latest)
![project is maintained](https://img.shields.io/maintenance/yes/2025.svg)
[![Documentation](https://img.shields.io/badge/docs-mkdocs-blue.svg)](https://kanopi.github.io/ddev-kanopi-drupal/)

A comprehensive DDEV add-on that provides Kanopi's battle-tested workflow for Drupal development. This add-on includes complete tooling for modern Drupal development with multi-provider hosting support.

---

## ✨ Features

- **27+ Custom Commands** - Complete Drupal development workflow
- **Multi-Provider Hosting** - Pantheon and Acquia support
- **Smart Database Refresh** - 12-hour backup age detection
- **Recipe Development** - Drupal 11 recipe creation and management
- **Theme Development** - Node.js/NPM integration with build tools
- **E2E Testing** - Cypress integration with user management
- **Performance Tools** - Critical CSS generation and optimization
- **Service Integration** - PHPMyAdmin, Redis/Memcached, and Solr support

---

## 📚 Documentation

**[📖 Complete Documentation](https://kanopi.github.io/ddev-kanopi-drupal/)**

### Quick Links

| Topic | Description |
|-------|-------------|
| **[🏁 Getting Started](https://kanopi.github.io/ddev-kanopi-drupal/)** | Installation and setup guide |
| **[⚙️ Custom Configuration](https://kanopi.github.io/ddev-kanopi-drupal/custom-configuration/)** | Common customization examples |
| **[🛠 Commands](https://kanopi.github.io/ddev-kanopi-drupal/commands/)** | Complete command reference |
| **[🔧 Troubleshooting](https://kanopi.github.io/ddev-kanopi-drupal/troubleshooting/)** | Common issues and solutions |

---

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

---

## 📋 Requirements

- DDEV v1.22.0 or higher
- **Existing DDEV project** - Must be configured before installing this add-on
- Drupal 8+ project
- Hosting provider account (Pantheon or Acquia) with appropriate credentials
- Node.js (managed by add-on via NVM)

---

## 🔧 Management

### Update
```bash
ddev add-on get kanopi/ddev-kanopi-drupal
```

### Remove
```bash
ddev add-on remove kanopi-drupal
```

---

## 🤝 Contributing

This add-on is maintained by [Kanopi Studios](https://kanopi.com). For issues, feature requests, or contributions, please visit our [GitHub repository](https://github.com/kanopi/ddev-kanopi-drupal).

---

## 📄 License

This project is licensed under the GNU General Public License v2 - see the [LICENSE](https://github.com/kanopi/ddev-kanopi-drupal/blob/main/LICENSE) file for details.
