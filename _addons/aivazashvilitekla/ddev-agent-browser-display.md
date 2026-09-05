---
title: "aivazashvilitekla/ddev-agent-browser-display"
github_url: "https://github.com/aivazashvilitekla/ddev-agent-browser-display"
description: "DDEV add-on: a persistent Xvfb display + noVNC in the web container so agent-browser can run headed and be watched."
user: "aivazashvilitekla"
repo: "ddev-agent-browser-display"
repo_id: 1356752708
default_branch: "main"
tag_name: "v1.0.1"
ddev_version_constraint: ">= v1.24.0"
dependencies: []
type: "contrib"
created_at: "2026-09-04"
updated_at: "2026-09-04"
workflow_status: "unknown"
stars: 0
---

# ddev-agent-browser-display

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on-blue)](https://addons.ddev.com)

A persistent virtual display inside the DDEV **web** container, shared over
VNC and viewable in a browser tab. Built for
[agent-browser](https://github.com/vercel-labs/agent-browser): run it with
`--headed` and watch, live, what your coding agent does in the browser.

## Install

```bash
ddev add-on get aivazashvilitekla/ddev-agent-browser-display
ddev restart
```

Then open `https://<project>.ddev.site:6080/vnc.html` - `<project>` is your
DDEV project name, so for a project called `mysite` that is
`https://mysite.ddev.site:6080/vnc.html` - and click **Connect**. An empty
grey desktop means the display is live.

## Use with agent-browser

Inside the web container (`ddev ssh`, or prefix with `ddev exec`):

```bash
agent-browser open https://<project>.ddev.site --headed
```

Headless stays the default. `--headed` (or `AGENT_BROWSER_HEADED=true`) puts
Chrome on the display; with `DISPLAY` set, agent-browser uses this display
instead of starting a private one.

Three things to know inside a DDEV web container:

- **Apple Silicon / ARM64 hosts:** `agent-browser install` cannot download
  Chrome for Testing (no Linux ARM64 builds). Install Debian's Chromium and
  point agent-browser at it:
  `sudo apt-get install -y chromium`, then add
  `--executable-path /usr/bin/chromium` to your commands.
- **Your site's URL:** use the URL `ddev describe` prints, including any
  non-default router port (e.g. `https://<project>.ddev.site:8443`).
- **DDEV's local certificate:** Chromium inside the container does not trust
  it. Add `--ignore-https-errors`, or trust the CA properly with
  `--ca-cert <path>` (see agent-browser's README).

## What it adds

- Debian packages `xvfb`, `x11vnc`, `novnc`, `websockify` in the web image.
- `DISPLAY=:99` in the web environment.
- Three supervised daemons (`ddev exec supervisorctl status 'webextradaemons:*'`):
  `Xvfb :99` at 1920x1080, `x11vnc` on 5900 (container-internal), and
  `websockify` serving noVNC on 6080.
- Router port 6080 (https) / 6081 (http) for the noVNC page.

## Native VNC viewer instead of noVNC

The raw VNC port is not published by default (a fixed host port collides
across projects, and x11vnc runs with no password). To use macOS Screen
Sharing, add `.ddev/docker-compose.vnc.yaml`:

```yaml
services:
  web:
    ports:
      - "127.0.0.1:5900:5900"
```

then `ddev restart` and open `vnc://127.0.0.1:5900`. The `127.0.0.1:` prefix
keeps the port off your other network interfaces; without it Docker binds
0.0.0.0 and anyone on your network could drive the display.

## Security note

The X display has no authentication and x11vnc runs with `-nopw`. Inside the
container that is fine: anything running in the web container can already
do anything to your site. Two things are reachable from outside it:

- **The noVNC page on router port 6080** has no password. Whoever can open
  it can watch AND drive the display, including the agent's logged-in
  browser. ddev-router binds to 127.0.0.1 unless you set
  `router_bind_all_interfaces: true`, so by default that is only you.
- **Port 5900** is not published unless you add the snippet above; if you do,
  keep the `127.0.0.1:` prefix.

Also note `DISPLAY=:99` is set for every process in the web container. If the
Xvfb daemon is not running (check `ddev exec supervisorctl status
'webextradaemons:*'`), any X client, agent-browser included, fails with
"cannot open display :99" rather than starting a private display.

## Remove

```bash
ddev add-on remove agent-browser-display
ddev restart
```
