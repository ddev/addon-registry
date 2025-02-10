---
title: torhoehn/ddev-opensearch
github_url: https://github.com/torhoehn/ddev-opensearch
description: "Opensearch add-on for ddev"
user: torhoehn
repo: ddev-opensearch
categories:
  - community
created_at: 2023-07-26
updated_at: 2023-09-07
stars: 0
---

[![tests](https://github.com/ddev/ddev-opensearch/actions/workflows/tests.yml/badge.svg)](https://github.com/ddev/ddev-opensearch/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2024.svg)

## Installation

Uses [opensearch official image](https://hub.docker.com/r/opensearchproject/opensearch)

`ddev get ddev/ddev-opensearch`

## Configuration

From within the container, the opensearch container is reached at hostname: "opensearch", port: 9200, so the server URL might be `http://opensearch:9200`. You can also use the "ddev.site" http and https urls to access it: `http://<projectname>.ddev.site:9200`, and `https://<projectname>.ddev.site:9201`

## Connection

You can access the OpenSearch server directly from the host for debugging purposes by visiting `http://<projectname>.ddev.site:9200`. Via https you can access OpenSearch via `https://<projectname>.ddev.site:9201`

## Memory Limit

This configuration limits memory usage to 512mb. This should be enough for most projects, but if your `opensearch` service stops with no obvious reason, increase your docker max memory and/or the service max memory via the ES_JAVA_OPTS variable:

```yaml
- "ES_JAVA_OPTS=-Xms512m -Xmx512m"` environment variable in `docker-compose.opensearch.yaml`
```

If you change this variable, make sure to remove the `#ddev-generated` line at the top of the file.

You can use `ddev logs -s opensearch` to investigate what the opensearch daemon has been up to, or if you have a RAM-related crash.

**Contributed and maintained by [@torhoehn](https://github.com/torhoehn) based on the original [ddev-contrib recipe](https://github.com/ddev/ddev-contrib/tree/master/docker-compose-services/RECIPE) by [@CONTRIBUTOR](https://github.com/CONTRIBUTOR)**

**Originally Contributed by [somebody](https://github.com/somebody) in <https://github.com/ddev/ddev-contrib/>**

