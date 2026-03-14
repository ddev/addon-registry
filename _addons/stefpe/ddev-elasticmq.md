---
title: stefpe/ddev-elasticmq
github_url: https://github.com/stefpe/ddev-elasticmq
description: "In-memory message queue with an Amazon SQS-compatible interface as ddev integration"
user: stefpe
repo: ddev-elasticmq
repo_id: 1180670132
default_branch: main
tag_name: 0.0.2
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2026-03-13
updated_at: 2026-03-13
workflow_status: unknown
stars: 0
---

[![add-on registry](https://img.shields.io/badge/DDEV-Add--on_Registry-blue)](https://addons.ddev.com)
[![tests](https://github.com/stefpe/ddev-elasticmq/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/stefpe/ddev-elasticmq/actions/workflows/tests.yml?query=branch%3Amain)
[![last commit](https://img.shields.io/github/last-commit/stefpe/ddev-elasticmq)](https://github.com/stefpe/ddev-elasticmq/commits)
[![release](https://img.shields.io/github/v/release/stefpe/ddev-elasticmq)](https://github.com/stefpe/ddev-elasticmq/releases/latest)

# DDEV ElasticMQ

## Overview

This add-on integrates [ElasticMQ](https://github.com/softwaremill/elasticmq) into your [DDEV](https://ddev.com/) project. ElasticMQ is an Amazon SQS-compatible in-memory message queue, perfect for local development and testing.

**Features:**
- SQS-compatible REST API via DDEV router (HTTPS on port 9326)
- Web UI for monitoring queues via DDEV router (HTTPS on port 9327)
- Support for custom queue configuration via HOCON config file
- Multiple projects can run ElasticMQ simultaneously without port conflicts
- Optional queue persistence

## Installation

```bash
ddev add-on get stefpe/ddev-elasticmq
ddev restart
```

After installation, the following file is created in your `.ddev` directory:
- `.ddev/elasticmq.conf` - Queue configuration file

Make sure to commit this file to version control.

## Usage

| Command | Description |
| ------- | ----------- |
| `ddev describe` | View service status for ElasticMQ |
| `ddev logs -s elasticmq` | Check ElasticMQ logs |
| `ddev exec curl http://elasticmq:9324/health` | Check ElasticMQ health |
| `ddev elasticmq` | Launch the ElasticMQ Web UI in your browser |

### Accessing the Services

The add-on uses DDEV's built-in router for HTTPS access, eliminating port conflicts between projects:

- **SQS API**: `https://<project>.ddev.site:9326`
- **Web UI**: `https://<project>.ddev.site:9327`

Launch the Web UI quickly with:

```bash
ddev elasticmq
```

### Connecting with AWS SDK

```javascript
const AWS = require('aws-sdk');

const sqs = new AWS.SQS({
  endpoint: 'https://<project>.ddev.site:9326',
  region: 'elasticmq',
  accessKeyId: 'x',
  secretAccessKey: 'x'
});
```

## Queue Configuration

ElasticMQ supports configuring queues via a HOCON configuration file. After installing the add-on, you can define custom queues by editing the `.ddev/elasticmq.conf` file.

### Basic Queue Configuration

The configuration file uses HOCON format (Human-Optimized Config Object Notation). Here's an example of defining a simple queue:

```hocon
include classpath("application.conf")

queues {
  my-queue {
    fifo = false
    defaultVisibilityTimeout = 30 seconds
  }
}
```

### Advanced Queue Configuration

You can configure queues with various options including dead letter queues, visibility timeouts, and tags:

```hocon
include classpath("application.conf")

queues {
  # Standard queue with custom settings
  admin-integration-development-messenger-transport {
    fifo = false
    contentBasedDeduplication = false
    defaultVisibilityTimeout = 15 seconds
    delay = 0 seconds
    receiveMessageWait = 0 seconds
  }

  # FIFO queue example
  my-fifo-queue {
    fifo = true
    contentBasedDeduplication = true
    defaultVisibilityTimeout = 60 seconds
  }

  # Queue with dead letter queue
  main-queue {
    defaultVisibilityTimeout = 30 seconds
    deadLettersQueue {
      name = "main-queue-dlq"
      maxReceiveCount = 3
    }
  }
  main-queue-dlq { }
}
```

### Configuration Options

| Option | Description | Default |
|--------|-------------|---------|
| `fifo` | Whether this is a FIFO queue | `false` |
| `contentBasedDeduplication` | Enable content-based deduplication | `false` |
| `defaultVisibilityTimeout` | Time before message becomes visible again | `30 seconds` |
| `delay` | Default delay for messages | `0 seconds` |
| `receiveMessageWait` | Long polling wait time | `0 seconds` |
| `deadLettersQueue.name` | Name of the dead letter queue | - |
| `deadLettersQueue.maxReceiveCount` | Max receives before moving to DLQ | `1-1000` |
| `copyTo` | Copy all messages to another queue | - |
| `moveTo` | Move all messages to another queue | - |
| `tags` | Key-value tags for the queue | - |

### Applying Configuration Changes

After modifying `.ddev/elasticmq.conf`:

```bash
ddev restart
```

The queues will be created automatically when ElasticMQ starts.

### Connecting to Your Queues

Once configured, connect to ElasticMQ using any SQS-compatible client:

**AWS CLI example:**
```bash
aws --endpoint-url=https://<project>.ddev.site:9326 --no-verify-ssl sqs list-queues
```

**AWS SDK (JavaScript) example:**
```javascript
const sqs = new AWS.SQS({
  endpoint: 'https://<project>.ddev.site:9326',
  region: 'elasticmq',
  accessKeyId: 'x',
  secretAccessKey: 'x'
});
```

## Multiple Projects Support

The dynamic port binding through DDEV's router means multiple projects can run ElasticMQ simultaneously without any port configuration. Each project gets its own isolated ElasticMQ instance accessible through project-specific URLs.

## Custom Docker Image

To use a different ElasticMQ image, set the environment variable in your `.ddev/config.yaml`:

```yaml
web_environment:
  - ELASTICMQ_DOCKER_IMAGE=softwaremill/elasticmq-native:latest
```

**Note:** The native image is smaller (~30MB vs ~240MB) and starts faster, but some logging features may not work.

## Advanced Topics

### Queue Persistence

To persist queue metadata across restarts, enable queues-storage in your `.ddev/elasticmq.conf`:

```hocon
queues-storage {
  enabled = true
  path = "/mnt/ddev_config/elasticmq-queues.conf"
}
```

### Message Persistence

To persist messages to a database (H2), enable messages-storage:

```hocon
messages-storage {
  enabled = true
  uri = "jdbc:h2:/mnt/ddev_config/elasticmq.db"
}
```

Make sure to add the storage files to your `.gitignore` if you don't want to commit them.

## Credits

**Contributed and maintained by [@stefpe](https://github.com/stefpe)**
