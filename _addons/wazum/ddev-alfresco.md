---
title: wazum/ddev-alfresco
github_url: https://github.com/wazum/ddev-alfresco
description: "Alfresco Community Edition integration for DDEV"
user: wazum
repo: ddev-alfresco
repo_id: 1070787060
ddev_version_constraint: ">= v1.21.0"
dependencies: []
type: contrib
created_at: 2025-10-06
updated_at: 2025-10-07
workflow_status: success
stars: 0
---

# ddev-alfresco

[![tests](https://github.com/wazum/ddev-alfresco/actions/workflows/tests.yml/badge.svg)](https://github.com/wazum/ddev-alfresco/actions/workflows/tests.yml)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Minimal Alfresco Community Edition integration for DDEV projects.

## What is this?

This add-on provides a lightweight Alfresco Community Edition setup optimized for local development of applications that integrate with Alfresco's REST API, CMIS API, or WebDAV interface. Perfect for developing custom connectors, file management integrations, or document-centric applications without the overhead of a full Alfresco deployment.

The setup runs Alfresco Content Repository with PostgreSQL, with search indexing and transform services disabled for minimal resource usage. This gives you a fast, API-ready Alfresco instance that starts in minutes and uses 4-6GB RAM instead of 16GB+.

## Installation

```bash
ddev add-on get ddev/ddev-alfresco
ddev restart
ddev alfresco-wait
```

The first startup takes 3-5 minutes while Alfresco initializes.

## Access

**Web Interface:**
```bash
ddev launch :8081/alfresco
```

**Note:** This addon provides a minimal API-focused Alfresco setup. Alfresco Share (the traditional web UI) is not included to keep resource usage low. For DMS integrations via REST API, CMIS, or WebDAV, Share is not needed. If you require the full Alfresco web interface, you would need to add a separate Share container to the docker-compose configuration.

**API Endpoints (external):**
- REST API: `https://<project>.ddev.site:8081/alfresco/api/-default-/public/alfresco/versions/1`
- CMIS: `https://<project>.ddev.site:8081/alfresco/api/-default-/public/cmis/versions/1.1/browser`
- WebDAV: `https://<project>.ddev.site:8081/alfresco/webdav`

**Authentication for REST API:**

Option 1 - Basic Authentication (easiest):
```bash
# Example: Get node information
curl -u admin:admin https://<project>.ddev.site:8081/alfresco/api/-default-/public/alfresco/versions/1/nodes/-root-
```

Option 2 - Ticket-based Authentication:
```bash
# 1. Get authentication ticket
TICKET=$(curl -X POST https://<project>.ddev.site:8081/alfresco/api/-default-/public/authentication/versions/1/tickets \
  -H "Content-Type: application/json" \
  -d '{"userId":"admin","password":"admin"}' | jq -r .entry.id)

# 2. Use ticket in subsequent requests
curl https://<project>.ddev.site:8081/alfresco/api/-default-/public/alfresco/versions/1/nodes/-root- \
  -H "Authorization: Basic $(echo -n $TICKET | base64)"
```

**API Endpoints (internal, from web container):**
- REST API: `http://alfresco:8080/alfresco/api/-default-/public/alfresco/versions/1`
- CMIS: `http://alfresco:8080/alfresco/api/-default-/public/cmis/versions/1.1/browser`

**Default credentials:**
- Username: `admin`
- Password: `admin`

**Default groups:**
- `ALFRESCO_ADMINISTRATORS` - Full system administration
- `SITE_ADMINISTRATORS` - Site management
- `EMAIL_CONTRIBUTORS` - Email integration access

## Commands

```bash
ddev alfresco-wait              # Wait for Alfresco to be ready
ddev alfresco-status            # Check service health
ddev alfresco [command]         # Execute command in Alfresco container
ddev logs -s alfresco           # View Alfresco logs (standard DDEV command)
ddev logs -s alfresco-postgres  # View PostgreSQL logs (standard DDEV command)
```

## REST API Capabilities for CMS/DMS Integration

The Alfresco REST API provides complete functionality for DMS integrations:

### Core Operations Supported
- **Folder Management**: Create, rename, move, delete folders with full hierarchy support
- **File Operations**: Upload, download, copy, move, rename, delete files with metadata
- **Navigation**: Browse folder structures, resolve paths, get parent/child relationships
- **Pagination**: Handle large directories with `skipCount` and `maxItems` parameters
- **Filtering**: Query by file type, folder type, or custom properties
- **Metadata**: Full support for custom properties, versioning, and permissions
- **Streaming**: Direct file content access via `/nodes/{nodeId}/content` endpoint

### Key REST Endpoints
- `GET/POST /nodes/{nodeId}` - Node CRUD operations
- `GET /nodes/{nodeId}/children` - List folder contents with pagination
- `GET /nodes/{nodeId}/content` - Stream file content
- `POST /nodes/{nodeId}/copy` - Copy files/folders
- `POST /nodes/{nodeId}/move` - Move files/folders
- `DELETE /nodes/{nodeId}` - Soft delete to trash
- `GET /deleted-nodes` - Access trash/recycle bin

All operations return standard HTTP status codes and JSON responses, making integration straightforward for any programming language or framework.

## Integration Examples

### PHP with Basic Authentication
```php
$config = [
    'alfresco_endpoint' => 'http://alfresco:8080/alfresco/api/-default-/public/alfresco/versions/1',
    'cmis_endpoint' => 'http://alfresco:8080/alfresco/api/-default-/public/cmis/versions/1.1/browser',
    'username' => 'admin',
    'password' => 'admin',
];

// Example: List root folder contents
$ch = curl_init($config['alfresco_endpoint'] . '/nodes/-root-/children');
curl_setopt($ch, CURLOPT_USERPWD, $config['username'] . ':' . $config['password']);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
$response = curl_exec($ch);
$data = json_decode($response, true);
```

### JavaScript/Node.js with Basic Authentication
```javascript
const config = {
    endpoint: 'http://alfresco:8080/alfresco/api/-default-/public/alfresco/versions/1',
    auth: Buffer.from('admin:admin').toString('base64')
};

// Example: Get node information
fetch(`${config.endpoint}/nodes/-root-`, {
    headers: {
        'Authorization': `Basic ${config.auth}`
    }
})
.then(response => response.json())
.then(data => console.log(data));
```

## Configuration

### Using Alternate Versions of Alfresco

If you need to use a different version of Alfresco or PostgreSQL, you can set environment variables in `.ddev/.env.alfresco`:

```bash
# Set the desired versions
ddev dotenv set .ddev/.env.alfresco --alfresco-image="alfresco/alfresco-content-repository-community:23.2.0"
ddev dotenv set .ddev/.env.alfresco --postgres-image="postgres:15"

# Remove old volumes and restart
ddev stop
docker volume rm ddev-${DDEV_SITENAME}-alfresco-data ddev-${DDEV_SITENAME}-alfresco-postgres-data
ddev start
```

The addon uses these default versions:
- Alfresco: `alfresco/alfresco-content-repository-community:23.2.1`
- PostgreSQL: `postgres:15.7`

**Note:** When changing versions, you may need to remove the old volumes to avoid compatibility issues. Make sure to commit the `.ddev/.env.alfresco` file to version control to share the configuration with your team.

## Automatic Initialization

To ensure Alfresco is ready after every `ddev start` or `ddev restart`, add this hook to your `.ddev/config.yaml`:

```yaml
hooks:
  post-start:
    - exec-host: ddev alfresco-wait
```

This automatically waits for Alfresco to fully initialize before returning control, so your application can immediately connect to the API.

## System Requirements

- DDEV v1.21 or higher
- 4GB RAM minimum (6GB recommended)
- Port 8081 available

## Services

- **Alfresco Content Repository** v23.x (Community Edition)
- **PostgreSQL** v15.7

Search indexing (Solr), document transformation, and message queue services are disabled for minimal resource usage. These can be added later if needed.

## Removal

```bash
ddev add-on remove ddev-alfresco
```

This removes the add-on and deletes all Alfresco data.

## License

Apache License 2.0 - Copyright © 2025 Wolfgang Klinger <wolfgang@wazum.com>
