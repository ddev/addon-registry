---
title: "aivazashvilitekla/ddev-agent-browser-display"
github_url: "https://github.com/aivazashvilitekla/ddev-agent-browser-display"
description: "DDEV add-on: a persistent Xvfb display + noVNC in the web container so agent-browser can run headed and be watched."
user: "aivazashvilitekla"
repo: "ddev-agent-browser-display"
repo_id: 1356752708
default_branch: "main"
tag_name: "v1.1.0"
ddev_version_constraint: ">= v1.24.0"
dependencies: []
type: "contrib"
created_at: "2026-09-04"
updated_at: "2026-09-05"
workflow_status: "success"
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

Ports 6080 and 6081 are published on the shared ddev-router, so they are
claimed for the whole machine and not just for this project. If anything else
on your host already listens on either one, ddev-router fails to start and
every DDEV project loses its URLs until you resolve the conflict. 6080 is the
conventional noVNC port, so check it first if you run other VNC tooling.

## Use with agent-browser

Inside the web container (`ddev ssh`, or prefix with `ddev exec`):

```bash
agent-browser open https://<project>.ddev.site --headed
```

Headless stays the default. `--headed` (or `AGENT_BROWSER_HEADED=true`) puts
Chrome on the display; with `DISPLAY` set, agent-browser uses this display
instead of starting a private one. The flip side is that a headed run fails
outright when the Xvfb daemon is down, rather than quietly falling back to a
private display. Check the daemons with
`ddev exec supervisorctl status 'webextradaemons:*'`.

Three things to know inside a DDEV web container:

- **Apple Silicon / ARM64 hosts:** `agent-browser install` cannot download
  Chrome for Testing, which publishes no Linux ARM64 build. Bake Debian's
  Chromium into the web image instead, so it survives every restart. Create
  `.ddev/config.chromium.yaml`:

  ```yaml
  webimage_extra_packages:
    - chromium
  ```

  then `ddev restart`. agent-browser finds `chromium` on `PATH` by itself, so
  no extra flag is needed; add `--executable-path /usr/bin/chromium` only if
  you want to pin it explicitly.
- **Your site's URL:** use the URL `ddev describe` prints, including any
  non-default router port, for example `https://<project>.ddev.site:8443`.
- **DDEV's local certificate:** Chromium inside the container does not trust
  it, so add `--ignore-https-errors`. agent-browser's `--ca-cert` is the
  stricter route, but it shells out to `certutil`, which the DDEV web image
  does not ship: add `libnss3-tools` to `webimage_extra_packages` first if you
  want to go that way.

## What it adds

- Debian packages `xvfb`, `x11vnc`, `novnc`, `websockify` in the web image.
- `DISPLAY=:99` in the web environment.
- Three supervised daemons (`ddev exec supervisorctl status 'webextradaemons:*'`):
  `Xvfb :99` at 1920x1080, `x11vnc` on the container's loopback at 5900, and
  `websockify` serving noVNC on 6080.
- Router ports 6080 (https) and 6081 (http) for the noVNC page. DDEV needs both
  declared or the router will not start.

## Native VNC viewer instead of noVNC

x11vnc listens on the container's loopback only, so a native client cannot
reach it as shipped. Using macOS Screen Sharing means opening it up and

1. Take ownership of the config so DDEV stops managing it: delete the
   `#ddev-generated` line at the top of
   `.ddev/config.agent-browser-display.yaml`.
2. In that same file, remove `-localhost` from the `x11vnc` command.
3. Add `.ddev/docker-compose.vnc.yaml`:

   ```yaml
   services:
     web:
       ports:
         - "127.0.0.1:5900:5900"
   ```

Then `ddev restart` and open `vnc://127.0.0.1:5900`. Keep the `127.0.0.1:`
prefix: without it Docker binds 0.0.0.0 and anyone on your network could drive
the display. Two consequences worth knowing. Step 2 also exposes 5900 to every
other container on the shared `ddev_default` network. And once the
`#ddev-generated` line is gone, `ddev add-on remove` and future upgrades no
longer manage that file.

## Security note

The X display has no authentication and x11vnc runs with `-nopw`. Inside the
container that is fine: anything running in the web container can already do
anything to your site. Outside it:

- **The noVNC page on router ports 6080 and 6081** has no password. Whoever
  opens it can watch AND drive the display, including the agent's logged-in
  browser. ddev-router binds to the Docker host IP, normally 127.0.0.1, unless
  you set `router_bind_all_interfaces: true`.
- **websockify listens on 0.0.0.0:6080 inside the container**, and every DDEV
  project's web container shares the `ddev_default` Docker network, so any
  container of any DDEV project on this machine can reach the noVNC endpoint
  too. Treat that network as trusted, or do not run this add-on beside
  containers you do not trust.
- **x11vnc's port 5900 is bound to the container's loopback**, so it is not
  reachable from other containers or from the host unless you follow the
  native-viewer steps above.

## Remove

```bash
ddev add-on remove agent-browser-display
ddev restart
```
