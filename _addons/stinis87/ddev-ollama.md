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
updated_at: 2025-03-25
workflow_status: failure
stars: 3
---

# DDEV Ollama Add-on

The **DDEV Ollama Add-on** allows you to integrate [Ollama](https://ollama.com/), a powerful AI model runner, into your [DDEV](https://ddev.readthedocs.io/en/stable/) development environment. This add-on makes it easy to deploy and run language models locally, without the need for cloud-based services.  

With this integration, developers working with DDEV can leverage AI models like **Llama 3**, **Mistral**, and others directly from their local machine. This is particularly useful for web developers who want to use large language models in their projects.

## Features

- **Seamless integration**: Easily install and configure Ollama within your existing DDEV setup.
- **Local AI model execution**: Run models like **Llama3**, **Mistral**, and more without relying on external services.
  
---

## Prerequisites

Before installing the add-on, ensure you have:

- [DDEV installed](https://ddev.readthedocs.io/en/stable/)
- A working Docker setup (since DDEV relies on Docker)
- Sufficient system resources for running AI models (depending on the complexity of the model)

---

## Installation

You can install the **DDEV Ollama Add-on** with a single command:

```bash
ddev get stinis87/ddev-ollama && ddev restart
```
---

## Usage

**Access the Ollama service:**

ddev ssh -s ollama

**Install and run a model:**

For example, to run the llama3.2:3b model:

ollama run llama3.2:3b

Ollama is now running on hostname: http://ollama, port: 11434
