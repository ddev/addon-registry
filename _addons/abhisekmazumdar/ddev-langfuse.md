---
title: "abhisekmazumdar/ddev-langfuse"
github_url: "https://github.com/abhisekmazumdar/ddev-langfuse"
description: "DDEV add-on for a local Langfuse observability stack"
user: "abhisekmazumdar"
repo: "ddev-langfuse"
repo_id: 1320349068
default_branch: "main"
tag_name: "v1.0.0"
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: "contrib"
created_at: "2026-08-02"
updated_at: "2026-08-02"
workflow_status: "success"
stars: 0
---

# ddev-langfuse

A [DDEV](https://ddev.com) add-on that spins up a local
[Langfuse](https://langfuse.com) instance — LLM observability, tracing,
prompt management and evals — on your project's own Docker network, with
zero manual setup.

Built for the Drupal AI ecosystem
([`drupal/ai`](https://www.drupal.org/project/ai),
[`drupal/ai_agents`](https://www.drupal.org/project/ai_agents),
[`drupal/ai_search`](https://www.drupal.org/project/ai_search),
[`drupal/ai_answers`](https://www.drupal.org/project/ai_answers),
[`drupal/langfuse`](https://www.drupal.org/project/langfuse)) but not
Drupal-specific — anything that can send OpenTelemetry-ish LLM traces to a
Langfuse-compatible endpoint can use it.

## What you get

Running `ddev add-on get` on this repo adds Langfuse's full self-hosted
stack (image tag `4`, matching [langfuse/langfuse's own
`docker-compose.yml`](https://github.com/langfuse/langfuse)) as extra
services inside your existing DDEV project:

| Service | Image | Purpose |
|---|---|---|
| `langfuse-web` | `langfuse/langfuse:4` | UI + API |
| `langfuse-worker` | `langfuse/langfuse-worker:4` | async ingestion / processing |
| `langfuse-postgres` | `postgres:17` | relational data |
| `langfuse-clickhouse` | `clickhouse/clickhouse-server:25.12` | trace/analytics store |
| `langfuse-redis` | `redis:7` | queues/cache |
| `langfuse-minio` | `minio/minio` | S3-compatible blob storage |

No lightweight single-container alternative exists for Langfuse's current
architecture — all six services are required.

On first install, the add-on:

1. **Generates every required secret** (`NEXTAUTH_SECRET`, `SALT`,
   `ENCRYPTION_KEY`, and the Postgres/ClickHouse/MinIO/Redis passwords) with
   `openssl rand`, into `.ddev/.env.langfuse`. You never hand-edit a
   `CHANGEME` placeholder.
2. **Headlessly bootstraps** an organization, project, admin user, and a
   working API keypair using Langfuse's `LANGFUSE_INIT_*` environment
   variables — no signup wizard to click through.
3. **Prints the admin login and API keys** at the end of `ddev add-on get`,
   and again any time via `ddev langfuse-credentials`.

## Installation

```bash
# Latest tagged release (recommended -- pinned, reproducible):
ddev add-on get abhisekmazumdar/ddev-langfuse
ddev restart

# A specific release:
ddev add-on get abhisekmazumdar/ddev-langfuse --version v1.0.0
ddev restart

# Latest commit on the default branch, if you want unreleased changes:
ddev add-on get abhisekmazumdar/ddev-langfuse --default-branch
ddev restart

# From a local checkout instead of GitHub:
ddev add-on get /absolute/path/to/ddev-langfuse
ddev restart
```

See [releases](https://github.com/abhisekmazumdar/ddev-langfuse/releases)
for available versions.

Commit the `.ddev` directory changes (`docker-compose.langfuse.yaml`,
`.env.langfuse.example`, `commands/host/langfuse*`, and the updated
`.ddev/.gitignore`) to your project's git repo as usual — **but not**
`.ddev/.env.langfuse`, which the add-on gitignores automatically because it
holds real secrets and API keys.

After install, watch the printed credentials block, or run:

```bash
ddev langfuse-credentials
```

## Opening the UI

```bash
ddev launch :3000     # or:
ddev langfuse
```

This works because `langfuse-web` is wired into `ddev-router` the same way
official add-ons like `ddev-solr` expose their admin UIs — via
`HTTP_EXPOSE`/`HTTPS_EXPOSE`/`VIRTUAL_HOST` on the service, not a raw
published host port. If you run several DDEV projects with this add-on
installed **at the same time**, they'll all claim host ports 3000/3443 for
their router-exposed UI, so only start one at a time, or override the ports
by setting `HTTP_EXPOSE`/`HTTPS_EXPOSE` in your project's
`.ddev/docker-compose.langfuse.yaml` override.

## Reaching Langfuse from your Drupal container (or any other service)

This is the actual point of the add-on. Because
`docker-compose.langfuse.yaml` lives under `.ddev/`, DDEV starts these
services on the **same Docker network as the rest of your project**. Your
`web` container (or any other DDEV service) reaches Langfuse by plain
service hostname:

```
http://langfuse-web:3000
```

No `host.docker.internal`, no `docker network connect`, no separately
launched compose project to keep track of. This is the same mechanism
`ddev-redis` and `ddev-solr` use to make `redis`/`solr` resolvable from
`web` — placing the compose file in `.ddev/` is the whole trick, DDEV does
the rest.

## Configuring the `drupal/langfuse` module

The [`drupal/langfuse`](https://www.drupal.org/project/langfuse) module's
config schema (`langfuse.settings.yml`) has:

- `langfuse_url`
- `auth_method` (`none` | `bearer` | `basic` | `key_pair`)
- `bearer_token`
- `basic_auth.username` / `basic_auth.password`
- `key_pair.public_key` / `key_pair.secret_key`

Point it at this add-on's stack with:

| Field | Value |
|---|---|
| `langfuse_url` | `http://langfuse-web:3000` |
| `auth_method` | `key_pair` |
| `key_pair.public_key` | printed by `ddev langfuse-credentials` (`LANGFUSE_INIT_PROJECT_PUBLIC_KEY`) |
| `key_pair.secret_key` | printed by `ddev langfuse-credentials` (`LANGFUSE_INIT_PROJECT_SECRET_KEY`) |

Because the project/org/keypair are bootstrapped headlessly at install
time, these keys work immediately — no need to log into the UI first to
generate them.

## Ports

| Service | Host-published? | Why |
|---|---|---|
| `langfuse-web` | Yes, via `ddev-router` (3000/3443) | so you can browse the dashboard from the host |
| everything else | No — `expose` only, same-network hostname access | avoids host port collisions across multiple DDEV projects running this add-on; nothing outside the project network needs to reach Postgres/ClickHouse/Redis/MinIO directly |

## Data persistence

Each backing service gets its own named Docker volume
(`langfuse_postgres`, `langfuse_clickhouse_data`, `langfuse_clickhouse_logs`,
`langfuse_redis`, `langfuse_minio`), scoped to the DDEV project. Data
survives `ddev stop`/`ddev start` and `ddev restart`. `ddev delete` removes
project volumes along with everything else, same as any other DDEV service.

## Secrets and `.env.langfuse`

`install.yaml`'s `pre_install_actions` generate `.ddev/.env.langfuse` once,
the first time you run `ddev add-on get`, and leave it alone on every
re-run/update (so reinstalling the add-on doesn't invalidate your existing
Langfuse data by rotating `ENCRYPTION_KEY` out from under it). The file is:

- Never part of `project_files` (so `ddev add-on get` never silently
  overwrites it).
- Automatically appended to `.ddev/.gitignore` so it can't be committed by
  accident.
- Documented, not duplicated, by the committed `.env.langfuse.example` —
  which shows the shape of the file without real values.

Want fresh secrets and a fresh admin login? Delete `.ddev/.env.langfuse`
and reinstall the add-on. Note this effectively invalidates old Langfuse
data (`ENCRYPTION_KEY` won't decrypt anything encrypted under the old key),
so treat it as a reset, not a rotation.

## Uninstalling

```bash
ddev add-on remove langfuse
```

This removes the compose file and custom commands. It deliberately leaves
`.ddev/.env.langfuse` and the Docker volumes in place, in case you reinstall
later and want your data back. To fully wipe everything:

```bash
rm -f .ddev/.env.langfuse
docker volume rm $(docker volume ls -q --filter label=com.docker.compose.project=ddev-$(ddev describe -j | jq -r .raw.name) --filter name=langfuse)
```

## Testing

```bash
bats ./tests/test.bats
```

Follows the standard `ddev/ddev-addon-template` pattern: installs the
add-on into a throwaway DDEV project, starts it, and checks that
`.env.langfuse` has real (non-placeholder) generated secrets, that
`langfuse-web`'s health endpoint responds, and that the custom
`langfuse-credentials` command works. The "install from release" test is
tagged `release` and will only pass once this repo has an actual GitHub
release to install from.

## Design notes

- **Why service-name networking instead of `host.docker.internal`**: a
  one-off developer setup can get away with reaching a separately-launched
  Langfuse compose project through the host, but it's fragile and doesn't
  generalize. Living under `.ddev/docker-compose.langfuse.yaml` is the
  actual DDEV addon mechanism for joining a project's network — it isn't a
  special case at all, it's just where DDEV looks.
- **Why namespaced service names** (`langfuse-postgres`, `langfuse-redis`,
  etc., not bare `postgres`/`redis`): a Docker Compose service name is also
  its DNS hostname on the project's network. Bare names would collide with
  other addons a user might also install (e.g. `ddev-redis` already claims
  `redis`).
- **Why `.env.langfuse` isn't in `project_files`**: files in `project_files`
  are treated as addon-owned and get refreshed on every `ddev add-on get`
  (that's what the `#ddev-generated` marker convention is for). Secrets
  need the opposite behavior — generate once, then leave alone — so they're
  produced by a `pre_install_actions` script with an existence check
  instead.

## Verification

`bats ./tests/test.bats --filter-tags '!release'` has been run end-to-end
against a real DDEV/Docker environment: `ddev add-on get` from a local
checkout, `ddev restart`, all six services reaching healthy, and
`ddev langfuse-credentials` printing real (non-placeholder) secrets. One
real bug surfaced and was fixed by this process: `langfuse-web` and
`langfuse-worker` both stayed `unhealthy` indefinitely even once fully
booted, because
1. Docker auto-populates the container's `HOSTNAME` env var from the
   compose `hostname:` field, and Next.js's standalone server binds to
   `$HOSTNAME` if set — so it ended up listening only on its own container
   IP instead of `0.0.0.0`, and
2. even after fixing that, the healthchecks' `wget http://localhost:PORT`
   failed too, because these images' `/etc/hosts` resolves `localhost` to
   `::1` only while the server listens IPv4-only.

Both are fixed in `docker-compose.langfuse.yaml` (`HOSTNAME=0.0.0.0` env
var override, healthchecks target `127.0.0.1` instead of `localhost`) — see
the comments next to those lines if this needs revisiting for a future
Langfuse image version.

## License

MIT — see [LICENSE](https://github.com/abhisekmazumdar/ddev-langfuse/blob/main/LICENSE).
