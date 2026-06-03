---
title: wazum/ddev-wiremock
github_url: https://github.com/wazum/ddev-wiremock
description: "WireMock mocking service add-on for DDEV"
user: wazum
repo: ddev-wiremock
repo_id: 1216113726
default_branch: main
tag_name: v0.2.0
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: contrib
created_at: 2026-04-20
updated_at: 2026-04-21
workflow_status: disabled
stars: 0
---

# ddev-wiremock

Run [WireMock](https://wiremock.org/) as a DDEV service, with host commands
that wrap the admin API.

[![tests](https://github.com/wazum/ddev-wiremock/actions/workflows/tests.yml/badge.svg)](https://github.com/wazum/ddev-wiremock/actions/workflows/tests.yml)
[![lint](https://github.com/wazum/ddev-wiremock/actions/workflows/lint.yml/badge.svg)](https://github.com/wazum/ddev-wiremock/actions/workflows/lint.yml)
[![release](https://img.shields.io/github/v/release/wazum/ddev-wiremock?sort=semver)](https://github.com/wazum/ddev-wiremock/releases)

## Install

```bash
ddev add-on get wazum/ddev-wiremock
ddev restart
```

WireMock is then reachable at `https://<project>.ddev.site:8443` with the
admin API under `/__admin`. Stubs live under `.ddev/wiremock/mappings/`.

## Getting started

Point your app's HTTP client at WireMock instead of the real upstream. In
the DDEV web container use `$DDEV_WIREMOCK_URL` (`http://wiremock:8080`);
from the host or a browser use `https://<project>.ddev.site:8443`.

Then add stubs one of three ways:

- **Scaffold quickly** - `ddev wiremock-add GET /users/42` writes a stub
  JSON file; `ddev wiremock-reload` makes it live without a restart.
- **Author by hand** - drop JSON files into `.ddev/wiremock/mappings/`.
  See `sample.json` there for the shape, then run `ddev wiremock-reload`.
- **Record from a live upstream** - see [Recording](#recording) below.

Every request your app makes to the WireMock URL is matched against a
stub. On a match, WireMock serves the stub response. On a miss, WireMock
answers with 404 - unless you are recording, in which case it proxies the
request to the configured upstream and persists the result as a new stub.

## Commands

| Command | What it does |
|---|---|
| `ddev wiremock-add <METHOD> <PATH>` | Scaffold a new stub JSON file. `--status N`, `--body JSON`, `--content-type TYPE`, `--force`. |
| `ddev wiremock-reload` | Re-read stub files from disk without restarting. Keeps the request journal. |
| `ddev wiremock-mappings` | List active stubs. `--id <uuid>` fetches one, `--json` dumps full JSON. |
| `ddev wiremock-requests` | Show the request journal. `--limit N`, `--unmatched`, `--json`. |
| `ddev wiremock-logs` | Tail the WireMock container logs. Passes flags through to `ddev logs`. |
| `ddev wiremock-reset` | Wipe runtime stubs and the request journal. File-backed stubs reload automatically. |
| `ddev wiremock-record <url>` | Start recording against an upstream URL. |
| `ddev wiremock-record-stop` | Stop recording; writes stubs into `mappings/`. |
| `ddev wiremock-snapshot` | Convert the current journal into stubs (no upstream needed). |

All commands take `--help`.

## Stubs and the request journal

**Stubs** are JSON files that tell WireMock how to respond to requests.
Each file describes one match (method, URL pattern, headers, body) and
the response to return. They live in `.ddev/wiremock/mappings/`, with
response bodies in `.ddev/wiremock/__files/`. Everything there is
committed - stubs are shared team state.

Stub syntax: <https://wiremock.org/docs/stubbing/>.

**The request journal** is an in-memory log of every HTTP call WireMock
has received during the current run: method, URL, headers, body, which
stub matched, and the response returned. `ddev wiremock-requests` reads
it; `ddev wiremock-snapshot` persists it as stubs. The journal clears on
`ddev wiremock-reset` and on container restart.

## Recording

WireMock can capture stubs from live upstream responses. Point it at the
real API, make the calls you want to capture, and WireMock writes a stub
for each response.

1. Start recording:

   ```bash
   ddev wiremock-record https://api.example.com
   ```

2. Issue the requests you want to record. Anything that hits WireMock and
   doesn't match an existing stub is forwarded to the upstream, and
   WireMock writes a new stub from the response. Pick whichever fits:

   ```bash
   # by hand from the host
   curl -k "https://<project>.ddev.site:8443/users/42"

   # from inside the web container (DDEV_WIREMOCK_URL is injected there)
   ddev exec "curl $DDEV_WIREMOCK_URL/users/42"

   # or run your app's integration tests - as long as its HTTP client is
   # configured to call WireMock's URL instead of the real API, every
   # request flows through.
   ```

3. Stop recording and review the new stub files:

   ```bash
   ddev wiremock-record-stop
   git diff .ddev/wiremock/mappings/
   ```

`ddev wiremock-snapshot` is the after-the-fact variant: if WireMock has
already been serving requests (for example, against a proxy stub you set
up manually), it dumps the current request journal into stub files in one
call - no upstream URL needed.

## Environment and configuration

Injected into the DDEV web container:

| Variable | Value |
|---|---|
| `DDEV_WIREMOCK_URL` | `http://wiremock:8080` |
| `DDEV_WIREMOCK_ADMIN_URL` | `http://wiremock:8080/__admin` |

Editable in `.ddev/.env.wiremock`:

| Key | Default | Effect |
|---|---|---|
| `WIREMOCK_TAG` | `3x` | WireMock image tag. |
| `WIREMOCK_HTTP_PORT` | `8080` | Router port for HTTP. |
| `WIREMOCK_HTTPS_PORT` | `8443` | Router port for HTTPS. |

## Remove

```bash
ddev add-on remove wiremock
```

User-authored files under `.ddev/wiremock/` are preserved. Delete that
directory manually if you no longer need the stubs.

## License

Apache-2.0. Copyright 2026 Wolfgang Klinger. WireMock is a separate project;
see <https://wiremock.org/> for its own license and credits.
