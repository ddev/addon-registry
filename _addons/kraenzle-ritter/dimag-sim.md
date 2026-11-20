---
title: "kraenzle-ritter/dimag-sim"
github_url: "https://github.com/kraenzle-ritter/dimag-sim"
description: "Simulated DIMAG digital archive (KM-SOAP + SFTP) for local development and CI, checked against recordings of a real DIMAG. Docker compose and DDEV add-on."
user: "kraenzle-ritter"
repo: "dimag-sim"
repo_id: 1463760288
default_branch: "main"
tag_name: ""
ddev_version_constraint: ">= v1.24.0"
dependencies: []
type: "contrib"
created_at: "2026-11-19"
updated_at: "2026-11-19"
workflow_status: "unknown"
stars: 0
---

# dimag-sim

[![Tests](https://github.com/kraenzle-ritter/dimag-sim/actions/workflows/test.yml/badge.svg)](https://github.com/kraenzle-ritter/dimag-sim/actions/workflows/test.yml)
[![License](https://img.shields.io/github/license/kraenzle-ritter/dimag-sim)](LICENSE)
[![PHP Version](https://img.shields.io/badge/PHP-8.2%2B-blue)](https://www.php.net/)
[![DIMAG](https://img.shields.io/badge/DIMAG-3.5.0-0b5394)](docs/behaviour.md)
[![DDEV add-on](https://img.shields.io/badge/DDEV-add--on-02a8e2)](https://ddev.com/)

A stand-in for the [DIMAG](https://www.dimag.la-bw.de/) digital archive for
local development and CI: the KM-SOAP web service of the DIMAG core module
plus the SFTP drop it reads from and writes to. It accepts control files
and primary data, keeps an object tree (S/O/R/F, protocol objects), exports,
moves, changes status and deletes, and answers with the same messages a real
DIMAG 3.5.0 gives.

Its behaviour is pinned to **recordings of a real DIMAG**: a scenario suite
runs the operations an AIS bridge needs against a real installation, the
responses are anonymised into fixtures, and the same suite runs against
the simulator as a contract test. What could not be recorded is modelled and
marked as such — see [docs/behaviour.md](https://github.com/kraenzle-ritter/dimag-sim/blob/main/docs/behaviour.md).

dimag-sim was written for [Inge](https://k-r.ch), the bridge between the
archive information system Anton and DIMAG, but it does not depend on
either. It is not affiliated with the Landesarchiv Baden-Württemberg or the
DIMAG network.

## Quick start (docker compose)

```bash
docker compose up -d --build
```

| | |
|---|---|
| SOAP | `http://localhost:8080/soap/webservice_3_5_0.php` |
| SFTP | `localhost:2222`, user `inge`, password `inge`, drop directory `Import` |
| DIMAG account | `inge` / `inge`, may write below `sim-2` |
| Admin account | `admin` / `admin` (status 40/50, sees everything) |
| Control API | `http://localhost:8080/_sim/state`, `/_sim/reset`, `/_sim/seed`, `/_sim/faults` |

A client that works with DIMAG needs nothing but these settings. For Inge,
an archive record with:

| Inge field | Value |
|---|---|
| `soap_url` | `http://localhost:8080/soap/webservice_3_5_0.php` (in DDEV: `http://dimag:8080/…`) |
| `soap_username` / `soap_password` | `inge` / `inge` |
| `sftp_address` / `sftp_port` | `localhost` / `2222` (in DDEV: `dimag-sftp` / `22`) |
| `sftp_username` / `sftp_password` | `inge` / `inge` |
| `sftp_path` | `Import` |
| `parent_aid` | `sim-2` |
| `SOAP_ADMIN_USERNAME` / `SOAP_ADMIN_PASSWORD` (env) | `admin` / `admin` |

[examples/inge-smoke.php](https://github.com/kraenzle-ritter/dimag-sim/blob/main/examples/inge-smoke.php) drives Inge's own
`SoapController` and `SftpController` through ingest, finalize, inventory,
export, move, md5 check, the md5 error and deletion:

```bash
php examples/inge-smoke.php /path/to/inge
```

## DDEV add-on

```bash
ddev add-on get kraenzle-ritter/dimag-sim
ddev restart
ddev dimag status
```

This adds two services to the project: `dimag` (SOAP on port 8080) and
`dimag-sftp` (port 22), sharing the drop volume. Inside the project:
SOAP `http://dimag:8080/soap/webservice_3_5_0.php`, SFTP `dimag-sftp:22`.
From another DDEV project on the same machine: `ddev-<project>-dimag:8080`
and `ddev-<project>-dimag-sftp:22`. The control API is also at
`https://<project>.ddev.site:8581/_sim/state`.

```bash
ddev dimag state                     # dump the object tree
ddev dimag reset                     # empty tree, seeded again
ddev dimag seed .ddev/dimag-seed.json
ddev dimag faults '{"down":true}'    # error switches, see below
ddev dimag faults-off
```

Settings go into `.ddev/.env` (e.g. `DIMAG_SIM_USERS=…`) or
`.ddev/docker-compose.dimag_extra.yaml`. After updating the add-on, rebuild
the image: `ddev debug rebuild -s dimag`.

## Configuration

| Variable | Default | |
|---|---|---|
| `DIMAG_SIM_VERSION` | `3.5.0` | reported by `getStatus`; endpoints `webservice_3_0_0` … `webservice_3_<minor>_0` are served, newer ones answer 404 |
| `DIMAG_SIM_PREFIX` | `sim` | installation prefix; generated AIDs are `sim-<n>` |
| `DIMAG_SIM_USERS` | `inge:inge:Test User:sim-2;admin:admin:Test Admin:*:admin` | `user:password:display name:root1\|root2[:admin]`; roots that do not exist are created as S below `sim-1` |
| `DIMAG_SIM_SEED` | — | JSON file with more objects, loaded on first start |
| `DIMAG_SIM_TIMEZONE` | `Europe/Berlin` | DIMAG works in local time; `getObjectsChangedSince` compares with it |
| `DIMAG_SIM_ARCHIVE_NAME`, `DIMAG_SIM_URL` | `Simulated Archive`, `http://dimag` | shown by `getStatus` |
| `DIMAG_SIM_ADMIN_API` | on | `off` disables `/_sim/*` |
| `DIMAG_SFTP_USERS` (SFTP container) | `inge:inge:1000:Import` | `user:password:uid[:dir]`; uid 1000 is the SOAP container's |

### Seeding

`POST /_sim/seed` (or `DIMAG_SIM_SEED`) takes a list of objects; parents
must exist or come earlier in the list. An F gets its primary data from
`content`; its md5 is calculated.

```json
[
  {"aid": "demo-o", "typ": "O", "pid": "sim-2", "titel": "Delivery 1", "fields": {"entstehungs-zeitraum": "2026"}},
  {"aid": "demo-o-r1", "typ": "R", "pid": "demo-o", "titel": "Files"},
  {"aid": "demo-media-1", "typ": "F", "pid": "demo-o-r1", "titel": "a.pdf", "content": "…", "status": 20}
]
```

### Error switches

`POST /_sim/faults` with any of `down`, `delayMs`, `deny`, `failExport`,
`corrupt`, `failMoveAt` — see [docs/behaviour.md](https://github.com/kraenzle-ritter/dimag-sim/blob/main/docs/behaviour.md#error-switches).
`{}` switches all off.

## Tests

```bash
composer install
docker compose up -d --build
export DIMAG_SOAP_URL=http://localhost:8080/soap/webservice_3_5_0.php \
  DIMAG_USER=inge DIMAG_PASSWORD=inge DIMAG_ADMIN_USER=admin DIMAG_ADMIN_PASSWORD=admin \
  DIMAG_SFTP_HOST=localhost DIMAG_SFTP_PORT=2222 DIMAG_SFTP_USER=inge DIMAG_SFTP_PASSWORD=inge DIMAG_SFTP_ROOT=Import
php tests/contract.php     # recorded DIMAG vs. simulator, step by step
php tests/sim.php          # seeding, error switches, versions, accounts
php tests/scrub-check.php  # fixtures carry no installation data
```

Without containers, the simulator also runs on PHP's built-in server with
the drop as a local directory (`DIMAG_SFTP_HOST=file:///path/to/drop-parent`):

```bash
DIMAG_SIM_SFTP_DIR=/tmp/sim/drop/Import DIMAG_SIM_STATE=/tmp/sim/state/state.json \
  php -S 127.0.0.1:8080 dimag-sim/public/index.php
```

## Recording against a real DIMAG

The scenarios in [recorder/scenarios](https://github.com/kraenzle-ritter/dimag-sim/blob/main/recorder/scenarios) are what the
contract test replays. To record them on another DIMAG (a new version, or
your own installation):

1. Use a **test** system. `read-only` changes nothing. `lifecycle` creates
   objects with the AID prefix `dimagsim-<run>-` below a parent you name,
   provokes errors, and deletes everything again (status 30, then 50 with
   an admin account). DIMAG keeps an "updated" entry in the protocol of the
   parent. Agree on the parent with whoever runs the system.
2. Run the recorder where the DIMAG is reachable, with credentials in the
   environment (`DIMAG_SOAP_URL`, `DIMAG_USER`, `DIMAG_PASSWORD`,
   `DIMAG_ADMIN_USER`, `DIMAG_ADMIN_PASSWORD`, `DIMAG_SFTP_HOST`,
   `DIMAG_SFTP_PORT`, `DIMAG_SFTP_USER`, `DIMAG_SFTP_PASSWORD`,
   `DIMAG_SFTP_ROOT`, optional `DIMAG_CA_FILE`). On a server running Inge,
   `--from-inge=/path/to/inge --archive=<slug>` reads them from Inge
   instead (in-process; nothing is printed):

   ```bash
   php recorder/record.php --scenario=recorder/scenarios/read-only.php --out=read-only.raw.json
   php recorder/record.php --scenario=recorder/scenarios/lifecycle.php --root=<parent AID> --out=lifecycle.raw.json
   ```

   Credentials are masked in the raw file, but it still holds hostnames,
   paths and titles of the installation: **do not commit it** (`*.raw.json`
   is ignored).
3. Normalise into fixtures, naming words that must disappear (archive name,
   hostnames), and check:

   ```bash
   php recorder/normalize.php lifecycle.raw.json --scrub=word1,word2 > fixtures/3.5.0/lifecycle.json
   DIMAG_SCRUB_WORDS=word1,word2 php tests/scrub-check.php
   ```

4. Run `php tests/contract.php --version=<version>` against the simulator
   and adjust it until the steps match, or record the difference with a
   reason in the scenario's `known_differences`.

## Documentation of DIMAG

- KM-SOAP interface, control files, protocol files:
  [dimag-soap.md](https://gitlab.la-bw.de/pub/dimag-info/-/blob/master/dimag-soap.md)
  (the 3.5.0 text is in the history of that file, August 2023)
- Metadata: [dimag-metadata-3.5.0.md](https://gitlab.la-bw.de/pub/dimag-info/-/blob/master/dimag-metadata-3.5.0.md)

These documents are not part of this repository.

## Licence

MIT, see [LICENSE](https://github.com/kraenzle-ritter/dimag-sim/blob/main/LICENSE). © 2026 Kränzle & Ritter.
