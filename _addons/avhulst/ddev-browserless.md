---
title: avhulst/ddev-browserless
github_url: https://github.com/avhulst/ddev-browserless
description: "DDEV add-on for browserless v2 (Chromium): PDFs, screenshots, scraping and browser automation on demand. Ships a Claude Code skill for driving it."
user: avhulst
repo: ddev-browserless
repo_id: 1319337407
default_branch: main
tag_name: v1.1.0
ddev_version_constraint: ">= v1.25.0"
dependencies: []
type: contrib
created_at: 2026-08-01
updated_at: 2026-08-03
workflow_status: success
stars: 0
---

# DDEV Browserless Add-on

A DDEV add-on that provides a [browserless](https://github.com/browserless/browserless)
**v2** service (`ghcr.io/browserless/chromium`) for headless Chromium automation:
PDF generation, screenshots, scraping and full browser automation over HTTP or
WebSocket.

The service runs behind a Docker Compose profile, so it consumes nothing until
you start it with `ddev browserless on`.

> **Coming from browserless v1?** v2 changed the interface in ways that break
> every v1 example:
>
> - Endpoints are **POST-only** by default (`ALLOW_GET=false`).
>   `GET /pdf?url=…` returns **404**.
> - Routes are browser-prefixed: `/chromium/pdf`, `/chromium/screenshot`, …
> - **Selenium/WebDriver support was removed.** `POST /webdriver` returns 404,
>   so Symfony Panther and `php-webdriver` cannot drive this service.

## Features

- On-demand Chromium instances, started only when you need them
- PDF generation from a URL or from raw HTML
- Screenshot capture, full-page or viewport
- Structured scraping via CSS selectors
- Server-side Puppeteer scripts via `/chromium/function`
- CDP and Playwright WebSocket endpoints
- DDEV-integrated via a `ddev browserless` command

## Installation

```bash
ddev add-on get avhulst/ddev-browserless
ddev restart
ddev browserless on
```

Pin a release, or install from a local checkout:

```bash
ddev add-on get avhulst/ddev-browserless --version v1.0.0
ddev add-on get /path/to/ddev-browserless
```

## Usage

```bash
ddev browserless on        # start the service
ddev browserless off       # stop it
ddev browserless restart   # stop and start again
ddev browserless status    # running? plus the URLs
ddev browserless url       # print the host-side base URL
ddev browserless meta      # browserless / Chromium / Playwright / Puppeteer versions
ddev browserless logs -f   # follow the service logs
```

## Access

The service is only reachable while it is running (`ddev browserless on`).

| From | Base URL | Env var |
|------|----------|---------|
| Web container (PHP, Node.js, `ddev ssh`) | `http://browserless:3000` | `BROWSERLESS_URL` |
| Web container, CDP WebSocket | `ws://browserless:3000/chromium` | `BROWSERLESS_WS_URL` |
| Your host (browser, host-side `curl`) | `http://<project>.ddev.site:3000` | `BROWSERLESS_HOST_URL` |
| Your host, TLS | `https://<project>.ddev.site:3001` | — |

- **API reference (OpenAPI/Redoc):** `http://<project>.ddev.site:3000/docs/` — no token required
- **Live debugger UI:** `http://<project>.ddev.site:3000/debugger/?token=<BROWSERLESS_TOKEN>`
- `GET /` itself returns 404 — there is no landing page.

> **Port note.** `HTTP_EXPOSE=3000:3000` binds router port 3000 project-wide.
> Two DDEV projects running this add-on at the same time will collide on that
> port; start only one of them at a time.

## Configuration

Create `.ddev/.env.browserless` and restart:

```bash
ddev dotenv set .ddev/.env.browserless --browserless-token "$(openssl rand -hex 32)"
ddev restart
```

Add `/.ddev/.env.browserless` to your project's **root** `.gitignore` —
`.ddev/.gitignore` is `#ddev-generated` and DDEV overwrites it.

### Service variables

| Variable | Default | Maps to | Description |
|----------|---------|---------|-------------|
| `BROWSERLESS_TOKEN` | `ddev-browserless-token` | `TOKEN` | API authentication token |
| `BROWSERLESS_TAG` | `latest` | image tag | Pin the image, e.g. `2.55.2` |
| `BROWSERLESS_TIMEOUT` | `120000` | `TIMEOUT` | Max session runtime in ms |
| `BROWSERLESS_CONCURRENT` | `2` | `CONCURRENT` | Concurrent browser sessions |
| `BROWSERLESS_QUEUED` | `10` | `QUEUED` | Requests queued beyond `CONCURRENT` |
| `BROWSERLESS_MAX_PAYLOAD_SIZE` | `20971520` | `MAX_PAYLOAD_SIZE` | Max request body in bytes |
| `BROWSERLESS_ALLOW_GET` | `false` | `ALLOW_GET` | Enable v1-style GET calls |

Per-request Chromium flags are **not** configured here — pass them as a `launch`
query parameter (see below).

### Web container variables

`config.browserless.yaml` exports `BROWSERLESS_TOKEN`, `BROWSERLESS_URL`,
`BROWSERLESS_WS_URL` and `BROWSERLESS_HOST_URL` into the web container.

> If your project also has a `.ddev/config.token.yaml` (or any other config file
> sorting after `config.browserless.yaml`) that sets `BROWSERLESS_TOKEN`, that
> value wins and this add-on's default never applies.

## Examples

All examples below were verified against `ghcr.io/browserless/chromium` **2.55.2**.
Run the `curl` examples inside the web container (`ddev ssh`), where
`$BROWSERLESS_URL` and `$BROWSERLESS_TOKEN` are already set.

Authentication works either as an `Authorization: Bearer` header (used below —
it keeps the token out of URLs and shell history) or as a `?token=` query
parameter (needed for WebSockets and for URLs you open in a browser).

### Rendering *your own* DDEV site

The browserless container does not trust DDEV's mkcert CA, so loading
`https://<project>.ddev.site` fails with `ERR_CERT_AUTHORITY_INVALID` and the API
returns **HTTP 500**. Pass `launch` as a **query parameter** — it is rejected
inside the JSON body (400):

```bash
curl -sS -X POST \
  "${BROWSERLESS_URL}/chromium/pdf?launch=%7B%22ignoreHTTPSErrors%22%3Atrue%7D" \
  -H "Authorization: Bearer ${BROWSERLESS_TOKEN}" \
  -H "Content-Type: application/json" \
  -d "{\"url\":\"${DDEV_PRIMARY_URL}/\",\"options\":{\"format\":\"A4\",\"printBackground\":true}}" \
  -o site.pdf
```

That URL-encoded blob is `{"ignoreHTTPSErrors":true}`. The browser reaches your
site through the DDEV router — that is what the
`external_links: ["ddev-router:${DDEV_SITENAME}.${DDEV_TLD}"]` entry in
`docker-compose.browserless.yaml` is for.

### HTTP API (cURL)

```bash
# Screenshot
curl -sS -X POST "${BROWSERLESS_URL}/chromium/screenshot" \
  -H "Authorization: Bearer ${BROWSERLESS_TOKEN}" \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com","options":{"fullPage":true,"type":"png"}}' \
  -o screenshot.png

# PDF from a URL
curl -sS -X POST "${BROWSERLESS_URL}/chromium/pdf" \
  -H "Authorization: Bearer ${BROWSERLESS_TOKEN}" \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com","options":{"format":"A4","printBackground":true,"margin":{"top":"1cm","bottom":"1cm"}}}' \
  -o example.pdf

# PDF from an HTML string (no URL needed)
curl -sS -X POST "${BROWSERLESS_URL}/chromium/pdf" \
  -H "Authorization: Bearer ${BROWSERLESS_TOKEN}" \
  -H "Content-Type: application/json" \
  -d '{"html":"<h1>Hello DDEV</h1>","options":{"format":"A4"}}' \
  -o hello.pdf

# Fully rendered HTML
curl -sS -X POST "${BROWSERLESS_URL}/chromium/content" \
  -H "Authorization: Bearer ${BROWSERLESS_TOKEN}" \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com","gotoOptions":{"waitUntil":"networkidle2"}}'

# Structured scraping — "elements" is REQUIRED (without it: HTTP 400)
curl -sS -X POST "${BROWSERLESS_URL}/chromium/scrape" \
  -H "Authorization: Bearer ${BROWSERLESS_TOKEN}" \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com","elements":[{"selector":"h1"},{"selector":"a"}]}'
```

### PHP (Symfony HttpClient)

```php
<?php
use Symfony\Component\HttpClient\HttpClient;

$client = HttpClient::create([
    'base_uri' => rtrim(getenv('BROWSERLESS_URL'), '/') . '/',
    'headers'  => ['Authorization' => 'Bearer ' . getenv('BROWSERLESS_TOKEN')],
    'timeout'  => 120,
]);

// PDF of a public page — POST, not GET
$pdf = $client->request('POST', 'chromium/pdf', [
    'json' => [
        'url'     => 'https://example.com',
        'options' => ['format' => 'A4', 'printBackground' => true],
    ],
])->getContent();
file_put_contents('example.pdf', $pdf);

// PDF of your own DDEV site: ignoreHTTPSErrors must go in the QUERY STRING
$pdf = $client->request('POST', 'chromium/pdf', [
    'query' => ['launch' => json_encode(['ignoreHTTPSErrors' => true])],
    'json'  => [
        'url'     => getenv('DDEV_PRIMARY_URL') . '/',
        'options' => ['format' => 'A4', 'printBackground' => true],
    ],
])->getContent();
file_put_contents('site.pdf', $pdf);

// Screenshot
$png = $client->request('POST', 'chromium/screenshot', [
    'json' => ['url' => 'https://example.com', 'options' => ['fullPage' => true, 'type' => 'png']],
])->getContent();
file_put_contents('screenshot.png', $png);

// Scraping — "elements" is required; the result is {"data": [...]}
$data = $client->request('POST', 'chromium/scrape', [
    'json' => ['url' => 'https://example.com', 'elements' => [['selector' => 'h1']]],
])->toArray();
```

In a Contao/Symfony application, inject `HttpClientInterface` instead of calling
`HttpClient::create()`. Guzzle works the same way — `request('POST', …)` becomes
`post(…)` and `getContent()` becomes `getBody()->getContents()`.

### PHP: full browser automation via `/chromium/function`

**Symfony Panther and `php-webdriver` do not work with browserless v2** —
Selenium/WebDriver support was removed and `POST /webdriver` returns 404. There
is no maintained PHP CDP client either. Instead, send a Puppeteer script to the
service and let it execute there:

```php
<?php
use Symfony\Component\HttpClient\HttpClient;

$script = <<<'JS'
export default async ({ page }) => {
  await page.goto('https://example.com', { waitUntil: 'networkidle2' });
  await page.type('#search', 'ddev');
  await page.click('button[type=submit]');
  await page.waitForSelector('.results');
  return {
    data: { title: await page.title(), url: page.url() },
    type: 'application/json',
  };
};
JS;

$result = HttpClient::create()
    ->request('POST', getenv('BROWSERLESS_URL') . '/chromium/function', [
        'query'   => ['launch' => json_encode(['ignoreHTTPSErrors' => true])],
        'headers' => [
            'Authorization' => 'Bearer ' . getenv('BROWSERLESS_TOKEN'),
            'Content-Type'  => 'application/javascript',
        ],
        'body'    => $script,
        'timeout' => 120,
    ])
    ->toArray();
```

The handler receives `{ page, browser, context }` and must return `{ data, type }`.
`type` becomes the response `Content-Type`, so you can also return `image/png` or
`application/pdf` buffers.

### Node.js: Puppeteer (CDP) or Playwright

```javascript
// Puppeteer — CDP endpoint
import puppeteer from 'puppeteer-core';

const browser = await puppeteer.connect({
  browserWSEndpoint:
    `${process.env.BROWSERLESS_WS_URL}?token=${process.env.BROWSERLESS_TOKEN}` +
    `&launch=${encodeURIComponent(JSON.stringify({ ignoreHTTPSErrors: true }))}`,
});
const page = await browser.newPage();
await page.goto(`${process.env.DDEV_PRIMARY_URL}/`);
await page.screenshot({ path: 'site.png' });
await browser.close();
```

```javascript
// Playwright — dedicated Playwright endpoint, NOT the CDP one
import { chromium } from 'playwright-core';

const browser = await chromium.connect(
  `ws://browserless:3000/chromium/playwright?token=${process.env.BROWSERLESS_TOKEN}`,
);
const page = await browser.newPage({ ignoreHTTPSErrors: true });
await page.goto('https://example.com');
await page.pdf({ path: 'example.pdf', format: 'A4' });
await browser.close();
```

> **Version pinning.** `chromium.connect()` fails unless your Playwright version
> is one the image supports. Ask the service instead of guessing:
>
> ```bash
> ddev browserless meta
> # {"version":"2.55.2","chromium":"151.0.7922.34","firefox":null,"webkit":null,
> #  "playwright":["1.62.0","1.61.1","1.60.0","1.59.1","1.58.2"],
> #  "puppeteer":["25.3.0"]}
> ```
>
> Pin `playwright-core` to one of the listed versions. This image is
> **chromium-only** — `/firefox/playwright` and `/webkit/playwright` return 404.
> Puppeteer speaks plain CDP and is far more version-tolerant.

## Claude Code plugin

This repository is also a [Claude Code](https://claude.com/claude-code)
marketplace. It ships one plugin, `ddev-browserless`, containing the
`browserless-ddev` skill: it teaches the agent to drive this service — verified
recipes for screenshots, PDFs, rendered HTML, scraping and Lighthouse audits,
plus the failure modes that are easy to hit blind (a missing `Content-Type`
yielding 404, a Lighthouse audit failing with 200, your own site needing
`ignoreHTTPSErrors`).

```bash
/plugin marketplace add avhulst/ddev-browserless
/plugin install ddev-browserless@ddev-browserless
```

The plugin and the add-on are released together and share this repository's
version tags, so a matching pair is whatever you last installed from the same
release.

**The skill needs the add-on.** It is not self-contained and assumes:

- the add-on installed and the service started (`ddev browserless on`);
- the agent running **inside the web container**, where `$BROWSERLESS_URL` and
  `$BROWSERLESS_TOKEN` are set — it deliberately cannot start the service, since
  there is no Docker socket there and DDEV's `ddev` shim refuses management
  commands;
- `curl`, `jq` and `node` in that container (DDEV's web image ships all three);
- a git-ignored, project-local directory for output — `var/browserless/` by
  default.

Installing the plugin without the add-on is harmless: the skill's preflight
detects the unset `$BROWSERLESS_URL` and says what to install.

Working on the skill itself: see
[`skills/browserless-ddev/README.md`](https://github.com/avhulst/ddev-browserless/blob/main/skills/browserless-ddev/README.md).

## Removing the Add-on

```bash
ddev add-on remove browserless
ddev restart
```

Removing the add-on leaves the Claude Code plugin in place — it is installed
through a separate channel:

```bash
/plugin uninstall ddev-browserless@ddev-browserless
```

## Requirements

- DDEV v1.25.0 or higher
- Docker running

## Troubleshooting

### HTTP 500 when rendering your own site

`ERR_CERT_AUTHORITY_INVALID` — the container does not trust DDEV's mkcert CA.
Add `?launch=%7B%22ignoreHTTPSErrors%22%3Atrue%7D` to the request URL. Note this
must be a query parameter; a `launch` key in the JSON body returns 400.

### HTTP 404 on `GET /pdf?url=…`

v2 is POST-only. Use `POST /chromium/pdf` with a JSON body. If you genuinely need
GET, set `BROWSERLESS_ALLOW_GET=true` in `.ddev/.env.browserless` and restart —
then verify with `curl -s "$BROWSERLESS_URL/config?token=$BROWSERLESS_TOKEN"`.

### HTTP 400 from `/chromium/scrape`

The `elements` array is required: `{"url":"…","elements":[{"selector":"h1"}]}`.

### HTTP 401

No token, or the wrong one. Every route except `/docs/` requires it.

### Service won't start

```bash
ddev browserless logs
ddev utility compose-config | grep -A20 browserless
```

### Connection refused right after starting

The container needs a few seconds. It has a healthcheck against `/docs/`; check
with `ddev browserless status`.

## License

This add-on is licensed under the MIT License. The browserless image has its own
[license](https://github.com/browserless/browserless/blob/main/LICENSE).

## Acknowledgments

- [DDEV](https://ddev.com) — the local development environment
- [browserless](https://github.com/browserless/browserless) — the headless browser service
