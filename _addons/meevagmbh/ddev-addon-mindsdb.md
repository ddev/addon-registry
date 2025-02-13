---
title: meevagmbh/ddev-addon-mindsdb
github_url: https://github.com/meevagmbh/ddev-addon-mindsdb
description: "MindsDB addon service for DDEV"
user: meevagmbh
repo: ddev-addon-mindsdb
type: contrib
created_at: 2024-12-12
updated_at: 2024-12-13
stars: 0
---

[![tests](https://github.com/meevagmbh/ddev-addon-mindsdb/actions/workflows/tests.yml/badge.svg)](https://github.com/meevagmbh/ddev-addon-mindsdb/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2024.svg)

# ddev-addon-mindsdb - MindsDB add-on for DDEV

This is a DDEV add-on for [MindsDB](https://mindsdb.com/), an open-source AI layer for existing databases that allows you to effortlessly develop, train, and deploy state-of-the-art machine learning models using SQL.

## Installation

For DDEV v1.23.5 or above run

```bash
ddev add-on get meevagmbh/ddev-addon-mindsdb
```

For earlier versions of DDEV run
```bash
ddev get meevagmbh/ddev-addon-mindsdb
```

Then restart your project
```bash
ddev restart
```

## Configuration

You can change the used Docker image by setting the `MINDSDB_DOCKER_TAG` to a supported Docker tag. The default is `lightwood`.

For all available tags see [MindsDB Docker Hub](https://hub.docker.com/r/mindsdb/mindsdb/tags).


## Usage

After installation, you can access the MindsDB web interface at `http://<project>.ddev.site:47334/`.

## License

[Apache 2.0](https://github.com/meevagmbh/ddev-addon-mindsdb/blob/main/LICENSE)

This project is not affiliated with or endorsed by [MindsDB](https://mindsdb.com/).

---

**Based on the original [ddev-addon-template](https://github.com/ddev/ddev-addon-template).**  
**Contributed and maintained by [meeva GmbH](https://meeva.de/).**

