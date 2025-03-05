---
title: stinis87/ddev-ollama
github_url: https://github.com/stinis87/ddev-ollama
description: "Run Ollama in DDEV"
user: stinis87
repo: ddev-ollama
repo_id: 933817805
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2025-02-16
updated_at: 2025-03-04
stars: 1
---

## Getting started

1. Install the Ollama ddev add-on and restart:

    ```shell
    ddev add-on get stinis87/ddev-ollama && ddev restart
    ```
2. Access Ollama service:

    ```shell
    ddev ssh -s ollama
    ```
3. Install and run a model. Example showcases llama but you can find the different models on https://ollama.com/search:

    ```shell
    ollama run llama3.2:3b
    ```
Ollama is now running on hostname: http://ollama, port: 11434
