---
title: jfastnacht/ddev-mitmproxy
github_url: https://github.com/jfastnacht/ddev-mitmproxy
description: "Use mitmproxy to catch and analyze outgoing HTTP/HTTPS traffic within ddev"
user: jfastnacht
repo: ddev-mitmproxy
repo_id: 1178691817
default_branch: main
tag_name: 1.0.0
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2026-03-11
updated_at: 2026-03-15
workflow_status: failure
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/jfastnacht/ddev-mitmproxy/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/jfastnacht/ddev-mitmproxy/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/jfastnacht/ddev-mitmproxy)](https://github.com/jfastnacht/ddev-mitmproxy/commits)
[![release](https://img.shields.io/github/v/release/jfastnacht/ddev-mitmproxy)](https://github.com/jfastnacht/ddev-mitmproxy/releases/latest)

# DDEV mitmproxy

## Overview

This add-on integrates mitmproxy into your [DDEV](https://ddev.com/) project.

## Installation

```bash
ddev add-on get jfastnacht/ddev-mitmproxy
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Configuration

### Basic

Set the proxy configuration for your web container in `docker-compose.web_extra.yaml` or for whatever
other service you run:

```yaml
services:
  web:
    environment:
      - http_proxy=http://mitmproxy:8080/
      - no_proxy=localhost,127.0.0.1
      - HTTP_PROXY=http://mitmproxy:8080/
      - HTTPS_PROXY=http://mitmproxy:8080/
      - SSL_CERT_FILE=/etc/ssl/certs/ca-certificates.crt
```

You might have to add more IPs/domains to your `no_proxy` configuration, if you don't want specific
traffic to go through `mitmproxy`.

### HTTPS

> [!IMPORTANT]
> If you have set the proxy configuration for HTTPS, it is necessary to run this hook before any
connection is made.

If you want to use HTTPS, you have to install the certificates on the web container, e.g. `config.mitmproxy.yaml`:

```yaml
hooks:
  post-start:
    - exec: "wget http://mitm.it/cert/pem -O /usr/local/share/ca-certificates/mitmproxy.crt"
    - exec: "sudo update-ca-certificates"
```

If you want to use HTTPS in any other container, you might have to adjust this example according to your needs:

```yaml
hooks:
  post-start:
    - exec-host: "ddev exec -s myservice wget http://mitm.it/cert/pem -O /usr/local/share/ca-certificates/mitmproxy.crt"
    - exec-host: "ddev exec -s myservice sudo update-ca-certificates"
```

### Application

You probably have to set up proxies for your PHP application or other services as well. A simple example would be:

```php
<?php
$ch = curl_init();
curl_setopt($ch, CURLOPT_PROXY, "mitmproxy:8080");
curl_setopt($ch, CURLOPT_URL, "https://github.com");
curl_exec($ch);
curl_close($ch);
```

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev describe` | View service status and used ports for mitmproxy |
| `ddev logs -s mitmproxy` | Check mitmproxy logs |

## Advanced Customization

To change the Docker image:

```bash
ddev dotenv set .ddev/.env.mitmproxy --mitmproxy-docker-image="mitmproxy/mitmproxy:latest"
ddev add-on get jfastnacht/ddev-mitmproxy
ddev restart
```

Make sure to commit the `.ddev/.env.mitmproxy` file to version control.

All customization options (use with caution):

| Variable | Flag | Default |
| -------- | ---- | ------- |
| `MITMPROXY_DOCKER_IMAGE` | `--mitmproxy-docker-image` | `mitmproxy/mitmproxy:latest` |
| `MITMPROXY_PASSWORD` | `--mitmproxy-password` | `mitm` |

## Credits

**Contributed and maintained by [@jfastnacht](https://github.com/jfastnacht)**
