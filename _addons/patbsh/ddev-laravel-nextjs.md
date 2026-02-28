---
title: patbsh/ddev-laravel-nextjs
github_url: https://github.com/patbsh/ddev-laravel-nextjs
description: "A ddev addon that sets up laravel api + nextjs frontend in an empty project."
user: patbsh
repo: ddev-laravel-nextjs
repo_id: 1020470173
default_branch: main
tag_name: 1.0.0
ddev_version_constraint: ">= v1.24.0"
dependencies: []
type: contrib
created_at: 2025-07-15
updated_at: 2025-07-16
workflow_status: success
stars: 1
---

# DDEV Laravel + Next.js Addon

A DDEV addon that sets up a Laravel backend with Next.js frontend in separate containers.

## What it does

- Creates a fresh Laravel project in `/backend` directory
- Creates a fresh Next.js project in `/frontend` directory  
- Configures environment files for both
- Sets up proper routing:
  - Frontend: `https://projectname.ddev.site`
  - Backend API: `https://api.projectname.ddev.site`

## Installation

In a new directory, initialize ddev

```bash
ddev config --project-type=php
```

then get the addon from the addon registry

```bash
ddev get ddev/ddev-laravel-nextjs
```

or from a local location

```bash
ddev get /path/to/this/addon
```

## Usage

After installation, bootstrap your project:

```bash
ddev start && ddev laravel-nextjs-setup
```

## Available Commands

- `ddev artisan` - Run Laravel artisan commands
- `ddev npm-frontend` - Run npm commands in frontend directory
- `ddev laravel-nextjs-setup [--force]` - Bootstrap the project (use --force to overwrite existing)

## Project Structure

```
project/
├── backend/          # Laravel application
├── frontend/         # Next.js application
└── .ddev/           # DDEV configuration
```

## Next Steps

1. Run `ddev artisan migrate` to set up your database
2. Start building your application!

## URLs

- Frontend: `https://projectname.ddev.site`
- Backend: `https://api.projectname.ddev.site`
- Database: Available to both containers as `db` host
