---
title: OpenForgeProject/ddev-woodoo-buildtools-magento
github_url: https://github.com/OpenForgeProject/ddev-woodoo-buildtools-magento
description: "ddev Addon to build Magento projects without headache"
user: OpenForgeProject
repo: ddev-woodoo-buildtools-magento
repo_id: 614074383
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2023-03-14
updated_at: 2025-05-22
workflow_status: success
stars: 11
---

# [DDEV] Woodoo Frontend Buildtools for Magento & Hyvä

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/ac65544f530a4240a282f4bf988abb90)](https://app.codacy.com/gh/OpenForgeProject/ddev-woodoo-buildtools-magento/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade) [![tests](https://github.com/OpenForgeProject/ddev-woodoo-buildtools-magento/actions/workflows/tests.yml/badge.svg)](https://github.com/OpenForgeProject/ddev-woodoo-buildtools-magento/actions/workflows/tests.yml) <img src="https://img.shields.io/github/last-commit/OpenForgeProject/ddev-woodoo-buildtools-magento" alt="Last commit date Badge">
<img src="https://img.shields.io/github/v/release/OpenForgeProject/ddev-woodoo-buildtools-magento" alt="Release-Badge"> <img src="https://img.shields.io/github/sponsors/dermatz" alt="Sponsors"> [<img src="https://img.shields.io/badge/Discord-Join%20Chat-orange" alt="Join our Discord">](https://discord.gg/H5CjMXQQHn)

---

This ddev add-on helps you to manage your theme build processes in Magento easily. You can run a build command to build all your themes or a specific theme. It doesn't matter if the theme is a Magento Standard Theme, a Custom Theme, or a Hyva theme. It automatically takes care of which type of theme you use and builds it for you. If a dependency is missing, Woodoo will install it.

_>> Build all your Magento themes at once, with Magic, build with Woodoo!_

| Supported Versions                                                                                                        | Status                                                                    |
| ------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| <img src="https://img.shields.io/badge/Magento_2_Hyvä_(all_Versions)-Supported-43A047" alt="All Hyvä Versions supported"> | ![project is maintained](https://img.shields.io/maintenance/yes/2025.svg) |
| <img src="https://img.shields.io/badge/Magento_2_Hyvä_Fallback-Supported-43A047" alt="Hyvä Fallback is supported">        | ![project is maintained](https://img.shields.io/maintenance/yes/2025.svg) |
| <img src="https://img.shields.io/badge/Magento_2.x-Supported-43A047" alt="Magento 2 Supported">                           | ![project is maintained](https://img.shields.io/maintenance/yes/2025.svg) |
| <img src="https://img.shields.io/badge/MageOS-Supported-43A047" alt="MageOS is supported">                                | ![project is maintained](https://img.shields.io/maintenance/yes/2025.svg) |
| <img src="https://img.shields.io/badge/Avanta B2B-Supported-43A047" alt="Avanta B2B is supported">                            | ![project is maintained](https://img.shields.io/maintenance/yes/2025.svg) |

---

## Features

- **Automatic Theme Detection**: Automatically detects and builds Magento Standard, Custom, and Hyvä themes.
- **Dependency Management**: Installs missing dependencies automatically.
- **Theme Configuration**: Easily configure theme paths in `.ddev/config-themes.yaml`.
- **Build Commands**: Build all themes or specific themes with simple commands.
- **Watch Mode**: Watch for CSS and JS changes in specific themes.
- **Self-Test**: Perform a self-check to identify potential problems.
- **Update Notifications**: Get notified about new releases and update easily.
- **Developer Preview**: Option to use the latest developer preview version.

## Getting Started

### Installation

**1. Get the latest Version with following command in your ddev project root**

##### Latest stable version (recommended)

For DDEV v1.23.5 or above run

```shell
ddev add-on get OpenForgeProject/ddev-woodoo-buildtools-magento
```

For earlier versions of DDEV run

```shell
ddev get OpenForgeProject/ddev-woodoo-buildtools-magento
```

##### Latest Developer-Preview (Main-Branch - unreleased) - can be unstable!

For DDEV v1.23.5 or above run

```shell
ddev add-on get https://github.com/OpenForgeProject/ddev-woodoo-buildtools-magento/archive/refs/heads/main.tar.gz
```

For earlier versions of DDEV run

```shell
ddev get https://github.com/OpenForgeProject/ddev-woodoo-buildtools-magento/archive/refs/heads/main.tar.gz
```

---

**2. Start Woodoo INIT to update your ./ddev/config.yaml**

Automatically connect to your ddev database to list all available themes and add these to `.ddev/config-themes.yaml`. Add theme path if possible (can be edited later)

```shell
ddev frontend init
```

---

**3. Edit theme path in .ddev/config-themes.yaml**
Open `.ddev/config-themes.yaml` and specify relative path to the theme based on your project structure. (E.g. `src/themeVendor/theme ...` or `app/design/frontend/themeVendor/theme/ ...`)

## Update Woodoo to latest stable Version

If a newer Release is available you will be notified in your CLI about new versions. To update use `ddev frontend-update` for the latest stable Version.

```shell
  ddev frontend-update
```

## Get Developer Preview

If you want to get the latest developer preview (can be unstable) use `ddev frontend-update -dev`

```shell
  ddev frontend-update -dev
```

## Usage

```shell
Usage: ddev frontend [command] [option or theme]

Command:
  init, install         Initial Woodoo Setup to read all themes from database and write to .ddev/config-themes.yaml
  help, -help, -h       Displays help page
  version, --version    Show version
  check, selftest       Frontend Self-Test
  frontend-update       Update to the latest Version (stable Version)
  frontend-update -dev  Update to the latest Main-Branch Version (not recommended)
  themes                List all available themes that are configured in .ddev/config-themes.yaml
  build                 Builds all themes that are configured in .ddev/config-themes.yaml
  build -f              Builds all configured themes without yes/no confirmation
  build theme           Build a specific theme
  watch theme           Watch for CSS and JS changes in a specific theme

Option:
  -f                    Force the build command to run without yes/no confirmation

Aliases:
  f, fe
```

## Usage Examples

Here are some usage examples to help you get started:

### Building All Themes

To build all themes configured in `.ddev/config-themes.yaml`, run:

```shell
ddev frontend build
```

### Building a Specific Theme

To build a specific theme, use the theme code as an argument:

```shell
ddev frontend build <theme_code>
```

### Watching for Changes

To watch for CSS and JS changes in a specific theme, use:

```shell
ddev frontend watch <theme_code>
```

## Troubleshooting Tips

-   Ensure that your theme paths in `.ddev/config-themes.yaml` are correct and relative to your project root.
-   If you encounter any issues, try running `ddev frontend selftest` to perform a self-check and identify potential problems.

## Feature request

Submit your [Feature Request](https://github.com/OpenForgeProject/ddev-woodoo-buildtools-magento/issues) by adding a new issue and add a label with `feature-request`.

---

## Contributing

Contributions are welcome! Please read the [Contributing Guidelines](https://github.com/OpenForgeProject/ddev-woodoo-buildtools-magento/blob/main/./CONTRIBUTING.md) first.
Feel free to send your improvements as [Pull request](https://github.com/OpenForgeProject/ddev-woodoo-buildtools-magento/pulls).

---

### Watch out for our Contributors, Testers & Supporters

-   [Mathias Elle | @dermatz](https://github.com/dermatz) (Maintainer)
-   [Thomas Hauschild | @Morgy93](https://github.com/Morgy93)
-   [e3n GmbH & Co. KG - E-Commerce Agency based in Mainz | e3n.de](https://e3n.de)
-   [Tom Muir | @t-muir](https://github.com/t-muir)

---

## License

[License](https://github.com/OpenForgeProject/ddev-woodoo-buildtools-magento/blob/main/./LICENSE) | by [Mathias Elle](https://www.linkedin.com/in/mathias-elle-842783102/)
