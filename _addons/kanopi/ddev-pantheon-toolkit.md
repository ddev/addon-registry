---
title: kanopi/ddev-pantheon-toolkit
github_url: https://github.com/kanopi/ddev-pantheon-toolkit
description: "A DDEV add-on providing commands to manage supporting organizations, sites, and people on Pantheon."
user: kanopi
repo: ddev-pantheon-toolkit
repo_id: 1135788933
default_branch: main
tag_name: 1.0.4
ddev_version_constraint: ">= v1.23.0"
dependencies: []
type: contrib
created_at: 2026-01-16
updated_at: 2026-01-16
workflow_status: unknown
stars: 0
---

# DDEV Pantheon Toolkit

A DDEV add-on providing commands to manage supporting organizations across Pantheon sites based on their upstream.

## Quick Start

```bash
# Install the DDEV add-on
ddev get kanopi/ddev-pantheon-toolkit

# Restart DDEV to activate the add-on
ddev restart

# Authenticate with Terminus (one-time setup)
ddev ssh
terminus auth:login
exit
```

## Available Commands

- `ddev find-upstreams` - Discover upstreams used across your organization
- `ddev find-people` - List team members across all sites
- `ddev add-org` - Add supporting organizations to sites (with filtering and dry-run options)
- `ddev site-inventory` - Generate comprehensive site inventory reports

## Documentation

For complete documentation, visit: **[https://kanopi.github.io/ddev-pantheon-toolkit](https://kanopi.github.io/ddev-pantheon-toolkit)**

- [Installation Guide](https://kanopi.github.io/ddev-pantheon-toolkit/installation)
- [Usage Guide](https://kanopi.github.io/ddev-pantheon-toolkit/usage)
- [Command Reference](https://kanopi.github.io/ddev-pantheon-toolkit/commands)
- [Example Workflows](https://kanopi.github.io/ddev-pantheon-toolkit/workflows)
- [Troubleshooting](https://kanopi.github.io/ddev-pantheon-toolkit/troubleshooting)

## Prerequisites

- **DDEV** - [Install DDEV](https://ddev.readthedocs.io/en/stable/)
- **Pantheon Account** - You must have access to a Pantheon organization
- **Terminus Authentication** - See [Installation Guide](https://kanopi.github.io/ddev-pantheon-toolkit/installation)

## Support

For issues with this DDEV add-on: [Open an issue on GitHub](https://github.com/kanopi/ddev-pantheon-toolkit/issues)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the GNU General Public License v2 - see the LICENSE file for details.

---

Originally Contributed by Kanopi Studios
