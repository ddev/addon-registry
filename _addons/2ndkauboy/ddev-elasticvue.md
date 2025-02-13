---
title: 2ndkauboy/ddev-elasticvue
github_url: https://github.com/2ndkauboy/ddev-elasticvue
description: "Elasticvue service for DDEV"
user: 2ndkauboy
repo: ddev-elasticvue
type: contrib
created_at: 2024-05-21
updated_at: 2024-10-24
stars: 0
---

[![tests](https://github.com/2ndkauboy/ddev-elasticvue/actions/workflows/tests.yml/badge.svg)](https://github.com/2ndkauboy/ddev-elasticvue/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2024.svg)

# DDEV Elasticvue

## What is ddev-elasticvue?

This repository allows you to quickly install Elasticvue into a [DDEV](https://ddev.readthedocs.io) project using just `ddev get 2ndkauboy/ddev-elasticvue`.

## Installation

For DDEV v1.23.5 or above run

```sh
ddev add-on get 2ndkauboy/ddev-elasticvue && ddev restart
```

For earlier versions of DDEV run

```sh
ddev get 2ndkauboy/ddev-elasticvue && ddev restart
```

You can then visit Elasticvue by running `ddev elasticvue` or visiting the URL shown in `ddev describe`.

## Configuration

The add-on assumes that the Elasticsearch service is named `elasticsearch`. If this is the case, you can import the predefined cluster on the welcome screen:

![Web-to-print settings menu](https://raw.githubusercontent.com/2ndkauboy/ddev-elasticvue/main/images/elasticvue-predefined-clusters-annotated.png)

1. Click on the `Predefined Clusters` button.
2. Click on the `Import 1 Cluster` button.

This only need to be done the first time you use this for a new project. If your clusters are named differently, you can use the "Add Elasticsearch Cluster" button. If you have other issues, refer to the official website or GitHub repository of Elasticvue listed below.


## Explanation

Elasticvue is a free and open-source gui for elasticsearch that you can use to manage the data in your cluster.

## Additional Resources

- [Elasticvue official website](https://elasticvue.com/).
- [Elasticvue GitHub repository](https://github.com/cars10/elasticvue).

**Elasticvue is maintained by [@cars10](https://github.com/cars10)**  
**DDEV Elasticvue is maintained by [@2ndkauboy](https://github.com/2ndkauboy)**

