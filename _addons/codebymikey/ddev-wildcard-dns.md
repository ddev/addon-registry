---
title: "codebymikey/ddev-wildcard-dns"
github_url: "https://github.com/codebymikey/ddev-wildcard-dns"
description: "A DDEV add-on that provides a local wildcard DNS resolver for DDEV containers"
user: "codebymikey"
repo: "ddev-wildcard-dns"
repo_id: 1385491386
default_branch: "main"
tag_name: "0.2.0"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-09-24"
updated_at: "2026-09-25"
workflow_status: "success"
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/codebymikey/ddev-wildcard-dns/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/codebymikey/ddev-wildcard-dns/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/codebymikey/ddev-wildcard-dns)](https://github.com/codebymikey/ddev-wildcard-dns/commits)
[![release](https://img.shields.io/github/v/release/codebymikey/ddev-wildcard-dns)](https://github.com/codebymikey/ddev-wildcard-dns/releases/latest)

# DDEV Wildcard DNS

## Overview

This add-on runs a project-local wildcard DNS resolver for your [DDEV](https://ddev.com/) project.
It maps the configured wildcard domain to the DDEV router, which allows tools running inside the
web container, such as Puppeteer and Playwright, to resolve DDEV hostnames properly.

Without it, only the hostnames declared in the project configuration resolve to the router from
inside the container. Any other `*.ddev.site` subdomain falls through to public DNS, which answers
`127.0.0.1`, so the request ends up at the container itself rather than at the router.

See the [blog post](https://ddev.com/blog/ddev-name-resolution-wildcards/) for additional context.

If you have a fixed set of subdomains that need supporting, you can add them explicitly to your
`additional_hostnames` saving the need to use this add-on, however a restart will be needed for them
to be picked up.

Whereas having the custom DNS should allow it to work for any new domains without issue.

## Installation

```bash
ddev add-on get codebymikey/ddev-wildcard-dns
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

The add-on starts a `wildcard-dns` service and automatically adds that service as a resolver in the
web container during the `post-start` hook. No manual DNS configuration is required.

The resolver waits for the DDEV router to come up before answering, and if the router is later
recreated on a different IP address, for example when another project starts, it picks up the new
address automatically.

## Usage

The default wildcard domain is the project's DDEV top-level domain, normally `ddev.site`. For
example, `demo.example.ddev.site` resolves to the DDEV router when the project uses `ddev.site`.

| Command                                      | Description                            |
|----------------------------------------------|----------------------------------------|
| `ddev describe`                              | View the wildcard DNS service          |
| `ddev logs -s wildcard-dns`                  | View the local dnsmasq resolver logs   |
| `ddev exec getent hosts demo.demo.ddev.site` | Test resolution from the web container |

## Registering the resolver in other services

The add-on registers the resolver automatically in the `web` service. To enable wildcard DNS
resolution in another service, add a `post-start` hook for that service in a project config file,
such as `.ddev/config.yaml`. The script writes `/etc/resolv.conf`, so it has to run as root or as
a user with `sudo`; when the service's default user is neither, run the hook as root with `user:`:

```yaml
hooks:
  post-start:
    - service: db
      exec: /mnt/ddev_config/wildcard-dns/register-resolver.sh
    # A service whose default user is unprivileged and has no sudo.
    - service: playwright-mcp
      user: root
      exec: /var/www/html/.ddev/wildcard-dns/register-resolver.sh playwright-mcp
```

The script accepts an optional label as its first argument, which is only used to prefix its log
output. It defaults to the container hostname.

For a custom service, make sure the project's `.ddev` directory is mounted at `/mnt/ddev_config`
(or adjust the path in the hook, for example to `/var/www/html/.ddev/...` when the service mounts
the project root instead). The service image must provide `getent`, `awk` and `mktemp`.

When the script cannot register the resolver (no `getent`, no root and no `sudo`, or `/etc/resolv.conf`
not writable) it prints the reason to stderr and exits `0`, so `ddev start` still succeeds and the
service keeps its default resolver.

Restart the project after adding the hook:

```bash
ddev restart
```

## Customization

Set options in the project's DDEV environment file, then restart the project. The environment file
is project-specific and should be committed with the rest of `.ddev`:

```bash
ddev dotenv set .ddev/.env.wildcard-dns \
  --ddev-wildcard-dns-debug=true
ddev restart
```

Supported variables:

| Variable                   | `ddev dotenv set` option     | Default                            |
|----------------------------|------------------------------|------------------------------------|
| `DDEV_WILDCARD_DNS_DOMAIN` | `--ddev-wildcard-dns-domain` | `$DDEV_TLD` (normally `ddev.site`) |
| `DDEV_WILDCARD_DNS_DEBUG`  | `--ddev-wildcard-dns-debug`  | `false`                            |

To limit wildcard resolution to the current project, set the domain to the project's full hostname:

```bash
ddev dotenv set .ddev/.env.wildcard-dns \
  --ddev-wildcard-dns-domain=my-project.ddev.site
ddev restart
```

Replace `my-project` with your DDEV project name.

When debug mode is enabled, dnsmasq logs DNS queries and debug messages. View them with:

```bash
ddev logs -s wildcard-dns
```

## Credits

**Contributed and maintained by [@codebymikey](https://github.com/codebymikey)**
