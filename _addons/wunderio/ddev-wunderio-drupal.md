---
title: wunderio/ddev-wunderio-drupal
github_url: https://github.com/wunderio/ddev-wunderio-drupal
description: "Wunderio’s DDEV add-on template for Drupal projects, providing custom commands, automation hooks, and optimized local environment configuration."
user: wunderio
repo: ddev-wunderio-drupal
repo_id: 1094069257
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2025-11-11
updated_at: 2025-11-18
workflow_status: unknown
stars: 0
---

# Wunder template for DDEV Drupal projects

This project extends the standard [DDEV](https://ddev.com/) setup with additional functionality and tools specifically
designed for Drupal development. It provides a set of custom commands, configurations, and automation
scripts to enhance your Drupal development workflow.

## Installation

### Requirements

- [DDEV >= 1.24.3](https://ddev.com/)

### Steps

1. Initialize your Drupal 10 project. Project name parameter is optional, but
it's advisable to use domain name as your project name as that's used for
the subdomain of ddev.site eg if project name is example.com, then localhost
URL will become example.com.ddev.site.

   ```bash
   ddev config --project-type=drupal10 --docroot=web --project-name=example.com
   ```

2. Install Wunderio DDEV Drupal as a DDEV add-on and restart DDEV:

   ```bash
   ddev add-on get wunderio/ddev-wunderio-drupal && ddev restart
   ```

3. Optionally if you have GrumPHP installed, update grumphp.yml:

   ```
   grumphp:
     git_hook_variables:
       EXEC_GRUMPHP_COMMAND: 'ddev php'
   ```

   and then re-init the hook:

   ```bash
   ddev grumphp git:init
   ```

4. Add changes to GIT (note that below command uses -p, so you need to say 'y'es or 'n'o if it asks what to commit):

   ```bash
   git add .ddev/ &&
   git add drush/sites/ &&
   git add -p web/sites/default/settings.php grumphp.yml &&
   git commit
   ```

   Also note that whenever you update wunderio/ddev-drupal add-on, you need to add everything under .ddev to GIT.

### Updating the add-on

- To update the add-on to the latest version:

  ```bash
  ddev add-on get wunderio/ddev-drupal --update
  ```

- Optional interactive update prompt can be enabled by setting `WUNDERIO_UPDATE_PROMPT=1` in your environment; it runs on `ddev start`.

## Features

### Custom DDEV Commands

- `pmu`: Runs drush pmu commands and creates dummy module folders if they don't exist.
  This helps to uninstall module that has gone missing for example during branch
  switching.
  ```bash
  ddev pmu module1 module2 ...
  ```
- `twig-debug`: Toggles Drupal Twig debugging on/off. Useful for template development.
  ```bash
  ddev twig-debug        # Enable Twig debugging
  ddev twig-debug off    # Disable Twig debugging
  ```
- `grumphp`: Runs GrumPHP commands
  ```bash
  ddev grumphp
  ```
- `phpunit`: Runs PHPUnit commands
  ```bash
  ddev phpunit
  ```
- `regenerate-phpunit-config`: Regenerates fresh PHPUnit configuration. Run this if you don't have phpunit configured in your project.
  ```bash
  ddev regenerate-phpunit-config
  ```
- `codecept`: Runs codeception commands.
  ```bash
  ddev codecept
  ```
- `phpcbf`: Runs phpcbf commands
  ```bash
  ddev phpcbf
  ```
- `phpcs`: Runs PHPcs commands
  ```bash
  ddev phpcs
  ```
- `phpstan`: Runs PHPStan commands. Usually, the directory to be scanned is web/modules/custom or a module in the said directory.
  ```bash
  ddev phpstan analyze <directory-or-module-to-be-scanned>
  ```
- `syncdb`: Synchronizes local database from desired environment.
  You should have aliases set in drush/sites/self.site.yml
  ```bash
  ddev syncdb      # Will give error as environment is not set
  ddev syncdb prod # Will fetch database from production.
  ```
- `yq`: Runs [yq](https://mikefarah.gitbook.io/yq) commands (YAML processor)
  It's available inside DDEV, but we expose it to host because why not :). It's required in syncdb script, but it could prove useful in day to day work.
  ```bash
  ddev yq
  ```

### Enhanced Configuration

1. **Custom DDEV Configuration**
   - Post-start scripts for both host and web containers - by default it gives you uli link.
   - Automatic update checks for this package

2. **Performance Optimizations**
   - Special `database_dumps/` directory for Mac users not to mount db dumps

### Automated Workflows

The project includes several automated workflows:

1. **Database Management**
   - Post-import database hooks (clears cache, sanitizes database)
   - Post-restore snapshot hooks (clears cache, sanitizes database)
   - Database synchronization from production

2. **Development Environment Setup**
   - Automatic composer installation on first start
   - Post-start hook that run drush uli
   - Integration with Wunderio's development tools eg grumphp, phpunit

Both custom commands and hooks are scripts under `~/.ddev/wunderio/core/` folder
(note it's your host home folder) and you can extend them if you copy particular
script to your project `.ddev/wunderio/custom/` folder. This folder is never
overwritten during autoupdate.

### Migration from Composer plugin (legacy)

Previously, this package was installed as a Composer plugin and deployed files into the project. To migrate:

1. Remove legacy managed files from previous versions if present and commit your cleanup.
2. Install the add-on via `ddev add-on get wunderio/ddev-drupal`.
3. Verify `.ddev/wunderio/custom/` overrides are preserved.
4. Commit updated `.ddev/` files.

5. Import database:

   ```bash
   ddev import-db --file=some-sql-or-sql.gz.file.sql.gz
   ```

   or install site:

   ```bash
   ddev drush si
   ```

6. Create admin link and login:

   ```bash
   ddev drush uli
   ```

## Performance Optimization

### Database Operations for Mac Users

**Important for Mac users:** When working with database imports and exports on macOS, you should store your database
dump files in the `database_dumps` directory at the project root. This directory is specially configured in this
template to provide specific performance benefits.

```
project-root/
├── database_dumps/    <- Place your .sql or .sql.gz files here
├── web/
├── .ddev/
└── ...
```

**Key benefits:**

1. **Faster DDEV startup times:** When large database files are stored in the standard project directories,
they can significantly slow down DDEV startup as Mutagen indexes and syncs these files. Using the `database_dumps`
directory avoids this overhead.

2. **Reduced virtual disk usage:** By excluding database dumps from Mutagen synchronization, your virtual disk
requires less space, preventing potential disk space issues.

This optimization is configured via `upload_dirs` in `.ddev/config.wunderio.yaml`:

```yaml
upload_dirs:
  - ../database_dumps
```

**Example usage:**
```bash
# Save your database dumps to the database_dumps directory
cp ~/Downloads/my-database-backup.sql.gz ./database_dumps/

# Then import using the path relative to your project
ddev import-db --file=database_dumps/my-database-backup.sql.gz
```

This improvement is particularly noticeable in projects with multiple or large database dumps, where
startup times can be reduced from minutes to seconds.

**Note for Linux users:** While this configuration doesn't provide performance improvements on Linux
systems (which don't use Mutagen), it's still good practice to store database dumps in the
dedicated `database_dumps` folder for consistent organization across team environments.
