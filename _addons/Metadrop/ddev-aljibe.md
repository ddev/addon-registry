---
title: Metadrop/ddev-aljibe
github_url: https://github.com/Metadrop/ddev-aljibe
description: "DDEV Aljibe (ddev-aljibe) is an add-on for DDEV to develop Drupal projects with many tools included out-of-the-box"
user: Metadrop
repo: ddev-aljibe
repo_id: 817303007
ddev_version_constraint: ">= v1.24.1"
dependencies: []
type: contrib
created_at: 2024-06-19
updated_at: 2025-05-28
workflow_status: disabled
stars: 9
---

[![tests](https://github.com/Metadrop/ddev-aljibe/actions/workflows/tests.yml/badge.svg)](https://github.com/Metadrop/ddev-aljibe/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2025.svg)
![GitHub Release](https://img.shields.io/github/v/release/Metadrop/ddev-aljibe)

# DDEV Aljibe
## About Aljibe
Aljibe is an add-on for DDEV for Drupal projects that adds several tools in a simple and fast way, leaving a new project ready for development in a few minutes.

Aljibe sits on top of DDEV and adds some containers, configuration and commands to make the development of Drupal projects faster and easier.

### Requirements

- [DDEV](https://ddev.readthedocs.io/en/stable/) v1.24.0 or higher
- [Docker](https://www.docker.com/) 24 or higher

### Included tools

- Behat: BDD and Acceptance testig
- [BackstopJS](https://github.com/Metadrop/ddev-backstopjs): Visual regression testing
- [Unlighthouse](https://github.com/Metadrop/ddev-unlighthouse): Audit all website quality
- [Pa11y](https://github.com/Metadrop/ddev-pa11y): Accesibility checks
- [MkDocs](https://github.com/Metadrop/ddev-mkdocs): Documentation wiki
- [Adminer](https://github.com/ddev/ddev-adminer): Database manager

> **Note:** This tool is based on DDEV, so any DDEV add-on can work with this: `ddev add-on list --all`

## How to use it

If you haven't already cloned or created an Aljibe project, please follow the [Setup guide](#setup-a-new-project-with-aljibe).

### Start working
Once the project has been configured, or if you have cloned an already setup Aljibe project, run this command to have the project ready to work with.

```sh
ddev setup [--all] [--no-install] [--no-themes]
```

Use:
- `--all` to install all sites on Multisite
- `--no-install` to prepare only the environment
- `--no-themes` if you don't need to transpile the CSS/JS of the themes.

After this, you are **ready to work on your project.**

#### Unique site install (Multisite)

If you have a multisite instalation, you can install only one site by running:

```sh
ddev site-install <site_name>
```
### Running tests

#### Behat tests

To launch local, or env tests, you can run:

```sh
ddev behat [local|pro|other_behat_folder] [suite]
```

#### Backstopjs tests

To launch backstopjs commands you can run:

```sh
ddev backstopjs [local|pro|other_backstop_folder] [backstop_command]
```

For example `ddev backstopjs local test` to run the tests on the local environment.
Or `ddev backstopjs pro reference` to create a reference of the production environment.


#### Run pa11y tests

To run pa11y against a site, you can run:

```sh
ddev pa11y [site_url]
``` 
For example `ddev pa11y http://web`

### Working with frontend themes
#### Process custom themes CSS

By default there is one theme defined in .ddev/aljibe.yml with the name "custom_theme". 
You can [add multiple themes](#advanced-configuration). To transpile them, run:

```sh
ddev frontend production [theme_name]
```

where theme_name is the key defined in .ddev/aljibe.yml. 

You can run a watch command to process the CSS on the fly:

```sh
ddev frontend watch [theme_name]
```

The **production** and **watch** parameters are scripts defined in the **package.json** of the default Aljibe theme.
Any scripts or commands defined there can be executed with the `ddev frontend [my_script]` command.

### Other commands
#### Create a secondary database

If you need to create a secondary database, you can run:

```sh
ddev create-database <db_name>
```

> **Note:** This command will create a database accesible with the same user and password from the main one. If you want to persist this across multiples setups, you can add this command to te pre-setup hooks in .ddev/aljibe.yml file.

#### Running drush on all sites

```sh
ddev all-sites-drush <drush_command>
```

#### Sync solr config

If you use ddev-solr addon and need to sync the solr config from the server, you can run:

```sh
ddev solr-sync
```

#### Power off ddev

```sh
ddev poweroff
```

## Setup a new project with Aljibe

Create a folder for your new project (e.g. `mkdir my-new-project`)
Configure a basic ddev project:

```sh
ddev config --auto
```

Install the Aljibe addon. This will install all the dependant addons too:

```sh
ddev add-on get metadrop/ddev-aljibe
```

Launch Aljibe Assistant. This will guide you throught the basic Drupal site instalation process:

```sh
ddev aljibe-assistant
```

You are ready! you will have a new Drupal project based on Aljibe ready for development!

## Add Aljibe to existing projects

To transform a legacy project to Ddev Aljibe, the following steps must be followed, always taking into account the particularities of each project:

1. Clone the project without install it and remove all docker related files
1. Run basic ddev-config:

    ```sh
    ddev config --auto
    ```

1. Install Aljibe:
    ```sh
    ddev add-on get metadrop/ddev-aljibe
    ```


1. Run again ddev config, but this time go throught the assistant to set project type to Drupal, docroot folder, etc...

    ```sh
    ddev config
    ```

1. Edit .ddev/config.yml to fine tune the environment.
1. Edit .ddev/aljibe.yml to set deault site name (the folder inside sites) and all themes to be transpiled
1. update .gitignore to look like [this](https://github.com/Metadrop/ddev-aljibe/blob/main/kickstart/common/.gitignore).

If you come from a [boilerplate](https://github.com/Metadrop/drupal-boilerplate) project:

- Remove from settings.local.php database, trusted host patterns and others that can conflict with settings.ddev.php.
- Adapt the drush alias to the new url.
- Review tests folder structure as in aljibe, all tests (behat.yml included) are inside tests folder and replace <http://apache> or <http://nginx> by <http://web>.
- Config also the nodejs_version in .ddev/config.yml with the same as the old project. Old version on .env file, variable **“NODE_TAG”**
- Adapt grumphp changing EXEC_GRUMPHP_COMMAND on grumphp.yml to “ddev exec”
- Launch ddev setup:
  - If monosite: `ddev setup`
  - If multisite:`ddev setup —all` or `ddev setup --sites=site1`

## Advanced Configuration

The `aljibe.yml` file allows you to customize various aspects of the Aljibe setup. Below are the available options and how to use them:

### `default_site`

This option sets the default site name to be installed on setup command. 
It is used when no specific site name is provided when running. For example,
with this configuration:

```yaml
default_site: my_site
```
`dev setup` will install the site "my_site", the same way as `dev setup my_site`.   

> **NOTE**: The site names must mustch the drush aliases. Names without dots are
> considered as ".local" aliases, but if you have a different alias, you can specify it
> and .local will not be appended.


### `theme_paths`

This section allows you to define paths to custom themes. Each theme should be 
listed with a unique key. Those themes works with the `ddev frontend` command.

```yaml
theme_paths:
  custom_theme: /var/www/html/web/themes/custom/custom_theme
```

### `hooks`

Hooks are commands that can be executed at different stages of the setup process. They are defined as lists of commands under various hook points.

- `pre_setup`: Commands to run before the setup process.
- `post_setup`: Commands to run after the setup process.
- `pre_site_install`: Commands to run before any type site installation.
- `post_site_install`: Commands to run after any type site installation.
- `pre_site_install_config`: Commands to run before the site installation from config.
- `post_site_install_config`: Commands to run after the site installation from config.
- `pre_site_install_db`: Commands to run before the site installation from database.
- `post_site_install_db`: Commands to run after the site installation from database.

Example:

```yaml
hooks:
  pre_setup:
    - echo "Aljibe pre setup hook"
    - ddev snapshot
  post_setup:
    - echo "Aljibe post setup hook"
    - ddev drush uli --uid=2
  pre_site_install: []
  post_site_install: []
    - ddev @${SITE_ALIAS} drush cr
  pre_site_install_config: []
  post_site_install_config: []
  pre_site_install_db: []
  post_site_install_db: []
```

Available variables for site_install hooks are:

- DRUPAL_PROFILE: The profile used to install the site.
- SITE_PATH: The path to the site to be installed relative to web/sites
- SITE_ALIAS: The drush alias of the site to be installed, without the @.

Available variables for setup hooks are:

- NO_THEMES: If the setup command was called with the --no-themes flag
- NO_INSTALL: If the setup command was called with the --no-install flag
- SITES: The sites to be installed
- CONFIG_DEFAULT_SITE: The default site to be installed

In addition, all the variables [provided by ddev](https://ddev.readthedocs.io/en/stable/users/extend/custom-commands/#command-line-completion) are available on all hooks.

#### `installable_sites_aliases`

You can add the names of the different sites you want install when running 
`ddev setup --all`. Additional sites can still be installed later using the `ddev site-install MYSITE` command.

```yaml
installable_sites_aliases:
  - site1
  - site2
  - site3.mylocal
```
> **NOTE**: The site names must mustch the drush aliases. Names without dots are 
> considered as ".local" aliases, but if you have a different alias, you can specify it
> and .local will not be appended.

## Get configuration
You can add any other configuration you need to the `aljibe.yml` file. This config can be obtained with the `ddev aljibe-config` command.

Example commands to obtain specific configurations:

- To obtain all theme paths to be processed:

  ```sh
  ddev aljibe-config theme_path
  ```

- To get the default site to be processed:

  ```sh
  ddev aljibe-config default_site
  ```

## Troubleshooting

### https not working

Follow ddev [install recomendations](https://ddev.readthedocs.io/en/stable/users/install/ddev-installation/#linux).
It is needed to install mkcert and libnss3-tools, and then run:

```sh
mkcert -install
```

### Can't debug with NetBeans

Until <https://github.com/apache/netbeans/issues/7562> is solved you need to create a file named `xdebug.ini` at `.ddev/php` with the following content:

```
[XDebug]
xdebug.idekey = netbeans-xdebug
```

NOTE: The `netbeans-xdebug` is the default Session ID value in the the Debugging tab in the PHP Netbeans' configuration dialog. If you have changed it do it in the `xdebug.ini` file as well.

### Xdebug profiler does not save the files

Follow the instructions from [ddev xprofiler documentation](https://ddev.readthedocs.io/en/stable/users/debugging-profiling/xdebug-profiling/#basic-usage)

```
[XDebug]
xdebug.mode=profile
xdebug.start_with_request=yes
# Set a ddev shared folder for the xprofile reports.
xdebug.output_dir=/var/www/html/tmp/xprofile
xdebug.profiler_output_name=trace.%c%p%r%u.out
```

Review the php info (/admin/reports/status/php) page to review that the xdebug variables are setup properly after run ddev xdebug on, restart the project if necessary.
