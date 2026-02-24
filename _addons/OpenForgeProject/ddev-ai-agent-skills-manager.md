---
title: OpenForgeProject/ddev-ai-agent-skills-manager
github_url: https://github.com/OpenForgeProject/ddev-ai-agent-skills-manager
description: "Manage Ai Skills from skills.sh via .env.skills file"
user: OpenForgeProject
repo: ddev-ai-agent-skills-manager
repo_id: 1160063128
default_branch: main
tag_name: v2.0.0
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2026-02-17
updated_at: 2026-02-23
workflow_status: success
stars: 0
---

<div style="width: 100%; display: block; text-align: center;">
    <img src="image/ddev-skills-logo.png" height="200px" />
</div>

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/OpenForgeProject/ddev-skills/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/OpenForgeProject/ddev-skills/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/OpenForgeProject/ddev-skills)](https://github.com/OpenForgeProject/ddev-skills/commits)
[![release](https://img.shields.io/github/v/release/OpenForgeProject/ddev-skills)](https://github.com/OpenForgeProject/ddev-skills/releases/latest)

# DDEV Skills

## Overview

This add-on integrates Skills into your [DDEV](https://ddev.com/) project. It allows you to manage and install [skills.sh](https://skills.sh) via a simple configuration file. It reads the skills you want to install from a `.env.skills` file and uses `npx` to install them inside your DDEV web container. This add-on is ideal for developers who want to easily manage and update their skills directly within their project's environment without requiring Node.js on their host machine.

## Prerequisites

- **DDEV**: You need a running DDEV environment (v1.24.3 or higher).
- **Node.js**: Node.js and npm must be available in your DDEV web container (usually included by default).

## Installation

```bash
ddev add-on get OpenForgeProject/ddev-skills
ddev restart
```

## Configuration

1.  **Edit your configuration**:
    Edit `.ddev/.env.skills` to configure agents and add the skills you want to install.

    **Example `.ddev/.env.skills`:**

    ```env
    # Get available skills from https://skills.sh

    # Define your skill-agents here, separated by spaces
    # AGENTS="github-copilot claude-code ..."

    # Define your skills here Name="RepoURL"
    MySkill="https://github.com/username/my-skill-repo"
    AnotherSkill="https://github.com/username/another-skill"
    ```

    *   **AGENTS**: (Optional) Space-separated list of agents to install skills for. Defaults to `*` (all agents).
    *   **Skills**: format is `SkillName="GitRepoURL"`.

2.  **Version Control**:
    Make sure to commit the `.ddev/.env.skills` file to version control so your team has the same skills configuration.

## Usage

Run the following command within your DDEV project to install or update the skills defined in your configuration:

```bash
ddev skills
```

This command runs inside the web container and will:
- Read your `.ddev/.env.skills` file.
- Check if `npx` is available in the container.
- Install or update the specified skills using `npx skills` inside the container.

## Commands

| Command | Description |
| ------- | ----------- |
| `ddev skills` | Installs or updates skills based on `.ddev/.env.skills` configuration inside the web container. |

## Credits

**Contributed and maintained by [@OpenForgeProject](https://github.com/OpenForgeProject)**
