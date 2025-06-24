---
title: andy-blum/ddev-websocket
github_url: https://github.com/andy-blum/ddev-websocket
description: "A simple, extensible node.js websocket server"
user: andy-blum
repo: ddev-websocket
repo_id: 972679449
ddev_version_constraint: ">= v1.24.3"
dependencies: []
type: contrib
created_at: 2025-04-25
updated_at: 2025-05-03
workflow_status: success
stars: 2
---

# DDEV Websocket

This add-on for [DDEV](https://ddev.readthedocs.io) provides a simple, extensible websocket server. This server allows developers to send/receive messages between the browser and the server without the need for a page refresh.

## Installation

To install this add-on, run the following command:

```bash
ddev add-on get ddev/ddev-websocket
ddev restart
```

## Plugins

Plugins are JavaScript files that can be added to the `.ddev/js/plugins` directory. Each plugin may export functions for the following events:

- `onConnection`: Called when a new connection is established.
- `onMessage`: Called when a message is received from the client.
- `onError`: Called when an error occurs.
- `onClose`: Called when the connection is closed.

Plugins _do not_ need to implement all of these functions, all are opt-in. Each function is called with the individual connection as `this`, and the WebSocketServer instance as the first argument.

A sample plugin `pong.example.js` is provided in the `.ddev/js/plugins` directory, and by default is loaded by the server on startup:

## Usage

The websocket server is available at `http://<project-name>.ddev.site:8080`, and is created using the [ws package](https://github.com/websockets/ws). Connections are initialized in the browser using the WebSocket constructor:

```javascript
const ws = new WebSocket('wss://<project-name>.ddev.site:8080');
```

To use the `pong` plugin, run the following command in your browser's console:

```js
const ws = new WebSocket('wss://<project-name>.ddev.site:8080');
ws.onmessage = ({data}) => console.log(data);
// "pong connected"

ws.send('ping');
// "pong"

ws.send('ping all');
// "pong"
// (appears in all connected clients)

ws.close();
// View ddev logs
```

To view the server logs as you communicate back and forth, run the following command:

```bash
ddev logs -s websocket -f
```
