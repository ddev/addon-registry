---
title: Pronovix/ddev-saml-idp
github_url: https://github.com/Pronovix/ddev-saml-idp
description: "A DDEV add-on that provisions a project-specific SimpleSAMLphp Identity Provider (IdP) as a dedicated, optional DDEV service."
user: Pronovix
repo: ddev-saml-idp
repo_id: 1259096372
default_branch: main
tag_name: 1.0.0-alpha1
ddev_version_constraint: ">= v1.25.2"
dependencies: []
type: contrib
created_at: 2026-06-04
updated_at: 2026-06-04
workflow_status: failure
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/Pronovix/ddev-saml-idp/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/Pronovix/ddev-saml-idp/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/Pronovix/ddev-saml-idp)](https://github.com/Pronovix/ddev-saml-idp/commits)
[![release](https://img.shields.io/github/v/release/Pronovix/ddev-saml-idp)](https://github.com/Pronovix/ddev-saml-idp/releases/latest)

# DDEV SAML Identity Provider (IdP) Add-on

A DDEV add-on that provisions a project-specific **SimpleSAMLphp Identity Provider (IdP)** as a dedicated, optional DDEV service. 

This add-on is designed to streamline local SAML integration and testing for web applications (especially Drupal sites utilizing the `samlauth` module) by providing a completely self-contained, fully customizable local SAML Identity Provider.

## Key Features

- **Optional & Resource-Friendly:** Defined via a Docker Compose profile (`saml-idp`), meaning the container only runs when you actively need to develop/test SAML authentication.
- **Predictable Dynamic Hostnames:** Accessible at `https://idp.<your-project>.ddev.site/simplesaml/`. Fully integrated with DDEV's router and trusting your local `mkcert` CA for automatic HTTPS termination.
- **Environment-Agnostic Setup:** No hardcoded URLs in your committed configuration. SimpleSAMLphp config and metadata templates resolve your specific DDEV project hostname dynamically at runtime via `getenv('SAML_IDP_PRIMARY_HOST')` and `getenv('SAML_SP_PRIMARY_HOST')`.
- **Automatic Cryptographic Initialization:** Self-signed certificates and private keys (`idp.crt`, `idp.key`, `sp.crt`, `sp.key`) are automatically initialized inside `.ddev/saml-idp/certs/` upon container boot. This folder is gitignored automatically.
- **Safe Add-on Upgrades:** Re-installing or updating the add-on (`ddev add-on get`) updates only the core container infrastructure and templates, **never** overwriting your custom configuration, metadata, or certificates in `.ddev/saml-idp/`.
- **Pre-seeded Test Personas:** Ships with ready-to-use user profiles under the `example-userpass` SimpleSAMLphp auth source, populated with realistic SAML user attributes.
- **Version Control Pinning:** Control base PHP version (`PHP_IMAGE_TAG`, default: `8.4`) and SimpleSAMLphp library version (`SSP_VERSION`, default: `^2.2`) via DDEV environment variables.
- **Native Multi-Architecture Support:** Built locally from source, running natively on both Apple Silicon (`linux/arm64`) and Intel/AMD (`linux/amd64`) machines.

## Installation & Setup

1. **Install the add-on:**
   ```bash
   ddev add-on get Pronovix/ddev-saml-idp
   ```

2. **Start/Run the service:**

    - **Persistent configuration (recommended):**
      If you want the SAML IdP service to always start automatically alongside your project, add the `saml-idp` profile to your `.ddev/config.local.yaml` (or `.ddev/config.yaml`) file:
      ```yaml
      profiles:
        - saml-idp
      ```
      Then restart your project:
      ```bash
      ddev restart
      ```

    - **On-demand (alternative):**
      If you prefer to start the SAML IdP only when needed, without persisting it to your DDEV configuration, run `ddev restart` first to ensure the router is ready, then start with the profile:
      ```bash
      ddev restart
      ddev start --profiles=saml-idp
      ```

   > [!NOTE]
   > On the very first start after installation, the container image is built from scratch. The `ddev restart` step above accounts for this. Subsequent starts work normally without any extra steps.

## Service Access & Endpoints

Once running, the SAML IdP is accessible via the following local URLs:

- **SimpleSAMLphp Admin UI:** `https://idp.<your-project>.ddev.site/simplesaml/`
- **IdP Metadata XML Endpoint:** `https://idp.<your-project>.ddev.site/simplesaml/module.php/saml/idp/metadata`
- **IdP Single Sign-On (SSO) Endpoint:** `https://idp.<your-project>.ddev.site/simplesaml/saml2/idp/SSOService.php`
- **IdP Single Logout (SLO) Endpoint:** `https://idp.<your-project>.ddev.site/simplesaml/saml2/idp/SingleLogoutService.php`
- **Default Admin Password:** `admin` *(configured via `SSP_ADMIN_PASSWORD`)*
- **Default Secret Salt:** `default_development_secret_salt_12345678_override_me` *(configured via `SSP_SECRET_SALT`)*

## Pre-seeded Test Personas

The default authentication source (`example-userpass`) comes preloaded with three test personas and standard SAML attributes (including `uid`, `email`, `givenName`, `sn`, `cn`, and `groups`):

| Username | Password | Display Name (`cn`) | Primary E-mail (`email`) | Roles / Affiliations (`groups`) |
| :--- | :--- | :--- | :--- | :--- |
| **`admin`** | `password` | Admin User | `admin@example.com` | `administrator`, `editor` |
| **`editor`** | `password` | Content Editor | `editor@example.com` | `editor` |
| **`user1`** | `password` | Test User 1 | `user1@example.com` | `member` |

## Step-by-Step Integration with Drupal `samlauth`

Setting up SAML integration in local Drupal development is incredibly easy with this add-on. Here is a step-by-step configuration guide for the Drupal **SAML Authentication (`samlauth`)** module:

### Step 1: Install `samlauth`
Run the following in your DDEV project to download and enable the module:
```bash
ddev composer require drupal/samlauth
ddev drush pm:enable samlauth -y
```

### Important: Automated Configuration vs. Manual Configuration

> [!NOTE]
> **Automatic Settings Overrides:**
> When you run `ddev add-on get Pronovix/ddev-saml-idp`, the installer automatically appends a block of configuration overrides to your `settings.local.php` (or `project.local.settings.php`) file between the markers:
> `// --- BEGIN DDEV SAML IDP OVERRIDES ---`
> and
> `// --- END DDEV SAML IDP OVERRIDES ---`
>
> This block automatically configures the Drupal `samlauth` module at runtime with optimal development values (pointing directly to the correct certificate files and SimpleSAMLphp endpoints). **This means you do not even need to configure these settings manually in the Admin UI!**
>
> **If you wish to configure or verify them manually via the Drupal Admin UI (`/admin/config/people/saml`), you MUST empty or comment out the content inside the `// --- BEGIN DDEV SAML IDP OVERRIDES ---` block first.** Otherwise, the PHP settings overrides take precedence and any settings you save in the UI will be ignored.

### Step 2: Configure Service Provider (SP) Keys in Drupal
The IdP requires your Service Provider (Drupal) to sign its authentication requests. The DDEV SAML IdP container automatically generates matching keys for this purpose on startup inside `.ddev/saml-idp/certs/`.

Since `.ddev/` is mounted inside the DDEV web container at `/mnt/ddev_config/`, your Drupal site can point directly to these keys on the local filesystem:
- **SP Private Key Path:** `/mnt/ddev_config/saml-idp/certs/sp.key` *(or with file prefix `file:/mnt/ddev_config/saml-idp/certs/sp.key`)*
- **SP X.509 Certificate Path:** `/mnt/ddev_config/saml-idp/certs/sp.crt` *(or with file prefix `file:/mnt/ddev_config/saml-idp/certs/sp.crt`)*

### Step 3: Configure `samlauth` Module Settings (Manual UI Verification)
If you have cleared/commented out the settings overrides block and want to set up `samlauth` manually in the Drupal UI at `/admin/config/people/saml`, enter the following values:

#### Service Provider (SP) Settings
- **SP Entity ID:** `[site:base-url]` *(or `https://<your-project>.ddev.site`)*
- **Type of storage:** `File`
- **Private key path:** `file:/mnt/ddev_config/saml-idp/certs/sp.key`
- **X.509 certificate path:** `file:/mnt/ddev_config/saml-idp/certs/sp.crt`

#### Identity Provider (IdP) Settings
- **IdP Entity ID:** `https://idp.<your-project>.ddev.site/simplesaml/saml2/idp/metadata`
- **Single Sign-On Service:** `https://idp.<your-project>.ddev.site/simplesaml/saml2/idp/SSOService.php`
- **Single Logout Service:** `https://idp.<your-project>.ddev.site/simplesaml/saml2/idp/SingleLogoutService.php`
- **X.509 Certificate:** Enter `file:/mnt/ddev_config/saml-idp/certs/idp.crt` into this field, OR copy and paste the entire PEM content of `.ddev/saml-idp/certs/idp.crt` (including the `-----BEGIN CERTIFICATE-----` and `-----END CERTIFICATE-----` lines).

#### User Attribute Mapping (on the **Login / Users** tab)
- **Unique ID source:** `Attribute`
- **Attribute name:** `uid`
- **User name attribute:** `uid` *(or `cn`)*
- **User mail attribute:** `email`
- Check **Create users from SAML data** *(to auto-provision accounts on first login)*
- Check **Attempt to link SAML data to existing Drupal users**
- Click **Save configuration**.

### Step 4: Test Login!
1. Open a new Incognito/Private browser window.
2. Go to `https://<your-project>.ddev.site/saml/login`.
3. You will be redirected to your local SimpleSAMLphp login page at `https://idp.<your-project>.ddev.site`.
4. Enter credentials for one of the pre-seeded users (e.g., username `user1`, password `password`) and click **Login**.
5. You will be successfully logged in and redirected back to your Drupal site with your new user account mapped!

## Advanced Customization

### Adding / Modifying Test Users
To add new test accounts or modify user attributes, open `.ddev/saml-idp/config/authsources.php` in your editor and add entries under the `example-userpass` array using standard SimpleSAMLphp conventions. Commit this file to git to share these test personas with your entire development team.

### Adding Additional Service Providers
If you need to connect more than one SP (for example, if you are developing a multi-site Drupal architecture), edit `.ddev/saml-idp/metadata/saml20-sp-remote.php` to register other remote SP metadata configurations.

### SimpleSAMLphp & PHP Version Pinning
You can easily customize the PHP version or SimpleSAMLphp version without modifying the add-on source. Add/update the following environment variables in `.ddev/.env` (or configure them via `.ddev/docker-compose.saml-idp.yaml` build arguments):

```env
# Pin the PHP version of the IdP container (e.g., 8.2, 8.3, 8.4)
PHP_IMAGE_TAG=8.3

# Pin the exact SimpleSAMLphp composer version or version constraint
SSP_VERSION=2.2.0
```
Run `ddev restart` to rebuild the container with your newly specified versions.

## Troubleshooting & Logs

- **Inspect IdP Container Logs:**
  ```bash
  ddev logs -s saml-idp
  ```
- **Reset/Regenerate Signing Certificates:**
  If you ever need to regenerate the certificates and private keys, simply delete them and restart/start the IdP service (using `ddev restart` if the profile is persistent, or `ddev start --profiles=saml-idp` if starting on-demand):
  ```bash
  rm -f .ddev/saml-idp/certs/*
  ddev restart
  # OR if running on-demand:
  # ddev start --profiles=saml-idp
  ```
  The entrypoint script will automatically detect the missing certificates and securely generate a brand-new, matching cryptographic key set.

## Credits

Contributed and maintained by **[@Pronovix](https://github.com/Pronovix)**.
