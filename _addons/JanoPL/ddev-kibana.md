---
title: JanoPL/ddev-kibana
github_url: https://github.com/JanoPL/ddev-kibana
description: "Kibana add-on for DDEV"
user: JanoPL
repo: ddev-kibana
repo_id: 530812551
ddev_version_constraint: ""
dependencies: ["elasticsearch"]
type: contrib
created_at: 2022-08-30
updated_at: 2024-10-27
workflow_status: disabled
stars: 1
---

[![tests](https://github.com/janopl/ddev-kibana/actions/workflows/tests.yml/badge.svg)](https://github.com/janopl/ddev-kibana/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2024.svg)

## Instalation

Uses [Kibana official image](https://registry.hub.docker.com/_/kibana)

For DDEV v1.23.5 or above run

```sh
ddev add-on get janopl/ddev-kibana
```

For earlier versions of DDEV run

```sh
ddev get janopl/ddev-kibana
```

## Configuration

From within the container, the kibana container is reached at hostname "kibana", port: 5601

### Kibana Version 
To adjust the version of your elastic search, you can use the new argument variable that docker compose provides for the version.

```docker-compose.kibana.yml```
```
services:
    kibana:
        build:
            ...
            args:
                - KIBANA_VERSION=7.17.14 // example: 8.10.2
        ...
        environment:
            KIBANA_VERSION: 7.17.14 // example: 8.10.2
```

OR 

After adding the add-on, run ```cp .ddev/elasticsearch/docker-compose.elasticsearch8.yaml .ddev/``` to enable Kibana 8.

### Configuration file
You can configure Kibana dashboard through the config file under: ```.ddev/kibana/config.yml```

## Connection

You can access the Kibana server directly from the host by visiting:

- `https://<projectname>.ddev.site:5601`
- `http://<projectname>.ddev.site:5600`

## Contribution

First off, thanks for taking the time to contribute! Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make will benefit everybody else and are **greatly appreciated**.


Please read [our contribution guidelines](https://github.com/JanoPL/ddev-kibana/blob/main/./docs/CONTRIBUTING.md), and thank you for being involved!

## License

This project is licensed under the **APACHE license**.

See [LICENSE](https://github.com/JanoPL/ddev-kibana/blob/main/LICENSE) for more information.
