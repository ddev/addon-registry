---
title: "wazum/ddev-maildev"
github_url: "https://github.com/wazum/ddev-maildev"
description: "DDEV add-on for MailDev (with MCP)"
user: "wazum"
repo: "ddev-maildev"
repo_id: 1374260028
default_branch: "main"
tag_name: "v1.0.0"
ddev_version_constraint: ">= v1.25.4"
dependencies: []
type: "contrib"
created_at: "2026-09-17"
updated_at: "2026-09-17"
workflow_status: "disabled"
stars: 0
---

<h1 align="center">ddev-maildev</h1>
<p align="center"><em>MailDev for DDEV. Your coding agent reads the dev inbox over MCP.</em></p>
<br>

<p align="center">
  <a href="https://addons.ddev.com"><img src="https://img.shields.io/badge/DDEV-Add--on_Registry-a8d8ea?style=for-the-badge&labelColor=24273a" alt="add-on registry"></a>
  <a href="https://github.com/wazum/ddev-maildev/actions/workflows/tests.yml?query=branch%3Amain"><img src="https://img.shields.io/github/actions/workflow/status/wazum/ddev-maildev/tests.yml?branch=main&style=for-the-badge&logo=githubactions&logoColor=white&label=tests&labelColor=24273a" alt="tests"></a>
  <a href="https://github.com/wazum/ddev-maildev/releases/latest"><img src="https://img.shields.io/github/v/release/wazum/ddev-maildev?style=for-the-badge&labelColor=24273a&color=ffdac1" alt="release"></a>
  <a href="https://github.com/wazum/ddev-maildev/commits"><img src="https://img.shields.io/github/last-commit/wazum/ddev-maildev?style=for-the-badge&labelColor=24273a&color=e2f0cb" alt="last commit"></a>
  <br>
  <a href="https://ddev.com"><img src="https://img.shields.io/badge/DDEV-v1.25.4%2B-b5ead7?style=for-the-badge&labelColor=24273a&logo=data:image/svg%2Bxml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAyNCAyNCI%2BPHBhdGggZmlsbD0iI2ZmZiIgZD0iTTEyIDIgMiA3djEwbDEwIDUgMTAtNVY3em0wIDIuMyA3LjUgMy43TDEyIDExLjcgNC41IDh6TTQgOS42bDcgMy41djdMNCAxNi42em05IDEwLjV2LTdsNy0zLjV2N3oiLz48L3N2Zz4%3D" alt="DDEV v1.25.4 or newer"></a>
  <a href="https://github.com/maildev/maildev"><img src="https://img.shields.io/badge/MailDev-3.0.0--rc.3-ffb997?style=for-the-badge&labelColor=24273a" alt="MailDev 3.0.0-rc.3"></a>
  <a href="https://modelcontextprotocol.io"><img src="https://img.shields.io/badge/MCP-enabled-c3b1e1?style=for-the-badge&logo=modelcontextprotocol&logoColor=white&labelColor=24273a" alt="MCP enabled"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/licence-Apache--2.0-ffc6d9?style=for-the-badge&logo=apache&logoColor=white&labelColor=24273a" alt="Apache-2.0 licence"></a>
</p>

## The idea

[DDEV](https://ddev.com) already catches mail with
[Mailpit](https://mailpit.axllent.org). Catching mail is not the point.

The point is that you cannot ask Mailpit anything. A test sends a password reset, you
want the link, so you open a browser and copy it out by hand.
**[MailDev](https://github.com/maildev/maildev) 3 answers over the
[Model Context Protocol](https://modelcontextprotocol.io), so your agent can search the
inbox and read the message itself.**

## Install

```bash
ddev add-on get wazum/ddev-maildev
ddev restart
```

That is all. PHP mail now goes to MailDev, and the `maildev` server is in your
`.mcp.json`. Claude Code reads that file on its next start and asks you to approve the
server once. Other MCP clients use their own config file, so copy the entry there.

```bash
ddev maildev     # open the inbox
```

The inbox is at `https://<project>.ddev.site:1081`. The login is in
`.ddev/.env.maildev.local`.

## Mail your framework sends itself

PHP's `mail()` is redirected for you. Frameworks that speak SMTP themselves never use
`sendmail_path`, so they need one line. The web container has `MAILDEV_SMTP_HOST` and
`MAILDEV_SMTP_PORT`, so nothing is hardcoded:

| Framework | Where | What to set |
|---|---|---|
| Symfony | `.env.local` | `MAILER_DSN=smtp://${MAILDEV_SMTP_HOST}:${MAILDEV_SMTP_PORT}` |
| Laravel | `.env` | `MAIL_MAILER=smtp`, `MAIL_HOST=${MAILDEV_SMTP_HOST}`, `MAIL_PORT=${MAILDEV_SMTP_PORT}` |
| TYPO3 | `AdditionalConfiguration.php` | `$GLOBALS['TYPO3_CONF_VARS']['MAIL']['transport'] = 'smtp';`<br>`…['transport_smtp_server'] = getenv('MAILDEV_SMTP_HOST') . ':' . getenv('MAILDEV_SMTP_PORT');` |

Turn encryption and credentials off for these. Restart the workers or clear the config
cache if your framework caches them.

## Why the inbox has a password

MailDev has no authentication by default. The add-on generates a password into
`.ddev/.env.maildev.local`, which DDEV's own `.gitignore` covers, and both the API and
`/mcp` then ask for it. Without a password the installation stops.

**The password also goes into `.mcp.json`, as an `Authorization` header.** A
`.mcp.json` the add-on creates is readable only by you. One you already had keeps your
permissions. People commit this file, so the installation warns you if Git tracks it.
Untrack it, or keep the entry in your own MCP config instead.

## How it works

MailDev listens on 1025 for SMTP and 1080 for the web interface. The DDEV router
publishes the interface on 1081 over HTTPS. SMTP stays inside Docker.

`.ddev/php/maildev.ini` changes `sendmail_path` to relay through the `mailpit` binary
that is already in the web image, so nothing new is installed. It uses the container
name, not the service name: DDEV puts every service on one shared network, where a
plain `maildev` is ambiguous as soon as a second project installs this add-on.

Mail lives in a Docker volume and survives `ddev restart`.

## What it does to your files

| File | What happens |
|---|---|
| `.mcp.json` | the `maildev` server is merged in, your other servers are left alone |
| `.ddev/php/maildev.ini` | generated, and skipped on reinstall if you remove its `#ddev-generated` marker |
| `.ddev/.env.web.maildev` | the SMTP host and port for the web container |
| `.ddev/.env.maildev.local` | the password, and your `MAILDEV_IMAGE` override if you set one |

The add-on records what it wrote in `.ddev/maildev/mcp-state.json`. Removal takes out
that entry only, and only while it still matches. An entry you edited stays, and so
does a `.mcp.json` that was there before you installed.

## Known limits

**The pin is a release candidate.** MailDev 3 is a rewrite and the only line with an
MCP server, so `3.0.0-rc.3` is what you get. Override it with `MAILDEV_IMAGE`.

**Mailpit keeps running and stays empty.** It cannot be removed from the web container,
so `ddev describe` lists both.

**Links in MCP replies point at `http://localhost:1080`.** MailDev builds them from its
own bind address and has no setting for the public URL, so they do not open. The message
content is correct.

**Two writes are not one transaction.** If writing `.mcp.json` fails after the state was
recorded, a reinstall or a removal recovers. Someone editing the file during that moment
is not detected.

**Anything on your machine can post to the inbox.** DDEV puts every service on one
shared Docker network, so another project's container can reach this SMTP port, and
MailDev accepts mail without a password, as mail servers do. Treat what comes out of the
inbox as data, not as instructions.

## Removing it

```bash
ddev add-on remove maildev
ddev restart
```

This removes the container, the volume, the generated files and the MCP entry. The
restart sends PHP mail back to Mailpit. `.ddev/.env.maildev.local` stays, in case you
put your own settings in it. Framework SMTP settings are yours to undo.

## Licence

[Apache-2.0](https://github.com/wazum/ddev-maildev/blob/main/LICENSE)
