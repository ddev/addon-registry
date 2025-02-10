---
title: b13/ddev-stirling-pdf
github_url: https://github.com/b13/ddev-stirling-pdf
description: "Adds Stirling PDF to ddev"
user: b13
repo: ddev-stirling-pdf
categories:
  - community
created_at: 2024-04-04
updated_at: 2025-02-08
stars: 0
---

# Stirling PDF

Stirling PDF is an open-source tool to create, modify and convert PDF files.

List of all features: https://github.com/Stirling-Tools/Stirling-PDF?tab=readme-ov-file#pdf-features

Stirling PDF website: https://github.com/Stirling-Tools/Stirling-PDF/

## Install

For DDEV v1.23.5 or above run

```bash
ddev add-on get b13/ddev-stirling-pdf && ddev restart
```

For earlier versions of DDEV run

```bash
ddev get b13/ddev-stirling-pdf && ddev restart
```

Add the `.ddev/stirling-pdf/extraConfigs/settings.yml` and `.ddev/stirling-pdf/extraConfigs/stirling-pdf-DB.mv.db` to the repository.
`stirling-pdf-DB.mv.db` contains the API key, so it is recommended to have it in git so it can be shared across a team.

* Stirling PDF UI: `https://<project-name>.ddev.site:8064/`
* SwaggerUI (API): `https://<project-name>.ddev.site:8064/swagger-ui/`

## Additional language packs

For OCR, you may want to add additional languages as described [here](https://github.com/Stirling-Tools/Stirling-PDF/blob/main/HowToUseOCR.md#language-packs).
Don't forget to include these language packs in git.

## API Key

Go to `https://<project-name>.ddev.site:8064/account` and copy the API key.
This will be stored in the database.

## Login

Authentication is enabled by default 
(`DOCKER_ENABLE_SECURITY=true` and `SECURITY_ENABLE_LOGIN=true`).

User: `stirling`

Password: `stirling`

**Maintained by [@b13](https://github.com/b13)**

