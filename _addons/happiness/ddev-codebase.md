---
title: happiness/ddev-codebase
github_url: https://github.com/happiness/ddev-codebase
description: "A DDEV plugin for fetching tickets from Codebase."
user: happiness
repo: ddev-codebase
repo_id: 1061634936
default_branch: main
tag_name: 1.0.0
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2025-09-22
updated_at: 2025-10-01
workflow_status: unknown
stars: 0
---

# ddev-codebase

This is a DDEV add-on for [Codebase](https://www.codebasehq.com), a project managment
software.

## Add-on installation

```
ddev add-on get happiness/ddev-codebase
```

## PhpStorm Integration

Configure a task server at _Tools > Tasks > Servers_, add a _Generic_ server with the
following config:

```xml
<component name="TaskManager">
  <servers>
    <Generic url="https://<PROJECT>.ddev.site:8081">
      <option name="loginAnonymously" value="true" />
      <option name="responseHandlers">
        <XPathResponseHandler />
        <JsonResponseHandler>
          <selectors>
            <selector name="tasks" path="$.tickets" />
            <selector name="id" path="id" />
            <selector name="summary" path="summary" />
            <selector name="description" path="" />
            <selector name="updated" path="updated" />
            <selector name="created" path="created" />
            <selector name="closed" />
            <selector name="issueUrl" path="issue_url" />
            <selector name="singleTask-id" path="$.ticket.id" />
            <selector name="singleTask-summary" path="$.ticket.summary" />
            <selector name="singleTask-description" path="" />
            <selector name="singleTask-updated" path="$.ticket.updated" />
            <selector name="singleTask-created" path="$.ticket.created" />
            <selector name="singleTask-closed" />
            <selector name="singleTask-issueUrl" path="$.ticket.issue_url" />
          </selectors>
        </JsonResponseHandler>
        <RegExResponseHandler />
      </option>
      <option name="singleTaskUrl" value="{serverUrl}/codebase.php?op=ticket&amp;id={id}" />
      <option name="tasksListUrl" value="{serverUrl}/codebase.php?op=tickets" />
    </Generic>
  </servers>
</component>
```

## .env

Create a `.ddev/.env` file with the following contents or append it if the file aready exists.

```.env
CODEBASE_API_KEY=<API_KEY>
CODEBASE_API_URL="https://api3.codebasehq.com"
CODEBASE_PROJECT=<PROJECT-ID>
CODEBASE_USERNAME=<API_USERNAME>
```
