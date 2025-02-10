---
title: machine-rc/ddev-frontend-nextjs
github_url: https://github.com/machine-rc/ddev-frontend-nextjs
description: "NextJS Drupal Frontend application for ddev stack "
user: machine-rc
repo: ddev-frontend-nextjs
categories:
  - community
created_at: 2024-06-08
updated_at: 2024-06-08
stars: 0
---

[![tests](https://github.com/ddev/ddev-addon-template/actions/workflows/tests.yml/badge.svg)](https://github.com/ddev/ddev-addon-template/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2024.svg)

# Next.js frontend for Drupal with DDEV

## Getting started

In the DDEV project directory launch the command:
```sh
ddev get machine-rc/ddev-frontend-nextjs
```
Restart the DDEV instance:
```sh
ddev restart
```
Access NextJS frontend application on defined port (or `3003` default) via the url: https://your-project-name.ddev.site:3003/

## Setup Instructions
### Installation
1. Clone the repository:
    ```sh
    git clone git@github.com:machine-rc/ddev-frontend-nextjs.git
    cd ddev-frontend-nextjs
    ```
2. Navigate to the project directory:
    ```sh
    cd frontend
    ```
   
3. Install the dependencies:
    ```sh
    npm install
    ```
   
4. Build the project:
    ```sh
    npm run build
    ```
   

### Docker Setup
#### Dockerfile

The project includes a Dockerfile to build a Docker image for the service.

## Customizations
Current structure allows support for multiple graphql services.
To add a new service:
- create a new directory in the `frontend` directory with the service name.
- copy the `Dockerfile` from the `frontend/nextjs` directory and modify it as needed.
- adjust `docker-compose.frontend-nextjs.yaml` to include the new service by duplicating the `dotnet-users` service and changing the service name
  - adjust the `environment` section to expose the new service on a different port


