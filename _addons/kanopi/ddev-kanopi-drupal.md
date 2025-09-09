---
title: kanopi/ddev-kanopi-drupal
github_url: https://github.com/kanopi/ddev-kanopi-drupal
description: "A comprehensive DDEV add-on that provides Kanopi's battle-tested workflow for Drupal development with Pantheon hosting. This add-on includes 17 custom commands, enhanced Pantheon provider, and complete tooling for modern Drupal development."
user: kanopi
repo: ddev-kanopi-drupal
repo_id: 1034499133
ddev_version_constraint: ">= v1.22.0"
dependencies: []
type: contrib
created_at: 2025-08-08
updated_at: 2025-09-06
workflow_status: unknown
stars: 1
---

# DDEV Kanopi Pantheon Drupal Add-on

A comprehensive DDEV add-on that provides Kanopi's battle-tested workflow for Drupal development with Pantheon hosting. This add-on includes 17 custom commands, enhanced Pantheon provider, and complete tooling for modern Drupal development.

## Features

- **🚀 Complete Development Workflow**: From project init to deployment
- **🏛️ Enhanced Pantheon Integration**: Smart backup management and seamless database/file syncing
- **🌐 Nginx Proxy Configuration**: Automatic proxy setup to Pantheon environment so you dont need to download assets.
- **🧪 Cypress Testing Support**: E2E testing with user management
- **🎨 Theme Development Tools**: Node.js/NPM integration with build tools
- **📦 Drupal Recipe Support**: Apply Drupal 11 recipes with cache management
- **🔄 Migration Utilities**: Tools for site migrations and database management
- **⚡ Performance Tools**: Critical CSS generation and asset optimization
- **🔍 Search Integration**: Redis and Solr add-ons with DDEV configuration

## Installation

### For Existing DDEV Projects

If you already have DDEV set up in your project:

```bash
# Install the add-on (includes interactive configuration)
ddev add-on get kanopi/ddev-kanopi-pantheon-drupal

# Restart DDEV to apply changes
ddev init
```

### For Projects Without DDEV

If your project doesn't have DDEV yet, follow these steps:

#### Step 1: Install DDEV (if not already installed)

**macOS:**
```bash
# Using Homebrew (recommended)
# Install DDEV
brew install ddev/ddev/ddev
mkcert -install

# Or using the installer script
curl -fsSL https://ddev.com/install.sh | bash
```

**Linux/Windows:** Follow instructions at [ddev.readthedocs.io](https://ddev.readthedocs.io/en/stable/users/install/ddev-installation/)

#### Step 2: Initialize DDEV in Your Project and run addon

Navigate to your existing project directory and configure DDEV:
```bash
cd /path/to/your-drupal-project

# Initialize DDEV configuration. Match your existing PHP/DB versions set in pantheon.yml
ddev config --project-type=drupal --docroot=web --php-version=8.3 --database=mariadb:10.6

# Run add-on
ddev add-on get kanopi/ddev-kanopi-pantheon-drupal
```

During installation, you'll be prompted to configure:

1. **Theme Settings**:
    - **THEME**: Path to your active Drupal theme (e.g., `themes/custom/mytheme`)
    - **THEMENAME**: Your theme name (e.g., `mytheme`)

2. **Pantheon Settings**:
    - **PANTHEON_SITE**: Your Pantheon project machine name (required)
    - **PANTHEON_ENV**: Default environment for database pulls (defaults to `dev`)

3. **Optional Migration Settings**:
    - **MIGRATE_DB_SOURCE**: Migration source Pantheon project (optional)
    - **MIGRATE_DB_ENV**: Migration source environment (optional)

The configuration is applied automatically during installation. You can modify these settings later using:
```bash
ddev config --web-environment-add THEME=path/to/your/theme
ddev config --web-environment-add THEMENAME=your-theme-name
# etc.
```

If you are running Solr, copy and paste the connection details and tweak as necessary.

#### Step 3: Spin up project

```bash
# Initialize
ddev init
```

## Interactive Installation

During the add-on installation process, you'll be prompted to configure your project settings:

```bash
ddev add-on get kanopi/ddev-kanopi-pantheon-drupal
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

## Search Integration

### Solr Configuration

The add-on installs Solr for search functionality. To connect your Drupal site to the DDEV Solr container, add this configuration to `web/sites/default/settings.php`:

```php
/**
 * DDEV Solr Configuration
 * Override Pantheon search configuration when in DDEV environment
 */
if (getenv('IS_DDEV_PROJECT') == 'true') {
  // Override any Pantheon search configuration for DDEV
  $config['search_api.server.pantheon_solr8']['backend_config']['connector_config']['host'] = 'solr';
  $config['search_api.server.pantheon_solr8']['backend_config']['connector_config']['port'] = '8983';
  $config['search_api.server.pantheon_solr8']['backend_config']['connector_config']['path'] = '/';
  $config['search_api.server.pantheon_solr8']['backend_config']['connector_config']['core'] = 'dev';
  
  // Alternative configuration if using different server name
  $config['search_api.server.solr']['backend_config']['connector_config']['host'] = 'solr';
  $config['search_api.server.solr']['backend_config']['connector_config']['port'] = '8983';
  $config['search_api.server.solr']['backend_config']['connector_config']['path'] = '/';
  $config['search_api.server.solr']['backend_config']['connector_config']['core'] = 'dev';
}
```

**Note**: Adjust the server machine name (`pantheon_solr8` or `solr`) to match your project's Search API server configuration.

### Redis Integration

Redis is automatically installed and configured for object caching. The configuration is applied automatically during add-on installation.

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
# Remove the add-on completely (includes Redis, Solr, and all 17 commands)
ddev add-on remove kanopi-pantheon-drupal

# Restart DDEV to apply changes
ddev restart
```

The removal process automatically:
- ✅ Uninstalls Redis add-on (`ddev-redis`)
- ✅ Uninstalls Solr add-on (`ddev-drupal-solr`) 
- ✅ Removes all 17 custom commands
- ✅ Removes nginx proxy configuration
- ✅ Cleans up command directories
- ✅ Preserves your environment variables (remove manually if needed)

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

## Testing

To test the add-on installation process:

```bash
# Run the automated test script
./test-install.sh
```

The test script will:
- Create a temporary DDEV project with real Drupal from git.drupalcode.org
- Test the non-interactive installation process with predefined values
- Validate that environment variables are set correctly in `config.yaml`
- Verify PHP and database versions from `pantheon.yml` are applied
- Check that Redis and Solr add-ons are installed
- Verify custom commands are available
- Test add-on removal
- Preserve test environment for inspection (use cleanup commands shown at end)

## Contributing

This add-on is maintained by [Kanopi Studios](https://kanopi.com). For issues, feature requests, or contributions, please visit our [GitHub repository](https://github.com/kanopi/ddev-kanopi-pantheon-drupal).

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/kanopi/ddev-kanopi-drupal/blob/main/LICENSE) file for details.
