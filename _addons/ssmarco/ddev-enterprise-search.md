---
title: ssmarco/ddev-enterprise-search
github_url: https://github.com/ssmarco/ddev-enterprise-search
description: "Elastic Enterprise Search service for DDEV"
user: ssmarco
repo: ddev-enterprise-search
categories:
  - community
created_at: 2024-01-30
updated_at: 2025-01-14
stars: 0
---

[![tests](https://github.com/ssmarco/ddev-enterprise-search/actions/workflows/tests.yml/badge.svg)](https://github.com/ssmarco/ddev-enterprise-search/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2024.svg)

# ddev-enterprise-search <!-- omit in toc -->

- [Introduction](#introduction)
- [Getting started](#getting-started)
- [How to debug in Github Actions](#how-to-debug-tests-github-actions)

## Introduction

ddev-enterprise-search is an un-official implementation of Elastic Enterprise Search service for DDEV based on their Docker guide\*.

Enterprise Search is an additional Elastic service that adds APIs and UIs to those already provided by Elasticsearch and Kibana.

Currently sitting at version 8.12.0, part of the implementation as a service for DDEV includes a Kibana container.
This means that to use this service, existing Kibana service needs to be uninstalled in your project and you should install the supported Elastic Search from DDEV.

From your DDEV project, install this by running `ddev get ssmarco/ddev-enterprise-search` followed by `ddev restart`.
This can take up to 30 minutes or so due to downloading the required Docker containers (Elastic Search, Kibana and Enterprise Search).

- [Reference](https://www.elastic.co/guide/en/enterprise-search/current/start.html)
- [Docker guide\*](https://www.elastic.co/guide/en/enterprise-search/current/docker.html)

## Getting started

1. In the DDEV project directory:

    For DDEV v1.23.5 or above run

    ```sh
    ddev add-on get ddev/ddev-elasticsearch
    ddev add-on get ssmarco/ddev-enterprise-search
    ```

    For earlier versions of DDEV run

    ```sh
    ddev get ddev/ddev-elasticsearch
    ddev get ssmarco/ddev-enterprise-search
    ```

2. Restart the DDEV instance:

    ```sh
    ddev restart
    ```

3. Get the URL of the Kibana dashboard (e.g. https://your-project-name.ddev.site:5602):

    ```sh
    ddev describe
    ```

4. Login with the username `elastic` and password `elastic`.

5. The API credentials can be obtained by following the URL https://YOUR-PROJECT-NAME.ddev.site:5602/app/enterprise_search/app_search/credentials (use your actual project name).

## Configuring your framework

### Silverstripe

1. Update your project's `.env` file. The API keys are found in the Enterprise Search section of Kibana dashboard.

    ```
    ENTERPRISE_SEARCH_ENGINE_PREFIX="my-index"
    ENTERPRISE_SEARCH_API_KEY="private-xxxxxxxxxxxx-change-this"
    ENTERPRISE_SEARCH_API_SEARCH_KEY="search-xxxxxxxxxxxx-change-this"
    ENTERPRISE_SEARCH_ENDPOINT="http://enterprisesearch:3002"
    ```

2. The Enterprise Search endpoint is `http://enterprisesearch:3002`

3. The following modules are tested to work out of the box in your composer.json file:

    ```
    "silverstripe/silverstripe-search-service": "^3.0",
    "silverstripe/silverstripe-search-service-elastic": "^1.0@beta",
    ```

## Recommendations

This is for local development purposes only. Testing large amount of data depends on the host computer's resources.
If you have a good amount of CPU's and memory, you can increase the value of `mem_limit` for each container or remove this attribute to assign more resources as needed.

We have found that 4 to 6GB allocated to the docker provider seems to provide the best experience.

The default mem_limit (2GB) is usually enough for smaller indexes, under or around 1,000 documents. However, you may might encounter memory exhausted errors after a few hours or running elastic search. 
This will happen when the elastic dashboard starts returning `502` errors. You can check the logs or the `ddev describe` to confirm.

If this happens you can either restart (`ddev restart`), or increase the `mem_limit` to a higher allocation. 

For more than 5k documents we recommend allocating your docker provider with 4 to 6 GB of memory, and increasing the default `mem_limit` from 2GB to 4GB.

### How to configure your docker provider:
**Docker desktop**
1. Open Docker Desktop app
2. Click Preferences
3. Click 'Advanced' tab, you should then be able to configure the memory / cpu limit 

**Colima**
You can change the memory and cpu by stopping and starting colima with preferred options
```
colima stop
colima start --cpu 4 --memory 6
```

### Increasing the mem_limit
In the  `.ddev/docker-compose.enterprise-search.yaml` you can increase the `mem_limit`.
``` yaml
services:
    web:
        # elastic / config / kibana
        enterprisesearch:
            mem_limit: 6g
```

This will require you to restart your ddev box `ddev restart`.



## Troubleshooting

1. Make sure all required containers are downloaded

    ```sh
    docker pull docker.elastic.co/elasticsearch/elasticsearch:8.12.0
    docker pull docker.elastic.co/kibana/kibana:8.12.0
    docker pull docker.elastic.co/enterprise-search/enterprise-search:8.12.0
    ```

2. Remove container volumes to restart from scratch

    List all existing volumes from your system:

    ```sh
    docker volume ls
    ```

    This will show example output below:

    ```
    DRIVER    VOLUME NAME
    local     ddev-your-project-name_elastic-certs
    local     ddev-your-project-name_elastic-data
    local     ddev-your-project-name_elastic-kibana
    local     ddev-your-project-name_enterprise-data
    ```

    Delete the volumes by running:

    ```
    docker volume rm ddev-your-project-name_elastic-certs \
    ddev-your-project-name_elastic-data \
    ddev-your-project-name_elastic-kibana \
    ddev-your-project-name_enterprise-data
    ```

3. Restart by `ddev restart`

4. Check the status of the project by `ddev status`

5. Check the logs

    The `elastic-config` container does the necessary prerequisites to glue together the other containers. It checks connection to Elastic Search and resets `kybana_system` password.

    ```
    ddev logs -s elastic-config
    ddev logs -s elasticsearch
    ddev logs -s kibana
    ddev logs -s enterprisesearch
    ```

6. Check job health

    You might need to install `jq` for better legibility of the output.

    ```
    docker inspect --format "{{json .State.Health }}" ddev-your-project-name-enterprisesearch | jq
    docker inspect --format "{{json .State.Health }}" ddev-your-project-name-kibana | jq
    docker inspect --format "{{json .State.Health }}" ddev-your-project-name-elasticsearch | jq
    ```

7. Check memory consumptions

    ```
    docker stats
    ```


## Contribute

- Anyone is welcome to submit a PR to this repo. See README.md at https://github.com/ddev/ddev-addon-template, the parent of this repo.

## Maintainer

- Contributed and maintained by [Marco Hermo](https://github.com/ssmarco).

