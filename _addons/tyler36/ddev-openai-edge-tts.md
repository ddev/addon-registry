---
title: tyler36/ddev-openai-edge-tts
github_url: https://github.com/tyler36/ddev-openai-edge-tts
description: "An OpenAI-compatible text-to-speech (TTS) API service for DDEV"
user: tyler36
repo: ddev-openai-edge-tts
repo_id: 1048287216
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2025-09-01
updated_at: 2025-09-02
workflow_status: success
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/tyler36/ddev-openai-edge-tts/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/tyler36/ddev-openai-edge-tts/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/tyler36/ddev-openai-edge-tts)](https://github.com/tyler36/ddev-openai-edge-tts/commits)
[![release](https://img.shields.io/github/v/release/tyler36/ddev-openai-edge-tts)](https://github.com/tyler36/ddev-openai-edge-tts/releases/latest)

# DDEV Openai Edge TTS

## Overview

This add-on integrates Openai Edge TTS into your [DDEV](https://ddev.com/) project.

[Openai Edge TTS](https://github.com/travisvn/openai-edge-tts) provides a local, OpenAI-compatible text-to-speech (TTS) API using edge-tts.
edge-tts uses Microsoft Edge's online text-to-speech service, so it is completely free.

## Installation

```bash
ddev add-on get tyler36/ddev-openai-edge-tts
ddev restart
```

After installation, make sure to commit the `.ddev` directory to version control.

## Features

This addon is a helper wrapper for <https://github.com/travisvn/openai-edge-tts>.
The following features are stated as available:

- OpenAI-Compatible Endpoint: `/v1/audio/speech` with similar request structure and behavior.
- SSE Streaming Support: Real-time audio streaming via Server-Sent Events when stream_format: "sse" is specified.
- Supported Voices: Maps OpenAI voices (`alloy`, `echo`, `fable`, `onyx`, `nova`, `shimmer`) to edge-tts equivalents.
- Flexible Formats: Supports multiple audio formats (`mp3`, `opus`, `aac`, `flac`, `wav`, `pcm`).
- Adjustable Speed: Option to modify playback speed (0.25x to 4.0x).
- Optional Direct Edge-TTS Voice Selection: Use either OpenAI voice mappings or specify any edge-tts voice directly.

To preview edge-tts voices, see [Voice Samples](https://tts.travisvn.com/).

## Usage

This addon provides the service via an API endpoint.

| Access    | Endpoint                                    |
| --------- | ------------------------------------------- |
| Host      | https://{PROJECT_NAME}:5050/v1/audio/speech |
| Container | openai-edge-tts:5050/v1/audio/speech        |

OpenAI-Compatible Edge-TTS API tries to maintain compatibility with the OpenAI `audio/speech` endpoint.

@see [OpenAI-Compatible Edge-TTS API](https://github.com/travisvn/openai-edge-tts) for all available options.

For non-English languages, set the voice to an appropriate language list on [Voice Samples](https://tts.travisvn.com/).

### Curl

Use `curl` to make request to the endpoint.

Below, is an example of a command from the host that:

- transforms the "Hello, I am your AI assistant!" `input`
- using the `echo` voice
- into a `mp3` format
- at `1.1` speed (faster than normal)
- into a file called `speech.mp3`.

```shell
ddev exec curl -X POST openai-edge-tts:5050/v1/audio/speech \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer your_api_key_here" \
    -d '{
      "input": "Hello, I am your AI assistant!",
      "voice": "echo",
      "response_format": "mp3",
      "speed": 1.1
    }' \
    --output speech.mp3
```

> [!Note] "your_api_key_here" is the default API KEY and does not need to be changed for local development.

### OpenAI-php

[OpenAI-php](https://github.com/openai-php/client) can be used with this addon.

- Use the `factory()` function to generate a new `$client` that points to this addon.
- Then, use the `$client` as normal.

```php
    $client = \OpenAI::factory()
        ->withApiKey('your_api_key_here'))
        ->withBaseUri('http://openai-edge-tts:5050/v1/')
        ->make();

    // This is a standard OpenAI TTS API call.
    $mp3 = $client->audio()->speech([
        'input' => "What can I help you with today?",
        'voice' => 'echo',
        'response_format' => 'mp3',
    ]);
```

### Other Commands

| Command                        | Description                                            |
| ------------------------------ | ------------------------------------------------------ |
| `ddev describe`                | View service status and used ports for Openai Edge Tts |
| `ddev logs -s openai-edge-tts` | Check Openai Edge Tts logs                             |

## Advanced Customization

To change the Docker image:

```bash
ddev dotenv set .ddev/.env.openai-edge-tts --openai-edge-tts-docker-image="ddev/ddev-utilities:latest"
ddev add-on get tyler36/ddev-openai-edge-tts
ddev restart
```

Make sure to commit the `.ddev/.env.openai-edge-tts` file to version control.

All customization options (use with caution):

| Variable                       | Flag                             | Default                      |
| ------------------------------ | -------------------------------- | ---------------------------- |
| `OPENAI_EDGE_TTS_DOCKER_IMAGE` | `--openai-edge-tts-docker-image` | `ddev/ddev-utilities:latest` |

### Ports

To change the HTTP/S ports exposed to the host system:

- Use DDEV dotenv command

    ```shell
    ddev dotenv set .ddev/.env --tts-http-port="${TTS_HTTP_PORT}"   # HTTP
    ddev dotenv set .ddev/.env --tts-https-port="${TTS_HTTPS_PORT}" # HTTPS
    ```

- or, update the `.ddev/.env` file directly.

    ```env
    TTS_HTTP_PORT=5010
    TTS_HTTPS_PORT=5011
    ```

Restart DDEV to apply the changes.

## Credits

**Contributed and maintained by [@tyler36](https://github.com/tyler36)**
