---
title: "Pronovix/ddev-oidc-provider-mock"
github_url: "https://github.com/Pronovix/ddev-oidc-provider-mock"
description: "A mock OpenID Provider server to test and develop OpenID Connect authentication."
user: "Pronovix"
repo: "ddev-oidc-provider-mock"
repo_id: 1325102421
default_branch: "main"
tag_name: "1.0.0-alpha1"
ddev_version_constraint: ">= v1.25.3"
dependencies: []
type: "contrib"
created_at: "2026-08-06"
updated_at: "2026-08-07"
workflow_status: "failure"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/Pronovix/ddev-oidc-provider-mock/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/Pronovix/ddev-oidc-provider-mock/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/Pronovix/ddev-oidc-provider-mock)](https://github.com/Pronovix/ddev-oidc-provider-mock/commits)
[![release](https://img.shields.io/github/v/release/Pronovix/ddev-oidc-provider-mock)](https://github.com/Pronovix/ddev-oidc-provider-mock/releases/latest)

# DDEV OpenID Connect (OIDC) Provider Mock Add-on (`ddev-oidc-provider-mock`)

This DDEV add-on integrates [`ghcr.io/geigerzaehler/oidc-provider-mock`](https://github.com/geigerzaehler/oidc-provider-mock) into your DDEV environment to provide a mock OpenID Connect (OIDC) identity provider out of the box for local development and testing.

## Overview & Features

- **Optional Profile**: Defined via Docker Compose profiles (`oidc-provider-mock`, `oidc`), so the container runs only when you actively need it.
- **Turnkey Setup**: Pre-configured mock users (`admin`, `alice`, `bob`) with zero mandatory setup steps.
- **Framework-Agnostic**: Works with any OIDC client (Drupal `openid_connect`, `oauth2_client`, custom PHP/JS applications, etc.).
- **Dual Routing**: Routed through DDEV's HTTPS router at `https://oidc.<site>.<tld>` (e.g. `https://oidc.mysite.ddev.site`) so both browser front-channel redirects and server-to-server back-channel API calls hit the exact same Issuer URL.
- **Accepts Any Credentials**: Accepts any Client ID and Client Secret by default (unless restricted via arguments).
- **Persistent & Versionable User List**: User definitions are stored in `.ddev/oidc-provider-mock/users.yaml` which can be customized and committed to your project's version control.

## Installation & Usage

1. **Install the add-on:**
   ```bash
   ddev add-on get Pronovix/ddev-oidc-provider-mock
   ```

2. **Start/Run the service:**

   - **Persistent configuration (recommended):**
     If you want the OIDC Provider Mock service to always start automatically alongside your project, add the `oidc-provider-mock` profile to your `.ddev/config.local.yaml` (or `.ddev/config.yaml`) file:

     ```yaml
     profiles:
       - oidc-provider-mock
     ```

     Then restart your project:

     ```bash
     ddev restart
     ```

   - **On-demand (alternative):**
     If you prefer to start the OIDC Provider Mock only when needed, without persisting it to your DDEV configuration, run `ddev restart` first to ensure the router is ready, then start with the profile:

     ```bash
     ddev restart
     ddev start --profiles=oidc-provider-mock
     ```

## Key Endpoints

> [!NOTE]
> `<site>` corresponds to your DDEV project name and `<tld>` corresponds to your configured `project_tld` (default: `ddev.site`). The hostname dynamically adapts to custom primary domains or TLDs configured in DDEV.

| Endpoint | URL | Description |
| --- | --- | --- |
| **Issuer URL** | `https://oidc.<site>.<tld>` | Base OpenID Provider URL |
| **Discovery** | `https://oidc.<site>.<tld>/.well-known/openid-configuration` | OpenID Connect discovery metadata |
| **Authorization** | `https://oidc.<site>.<tld>/oauth2/authorize` | HTML authorization form / redirect |
| **Token** | `https://oidc.<site>.<tld>/oauth2/token` | Exchange code for tokens |
| **Userinfo** | `https://oidc.<site>.<tld>/oauth2/userinfo` | Fetch authenticated user claims |
| **Client Registration** | `POST https://oidc.<site>.<tld>/oauth2/clients` | Dynamic client registration (optional) |
| **Dynamic Claims API** | `PUT https://oidc.<site>.<tld>/users/{sub}` | Inject/update user claims during test runs |

## Usage & Commands

| Command | Description |
| ------- | ----------- |
| `ddev describe` | View service status and exposed endpoints |
| `ddev logs -s oidc-provider-mock` | Check OIDC Provider Mock logs |

## Customization Guide

### Editing Mock Users
Edit `.ddev/oidc-provider-mock/users.yaml` in your project to modify existing users (`admin`, `alice`, `bob`) or add custom users and claims (such as `given_name`, `family_name`, and `roles`). Run `ddev restart` after making changes.

Example user configuration:

```yaml
- sub: alice
  email: alice@example.com
  name: Alice Smith
  given_name: Alice
  family_name: Smith
  preferred_username: alice
  roles:
    - editor
```

### Additional Server Arguments
To pass extra flags to `oidc-provider-mock` (such as `--require-registration` or `--require-nonce`), add `OIDC_PROVIDER_MOCK_EXTRA_ARGS` to `.ddev/.env.oidc-provider-mock`:

```env
OIDC_PROVIDER_MOCK_EXTRA_ARGS="--require-registration --require-nonce"
```

Then run `ddev restart`.

### Changing Container Image
To override the Docker image used by the service, set `OIDC_PROVIDER_MOCK_DOCKER_IMAGE` in `.ddev/.env.oidc-provider-mock`:

```env
OIDC_PROVIDER_MOCK_DOCKER_IMAGE="ghcr.io/geigerzaehler/oidc-provider-mock:latest"
```

## Generic OIDC Client Configuration

To connect any OIDC client module or application to this mock provider:

1. **Issuer / Base URL**: `https://oidc.<site>.<tld>` (or use the environment variable `$OIDC_ISSUER_URL` inside the `web` container)
2. **Client ID**: Any string (e.g. `my-client-id`)
3. **Client Secret**: Any string (e.g. `my-client-secret`)
4. **Scopes**: `openid email profile`
5. **Authorization Endpoint**: `https://oidc.<site>.<tld>/oauth2/authorize`
6. **Token Endpoint**: `https://oidc.<site>.<tld>/oauth2/token`
7. **Userinfo Endpoint**: `https://oidc.<site>.<tld>/oauth2/userinfo`

## Credits

**Contributed and maintained by [@Pronovix](https://github.com/Pronovix)**
