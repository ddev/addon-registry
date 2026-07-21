---
title: RandalVanheede/secretvault-ddev
github_url: https://github.com/RandalVanheede/secretvault-ddev
description: ""
user: RandalVanheede
repo: secretvault-ddev
repo_id: 1245390793
default_branch: main
tag_name: v0.1.6
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2026-05-21
updated_at: 2026-05-22
workflow_status: unknown
stars: 0
---

# ddev-secret-vault

A DDEV addon that provides a centralized, encrypted local vault for project secrets. Secrets are stored **outside your project directory** (at `~/.ddev/secret-vault/`) so they can never be accidentally committed to git, and are automatically injected into `.ddev/.env` when DDEV starts.

---

## Requirements

- DDEV v1.22+
- `openssl` (standard on macOS and most Linux distros)
- `python3` (standard on macOS and most Linux distros)
- macOS or Linux (including WSL)

## Installation

```bash
ddev add-on get RandalVanheede/secretvault-ddev
```

---

## Quick Start

```bash
# 1. Initialize a vault for your project
ddev vault init

# 2. Import existing secrets automatically
ddev vault import

# 3. Or add secrets manually
ddev vault set STRIPE_SECRET_KEY sk_test_abc123

# 4. Start DDEV — secrets are injected automatically
ddev start
```

---

## Commands

### Core

| Command | Description |
|:---|:---|
| `ddev vault init` | Initialize an encrypted vault for this project |
| `ddev vault set KEY VALUE` | Store or update a secret |
| `ddev vault get KEY` | Print a secret's value |
| `ddev vault list` | List all secret keys |
| `ddev vault list --values` | List keys with masked values |
| `ddev vault remove KEY` | Delete a secret |
| `ddev vault inject` | Write secrets to `.ddev/.env` (also runs on `ddev start`) |

### Import

| Command | Description |
|:---|:---|
| `ddev vault import` | Auto-detect `.env` and `settings.local.php` and import |
| `ddev vault import --from=.env` | Import from a specific file |
| `ddev vault import --from=web/sites/default/settings.local.php` | Import from a PHP file |
| `ddev vault import --clean` | Import AND sanitize source files (backup created first) |
| `ddev vault import --skip-hash-salt` | Import without extracting `DRUPAL_HASH_SALT` (one-off) |
| `ddev vault config skip-hash-salt true` | Persist skip-hash-salt preference for all future imports |
| `ddev vault config skip-hash-salt false` | Re-enable hash salt import |
| `ddev vault config skip-hash-salt` | Show current skip-hash-salt setting |

### Utility

| Command | Description |
|:---|:---|
| `ddev vault export` | Print all secrets as `.env` format |
| `ddev vault export secrets.env` | Export to a file |
| `ddev vault status` | Show vault info and injection status |
| `ddev vault change-password` | Change the master password |

---

## How It Works

### Storage

Vaults live at `~/.ddev/secret-vault/<project-name>.vault` — outside your project, unreachable by git.

```
~/.ddev/secret-vault/
├── my-drupal-site.vault   # AES-256-CBC encrypted
└── other-project.vault
```

### Encryption

Secrets are encrypted with `openssl aes-256-cbc -pbkdf2`. On macOS, the master password is offered to be saved in the macOS Keychain so you aren't prompted on every `ddev start`. On Linux, `secret-tool` (libsecret) is used if available.

### Automatic Injection

A `pre-start` hook runs `ddev vault inject --quiet` before containers start. Secrets are written to a clearly-marked block in `.ddev/.env`:

```bash
# --- secret-vault: BEGIN ---
DB_PASSWORD="super_secret"
DRUPAL_HASH_SALT="abc123..."
STRIPE_SECRET_KEY="sk_test_..."
# --- secret-vault: END ---
```

The block is replaced on every inject — no stale secrets accumulate.

### `.gitignore` Safety

`ddev vault init` automatically adds `.ddev/.env` and `.ddev/.env.*` to your project's `.gitignore`.

---

## Importing from Drupal Projects

### From `.env`

```bash
ddev vault import --from=.env
```

Parses `KEY=VALUE` lines, strips quotes, skips comments.

### From `settings.local.php`

```bash
ddev vault import --from=web/sites/default/settings.local.php
```

Extracts:
- `$settings['hash_salt']` → `DRUPAL_HASH_SALT`
- `$databases[...]['password']` → `DB_PASSWORD`
- `$settings['smtp_password']` → `SMTP_PASSWORD`
- Google Maps, Stripe, Mailgun, SendGrid, reCAPTCHA keys
- Any `$settings` key matching common sensitive patterns (`api_key`, `auth_token`, etc.)

When DDEV is running, the PHP-based extractor (`extract-secrets.php`) runs inside the container for more accurate extraction. Otherwise, a pure-Python regex extractor is used as fallback.

### `--clean`: Sanitize Source Files

```bash
ddev vault import --from=web/sites/default/settings.local.php --clean
```

After importing, shows a diff and (with confirmation) replaces hardcoded values:

```diff
-$settings['hash_salt'] = 'abc123secrethash';
+$settings['hash_salt'] = getenv('DRUPAL_HASH_SALT');

-  'password' => 'super_secret_password',
+  'password' => getenv('DB_PASSWORD'),
```

A `.bak` backup is always created before any modifications.

### `--skip-hash-salt`: Exclude the Drupal Hash Salt

The Drupal hash salt is extracted by default because it can be exploited if leaked. If you manage it separately (e.g. via a different secrets manager) you can opt out:

**One-off:**
```bash
ddev vault import --skip-hash-salt
```

**Persisted for all future imports on this project:**
```bash
ddev vault config skip-hash-salt true
```

When `skip-hash-salt` is active and `DRUPAL_HASH_SALT` already exists in the vault, you will be prompted whether to remove it.

---

## Vault File Format (decrypted)

```json
{
  "version": 1,
  "project": "my-drupal-site",
  "created": "2026-05-20T13:30:00Z",
  "updated": "2026-05-20T14:00:00Z",
  "secrets": {
    "DRUPAL_HASH_SALT": "abc123...",
    "SMTP_PASSWORD": "hunter2"
  },
  "metadata": {
    "DRUPAL_HASH_SALT": {
      "imported_from": "settings.local.php",
      "imported_at": "2026-05-20T13:30:00Z",
      "updated_at": "2026-05-20T13:30:00Z"
    }
  }
}
```

---

## File Structure

```
ddev-secret-vault/
├── install.yaml                           # DDEV addon manifest
├── config.secret-vault.yaml              # pre-start hook
├── commands/host/vault                   # Main host command
├── secret-vault-helpers/
│   ├── crypto.sh                         # Encryption / keychain helpers
│   ├── ui.sh                             # Colors, prompts, diff output
│   ├── import-env.sh                     # .env import logic
│   ├── import-settings-local.sh         # settings.local.php import logic
│   └── extract-secrets.php              # PHP extractor (runs in container)
└── tests/
    ├── test-import-env.sh
    └── test-import-settings.sh
```

---

## Running Tests

```bash
bash tests/test-import-env.sh
bash tests/test-import-settings.sh
```

Tests run entirely without DDEV — they create temporary vaults, import from fixture files, and assert on extracted values.

---

## Security Notes

- Encryption is `AES-256-CBC` with PBKDF2 (100,000 iterations) — adequate for local development secrets
- Vault files are `chmod 600`; the vault directory is `chmod 700`
- The master password is never written to disk by this tool (only optionally stored in OS keychain)
- This is designed to prevent casual exposure (filesystem browsing, backup services) — not adversarial attacks on a compromised machine
