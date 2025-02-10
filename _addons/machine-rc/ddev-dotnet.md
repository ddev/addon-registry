---
title: machine-rc/ddev-dotnet
github_url: https://github.com/machine-rc/ddev-dotnet
description: ".NET service for ddev stack "
user: machine-rc
repo: ddev-dotnet
categories:
  - community
created_at: 2024-06-08
updated_at: 2024-06-08
stars: 0
---

[![tests](https://github.com/ddev/ddev-addon-template/actions/workflows/tests.yml/badge.svg)](https://github.com/ddev/ddev-addon-template/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2024.svg)

# .NET ddev Add-on

## Getting started

In the DDEV project directory launch the command:
```sh
ddev get machine-rc/ddev-dotnet
```
Restart the DDEV instance:
```sh
ddev restart
```
Access .NET application on defined port (or `8080` default) via the url: https://your-project-name.ddev.site:8080/

## Setup Instructions
### Installation
1. Clone the repository:
    ```sh
    git clone git@github.com:machine-rc/ddev-dotnet.git
    cd ddev-dotnet
    ```

2. Navigate to the project directory:
    ```sh
    cd apis/Users.API
    ```

3. Restore the .NET dependencies:
    ```sh
    dotnet restore
    ```

4. Build the project:
    ```sh
    dotnet build
    ```

### Running the Project Locally
1. To run the project locally, use the following command:
    ```sh
    dotnet run
    ```
2. Open your browser and navigate to `http://localhost:8080` or `http://localhost:8081`.

### Docker Setup
#### Dockerfile

The project includes a Dockerfile to build a Docker image for the service.

## Customizations
Current structure allows support for multiple graphql services.
To add a new service:
- create a new directory in the `dotnet` directory with the service name
- copy the `Dockerfile` from the `dotnet/Users.API` directory
- adjust `docker-compose.dotnet.yaml` to include the new service by duplicating the `dotnet-users` service and changing the service name
  - adjust the `environment` section to expose the new service on a different port


