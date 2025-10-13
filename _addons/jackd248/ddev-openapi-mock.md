---
title: jackd248/ddev-openapi-mock
github_url: https://github.com/jackd248/ddev-openapi-mock
description: "🔌 Provides a fake OpenAPI interface for local developing & testing. "
user: jackd248
repo: ddev-openapi-mock
repo_id: 1031875939
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2025-08-04
updated_at: 2025-08-04
workflow_status: disabled
stars: 1
---

<div align="center">

# DDEV OpenAPI Mock

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/jackd248/ddev-openapi-mock/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/jackd248/ddev-openapi-mock/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/jackd248/ddev-openapi-mock)](https://github.com/jackd248/ddev-openapi-mock/commits)
[![release](https://img.shields.io/github/v/release/jackd248/ddev-openapi-mock)](https://github.com/jackd248/ddev-openapi-mock/releases/latest)
</div>

## ✨ Overview

This add-on integrates [OpenAPI Mock](https://github.com/muonsoft/openapi-mock) into your [DDEV](https://ddev.com/) project to easily test and interact with a fake OpenAPI interface.

For an integrated OpenAPI documentation, the [Swagger UI](https://swagger.io/tools/swagger-ui/) is also available.

## 🔥 Installation

```bash
ddev add-on get jackd248/ddev-openapi-mock
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## 📂 Configuration

By default, the OpenAPI [petstore specification](https://github.com/readmeio/oas-examples/blob/main/3.0/json/petstore.json) is used for demo purpose.

> [!TIP]
> You can change the location of the OpenAPI specification by setting an DDEV environment variable. See the [Advanced Customization](#advanced-customization) section below for more details.

## 📊 Usage

Regarding your used OpenAPI specification, you can use the following commands to interact with the OpenAPI Mock service, e.g. list all pets:

```bash
# List all pets
ddev launch :8080/v1/pets

# Open Swagger UI
ddev launch :8081
```

## 🗂️ Advanced Customization

### OpenAPI Mock Docker Image

To change the Docker image:

```bash
ddev dotenv set .ddev/.env.openapi-mock --openapi-mock-specification="file:/data/openapi.json"
ddev restart
```

Make sure to commit the `.ddev/.env.openapi-mock` file to version control.

All customization options (use with caution):

| Variable | Flag | Default                                                                                            |
| -------- | ---- |----------------------------------------------------------------------------------------------------|
| `OPENAPI_MOCK_DOCKER_IMAGE` | `--openapi-mock-docker-image` | `muonsoft/openapi-mock:latest`                                                                     |
| `OPENAPI_MOCK_PORT` | `--openapi-mock-port` | `8080`                                                                                             |
| `OPENAPI_MOCK_SPECIFICATION` | `--openapi-mock-specification` | `https://raw.githubusercontent.com/muonsoft/openapi-mock/refs/heads/master/examples/petstore.yaml` |
| `OPENAPI_MOCK_USE_EXAMPLES` | `--openapi-mock-use-examples` | `if_present`                                                                                       |
| `OPENAPI_MOCK_CORS_ENABLED` | `--openapi-mock-cors-enabled` | `true`                                                                                             |


> [!NOTE]
> If you're using a local OpenAPI specification file, make sure to adjust the `OPENAPI_MOCK_SPECIFICATION` variable accordingly, e.g. `file:/data/openapi.json`.
> Don't forget to update the `OPENAPI_SWAGGER_SPECIFICATION` variable to a similar value, otherwise the Swagger UI will not work properly.

### Swagger UI Docker Image

To change the Swagger UI Docker image:

```bash
ddev dotenv set .ddev/.env.swagger-ui --openapi-swagger-specification="/data/openapi.json"
ddev restart
```

Make sure to commit the `.ddev/.env.swagger-ui` file to version control.

All customization options (use with caution):

| Variable | Flag                                 | Default                                                                                            |
| -------- |--------------------------------------|----------------------------------------------------------------------------------------------------|
| `OPENAPI_SWAGGER_SPECIFICATION` | `--openapi-swagger-specification` | `https://raw.githubusercontent.com/muonsoft/openapi-mock/refs/heads/master/examples/petstore.yaml` |


> [!NOTE]
> If you're using a local OpenAPI specification file, make sure to adjust the `OPENAPI_SWAGGER_SPECIFICATION` variable accordingly, e.g. `/data/openapi.json`.

## 💎 Credits

**Contributed and maintained by [@jackd248](https://github.com/jackd248)**
