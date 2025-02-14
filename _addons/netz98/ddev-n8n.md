---
title: netz98/ddev-n8n
github_url: https://github.com/netz98/ddev-n8n
description: "ddev addon to run n8n low-code platform (middleware)"
user: netz98
repo: ddev-n8n
repo_id: 641881943
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2023-05-17
updated_at: 2025-02-06
stars: 5
---

# DDEV n8n Addon

This addon sets up a n8n instance for your DDEV project. n8n is a free and open fair-code licensed node based Workflow Automation Tool.

## Installation

1. Run `ddev get netz98/ddev-n8n` to install the addon in your exiting ddev project.
2. `ddev restart` to restart your project.

## Usage

After installation, you can access the n8n instance by visiting `https://<yourname>.ddev.site:5678` and using the username and password you set.

Run `ddev describe` to list your project's services and their URLs.

## Configuration

### Docker

To modify the build of the used n8n image for the container there are dotenv variables available.

- `N8N_TAG` - The used n8n image tag. Default is `latest`.
- `N8N_TIMEZONE` - The timezone of the n8n container. Default is `Europe/Berlin`.
- `N8N_BASIC_AUTH_ACTIVE` - Enable basic authentication. Default is `false`.
- `N8N_BASIC_AUTH_USER` - The basic authentication username. Default is `n8n`.
- `N8N_BASIC_AUTH_PASSWORD` - The basic authentication password. Default is `n8n`.
- `N8N_ENCRYPTION_KEY` - The encryption key for sensitive data. Default is `n8n`.

Use the ddev dotenv command to set these variables.

```bash
ddev dotenv set .ddev/.env.n8n --n8n-tag=1.50.0
```

### n8n Settings

Settings are in general defined in the file `.ddev/docker-compose.n8n.yaml` via environment variables.

A list of all available environment variables can be found in the reference here:
https://docs.n8n.io/hosting/environment-variables/environment-variables/

The timezone is set to "Europe/Berlin". Why? Because we are from Germany.
Change the environment variable `GENERIC_TIMEZONE` to your timezone.

### Database

#### Standard SQLite

The n8n service uses the standard SQLite database.
The SQLite database file will be created as file `.ddev/n8n/data/.n8n/database.sqlite` during the n8n container startup.

## Logging

n8n logs are directed to the container's stdout. You can view the logs with `ddev logs -s n8n`.

## Security

The n8n service access can be protected with basic authentication.
It's disabled by default.

You can enable it by changing the variable `N8N_BASIC_AUTH_ACTIVE` in `.ddev/docker-compose.n8n.yaml` to `true`. 
The username and password can be configured in the same file.
