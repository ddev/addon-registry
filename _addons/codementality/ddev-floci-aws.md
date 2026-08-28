---
title: "codementality/ddev-floci-aws"
github_url: "https://github.com/codementality/ddev-floci-aws"
description: "DDEV Floci emulator for AWS services"
user: "codementality"
repo: "ddev-floci-aws"
repo_id: 1347075975
default_branch: "main"
tag_name: "0.5.2"
ddev_version_constraint: ">= v1.24.10"
dependencies: ["codementality/ddev-floci-ui"]
type: "contrib"
created_at: "2026-08-26"
updated_at: "2026-08-26"
workflow_status: "success"
stars: 1
---

# ddev-floci-aws <!-- omit in toc -->

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/codementality/ddev-floci-aws/actions/workflows/tests.yml/badge.svg)](https://github.com/codementality/ddev-floci-aws/actions/workflows/tests.yml)
[![last commit](https://img.shields.io/github/last-commit/codementality/ddev-floci-aws)](https://github.com/codementality/ddev-floci-aws/commits)
[![release](https://img.shields.io/github/v/release/codementality/ddev-floci-aws)](https://github.com/codementality/ddev-floci-aws/releases/latest)

A [DDEV](https://ddev.com) add-on that runs [Floci](https://floci.io/aws/) — a
free, open-source local AWS emulator — alongside your project, and points your
web container at it.

Your application talks to `http://floci-aws:4566` instead of AWS. S3 buckets,
DynamoDB tables, SQS queues, SNS topics, Secrets Manager secrets, Lambda
functions and about seventy other services exist locally, with no AWS account,
no auth token, and no feature gates.

- [Why this add-on](#why-this-add-on)
- [Running alongside the Azure and GCP emulators](#running-alongside-the-azure-and-gcp-emulators)
- [Installation](#installation)
- [Using it](#using-it)
- [Reaching Floci from the host](#reaching-floci-from-the-host)
- [The two default S3 buckets](#the-two-default-s3-buckets)
- [Creating resources on every start](#creating-resources-on-every-start)
- [Persistence](#persistence)
- [Docker-backed services and the Docker socket](#docker-backed-services-and-the-docker-socket)
- [Migrating from LocalStack](#migrating-from-localstack)
- [Configuration](#configuration)
- [The `ddev floci-aws` command](#the-ddev-floci-aws-command)
- [Removing the add-on](#removing-the-add-on)
- [Licensing](#licensing)
- [Credits](#credits)

## Why this add-on

LocalStack's community edition
[sunset in March 2026](https://blog.localstack.cloud/the-road-ahead-for-localstack/):
it now requires an auth token and its security updates are frozen. Floci is a
drop-in replacement under the MIT license — same port, same credentials model,
same SDK configuration — and it is meaningfully lighter, starting in
milliseconds rather than seconds.

Running it as a DDEV add-on rather than as a separate `docker compose` stack
buys you three things that are fiddly to arrange by hand:

- **The web container is configured for you.** `AWS_ENDPOINT_URL`,
  credentials, region and path-style addressing are set in the web container's
  environment, so any AWS SDK finds the emulator without an application change.
- **Returned URLs resolve.** SQS queue URLs, S3 locations and Lambda function
  URLs name the `floci-aws` container rather than `localhost`, which is what makes
  them usable from your PHP/Node/Python code. Getting this wrong is the single
  most common way a local AWS emulator appears to work and then doesn't.
- **Docker-backed services land on DDEV's network.** Lambda, RDS and
  ElastiCache containers that Floci starts join the project's own network, so your web
  container can reach them and they can reach Floci back.

## Running alongside the Azure and GCP emulators

Everything this add-on installs is named for the cloud it emulates rather than
for Floci: the service and its hostname are `floci-aws`, the container is
`ddev-<project>-floci-aws`, the state volume is `floci-aws`, and the command is
`ddev floci-aws`. Nothing here claims the bare `floci` name, so the equivalent
Azure and GCP add-ons can be installed in the same project without overwriting
this one's files or shadowing its command.

Your application can then tell them apart by endpoint:
`http://floci-aws:4566`, `http://floci-az:4577`, `http://floci-gcp:4588`.

### Installing it in several projects at once

State is per project, not shared. The `floci-aws` volume is a Compose named
volume, so Docker namespaces it by Compose project — DDEV gives each project
`ddev-<project>_floci-aws`. Two projects each get their own buckets, tables and
queues, `ddev floci-aws reset` in one leaves the other untouched, and deleting
one project's volume cannot corrupt another's. Both projects can serve on
router port 4567 at once, because the router routes by hostname.

The exception is **Docker-backed services**, which are not namespaced upstream.
Floci names some of the containers it starts globally — the ECR registry is
`floci-ecr-registry`, with no project in the name — so the first project to ask
for one gets it, and a second project asking for the same service silently
reuses that container even though it is attached to the first project's network
and therefore unreachable. If two projects both need ECR, RDS or another
Docker-backed service, run them one at a time, or keep the Docker-backed work in
a single project. Everything in-process (S3, DynamoDB, SQS, SNS, IAM, Secrets
Manager, …) is unaffected.

## Installation

```bash
ddev add-on get codementality/ddev-floci-aws
ddev restart
```

Then check it:

```bash
ddev floci-aws aws s3 mb s3://uploads
ddev floci-aws aws s3 ls
```

## Using it

Nothing in your application needs to know it is talking to an emulator. The
web container already has the environment set:

| Variable | Value |
|---|---|
| `AWS_ENDPOINT_URL`, `AWS_ENDPOINT` | `http://floci-aws:4566` |
| `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY` | `test` / `test` |
| `AWS_DEFAULT_REGION`, `AWS_REGION` | `us-east-1` |
| `AWS_USE_PATH_STYLE_ENDPOINT`, `AWS_S3_FORCE_PATH_STYLE` | `true` |
| `AWS_BUCKET`, `AWS_PUBLIC_BUCKET` | `public` |
| `AWS_PRIVATE_BUCKET` | `private` |
| `AWS_PUBLIC_BUCKET_URL` | `https://<project>.ddev.site:4567/public` — browser-facing |

**PHP** (`aws/aws-sdk-php` 3.300+ reads `AWS_ENDPOINT_URL` on its own):

```php
$s3 = new Aws\S3\S3Client([
  'version' => 'latest',
  'region'  => getenv('AWS_DEFAULT_REGION'),
  // Older SDK versions need these two spelled out:
  'endpoint' => getenv('AWS_ENDPOINT_URL'),
  'use_path_style_endpoint' => true,
]);
$s3->putObject(['Bucket' => 'uploads', 'Key' => 'a.txt', 'Body' => 'hi']);
```

For Drupal with `s3fs` or `flysystem` modules, point the bucket at `uploads` and the
custom host at `floci-aws:4566` with path-style addressing on.

**Node / Python / Go**: the v3 JS clients, botocore ≥ 1.31 and the Go v2 SDK all
honour `AWS_ENDPOINT_URL` with no configuration at all.

**Terraform / OpenTofu / CDK** work too — point the provider's `endpoints` block
at the same URL.

## Reaching Floci from the host

There are two endpoints, and they are not interchangeable:

```bash
ddev floci-aws url
# From the project (web container, Lambda, RDS): http://floci-aws:4566
# From the host (aws CLI, Terraform, your IDE):   https://<project>.ddev.site:4567
```

The in-project URL is the one baked into service URLs Floci returns, because
that is what your application code has to be able to reach. A host-side AWS CLI
pointed at `https://<project>.ddev.site:4567` will therefore get back queue URLs
naming `floci-aws:4566`, which the host cannot resolve. That is a real limitation,
not an oversight — you cannot have both without a hostname that resolves in
both places.

If your workflow is host-first (running Terraform from your terminal rather than
from `ddev ssh`), flip it:

```bash
ddev dotenv set .ddev/.env.floci-aws \
  --floci-aws-base-url=https://myproject.ddev.site:4567
ddev restart
```

The simplest way to avoid the question entirely is `ddev floci-aws aws ...`, which
runs the CLI inside the container where both halves agree.

## The two default S3 buckets

Out of the box the add-on creates two buckets on every start, with genuinely
different permissions:

| Bucket | Default name | Unsigned read | Unsigned write | Unsigned list |
|---|---|---|---|---|
| Public | `public` | ✅ objects readable | ❌ 403 | ❌ 403 |
| Private | `private` | ❌ 403 | ❌ 403 | ❌ 403 |

The **public** bucket is for the 90% case — images, document uploads, anything
meant to be served straight to a browser once it lands. It gets a bucket policy
granting `s3:GetObject` to `*`, so an object is readable the moment it is
written and your application never has to remember to set `public-read` on each
upload. Writing still requires credentials.

Listing is deliberately *not* public by default. A publicly listable bucket is
a misconfiguration rather than a feature, and nothing needs it to serve an
`<img src>`.

**This does not affect application code.** Only *anonymous* listing is blocked.
A signed `ListObjectsV2` — what the AWS SDK inside Drupal issues, including the
prefix/delimiter form that drives folder-style browsing in Media Library and
s3fs — works on both buckets, because the web container already holds
credentials. Only an unauthenticated browser hitting the bucket root gets a 403.

If you do find something that needs anonymous listing, it is a switch rather
than a hook edit:

```bash
ddev dotenv set .ddev/.env.floci-aws --floci-aws-public-bucket-list=true
ddev restart
```

That adds `s3:ListBucket` on the bucket ARN to the policy. Unsigned *writes*
stay refused either way.

The **private** bucket has no ACL and no policy — S3's owner-only default.
Nothing gets in without credentials or a presigned URL, which is what makes it
useful for testing Flysystem's presigned links and private uploads.

> *Measured against `floci/floci:1.7.0` on 2026-08-26. Floci is under active development, so this may already be out of date.*


The web container gets all of this in its environment:

| Variable | Value |
|---|---|
| `AWS_PUBLIC_BUCKET` | `public` |
| `AWS_PRIVATE_BUCKET` | `private` |
| `AWS_BUCKET` | `public` — the 90% case, so config can stay short |
| `AWS_PUBLIC_BUCKET_URL` | `https://<project>.ddev.site:4567/public` |

`AWS_PUBLIC_BUCKET_URL` is the one worth knowing about. A URL built from
`AWS_ENDPOINT_URL` names `floci-aws:4566`, which resolves only inside Docker —
paste it into an `<img src>` and the browser shows a broken image. This variable
is the same object as ddev-router serves it, which is what actually loads:

```php
$src = getenv('AWS_PUBLIC_BUCKET_URL') . '/' . $key;   // renders
$src = getenv('AWS_ENDPOINT_URL') . '/public/' . $key; // broken image
```

Rename either bucket, or set one to an empty string to skip it:

```bash
ddev dotenv set .ddev/.env.floci-aws --floci-aws-public-bucket=assets
ddev dotenv set .ddev/.env.floci-aws --floci-aws-private-bucket=""
ddev restart
```

### This depends on auth enforcement being on

`FLOCI_AWS_S3_ENFORCE_AUTH` defaults to **`true`** in this add-on — unlike
upstream Floci and unlike LocalStack Community, which both default it off.

That is a deliberate difference, because with enforcement off S3 ignores
signatures entirely: every bucket is world-readable *and* world-writable, the
policy above is decorative, and the two bucket names imply a boundary that is
not there. It also makes local behaviour diverge from real AWS in the direction
that hides bugs — code that forgets to sign works locally and fails in
production.

If something in your stack genuinely cannot sign requests, turn it off and
accept that both buckets are then equally open:

```bash
ddev dotenv set .ddev/.env.floci-aws --floci-aws-s3-enforce-auth=false
```

The init hook prints a warning on every start while it is off, so this cannot
be true by accident.

### Presigned URLs

Presigned URLs work against the private bucket, and the signature is really
checked — a tampered one gets a 403. The catch is the one from
[Reaching Floci from the host](#reaching-floci-from-the-host): the signature
covers the `Host` header, so a URL is only valid at the endpoint it was signed
for.

- **Fetched server-side** (your PHP code following the link): sign with the
  default `AWS_ENDPOINT_URL`, giving a URL naming `floci-aws:4566`. Works from
  the web container, unusable from your browser.
- **Handed to a browser** (the normal reason to presign): sign with the router
  URL instead, so the link names a host your machine can resolve.

```php
$s3 = new Aws\S3\S3Client([
  'version'  => 'latest',
  'region'   => getenv('AWS_DEFAULT_REGION'),
  // Not AWS_ENDPOINT_URL — the browser has to be able to resolve this host.
  'endpoint' => 'https://' . getenv('DDEV_HOSTNAME') . ':4567',
  'use_path_style_endpoint' => true,
]);
$cmd = $s3->getCommand('GetObject', [
  'Bucket' => getenv('AWS_PRIVATE_BUCKET'),
  'Key'    => $key,
]);
$url = (string) $s3->createPresignedRequest($cmd, '+20 minutes')->getUri();
```

Or flip it project-wide with `--floci-aws-base-url=https://<project>.ddev.site:4567`
if your whole workflow is host-first.

The buckets are *created* on start, not emptied: with the default `hybrid`
storage mode their contents survive restarts. Renaming creates the new bucket
and leaves the old one alone, so nothing is lost by changing a setting. Use
`ddev floci-aws reset` to start over.

The hook itself is
[`floci-aws/init/ready.d/10-default-buckets.sh`](https://github.com/codementality/ddev-floci-aws/blob/main/floci-aws/init/ready.d/10-default-buckets.sh),
an ordinary init script with nothing special about it. Delete it if you would
rather do your own thing, or read it as the worked example for the section
below.

## Creating resources on every start

Drop scripts in `.ddev/floci-aws/init/ready.d/`. They run once the AWS API is
accepting requests, on every container start, and the `aws` CLI inside the
container already points at the emulator:

```sh
#!/bin/sh
set -eu
aws s3api create-bucket --bucket uploads 2>/dev/null || true
aws sqs create-queue --queue-name jobs 2>/dev/null || true
```

`chmod +x` it and `ddev restart`. Make them idempotent — with the default
storage mode the resources they create survive restarts, so a script that fails
on "already exists" will break your next start. The shipped
`10-default-buckets.sh` shows the pattern: `head-bucket` first, create only if
that misses, so a genuine failure is still reported instead of swallowed by
`|| true`.

There are four phases (`boot.d`, `start.d`, `ready.d`, `stop.d`); see
[`floci-aws/init/README.md`](https://github.com/codementality/ddev-floci-aws/blob/main/floci-aws/init/README.md). `.py` scripts work as
well as `.sh`.

## Persistence

The default storage mode is `hybrid`: in-memory speed with an asynchronous flush
every few seconds, so your buckets and tables survive `ddev restart`. State
lives in a Docker volume, not in your project directory.

```bash
# Start clean.
ddev floci-aws reset

# Or make every start clean.
ddev dotenv set .ddev/.env.floci-aws --floci-aws-storage-mode=memory
ddev restart
```

Other modes are `persistent` (flush on every write) and `wal` (write-ahead log,
maximum durability). Use `memory` in CI.

## Docker-backed services and the Docker socket

Floci runs Lambda, RDS, Neptune, ElastiCache, MSK, ECS, EC2, EKS, OpenSearch,
DocumentDB and CodeBuild as **real containers** rather than mocking their
responses — that is what makes a Lambda actually execute your handler and an
RDS instance actually speak Postgres. Doing so requires the host's Docker
socket, which `docker-compose.floci-aws-docker.yaml` mounts.

Mounting the Docker socket into a container gives it control of the host's
Docker daemon, which is effectively root on the host. It is the same bargain
LocalStack and Testcontainers ask for, but it is worth making deliberately. If
this project only uses the in-process services — S3, DynamoDB, SQS, SNS, IAM,
KMS, Secrets Manager, SSM, EventBridge, Step Functions, API Gateway, Cognito,
Kinesis, SES and the rest — you can decline it:

```bash
rm .ddev/docker-compose.floci-aws-docker.yaml
ddev restart
```

Everything in-process keeps working; the Docker-backed services do not start.

One wrinkle worth knowing: containers Floci starts are its own, not DDEV's, so
`ddev stop` leaves them running. `ddev floci-aws reset` removes the ones belonging
to this project along with the state.

## Migrating from LocalStack

LocalStack parity is on by default, so an existing setup mostly moves across
unchanged:

- `LOCALSTACK_HOST`, `PERSISTENCE=1`, `LAMBDA_DOCKER_NETWORK`,
  `LAMBDA_REMOVE_CONTAINERS=1` and `DEBUG=1` are translated automatically.
- `/_localstack/health` and `/_localstack/init` are still served.
- Init scripts under `/etc/localstack/init/` still run — though this add-on
  mounts the native path, so put yours in `.ddev/floci-aws/init/ready.d/`.
- The startup log still ends with LocalStack's `Ready.` line, so Testcontainers'
  `LocalStackContainer` wait strategy works.

Turn the translation off with
`ddev dotenv set .ddev/.env.floci-aws --floci-aws-localstack-parity=false`.

## Configuration

Every setting is a `ddev dotenv set` away, followed by `ddev restart`:

```bash
ddev dotenv set .ddev/.env.floci-aws --floci-aws-region=eu-west-1
```

| Variable | Default | Purpose |
|---|---|---|
| `FLOCI_AWS_IMAGE` | `floci/floci:latest-compat` | Image tag. `floci/floci:latest` is the lean ~90 MB native build **without** the AWS CLI or boto3 |
| `FLOCI_AWS_HOSTNAME` | `floci-aws` | Hostname on DDEV's network, and the host in returned service URLs |
| `FLOCI_AWS_BASE_URL` | `http://floci-aws:4566` | Base URL Floci uses when it builds service URLs |
| `FLOCI_AWS_HTTP_PORT` | `4566` | Router port (http) |
| `FLOCI_AWS_HTTPS_PORT` | `4567` | Router port (https) |
| `FLOCI_AWS_REGION` | `us-east-1` | Default region |
| `FLOCI_AWS_ACCOUNT_ID` | `000000000000` | Default account id |
| `FLOCI_AWS_STORAGE_MODE` | `hybrid` | `memory`, `persistent`, `hybrid` or `wal` |
| `FLOCI_AWS_DOCKER_NETWORK` | `ddev-<project>_default` | Network for containers Floci starts |
| `FLOCI_AWS_LOCALSTACK_PARITY` | `true` | Accept LocalStack env vars and endpoints |
| `FLOCI_AWS_S3_ENFORCE_AUTH` | `true` | Enforce S3 public/private access and reject unsigned requests. **On by default here**, unlike upstream |
| `FLOCI_AWS_ACCESS_KEY_ID` | `test` | Credentials handed to the web container |
| `FLOCI_AWS_SECRET_ACCESS_KEY` | `test` | Credentials handed to the web container |
| `FLOCI_AWS_S3_PATH_STYLE` | `true` | Path-style S3 addressing in the web container |
| `FLOCI_AWS_PUBLIC_BUCKET` | `public` | Bucket created on every start with a public-read policy. Empty string to skip |
| `FLOCI_AWS_PRIVATE_BUCKET` | `private` | Owner-only bucket created on every start. Empty string to skip |
| `FLOCI_AWS_PUBLIC_BUCKET_LIST` | `false` | Allow *anonymous* listing of the public bucket. Signed listing works regardless |

Floci reads any `FLOCI_*` variable directly; for anything not listed above —
`FLOCI_SERVICES_RDS_DEFAULT_POSTGRES_IMAGE`, say — add it to the `environment:`
block of `.ddev/docker-compose.floci-aws.yaml`, or override it in a compose file
of your own. See the
[upstream configuration reference](https://floci.io/floci/configuration/advanced/application-yml).

**Multi-account isolation** works out of the box: set `AWS_ACCESS_KEY_ID` to
exactly twelve digits and Floci treats it as the account id, so resources from
one account are invisible to another.

## The `ddev floci-aws` command

| Command | Does |
|---|---|
| `ddev floci-aws aws <args>` | Runs the AWS CLI inside the container, already pointed at the emulator |
| `ddev floci-aws url` | Prints both endpoints (`--internal` / `--host` for just one) |
| `ddev floci-aws health` | Hits `/_floci/health` |
| `ddev floci-aws services` | Per-service status from `/_localstack/health` |
| `ddev floci-aws logs [-f]` | Container logs |
| `ddev floci-aws shell` | A shell in the container |
| `ddev floci-aws reset` | Deletes all emulated AWS state, and any containers Floci started for this project |

`ddev floci-aws aws` needs the AWS CLI, which only the `-compat` image ships. If you
switch `FLOCI_AWS_IMAGE` to `floci/floci:latest` for the smaller image, use your
own AWS CLI against `ddev floci-aws url --host` instead.

## Removing the add-on

Clean up first, remove second. `ddev add-on remove` deletes the `ddev floci-aws`
command along with everything else, so `reset` is no longer available once the
add-on is gone:

```bash
ddev floci-aws reset          # containers Floci started, plus the state volume
ddev add-on remove floci-aws
```

`ddev add-on remove` deliberately leaves three things behind, because none of
them are the add-on's to throw away:

- **The state volume**, holding your emulated buckets, tables and queues. DDEV
  does not delete it, and neither does this add-on. The removal output prints
  the exact command; it is `ddev-` plus your project name, then `_floci-aws`:

  ```bash
  docker volume rm ddev-<project>_floci-aws
  ```

  Note the doubled prefix when the project name itself starts with `ddev-`. A
  project named `ddev-floci-aws` gives `ddev-ddev-floci-aws_floci-aws`. If
  you're unsure of the name:

  ```bash
  docker volume ls --filter label=com.docker.compose.volume=floci-aws
  ```

- **Your init scripts** in `.ddev/floci-aws/init/`, if you wrote any. The
  directory is removed only when it still looks untouched.
- **`.ddev/.env.floci-aws`**, so your settings survive a reinstall.

If you skipped the `reset` and the add-on is already gone, the leftover
containers Floci started can be cleared by hand. Keep the network filter — the
label alone matches every project's Floci containers, not just this one's:

```bash
docker rm -f $(docker ps -aq \
  --filter label=floci_emulator=floci-aws \
  --filter network=ddev-<project>_default)
```

## Licensing

Two different licences are in play here, and it is worth being explicit about
which covers what:

- **This add-on** — the compose files, the `ddev floci-aws` command, the init
  scripts and the documentation — is licensed **Apache 2.0**, matching the DDEV
  add-on template and DDEV's own add-ons. See [LICENSE](https://github.com/codementality/ddev-floci-aws/blob/main/LICENSE).
- **Floci itself** is licensed **MIT** by the
  [floci-io](https://github.com/floci-io/floci) project, and is entirely separate work.

**No Floci source is vendored in this repository.** The add-on only references
the published `floci/floci` image, which Docker pulls at runtime; nothing
here redistributes Floci code or binaries. If you redistribute this add-on you
are redistributing Apache-2.0 material only.

## Credits

Floci is by the [floci-io](https://github.com/floci-io/floci) project, MIT licensed. This add-on packages
it under Apache 2.0; see [Licensing](#licensing) above.

Behavioural notes in this README were verified against `floci/floci:1.7.0` on 2026-08-26.
Upstream ships frequently; the nightly workflow in `.github/workflows/tests.yml`
re-runs the suite against `latest` every day, which is what catches drift.

**Contributed and maintained by [@codementality](https://github.com/codementality)**
