---
title: kanopi/ddev-kanopi-pantheon-drupal
github_url: https://github.com/kanopi/ddev-kanopi-pantheon-drupal
description: "A comprehensive DDEV add-on that provides Kanopi's battle-tested workflow for Drupal development with Pantheon hosting. This add-on includes 17 custom commands, enhanced Pantheon provider, and complete tooling for modern Drupal development."
user: kanopi
repo: ddev-kanopi-pantheon-drupal
repo_id: 1034499133
ddev_version_constraint: ">= v1.22.0"
dependencies: []
type: contrib
created_at: 2025-08-08
updated_at: 2025-08-08
workflow_status: unknown
stars: 1
---

# DDEV Kanopi Pantheon Drupal Add-on

A comprehensive DDEV add-on that provides Kanopi's battle-tested workflow for Drupal development with Pantheon hosting. This add-on includes 17 custom commands, enhanced Pantheon provider, and complete tooling for modern Drupal development.

## Features

- **🚀 Complete Development Workflow**: From project init to deployment
- **🏛️ Enhanced Pantheon Integration**: Smart backup management and seamless database/file syncing
- **🧪 Cypress Testing Support**: E2E testing with user management
- **🎨 Theme Development Tools**: Node.js/NPM integration with build tools
- **📦 Drupal Recipe Support**: Apply Drupal 11 recipes with cache management
- **🔄 Migration Utilities**: Tools for site migrations and database management
- **⚡ Performance Tools**: Critical CSS generation and asset optimization

## Installation

### For Existing DDEV Projects

If you already have DDEV set up in your project:

```bash
# Install the add-on
ddev add-on get kanopi/ddev-kanopi-pantheon-drupal

# Configure your Pantheon machine token globally
ddev config global --web-environment-add=TERMINUS_MACHINE_TOKEN=your_token_here

# Restart DDEV to apply changes
ddev restart
```

### For Projects Without DDEV

If your project doesn't have DDEV yet, follow these steps:

#### Step 1: Install DDEV (if not already installed)

**macOS:**
```bash
# Using Homebrew (recommended)
brew install ddev

# Or using the installer script
curl -fsSL https://ddev.com/install.sh | bash
```

**Linux/Windows:** Follow instructions at [ddev.readthedocs.io](https://ddev.readthedocs.io/en/stable/users/install/ddev-installation/)

#### Step 2: Initialize DDEV in Your Project

Navigate to your existing project directory and configure DDEV:
```bash
cd /path/to/your-drupal-project

# Initialize DDEV configuration
ddev config --project-type=drupal --docroot=web --php-version=8.3

# Add Pantheon machine token
ddev config global --web-environment-add=TERMINUS_MACHINE_TOKEN=your_pantheon_token

# Start DDEV
ddev start
```

#### Step 3: Install This Add-on

```bash
# Install the Kanopi Pantheon Drupal add-on
ddev add-on get kanopi/ddev-kanopi-pantheon-drupal

# Restart DDEV
ddev restart
```

#### Step 4: Configuration

Configure the add-on for your project:

1. **Configure your Pantheon project** in `.ddev/providers/pantheon.yaml`:
   ```yaml
   environment_variables:
     project: your-site-name.env
   ```

2. **Configure Pantheon and Migration variables** in `.ddev/env.web`:
   ```yaml
THEME=themes/custom/your-theme
THEMENAME=your-theme
hostingsite=your-pantheon-site
hostingenv=dev
MIGRATE_DB_SOURCE=pantheon_source_project
MIGRATE_DB_ENV=live
   ```

3. **Restart DDEV** to apply changes:
   ```bash
   ddev restart
   ```

#### Step 5: Complete Setup

```bash
# Run complete initialization (installs all dependencies and tools)
ddev init

# Open your site
ddev open
```

### Quick Start Summary

For the performant developer:
```bash
# 1. Navigate to your Drupal project
cd your-drupal-project

# 2. Configure DDEV
ddev config --project-type=drupal --docroot=web --php-version=8.3

# 3. Add Pantheon token
ddev config global --web-environment-add=TERMINUS_MACHINE_TOKEN=your_token

# 4. Start DDEV and add add-on
ddev start
ddev add-on get kanopi/ddev-kanopi-pantheon-drupal

# 5. Update variables in .ddev/env.web and .ddev/providers/pantheon.yaml

# 6. Restart and initialize
ddev restart
ddev init

# 7. You're ready to develop!
ddev open
```


## Available Commands

This add-on provides 17 custom commands organized by where they execute (host system vs. web container):

| Command | Type | Description | Example |
|---------|------|-------------|---------|
| `ddev cypress <command>` | Host | Run Cypress commands with environment support | `ddev cypress open` |
| `ddev cypress-users` | Host | Create default admin user for Cypress testing | `ddev cypress-users` |
| `ddev init` | Host | Complete project initialization with all dependencies | `ddev init` |
| `ddev install-cypress` | Host | Install Cypress E2E testing dependencies | `ddev install-cypress` |
| `ddev open` | Host | Open project URL in browser | `ddev open` |
| `ddev phpmyadmin` | Host | Launch PhpMyAdmin | `ddev phpmyadmin` |
| `ddev rebuild` | Host | Composer install + database refresh | `ddev rebuild` |
| `ddev refresh [env] [-f]` | Host | Smart database refresh from Pantheon with backup management | `ddev refresh test -f` |
| `ddev testenv <name> [profile]` | Host | Create isolated testing environment | `ddev testenv my-test minimal` |
| `ddev install-critical-tools` | Web | Install Critical CSS generation tools | `ddev install-critical-tools` |
| `ddev install-theme-tools` | Web | Set up Node.js, NPM, and build tools for theme development | `ddev install-theme-tools` |
| `ddev migrate-prep-db` | Web | Create secondary database for migrations | `ddev migrate-prep-db` |
| `ddev npm <command>` | Web | Run NPM commands in theme directory | `ddev npm run build` |
| `ddev npx <command>` | Web | Run NPX commands in theme directory | `ddev npx webpack` |
| `ddev recipe-apply <path>` | Web | Apply Drupal recipe with cache management | `ddev recipe-apply core/recipes/standard` |
| `ddev tickle` | Web | Keep Pantheon environment awake (useful for migrations) | `ddev tickle` |
| `ddev uuid-rm <path>` | Web | Remove UUIDs from config files (recipe development) | `ddev uuid-rm config/sync` |

## Smart Database Refresh

The enhanced `ddev refresh` command includes intelligent backup management:

- **Automatic Backup Age Detection**: Checks if backups are older than 12 hours
- **Force Flag Support**: Use `-f` to create new backup regardless of age
- **Environment Support**: Refresh from any Pantheon environment (dev, test, live, multidev)
- **Integrated User Creation**: Automatically creates Cypress test users after refresh

```bash
# Refresh from dev (default)
ddev refresh

# Refresh from live environment  
ddev refresh live

# Force new backup creation
ddev refresh -f

# Refresh from multidev environment
ddev refresh pr-123
```

## Theme Development Workflow

1. **Setup**: `ddev install-theme-tools`
2. **Development**: `ddev npm run watch` 
3. **Build**: `ddev npm run build`
4. **Critical CSS**: `ddev install-critical-tools`

## Recipe Development Workflow

1. **Apply Recipe**: `ddev recipe-apply ../recipes/my-recipe`
2. **Clean Config**: `ddev uuid-rm config/sync`
3. **Export Config**: `ddev drush config:export`

## Environment Variables

The add-on automatically configures these environment variables:

- `THEME`: Path to custom theme directory
- `THEMENAME`: Theme name for development tools
- `hostingsite`: Pantheon site identifier
- `hostingenv`: Current environment (dev, test, live)
- `MIGRATE_DB_SOURCE`: Source database for migrations
- `MIGRATE_DB_ENV`: Source environment for migrations

## Managing This Add-on

### View Installed Add-ons
```bash
# List all installed add-ons
ddev add-on list --installed
```

### Update the Add-on
```bash
# Update to the latest version
ddev add-on get kanopi/ddev-kanopi-pantheon-drupal
```

### Remove the Add-on
```bash
# Remove the add-on completely
ddev add-on remove kanopi-pantheon-drupal
```

## Troubleshooting

### Pantheon Authentication Issues
```bash
# Check if token is set
ddev exec printenv TERMINUS_MACHINE_TOKEN

# Re-authenticate manually
ddev exec terminus auth:login --machine-token="your_token"
```

### Theme Build Issues
```bash
# Check Node.js version
ddev exec node --version

# Reinstall dependencies
ddev install-theme-tools
```

### Database Refresh Issues
```bash
# Check Pantheon connection
ddev exec terminus site:list

# Force new backup
ddev refresh -f
```

## Contributing

This add-on is maintained by [Kanopi Studios](https://kanopi.com). For issues, feature requests, or contributions, please visit our [GitHub repository](https://github.com/kanopi/ddev-kanopi-pantheon-drupal).

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/kanopi/ddev-kanopi-pantheon-drupal/blob/main/LICENSE) file for details.
