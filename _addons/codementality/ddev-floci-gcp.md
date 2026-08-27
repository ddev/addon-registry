---
title: "codementality/ddev-floci-gcp"
github_url: "https://github.com/codementality/ddev-floci-gcp"
description: "DDEV Floci emulator for GCP services."
user: "codementality"
repo: "ddev-floci-gcp"
repo_id: 1347148032
default_branch: "main"
tag_name: "0.5.1"
ddev_version_constraint: ">= v1.24.10"
dependencies: ["codementality/ddev-floci-ui"]
type: "contrib"
created_at: "2026-08-26"
updated_at: "2026-08-26"
workflow_status: "unknown"
stars: 0
---

# ddev-floci-gcp <!-- omit in toc -->

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/codementality/ddev-floci-gcp/actions/workflows/tests.yml/badge.svg)](https://github.com/codementality/ddev-floci-gcp/actions/workflows/tests.yml)
[![last commit](https://img.shields.io/github/last-commit/codementality/ddev-floci-gcp)](https://github.com/codementality/ddev-floci-gcp/commits)
[![release](https://img.shields.io/github/v/release/codementality/ddev-floci-gcp)](https://github.com/codementality/ddev-floci-gcp/releases/latest)

A [DDEV](https://ddev.com) add-on that runs [floci-gcp](https://floci.io/gcp/) —
a free, open-source local Google Cloud emulator — alongside your project, and
points your web container at it.

Your application talks to `floci-gcp:4588` instead of Google. GCS buckets,
Pub/Sub topics, Firestore and Datastore documents, Secret Manager secrets,
BigQuery datasets, Cloud Tasks queues and about twenty other services exist
locally, with no GCP project, no service-account key, and no billing account.

- [Why this add-on](#why-this-add-on)
- [Running alongside the AWS and Azure emulators](#running-alongside-the-aws-and-azure-emulators)
- [Installation](#installation)
- [Using it](#using-it)
- [Which services configure themselves, and which need a nudge](#which-services-configure-themselves-and-which-need-a-nudge)
- [Reaching Floci from the host](#reaching-floci-from-the-host)
- [The two default buckets](#the-two-default-buckets)
- [Creating resources on every start](#creating-resources-on-every-start)
- [Persistence](#persistence)
- [Docker-backed services and the Docker socket](#docker-backed-services-and-the-docker-socket)
- [Projects are the isolation boundary](#projects-are-the-isolation-boundary)
- [Migrating from the gcloud emulators](#migrating-from-the-gcloud-emulators)
- [Configuration](#configuration)
- [The `ddev floci-gcp` command](#the-ddev-floci-gcp-command)
- [Removing the add-on](#removing-the-add-on)
- [Licensing](#licensing)
- [Credits](#credits)

## Why this add-on

Google ships a separate emulator per service — `gcloud beta emulators pubsub`,
`firestore`, `datastore`, `bigtable`, `spanner` — each its own process on its
own port, each needing the Cloud SDK installed, and most of GCP has no emulator
at all. floci-gcp replaces the lot with one ~90 MB container on one port, and
covers GCS, Secret Manager, KMS, BigQuery, Cloud Tasks, Scheduler, Eventarc,
IAM, Cloud Run, Cloud SQL and more besides. It is MIT licensed with no auth
token and no feature gates.

Running it as a DDEV add-on rather than as a separate `docker compose` stack
buys you three things that are fiddly to arrange by hand:

- **The web container is configured for you.** `PUBSUB_EMULATOR_HOST`,
  `FIRESTORE_EMULATOR_HOST`, `STORAGE_EMULATOR_HOST` and the rest are set in the
  web container's environment, and a GCP SDK that sees them talks to the
  emulator *and* skips credential discovery entirely — so no service-account
  JSON, no `gcloud auth`, no application change.
- **Returned URLs resolve.** GCS media links, Cloud Run service URLs and LRO
  self links name the `floci-gcp` container rather than `localhost`, which is
  what makes them usable from your PHP/Node/Python code. Upstream's default is
  `localhost`, and getting this wrong is the single most common way a local
  cloud emulator appears to work and then doesn't.
- **gRPC works from both sides.** Pub/Sub, Firestore and Secret Manager speak
  gRPC and nothing else. In-project that is plain HTTP/2 over DDEV's network;
  from your host it needs a direct port, because a TLS-terminating router is not
  something a `*_EMULATOR_HOST` client knows how to talk to. Both are set up
  here.

## Running alongside the AWS and Azure emulators

Everything this add-on installs is named for the cloud it emulates rather than
for Floci: the service and its hostname are `floci-gcp`, the container is
`ddev-<project>-floci-gcp`, the state volume is `floci-gcp`, and the command is
`ddev floci-gcp`. Nothing here claims the bare `floci` name, so the equivalent
AWS and Azure add-ons can be installed in the same project without overwriting
this one's files or shadowing its command.

Your application can then tell them apart by endpoint:
`http://floci-aws:4566`, `http://floci-az:4577`, `http://floci-gcp:4588`.

## Installation

```bash
ddev add-on get codementality/ddev-floci-gcp
ddev restart
```

Then check it:

```bash
ddev floci-gcp services
ddev floci-gcp python -c 'from google.cloud import storage
c = storage.Client(project="floci-local")
c.create_bucket("uploads")
print([b.name for b in c.list_buckets()])'
```

## Using it

Nothing in your application needs to know it is talking to an emulator. The web
container already has the environment set:

| Variable | Value |
|---|---|
| `PUBSUB_EMULATOR_HOST`, `FIRESTORE_EMULATOR_HOST`, `DATASTORE_EMULATOR_HOST` | `floci-gcp:4588` |
| `SECRET_MANAGER_EMULATOR_HOST`, `FIREBASE_AUTH_EMULATOR_HOST` | `floci-gcp:4588` |
| `STORAGE_EMULATOR_HOST`, `FLOCI_GCP_ENDPOINT` | `http://floci-gcp:4588` |
| `GOOGLE_CLOUD_PROJECT`, `GCLOUD_PROJECT`, `CLOUDSDK_CORE_PROJECT` | `floci-local` |
| `CLOUDSDK_AUTH_DISABLE_CREDENTIALS` | `true` |

Note the two shapes, which are upstream's convention rather than a typo:
`STORAGE_EMULATOR_HOST` is a full URL, the gRPC ones are bare `host:port`.

**PHP** (`google/cloud-storage`, `google/cloud-pubsub` — the emulator variables
are read by the library itself):

```php
$storage = new Google\Cloud\Storage\StorageClient([
  'projectId' => getenv('GOOGLE_CLOUD_PROJECT'),
]);
$storage->createBucket('uploads');
$storage->bucket('uploads')->upload('hi', ['name' => 'a.txt']);

$pubsub = new Google\Cloud\PubSub\PubSubClient([
  'projectId' => getenv('GOOGLE_CLOUD_PROJECT'),
]);
$pubsub->createTopic('jobs');
```

For Drupal with the `google_cloud` or a `flysystem`-backed module, point the
bucket at `uploads` and leave credentials empty — the emulator accepts requests
with no credential at all.

**Python / Node / Java**: `google-cloud-*`, `@google-cloud/*` and the Java
libraries all honour the same variables, with the one exception below.

## Which services configure themselves, and which need a nudge

Worth knowing before you debug something that looks broken:

| Service | Configures itself from the environment? |
|---|---|
| Cloud Storage | Yes — `STORAGE_EMULATOR_HOST` |
| Pub/Sub | Yes — `PUBSUB_EMULATOR_HOST` |
| Firestore | Yes — `FIRESTORE_EMULATOR_HOST` |
| Datastore | Yes — `DATASTORE_EMULATOR_HOST` |
| Firebase Auth | Yes — `FIREBASE_AUTH_EMULATOR_HOST` |
| Secret Manager | **No** in the Python client — see below |
| BigQuery, KMS, Cloud Tasks, Scheduler, Eventarc, Logging, Monitoring, IAM, Cloud Run, Cloud SQL | No — pass an endpoint explicitly |

The services in the last two rows have no `*_EMULATOR_HOST` convention their
client libraries agree on, so the endpoint goes in the client constructor.
`FLOCI_GCP_ENDPOINT` is in the web container's environment for exactly this —
nothing reads it automatically, it is there for your code to read:

```python
import grpc
from google.cloud import secretmanager
from google.cloud.secretmanager_v1.services.secret_manager_service.transports \
    import SecretManagerServiceGrpcTransport

client = secretmanager.SecretManagerServiceClient(
    transport=SecretManagerServiceGrpcTransport(
        channel=grpc.insecure_channel("floci-gcp:4588")))
```

Secret Manager is called out because upstream documents
`SECRET_MANAGER_EMULATOR_HOST` and the Python client ignores it — it dials
`secretmanager.googleapis.com` and then fails on missing credentials. An
explicit insecure channel is the fix. Java and Node take the equivalent
`NoCredentialsProvider` / `apiEndpoint` options.

## Reaching Floci from the host

There are three endpoints, and which one is right depends on where the caller
runs and whether it speaks gRPC:

```bash
ddev floci-gcp url
# From the project (web container, Cloud Run, Cloud SQL): http://floci-gcp:4588
# From the host, REST — stable, shared between projects:  http://<project>.ddev.site:4588
#                      (TLS: https://<project>.ddev.site:4589)
# From the host, REST and gRPC:                           http://127.0.0.1:63756
```

**For REST, use the router endpoint.** ddev-router routes by hostname, so port
4588 is *shared* between every project running this add-on rather than owned by
one of them — each project's requests reach its own emulator. It is stable
across restarts, so it is the one worth writing into config.

**For gRPC, use the direct port.** Pub/Sub, Firestore and Secret Manager speak
nothing else, and gRPC does not survive the router: DDEV wires services to
Traefik as HTTP/1.1 backends, so a gRPC client pointed at the router endpoint
hangs and eventually reports `No status received`. The direct port is raw TCP
and carries it.

That port is assigned by Docker rather than fixed — the same way DDEV publishes
the web and db containers — so two projects can never collide over it. The cost
is that it changes on every `ddev restart`, so ask for it instead of writing it
down:

```bash
eval "$(ddev floci-gcp env)"     # re-run after a restart
ddev floci-gcp url --host
```

`ddev floci-gcp env` splits the difference for you: the `*_EMULATOR_HOST`
variables get the direct port because they carry gRPC, while `STORAGE_EMULATOR_HOST`
gets the stable router URL because GCS is REST.

If nothing on your host needs gRPC, you can drop the direct port entirely and
lose nothing else:

```bash
rm .ddev/docker-compose.floci-gcp-hostport.yaml && ddev restart
```

### Running several projects at once

This is the reason for the split. Two DDEV projects with this add-on installed
start side by side with no configuration: they share router ports 4588/4589 by
hostname, and Docker hands each one its own ephemeral gRPC port. State is per
project — the `floci-gcp` volume is namespaced `ddev-<project>_floci-gcp` — and
`FLOCI_GCP_DOCKER_RESOURCE_NAMESPACE` scopes every sidecar container Floci
spawns to its own project, so `ddev floci-gcp reset` or a full
`ddev delete` in one leaves the other untouched.

## The two default buckets

Out of the box the add-on creates two GCS buckets on every start:

| Bucket | Default name | Intent |
|---|---|---|
| Public | `public` | IAM binding: `roles/storage.objectViewer` for `allUsers` |
| Private | `private` | No bindings — owner only |

The web container gets `GCS_PUBLIC_BUCKET`, `GCS_PRIVATE_BUCKET`, `GCS_BUCKET`
(the public one) and `GCS_PUBLIC_BUCKET_URL` — the last being the router URL a
browser can actually resolve, since a URL built from `floci-gcp:4588` only works
inside Docker.

Rename either, or set one to an empty string to skip it:

```bash
ddev dotenv set .ddev/.env.floci-gcp --floci-gcp-public-bucket=assets
ddev dotenv set .ddev/.env.floci-gcp --floci-gcp-private-bucket=""
ddev restart
```

### Bucket permissions are not enforced

**The two buckets are equally open at runtime.** floci-gcp stores the IAM policy
and hands it back — `getIamPolicy` on `public` returns the `allUsers` binding,
and `private` returns none — but it does not enforce it. Measured on 0.7.0:

| | `public` | `private` |
|---|---|---|
| Unsigned read of an object | 200 | **200** |
| Unsigned bucket listing | 200 | **200** |
| Unsigned write | 200 | **200** |

There is no setting to change this. Unlike the AWS emulator, which has
`FLOCI_SERVICES_S3_ENFORCE_AUTH`, floci-gcp's configuration has no auth or ACL
enforcement option at all — the emulator accepts every request on every bucket.

So the split exists to make your application and IaC exercise the right shapes —
the right bucket for the right data, the right IAM policy in Terraform — and so
it is already in place if upstream starts enforcing. Treat *both* buckets as
world-readable and world-writable locally, and do not rely on the emulator to
catch a permissions mistake; only real GCP will.

> *Measured against `floci/floci-gcp:0.7.0` on 2026-08-26. Floci is under active development, so this may already be out of date.*


### Bucket listing is empty after a restart

A second upstream limitation, worth knowing before it confuses you. With the
default `hybrid` storage, buckets survive a restart and remain fully usable by
name — but they disappear from the bucket *listing*:

```bash
curl "$STORAGE_EMULATOR_HOST/storage/v1/b/public"      # 200, still there
curl "$STORAGE_EMULATOR_HOST/storage/v1/b?project=..." # {"kind":"storage#buckets"} — empty
```

Buckets restored from disk are never re-added to the project's bucket index, and
nothing repairs it: re-creating returns `409 Bucket already exists`, and neither
a metadata `PATCH` nor a fresh `setIamPolicy` puts them back in the list.

What still works after a restart: reading, writing and **listing objects inside
a bucket**, which is what most applications and the Media Library style of
browsing actually do. What breaks: `list_buckets()`, and the storage view in the
Floci console, which both come back empty.

If bucket listing matters to you, run in memory mode — every start is fresh, so
the index is always correct, at the cost of losing state between restarts:

```bash
ddev dotenv set .ddev/.env.floci-gcp --floci-gcp-storage-mode=memory
ddev restart
```

> *Measured against `floci/floci-gcp:0.7.0` on 2026-08-26. Floci is under active development, so this may already be out of date. The add-on's test suite carries a canary — `buckets survive a restart by name, even though listing does not` — that goes red when this changes, so a nightly CI failure there is good news, not a regression.*


## Creating resources on every start

Drop scripts in `.ddev/floci-gcp/init/ready.d/`. They run once the GCP API is
accepting requests, on every container start.

**Write them as `.sh`, not `.py`** — see the warning below, it is not a style
preference. The shape to copy:

```sh
#!/bin/sh
set -eu
PY="$(command -v python3 || command -v python3.11 || true)"
[ -n "${PY}" ] || exit 0

"${PY}" - <<'PYEOF' || echo "hook failed; emulator unaffected" >&2
import os
os.environ.setdefault("STORAGE_EMULATOR_HOST", "http://localhost:4588")
from google.cloud import storage
try:
    storage.Client(project="floci-local").create_bucket("uploads")
except Exception:
    pass
PYEOF
```

`chmod +x` it and `ddev restart`. Make them idempotent — with the default
storage mode the resources they create survive restarts, so a script that fails
on "already exists" will break your next start. A hook is killed after 30
seconds.

**Do not name a hook `.py`.** floci-gcp runs a `.py` hook by exec'ing
`python3`, and the compat image installs `python3.11` with no `python3` symlink
— so the exec fails, and floci-gcp treats a failed startup hook as fatal and
shuts the whole emulator down. Write a `.sh` hook that finds the interpreter
itself; the shipped `10-default-buckets.sh` shows the pattern:

```sh
PY="$(command -v python3 || command -v python3.11 || true)"
[ -n "${PY}" ] && "${PY}" - <<'PYEOF'
# ... your Python here; the compat image has the google-cloud-* libraries
PYEOF
```

Also wrap the body so a failure cannot take the emulator with it — the shipped
hook ends its Python block with `|| echo ... >&2` for exactly that reason.

> *Measured against `floci/floci-gcp:0.7.0` on 2026-08-26. Floci is under active development, so this may already be out of date. The add-on's test suite carries a canary — `a .py init hook would kill the emulator, so the shipped example is .sh` — that goes red when this changes, so a nightly CI failure there is good news, not a regression.*


There are four phases (`boot.d`, `start.d`, `ready.d`, `stop.d`); see
[`floci-gcp/init/README.md`](https://github.com/codementality/ddev-floci-gcp/blob/main/floci-gcp/init/README.md) and
`ready.d/10-example.sh.disabled`.

## Persistence

Upstream defaults to `memory`; this add-on defaults to `hybrid` — in-memory
speed with an asynchronous flush every five seconds, so your buckets and topics
survive `ddev restart`. State lives in a Docker volume, not in your project
directory.

```bash
# Wipe everything in place, no restart — the one for between test runs.
ddev floci-gcp flush

# Delete the state volume and any containers Floci spawned.
ddev floci-gcp reset

# Or make every start clean.
ddev dotenv set .ddev/.env.floci-gcp --floci-gcp-storage-mode=memory
ddev restart
```

Other modes are `persistent` (flush on every write) and `wal` (write-ahead log,
maximum durability). Use `memory` in CI.

## Docker-backed services and the Docker socket

floci-gcp runs **Cloud Run**, **Cloud Functions**, **Cloud SQL for PostgreSQL**,
**Managed Kafka** and **GKE** as real containers rather than mocking their
responses — a real Postgres, a real Redpanda broker, a real k3s cluster, your
actual Cloud Run image executing. Doing so requires the host's Docker socket,
which `docker-compose.floci-gcp-docker.yaml` mounts. The socket is also what
powers Floci's embedded DNS.

Mounting the Docker socket into a container gives it control of the host's
Docker daemon, which is effectively root on the host. It is the same bargain
LocalStack and Testcontainers ask for, but it is worth making deliberately. If
this project only uses the in-process services — GCS, Pub/Sub, Firestore,
Datastore, Secret Manager, IAM, KMS, BigQuery, Cloud Tasks, Scheduler, Eventarc,
Logging, Monitoring, Firebase Auth, STS and the rest — you can decline it:

```bash
rm .ddev/docker-compose.floci-gcp-docker.yaml
ddev restart
```

You can also keep the socket and stop one service from spawning anything, by
putting it in control-plane-only mock mode:

```bash
ddev dotenv set .ddev/.env.floci-gcp --floci-gcp-services-cloudsql-mock=true
```

One wrinkle worth knowing: containers Floci starts are its own, not DDEV's, so
`ddev stop` leaves them running. `ddev floci-gcp reset` removes the ones
belonging to this project along with the state.

## Projects are the isolation boundary

Two things called "project" meet here, and it helps to keep them apart.

**Your DDEV project.** State is per DDEV project, not shared. The `floci-gcp`
volume is a Compose named volume, so Docker namespaces it as
`ddev-<project>_floci-gcp`. This add-on also sets
`FLOCI_GCP_DOCKER_RESOURCE_NAMESPACE` to the DDEV project name, which names and
labels every sidecar container and volume Floci spawns — so two DDEV projects
can each run a Cloud SQL instance at the same time without claiming the same
container, and `ddev floci-gcp reset` in one leaves the other untouched. Both
projects can serve on router port 4589 at once, because the router routes by
hostname; only the direct host port needs to differ.

**Your GCP project id.** Inside the emulator, the GCP project id is the
multi-tenancy boundary: `projects/a/topics/jobs` and `projects/b/topics/jobs`
are two different topics, and one is invisible to the other. The default id is
`floci-local`; change it with
`ddev dotenv set .ddev/.env.floci-gcp --floci-gcp-default-project-id=my-project`,
which also updates `GOOGLE_CLOUD_PROJECT` in the web container.

## Migrating from the gcloud emulators

Point the same variables you already set at port 4588 and delete the
`gcloud beta emulators ... start` processes:

| gcloud emulator | floci-gcp |
|---|---|
| `PUBSUB_EMULATOR_HOST=localhost:8085` | `PUBSUB_EMULATOR_HOST=floci-gcp:4588` |
| `FIRESTORE_EMULATOR_HOST=localhost:8080` | `FIRESTORE_EMULATOR_HOST=floci-gcp:4588` |
| `DATASTORE_EMULATOR_HOST=localhost:8081` | `DATASTORE_EMULATOR_HOST=floci-gcp:4588` |
| `FIREBASE_AUTH_EMULATOR_HOST=localhost:9099` | `FIREBASE_AUTH_EMULATOR_HOST=floci-gcp:4588` |
| *(no official emulator)* | `STORAGE_EMULATOR_HOST=http://floci-gcp:4588` |

This add-on has already set all of them in the web container, so in practice
migrating means deleting the ones you were setting yourself.

## Configuration

These are Floci's own variable names, not add-on-specific aliases, because
floci-gcp already namespaces everything under `FLOCI_GCP_*`. Anything in the
[upstream reference](https://floci.io/floci-gcp/configuration/environment-variables/)
works here verbatim — set it and restart:

```bash
ddev dotenv set .ddev/.env.floci-gcp --floci-gcp-services-bigquery-enabled=false
ddev restart
```

The ones this add-on gives a different default to, or invents:

| Variable | Default | Purpose |
|---|---|---|
| `FLOCI_GCP_IMAGE` | `floci/floci-gcp:latest-compat` | Image tag. `floci/floci-gcp:latest` is the lean ~90 MB native build **without** Python or the client libraries |
| `FLOCI_GCP_HOSTNAME` | `floci-gcp` | Hostname on DDEV's network, and the host in returned URLs |
| `FLOCI_GCP_BASE_URL` | `http://floci-gcp:4588` | Base URL Floci falls back to when building URLs (upstream: `http://localhost:4588`) |
| `FLOCI_GCP_HTTP_PORT` | `4588` | ddev-router port (http), shared between projects |
| `FLOCI_GCP_HTTPS_PORT` | `4589` | ddev-router port (https), shared between projects |
| `FLOCI_GCP_DEFAULT_PROJECT_ID` | `floci-local` | Default GCP project id |
| `FLOCI_GCP_PUBLIC_BUCKET` | `public` | Bucket created on every start, with an advisory public-read IAM binding. Empty string to skip |
| `FLOCI_GCP_PRIVATE_BUCKET` | `private` | Owner-only bucket created on every start. Empty string to skip |
| `FLOCI_GCP_STORAGE_MODE` | `hybrid` | `memory`, `persistent`, `hybrid` or `wal` (upstream: `memory`) |
| `FLOCI_GCP_SERVICES_DOCKER_NETWORK` | `ddev-<project>_default` | Network for containers Floci spawns |
| `FLOCI_GCP_DOCKER_RESOURCE_NAMESPACE` | `<project>` | Names and labels spawned containers and volumes, so projects do not collide |
| `FLOCI_GCP_SERVICES_GKE_ENDPOINT_MODE` | `host` | `host` for a kubectl on your machine, `network` for one inside the project |
| `FLOCI_GCP_DNS_CONTAINER_FALLBACK_ENABLED` | `true` | Public resolvers in spawned containers; set `false` on an offline network |

Anything else — `FLOCI_GCP_SERVICES_CLOUDSQL_POSTGRES17_IMAGE`, say — can go
straight into `.ddev/.env.floci-gcp` and be added to the `environment:` block of
`.ddev/docker-compose.floci-gcp.yaml`, or set in a compose file of your own.

## The `ddev floci-gcp` command

| Command | Does |
|---|---|
| `ddev floci-gcp gcloud <args>` | Runs **your host's** gcloud against the emulator |
| `ddev floci-gcp python <args>` | Runs Python inside the container, client libraries and endpoints already set |
| `ddev floci-gcp env [--unset]` | Prints `export` lines for your shell; `eval "$(...)"` to apply |
| `ddev floci-gcp url` | Prints all three endpoints (`--internal` / `--host` / `--router` for one) |
| `ddev floci-gcp health` | Hits `/_floci-gcp/health` |
| `ddev floci-gcp info` | Version, port and default project |
| `ddev floci-gcp services` | Per-service status |
| `ddev floci-gcp logs [-f]` | Container logs |
| `ddev floci-gcp shell` | A shell in the container |
| `ddev floci-gcp flush` | Wipes all emulated state in place, no restart |
| `ddev floci-gcp reset` | Deletes the state volume, and any containers Floci spawned for this project |

`ddev floci-gcp gcloud` runs the gcloud on your machine, not one in the
container — neither Floci image ships it, since the Cloud SDK is a ~200 MB
Python distribution that would dwarf the emulator. It needs the direct host port
(see above). `ddev floci-gcp python` needs the `-compat` image, which is the
default.

## Removing the add-on

Clean up first, remove second. `ddev add-on remove` deletes the `ddev floci-gcp`
command along with everything else, so `reset` is no longer available once the
add-on is gone:

```bash
ddev floci-gcp reset          # containers Floci spawned, plus the state volume
ddev add-on remove floci-gcp
```

`ddev add-on remove` deliberately leaves three things behind, because none of
them are the add-on's to throw away:

- **The state volume**, holding your emulated buckets, topics and secrets. DDEV
  does not delete it, and neither does this add-on. The removal output prints
  the exact command; it is `ddev-` plus your project name, then `_floci-gcp`:

  ```bash
  docker volume rm ddev-<project>_floci-gcp
  ```

  Note the doubled prefix when the project name itself starts with `ddev-`. A
  project named `ddev-floci-gcp` gives `ddev-ddev-floci-gcp_floci-gcp`. If
  you're unsure of the name:

  ```bash
  docker volume ls --filter label=com.docker.compose.volume=floci-gcp
  ```

- **Your init scripts** in `.ddev/floci-gcp/init/`, if you wrote any. The
  directory is removed only when it still looks untouched.
- **`.ddev/.env.floci-gcp`**, so your settings survive a reinstall.

If you skipped the `reset` and the add-on is already gone, the leftover
containers Floci spawned can be cleared by hand. Keep the namespace filter — the
emulator label alone matches every project's floci-gcp containers, not just this
one's:

```bash
docker rm -f $(docker ps -aq \
  --filter label=floci_emulator=floci-gcp \
  --filter label=floci_namespace=<project>)
```

## Licensing

Two different licences are in play here, and it is worth being explicit about
which covers what:

- **This add-on** — the compose files, the `ddev floci-gcp` command, the init
  scripts and the documentation — is licensed **Apache 2.0**, matching the DDEV
  add-on template and DDEV's own add-ons. See [LICENSE](https://github.com/codementality/ddev-floci-gcp/blob/main/LICENSE).
- **floci-gcp itself** is licensed **MIT** by the
  [floci-io](https://github.com/floci-io/floci-gcp) project, and is entirely separate work.

**No Floci source is vendored in this repository.** The add-on only references
the published `floci/floci-gcp` image, which Docker pulls at runtime; nothing
here redistributes Floci code or binaries. If you redistribute this add-on you
are redistributing Apache-2.0 material only.

## Credits

floci-gcp is by the [floci-io](https://github.com/floci-io/floci-gcp) project, MIT licensed. This add-on packages
it under Apache 2.0; see [Licensing](#licensing) above.

Behavioural notes in this README were verified against `floci/floci-gcp:0.7.0` on 2026-08-26.
Upstream ships frequently; the nightly workflow in `.github/workflows/tests.yml`
re-runs the suite against `latest` every day, which is what catches drift.

**Contributed and maintained by [@codementality](https://github.com/codementality)**
