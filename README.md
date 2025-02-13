[![Deploy to GitHub Pages](https://github.com/ddev/addon-registry/actions/workflows/deploy-to-github-pages.yml/badge.svg)](https://ddev.github.io/addon-registry/)

# DDEV Add-on Registry

A registry for DDEV add-ons where users can discover, explore, and leave comments on available add-ons.

- [Jekyll](https://jekyllrb.com/) is used under the hood to generate the site.
- Add-ons in the [_addons](./_addons) directory are automatically updated using [Golang](./go/main.go).
- Manual updates can be performed using:
    ```bash
    export DDEV_ADDON_REGISTRY_TOKEN=your-classic-github-token-no-privileges
    cd go
    go mod vendor
    go run main.go
    ```

## Local Development Setup

DDEV already has all the dependencies included.

1. Run `ddev start` to start and set up the project's dependencies (the first start is slow).
2. Run `ddev launch :4000` to open the development server with LiveReload.
