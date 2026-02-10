---
title: sdubois/ddev-sms-inbox
github_url: https://github.com/sdubois/ddev-sms-inbox
description: "A Mailpit-like viewer for SMS messages. Catches and displays SMS messages sent by your application during development and testing, so you never accidentally send real texts."
user: sdubois
repo: ddev-sms-inbox
repo_id: 1153838207
default_branch: main
tag_name: 
ddev_version_constraint: ">= v1.24.10"
dependencies: []
type: contrib
created_at: 2026-02-09
updated_at: 2026-02-09
workflow_status: unknown
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)

# DDEV SMS Inbox

## Overview

A [Mailpit](https://mailpit.axigen.com/)-like viewer for SMS messages. Catches and displays SMS messages sent by your application during development and testing, so you never accidentally send real texts.

Your application writes messages to a JSON file instead of calling a real SMS provider (e.g. Twilio). SMS Inbox picks them up and displays them in a retro web UI.

## Installation

```bash
ddev add-on get sdubois/sms-inbox
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Configuration

Add the `SMS_DRY_RUN` environment variable to your DDEV project's `config.yaml` so your application can check whether to capture SMS locally or send via a live provider:

```yaml
web_environment:
  - SMS_DRY_RUN=1
```

Then have your application check for this variable before sending (see [Examples](#examples) below).

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev sms-inbox` | Open SMS Inbox in your browser |
| `ddev describe` | View service status and used ports |
| `ddev logs -s sms-inbox` | Check SMS Inbox logs |

## Writing Messages

Your application should write messages to `.ddev/sms-inbox/sms-inbox-data/messages.json` (which is `/mnt/sms-inbox-data/messages.json` inside the DDEV web container). Each message is an object in a JSON array:

```json
[
  {
    "id": "unique-id",
    "from": "+15551234567",
    "to": "+15559876543",
    "body": "Your verification code is 123456",
    "timestamp": 1707500000,
    "segments": 1,
    "sid": "optional-reference-id"
  }
]
```

| Field | Type | Description |
|-------|------|-------------|
| `id` | string | Unique message identifier |
| `from` | string | Sender phone number |
| `to` | string | Recipient phone number |
| `body` | string | SMS message text |
| `timestamp` | int | Unix timestamp |
| `segments` | int | Number of SMS segments |
| `sid` | string | Optional external reference ID |

## Examples

### PHP

```php
function sendSms(string $to, string $body, string $from = '+15551234567'): void
{
    // Capture locally when SMS_DRY_RUN is set, otherwise send via a live provider
    if (getenv('SMS_DRY_RUN')) {
        $dataFile = '/mnt/sms-inbox-data/messages.json';

        $messages = [];
        if (file_exists($dataFile)) {
            $messages = json_decode(file_get_contents($dataFile), true) ?: [];
        }

        $messages[] = [
            'id'        => bin2hex(random_bytes(16)),
            'from'      => $from,
            'to'        => $to,
            'body'      => $body,
            'timestamp' => time(),
            'segments'  => (int) ceil(mb_strlen($body) / 160),
            'sid'       => 'local_' . bin2hex(random_bytes(8)),
        ];

        file_put_contents($dataFile, json_encode($messages, JSON_PRETTY_PRINT));
        return;
    }

    // Send via your real SMS provider (e.g. Twilio) here...
}

// Usage
sendSms('+15559876543', 'Your verification code is 123456');
```

### Python

```python
import json
import os
import secrets
import math
import time

def send_sms(to: str, body: str, from_number: str = "+15551234567") -> None:
    # Capture locally when SMS_DRY_RUN is set, otherwise send via a live provider
    if os.environ.get("SMS_DRY_RUN"):
        data_file = "/mnt/sms-inbox-data/messages.json"

        messages = []
        if os.path.exists(data_file):
            with open(data_file) as f:
                messages = json.load(f)

        messages.append({
            "id": secrets.token_hex(16),
            "from": from_number,
            "to": to,
            "body": body,
            "timestamp": int(time.time()),
            "segments": math.ceil(len(body) / 160),
            "sid": f"local_{secrets.token_hex(8)}",
        })

        with open(data_file, "w") as f:
            json.dump(messages, f, indent=2)
        return

    # Send via your real SMS provider (e.g. Twilio) here...

# Usage
send_sms("+15559876543", "Your verification code is 123456")
```

## API

| Endpoint | Description |
|----------|-------------|
| `/` | Web UI |
| `/?action=api` | Returns all messages as JSON |
| `/?action=delete&id=ID` | Delete a single message |
| `/?action=clear` | Delete all messages |

## License

MIT
