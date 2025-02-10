---
title: thursdaybw/ddev-tightvnc
github_url: https://github.com/thursdaybw/ddev-tightvnc
description: "TightVNC server for DDEV"
user: thursdaybw
repo: ddev-tightvnc
categories:
  - community
created_at: 2024-06-23
updated_at: 2024-06-30
stars: 0
---


# ddev-tightvnc

[![tests](https://github.com/thursdaybw/ddev-tightvnc/actions/workflows/tests.yml/badge.svg)](https://github.com/thursdaybw/ddev-tightvnc/actions/workflows/tests.yml) ![project is maintained](https://img.shields.io/maintenance/yes/2024.svg)

## Table of Contents
- [What is ddev-tightvnc?](#what-is-ddev-tightvnc)
- [Installation](#installation)
- [Usage](#usage)
- [Testing](#testing)
- [Debugging](#debugging)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgements](#acknowledgements)

## What is ddev-tightvnc?
`ddev-tightvnc` is a DDEV add-on that provides VNC (Virtual Network Computing) capabilities for DDEV projects via tightvncserver. This allows developers to access a graphical desktop environment within their DDEV containers. The primary motivation for this add-on is to enable developers to visually monitor browsers, such as Chrome, being controlled by automation tools like Puppeteer. However, it is versatile enough to be used for any purpose requiring a graphical interface within a DDEV container.

### Key Features:
- **Monitor Headless Browsers**: Ideal for witnessing Chrome being controlled by Puppeteer.
- **Graphical Interface**: Access a minimal graphical desktop environment.
- **Simple Setup**: Easy to install and configure within any DDEV project.

## Installation
### Using `ddev get`
To install `ddev-tightvnc`, use the following commands:
```bash
ddev get thursdaybw/ddev-tightvnc
ddev restart
```

### Accessing the VNC Server
1. Use a VNC client to connect to `localhost:5901`; For example `vncviewer localhost:5901`
2. Enter 'password' which has been preset during configuration. 

### Accessing the VNC Server
1. Use a VNC client to connect to `localhost:5901`; For example `vncviewer localhost:5901`
2. Enter 'password' which has been preset during configuration. 

#### To connect using TigerVNC client without having to enter a password:
```
mkdir -p "~/.vnc"
echo "password" | vncpasswd -f > "${VNC_DIR}/passwd
vncviewer -passwd "~/.vnc/passwd" localhost:5901
```

### Example Usage
To visually monitor Puppeteer controlling Chrome, follow these steps:
1. Ensure your Puppeteer script is set up to run in a non-headless mode.
2. Start your DDEV project and ensure the VNC server is running.
3. Connect to the VNC server using a VNC client to watch the automation in action.

## Testing
### Running Tests Locally
To run tests locally, use:
```bash
bats ./tests/test.bats
```

### Running Tests in GitHub Actions
Tests are automatically run on push and pull request events.

## Debugging
If you encounter any issues, check the logs for the VNC server and your DDEV project. Logs can be found in the following locations:
- VNC server logs: `/var/log/supervisor/start-vnc.log` and `/var/log/supervisor/start-vnc_err.log`
- DDEV project logs: As configured in your DDEV setup

## License
This project is licensed under the MIT License - see the [LICENSE](https://github.com/thursdaybw/ddev-tightvnc/blob/main/LICENSE) file for details.

## Acknowledgements
- Thanks to the DDEV team for their support and contributions.

