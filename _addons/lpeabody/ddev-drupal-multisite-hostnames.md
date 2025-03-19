---
title: lpeabody/ddev-drupal-multisite-hostnames
github_url: https://github.com/lpeabody/ddev-drupal-multisite-hostnames
description: "A small DDEV addon for dynamically generating hostnames based on Drupal multisite directory names."
user: lpeabody
repo: ddev-drupal-multisite-hostnames
repo_id: 950388191
ddev_version_constraint: ">= v1.24.2"
dependencies: []
type: contrib
created_at: 2025-03-18
updated_at: 2025-03-18
stars: 0
---

# DDEV Drupal Multisite Hostnames

This addon for Drupal projects creates a `$sites` entry for each multisite
directory in the project. The entry is based on the name of the multisite
directory and the DDEV project name.

For example, if your project name is `my-ddev-project`, and you have two
multisite directories:

- default
- other_site

Then the `$sites` array will look like the following:

```php
$sites = [
  'other-site.my-ddev-project.ddev.site' => 'other_site',
];
```

## Why use this?

This addon was born of the desire to be able to run multiple copies of the
same DDEV project on the same host simultaneously. In order to do this, you
must have unique hostnames for each DDEV project/multisite directory
combination.

If you do not need to accommodate this workflow, you may not need this addon,
and can instead manually maintain `additional_hostnames` for your site
directories. This workflow requires that you only have one instance of the same
DDEV project running on your host at any given time.

## Installation

```bash
ddev add-on get lpeabody/ddev-drupal-multisite-hostnames
```

- `drupal_multisite/sites.ddev_multisite.php` is copied to `sites/sites.ddev_multisite.php`
- A command `multisite-start` is added to `.ddev/commands/host/multisite-start`
- `config.drupal_multisite_hosts.yaml` is added to `.ddev`

## Usage

Just install the addon and then run `ddev start` to apply the new hostname
pattern, which by default is calculated using the pattern `*.$DDEV_SITENAME`.

1. Install addon.
2. `ddev start`
3. `ddev status` and confirm that the wildcard site name URL was added.

Referring back to the example above, if you have the `other_site` directory
then you should be able to now access that site via 
https://other-site.my-ddev-project.ddev.site.

> [!IMPORTANT]
> If you change the project name, and therefore the site name, you must run
> `ddev multisite-start` to regenerate the additional hostname file and apply
> it.

> [!NOTE]
> Site directories named `default` and `settings` are ignored.
> The `default` directory is not necessary to set in most, if not all,
> circumstances in sites.php. The `settings` directory is conventionally used
> not as a site directory but as a location to store global settings files used
> by all sites.


### Customizing the hostname pattern

You can manually edit the `sites/sites.ddev_multisite.php` file and the
`multisite-start` command to customize the hostname pattern. When you update
the addon you might have to reapply your changes.
