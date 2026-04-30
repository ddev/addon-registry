---
title: xima-media/ddev-playwright
github_url: https://github.com/xima-media/ddev-playwright
description: "Playwright add-on for ddev"
user: xima-media
repo: ddev-playwright
repo_id: 1223285826
default_branch: main
tag_name: 3.0.4
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: contrib
created_at: 2026-04-28
updated_at: 2026-04-29
workflow_status: unknown
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/xima-media/ddev-playwright/actions/workflows/tests.yml/badge.svg)](https://github.com/xima-media/ddev-playwright/actions/workflows/tests.yml)
[![last commit](https://img.shields.io/github/last-commit/xima-media/ddev-playwright)](https://github.com/xima-media/ddev-playwright/commits)
[![Version](https://img.shields.io/github/v/release/xima-media/ddev-playwright)](https://github.com/xima-media/ddev-playwright/releases)

# ddev-playwright

This fork uses the [xima-media/playwright](https://github.com/xima-media/typo3-docker/pkgs/container/playwright) docker image, which bundles the playwright dependencies.

## Installation

```bash
ddev add-on get xima-media/ddev-playwright
```

## Quickstart

Generate example `playwright.config.js` in your project root directory with this command:

```bash
ddev playwright-init
```

(@TODO #2)

## Usage

* `ddev playwright test`
* `ddev playwright show-report`

## Extras

This add-on installs `mariadb-client` into the ddev image to populate test fixtures from inside the container. Configure [globalSetup](https://playwright.dev/docs/test-global-setup-teardown#option-2-configure-globalsetup-and-globalteardown) to truncate and load sql files + copy files.

```ts
// playwright.config.js
export default defineConfig({
  globalSetup: "./Tests/Playwright/global-setup.ts",
  ...
})
```

Example setup
 * Truncate + import sql in the format `Tests/Fixtures/<tablename>.sql`.
 * Copy directory `Tests/Fixtures/Files` to a fixed path.

```ts
// global-setup.ts
import {execSync} from 'child_process';
import * as path from 'path';
import * as fs from 'fs';

const DB_HOST = 'db';
const DB_USER = 'db';
const DB_PASS = 'db';
const DB_NAME = 'db';

const FIXTURE_PATH = '/var/www/html/Tests/Fixtures';

function mysql(sql: string): void {
  execSync(`mysql -h${DB_HOST} -u${DB_USER} -p${DB_PASS} ${DB_NAME}`, {input: sql, stdio: ['pipe', 'inherit', 'inherit']});
}

function mysqlFile(filePath: string): void {
  execSync(`mysql -h${DB_HOST} -u${DB_USER} -p${DB_PASS} ${DB_NAME} < "${filePath}"`, {shell: '/bin/bash', stdio: 'inherit'});
}

export default async function globalSetup(): Promise<void> {
  console.log('Setting up test fixtures...');

  const tables = fs.readdirSync(FIXTURE_PATH)
    .filter(f => f.endsWith('.sql'))
    .map(f => path.basename(f, '.sql'));

  const truncateStatements = tables.map(t => `TRUNCATE TABLE \`${t}\`;`).join(' ');
  mysql(`SET FOREIGN_KEY_CHECKS = 0; ${truncateStatements} SET FOREIGN_KEY_CHECKS = 1;`);

  for (const table of tables) {
    const file = path.join(FIXTURE_PATH, `${table}.sql`);
    console.log(`Importing ${table}.sql...`);
    mysqlFile(file);
  }

  const filesSource = path.join(FIXTURE_PATH, 'Files');
  const filesDest = '/var/www/html/public/fileadmin/Files';
  if (fs.existsSync(filesSource)) {
    execSync(`cp -r "${filesSource}" "${filesDest}"`, {stdio: 'inherit'});
  }

  console.log('Fixture setup complete.');
}
```
