[![Deploy to GitHub Pages](https://github.com/ddev/addon-registry/actions/workflows/deploy-to-github-pages.yml/badge.svg)](https://ddev.github.io/addon-registry/)

# DDEV Add-on Registry

A registry for DDEV add-ons where users can discover, explore, and leave comments on available add-ons.

While `ddev add-on list` is a fantastic tool, it only scratches the surface when it comes to discovering and exploring add-ons.

The need for a dedicated registry became clear as we sought to streamline access to add-ons, reduce API limitations, and lessen the maintenance burden.

For inspiration, [@stasadev](https://github.com/stasadev) looked to the simplicity and functionality of projects like [GTFOBins](https://gtfobins.github.io/), which is built with Jekyll.

## Used Tools

Here are the key tools used to build the DDEV Add-on Registry:

- [GitHub Pages and Jekyll](https://docs.github.com/en/pages/setting-up-a-github-pages-site-with-jekyll/about-github-pages-and-jekyll): Jekyll powers the static site, while GitHub Pages hosts it.
- [DataTables](https://datatables.net/): For sorting and searching add-on entries.
- Add-on comments: [giscus integration](https://github.com/ddev/giscus-comments).
- Golang: Used to aggregate add-on data into Markdown files, which are transformed into Liquid templates for Jekyll.

## How to Set Up the Registry Locally

Getting started with the DDEV Add-on Registry locally is straightforward. Here's how:

1. Clone the repository: <https://github.com/ddev/addon-registry>
2. Run `ddev start` to spin up the environment
3. Then, launch it by running `ddev launch :4000`
4. To update add-ons manually:
    ```bash
    export DDEV_ADDON_REGISTRY_TOKEN=your-classic-github-token-no-privileges
    cd go
    go mod vendor
    go run main.go
    ```

You'll now be able to explore the add-ons within the registry right on your local machine.

## Key Components and Structure

Here's a breakdown of where important content and configuration files live:

- **`.bundle`**: A script for installing Jekyll dependencies.
- **`.ddev`**: The DDEV configuration directory.
- **`.github`**: The GitHub workflows that handle the deployment process.
- **`_addons`**: Custom Jekyll collection that holds all the add-ons fetched from the community.
- **`_data`**: User-defined Jekyll data types.
- **`_includes`**: HTML partials used across the site.
- **`_layouts`**: The layout templates for Jekyll pages.
- **`_pages`**: Jekyll pages overrides.
- **`assets`**: Contains styles and static assets.
- **`go`**: The Golang app that collects add-on data and populates `_addons`.
- **`Gemfile`**: The Ruby equivalent of `composer.json`, managing dependencies.
- **`_config.yml`**: The main configuration file for the Jekyll site.
- **`addons.json`**: A JSON representation of all the DDEV add-ons.
- **`index.html`**: The homepage of the registry.
