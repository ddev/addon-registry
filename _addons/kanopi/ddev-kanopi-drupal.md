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
updated_at: 2025-09-12
workflow_status: success
stars: 1
---

# DDEV Kanopi Drupal Add-on

[![tests](https://github.com/kanopi/ddev-kanopi-drupal/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/kanopi/ddev-kanopi-drupal/actions/workflows/test.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/kanopi/ddev-kanopi-drupal)](https://github.com/kanopi/ddev-kanopi-drupal/commits)
[![release](https://img.shields.io/github/v/release/kanopi/ddev-kanopi-drupal)](https://github.com/kanopi/ddev-kanopi-drupal/releases/latest)

A comprehensive DDEV add-on that provides Kanopi's battle-tested workflow for Drupal development. This add-on includes complete tooling for modern Drupal development.

## Features

- **🚀 Complete Development Workflow**: From project init to deployment
- **🏛️ Multi-Platform Hosting Integration**: 
  - **Pantheon**: Smart backup management and seamless database/file syncing
  - **Acquia**: Multi-database support with intelligent backup age detection
- **🌐 Smart Proxy Configuration**: 
  - **Nginx**: Automatic proxy setup for Pantheon environments
  - **Apache-FPM**: Native Apache configuration for Acquia environments
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
ddev add-on get kanopi/ddev-kanopi-drupal

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

#### Step 2: Initialize DDEV in Your Project

**Important**: If you're migrating from Docksal or Lando, stop those services first to avoid port conflicts:
```bash
# Stop Docksal
fin system stop

# Or stop Lando
lando poweroff
```

Navigate to your existing project directory and configure DDEV:
```bash
cd /path/to/your-drupal-project

# Initialize DDEV configuration. Match your existing PHP/DB versions set in pantheon.yml
ddev config --project-type=drupal --docroot=web --php-version=8.3 --database=mariadb:10.6

# Run add-on
ddev add-on get kanopi/ddev-kanopi-drupal
```

During installation, you'll be prompted to configure:

1. **Theme Settings**:
    - **THEME**: Path to your active Drupal theme (e.g., `themes/custom/mytheme`)
    - **THEMENAME**: Your theme name (e.g., `mytheme`)

2. **Hosting Provider Settings**:
    - **HOSTING_PROVIDER**: Choose between `pantheon` or `acquia`
    
    **For Pantheon**:
    - **PANTHEON_SITE**: Your Pantheon project machine name (required)
    - **PANTHEON_ENV**: Default environment for database pulls (defaults to `dev`)
    
    **For Acquia**:
    - **HOSTING_SITE**: Your Acquia application UUID (required)
    - **HOSTING_ENV**: Default environment for database pulls (defaults to `prod`)
    - **HOSTING_DOMAIN**: Your Acquia domain for file proxy (optional)

3. **Optional Migration Settings**:
    - **MIGRATE_DB_SOURCE**: Migration source project (optional)
    - **MIGRATE_DB_ENV**: Migration source environment (optional)

The configuration is applied automatically during installation. You can modify these settings later using:
```bash
ddev config --web-environment-add THEME=path/to/your/theme
ddev config --web-environment-add THEMENAME=your-theme-name
ddev config --web-environment-add HOSTING_PROVIDER=pantheon
ddev config --web-environment-add PANTHEON_SITE=your-site-name
# For Acquia:
ddev config --web-environment-add HOSTING_PROVIDER=acquia
ddev config --web-environment-add HOSTING_SITE=your-app-uuid
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
ddev add-on get kanopi/ddev-kanopi-drupal
```

## Available Commands

This add-on provides 17 custom commands organized by where they execute (host system vs. web container):

### Host Commands (Execute on your machine)
| Command | Description | Example |
|---------|-------------|---------|
| `ddev cypress <command>` | Run Cypress commands with environment support | `ddev cypress open` |
| `ddev cypress-users` | Create default admin user for Cypress testing | `ddev cypress-users` |
| `ddev init` | Complete project initialization with dependencies, Lefthook, NVM, Cypress, and database refresh | `ddev init` |
| `ddev install-cypress` | Install Cypress E2E testing dependencies | `ddev install-cypress` |
| `ddev open` | Open project URL in browser | `ddev open` |
| `ddev phpmyadmin` | Launch PhpMyAdmin database interface | `ddev phpmyadmin` |
| `ddev rebuild` | Run composer install followed by database refresh | `ddev rebuild` |
| `ddev refresh [env] [-f]` | Smart database refresh from Pantheon with 12-hour backup age detection | `ddev refresh live -f` |
| `ddev testenv <name> [profile]` | Create isolated testing environment with optional install profile | `ddev testenv my-test minimal` |

### Web Commands (Execute inside DDEV container)
| Command | Description | Example |
|---------|-------------|---------|
| `ddev install-critical-tools` | Install Critical CSS generation tools | `ddev install-critical-tools` |
| `ddev install-theme-tools` | Set up Node.js, NPM, and build tools using .nvmrc | `ddev install-theme-tools` |
| `ddev migrate-prep-db` | Create secondary database for migrations | `ddev migrate-prep-db` |
| `ddev npm <command>` | Run NPM commands in theme directory specified by THEME env var | `ddev npm run build` |
| `ddev npx <command>` | Run NPX commands in theme directory | `ddev npx webpack --watch` |
| `ddev recipe-apply <path>` | Apply Drupal recipe with automatic cache clearing | `ddev recipe-apply ../recipes/my-recipe` |
| `ddev tickle` | Keep Pantheon environment awake during long operations | `ddev tickle` |
| `ddev uuid-rm <path>` | Remove UUIDs from config files for recipe development | `ddev uuid-rm config/sync` |

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

Redis is automatically installed and configured for object caching. The configuration is applied automatically during add-on installation through the `settings.ddev.redis.php` file.

If you need custom Redis configuration, add it within your DDEV settings block:

```php
/**
 * DDEV Configuration
 * Override Pantheon configurations when in DDEV environment
 */
if (getenv('IS_DDEV_PROJECT') == 'true') {
  // Redis configuration (if needed for custom settings)
  $settings['redis.connection']['interface'] = 'PhpRedis';
  $settings['redis.connection']['host'] = 'redis';
  $settings['redis.connection']['port'] = 6379;

  // Cache backend configuration
  $settings['cache']['default'] = 'cache.backend.redis';
  $settings['cache_prefix']['default'] = 'ddev_' . hash('sha256', $app_root);
}
```

### Additional DDEV Settings Recommendations

Here are additional settings you may want to include in your DDEV configuration block:

```php
/**
 * DDEV Configuration
 * Override Pantheon configurations when in DDEV environment
 */
if (getenv('IS_DDEV_PROJECT') == 'true') {
  // Database configuration
  $databases['default']['default'] = [
    'database' => 'db',
    'username' => 'db',
    'password' => 'db',
    'host' => 'db',
    'driver' => 'mysql',
    'port' => 3306,
    'prefix' => '',
  ];

  // Trusted host patterns
  $settings['trusted_host_patterns'] = [
    '^.+\\\\.ddev\\\\.site$',
  ];

  // Development-friendly settings
  $config['system.performance']['css']['preprocess'] = FALSE;
  $config['system.performance']['js']['preprocess'] = FALSE;
  $config['system.logging']['error_level'] = 'verbose';

  // Disable render cache and page cache for development
  $settings['cache']['bins']['render'] = 'cache.backend.null';
  $settings['cache']['bins']['page'] = 'cache.backend.null';
  $settings['cache']['bins']['dynamic_page_cache'] = 'cache.backend.null';

  // Mail configuration (use Mailpit)
  $config['system.mail']['interface']['default'] = 'test_mail_collector';

  // File system paths
  $settings['file_public_path'] = 'sites/default/files';
  $settings['file_private_path'] = '../private';
  $settings['file_temp_path'] = '/tmp';

  // Include local development settings if they exist
  if (file_exists($app_root . '/' . $site_path . '/settings.local.php')) {
    include $app_root . '/' . $site_path . '/settings.local.php';
  }
}
```

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
ddev add-on get kanopi/ddev-kanopi-drupal
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

## Post-Installation Setup

### Required Configuration Steps

After installing the add-on, complete these essential setup tasks to ensure proper functionality:

#### 1. Configure Hosting Provider Authentication

**For Pantheon Projects:**
Set your Pantheon Terminus machine token globally (required for database refreshes and Pantheon integration):

```bash
# Set the token globally for all DDEV projects
ddev config global --web-environment-add=TERMINUS_MACHINE_TOKEN=your_token_here
```

**Note**: Replace `your_token_here` with your actual Pantheon machine token. You can generate one at [https://dashboard.pantheon.io/machine-token/create](https://dashboard.pantheon.io/machine-token/create).

**For Acquia Projects:**
Set your Acquia API credentials globally (required for database refreshes and Acquia integration):

```bash
# Set Acquia API credentials globally for all DDEV projects
ddev config global --web-environment-add=ACQUIA_API_KEY=your_api_key_here
ddev config global --web-environment-add=ACQUIA_API_SECRET=your_api_secret_here
```

**Note**: Replace the placeholder values with your actual Acquia API credentials. You can generate these at [https://cloud.acquia.com/a/profile/tokens](https://cloud.acquia.com/a/profile/tokens).

#### 2. Stop Conflicting Development Tools

**If using Docksal**: Stop Docksal services before using DDEV to avoid port conflicts:
```bash
fin system stop
```

**If using Lando**: Stop Lando services:
```bash
lando poweroff
```

#### 3. Configure Settings.php for DDEV

Add DDEV-specific configuration to your `web/sites/default/settings.php` file. Place this at the end of your settings file:

```php
/**
 * DDEV Configuration
 * Override Pantheon configurations when in DDEV environment
 */
if (getenv('IS_DDEV_PROJECT') == 'true') {
  // Override any existing database configuration for DDEV
  $databases['default']['default'] = [
    'database' => 'db',
    'username' => 'db',
    'password' => 'db',
    'host' => 'db',
    'driver' => 'mysql',
    'port' => 3306,
    'prefix' => '',
  ];

  // Set local development settings
  $settings['trusted_host_patterns'] = [
    '^.+\\.ddev\\.site$',
  ];

  // Disable CSS/JS aggregation for development
  $config['system.performance']['css']['preprocess'] = FALSE;
  $config['system.performance']['js']['preprocess'] = FALSE;

  // Enable local development modules if they exist
  if (file_exists($app_root . '/' . $site_path . '/settings.local.php')) {
    include $app_root . '/' . $site_path . '/settings.local.php';
  }
}
```

**Migrate Docksal/Lando Settings**: If you have existing settings.php customizations for Docksal or Lando, review and adapt them within the `if (getenv('IS_DDEV_PROJECT') == 'true')` block.

#### 4. Review and Update Theme Tools Command

Compare the provided `install-theme-tools` command with your project's current build process:

```bash
# Review the command
ddev help install-theme-tools

# Test the command in your theme directory
ddev install-theme-tools
```

The command expects:
- A `.nvmrc` file in your theme directory specifying the Node.js version
- A `package.json` with build scripts (like `npm run build` and `npm run watch`)
- Standard NPM package structure

If your theme build process differs, you may need to:
- Update your theme's `package.json` scripts
- Modify the `commands/web/install-theme-tools` file to match your workflow
- Create a `.nvmrc` file with your preferred Node.js version

#### 5. Convert Existing Custom Commands

If you have custom commands from Docksal (`.docksal/commands/`) or Lando (in your `.lando.yml` tooling section):

1. **Host commands** (run on your machine): Place in `.ddev/commands/host/`
2. **Container commands** (run inside DDEV): Place in `.ddev/commands/web/`

Use the existing commands as templates:
```bash
# View existing command structure
ls -la .ddev/commands/host/
ls -la .ddev/commands/web/

# Copy a similar command as a template
cp .ddev/commands/host/rebuild .ddev/commands/host/your-new-command
```

#### 6. Update Project Documentation

Update your project's README to document the DDEV setup. Here's a template section:

```markdown
## Local Development with DDEV

This project supports both DDEV and Docksal for local development.

### DDEV Setup

1. **Install DDEV**: Follow [DDEV Installation Guide](https://docs.ddev.com/en/stable/users/install/ddev-installation/)

2. **Configure Terminus Token**:
   ```bash
   ddev config global --web-environment-add=TERMINUS_MACHINE_TOKEN=your_token_here
   ```

3. **Stop Docksal** (if running):
   ```bash
   fin system stop
   ```

4. **Initialize Project**:
   ```bash
   ddev start
   ddev init
   ```

### Available DDEV Commands

[Include the commands table from the add-on README]
```

#### 7. Initial Project Setup

After configuration, initialize your project:

```bash
# Start DDEV and run initialization
ddev start
ddev init

# This will:
# - Install Lefthook git hooks
# - Set up NVM for Node.js
# - Install Cypress for testing
# - Refresh database from Pantheon
# - Create Cypress test users
```

### Verification Steps

1. **Test database refresh**: `ddev refresh`
2. **Test theme tools**: `ddev install-theme-tools`
3. **Verify Pantheon connection**: `ddev exec terminus site:list`
4. **Test proxy setup**: Visit your local site and check if assets load from Pantheon

## Quick Reference

### Common Workflow
```bash
# Daily development workflow
ddev init

# or individually
ddev start                    # Start DDEV
ddev refresh                  # Get latest database
ddev install-theme-tools      # Set up theme tools (first time)
ddev npm run watch           # Start theme development

# Testing workflow
ddev install-cypress         # Set up Cypress (first time)
ddev cypress-users          # Create test users
ddev cypress open           # Open Cypress

# Deployment preparation
ddev npm run build          # Build theme assets
ddev drush cache:rebuild    # Clear Drupal caches
```

### Key Environment Variables
- `THEME`: Path to your theme directory (e.g., `themes/custom/mytheme`)
- `THEMENAME`: Theme name for development tools
- `HOSTING_PROVIDER`: Hosting platform (`pantheon` or `acquia`)

**For Pantheon:**
- `PANTHEON_SITE`: Pantheon project machine name
- `PANTHEON_ENV`: Default environment for database pulls
- `TERMINUS_MACHINE_TOKEN`: Pantheon API access token (set globally)

**For Acquia:**
- `HOSTING_SITE`: Acquia application UUID
- `HOSTING_ENV`: Default environment for database pulls (typically `prod`)
- `HOSTING_DOMAIN`: Acquia domain for file proxy (optional)
- `ACQUIA_API_KEY`: Acquia API key (set globally)
- `ACQUIA_API_SECRET`: Acquia API secret (set globally)

## Troubleshooting

### Pantheon Authentication Issues
```bash
# Check if token is set
ddev exec printenv TERMINUS_MACHINE_TOKEN

# Re-authenticate manually
ddev exec terminus auth:login --machine-token="your_token"
```

### Acquia Authentication Issues
```bash
# Check if API credentials are set
ddev exec printenv ACQUIA_API_KEY
ddev exec printenv ACQUIA_API_SECRET

# Test Acquia CLI authentication
ddev exec acli auth:login --key="your_key" --secret="your_secret"

# List Acquia applications to verify access
ddev exec acli api:applications:find
```

### Theme Build Issues
```bash
# Check Node.js version
ddev exec node --version

# Reinstall dependencies
ddev install-theme-tools
```

### Database Refresh Issues

**For Pantheon:**
```bash
# Check Pantheon connection
ddev exec terminus site:list

# Force new backup
ddev refresh -f
```

**For Acquia:**
```bash
# Check Acquia connection and applications
ddev exec acli api:applications:find

# List environments for your application
ddev exec acli api:environments:find HOSTING_SITE

# Force new backup
ddev refresh -f
```

## Platform-Specific Configurations

### Webserver Configuration

The add-on automatically configures the appropriate webserver based on your hosting provider:

**Pantheon Projects** (Nginx):
- Uses Nginx configuration optimized for Pantheon environments
- Includes automatic file proxy to Pantheon environment for assets
- Optimized for Drupal with clean URLs and caching headers
- Configuration file: `.ddev/nginx_full/nginx-site.conf`

**Acquia Projects** (Apache-FPM):
- Uses Apache with PHP-FPM configuration matching Acquia Cloud
- No additional Nginx configuration applied
- Native Apache performance with mod_rewrite for clean URLs
- Compatible with Acquia Cloud's Apache-based infrastructure

The webserver configuration is automatically selected during installation based on your `HOSTING_PROVIDER` setting. No manual configuration is required.

### File Proxy Configuration

**Pantheon**: Automatic proxy to `PANTHEON_ENV-PANTHEON_SITE.pantheonsite.io` for missing assets.

**Acquia**: Configurable proxy to `HOSTING_ENV-HOSTING_SITE.HOSTING_DOMAIN` (requires `HOSTING_DOMAIN` environment variable).

## Testing

This add-on uses the official DDEV testing framework combined with comprehensive integration testing:

### 1. Automated CI/CD Testing

The add-on is automatically tested using the official DDEV add-on testing framework:

- **GitHub Actions**: Tests run on every push and pull request
- **Matrix Testing**: Validates against both stable and HEAD DDEV versions  
- **Official Framework**: Uses `ddev/github-action-add-on-test@v2`
- **Reliable & Fast**: Proven testing patterns from the DDEV team

### 2. Local Development Testing

For local development and debugging, use the comprehensive integration test:

```bash
# Run comprehensive integration test
./tests/test-install.sh
```

**What this test does:**
- Creates a temporary DDEV project with real Drupal from git.drupalcode.org
- Tests the interactive installation process with automated responses
- Validates environment variables are set correctly in `config.yaml`
- Verifies PHP and database versions from `pantheon.yml` are applied
- Checks that Redis and Solr add-ons are installed and configured
- Validates all 17 custom commands are available and functional
- Tests add-on removal and cleanup
- Preserves test environment for inspection and debugging

**Test Results:**
- ✅ Pass: All functionality works correctly
- ❌ Fail: Issues found (test environment preserved for debugging)
- 🔍 Debug: Use provided commands to inspect the test environment

### 3. Component Testing

For testing specific functionality during development:

```bash
# Install bats if not already installed
# macOS: brew install bats-core
# Linux: See https://bats-core.readthedocs.io/en/stable/installation.html

# Run component tests
bats tests/test.bats
```

### Testing Strategy

| Test Level | Purpose | When to Use |
|------------|---------|-------------|
| **GitHub Actions** | Automated validation | Every push/PR |
| **Integration Test** | Comprehensive workflow | Local development, debugging |
| **Bats Tests** | Component functionality | Feature development |

### Debugging Failed Tests

If the integration test fails, inspect the preserved environment:

```bash
# Navigate to test environment
cd test/test-install/drupal

# Check DDEV status
ddev describe

# Verify environment variables
ddev exec env | grep -E "(THEME|PANTHEON)"

# Check installed add-ons
ddev add-on list --installed

# Cleanup when done
ddev delete -Oy test-kanopi-addon && cd ../.. && rm -rf test/
```

## Contributing

This add-on is maintained by [Kanopi Studios](https://kanopi.com). For issues, feature requests, or contributions, please visit our [GitHub repository](https://github.com/kanopi/ddev-kanopi-drupal).

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/kanopi/ddev-kanopi-drupal/blob/main/LICENSE) file for details.
