[![Deploy to GitHub Pages](https://github.com/stasadev/ddev-addon-registry/actions/workflows/deploy-to-github-pages.yml/badge.svg)](https://stasadev.github.io/ddev-addon-registry/)

# DDEV Add-on Registry

A registry for DDEV add-ons where users can discover, explore, and leave comments on available add-ons.

- [Jekyll](https://jekyllrb.com/) is used under the hood to generate the site.
- TODO: DDEV add-ons are updated on daily basis.

# Local usage

1. `ddev start`
2. `ddev logs -s jekyll -f` (wait until the server is ready)
3. `ddev launch :4000/ddev-addon-registry/`
