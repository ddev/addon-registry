---
title: barbieswimcrew/ddev-imgproxy
github_url: https://github.com/barbieswimcrew/ddev-imgproxy
description: "imgproxy add-on for DDEV"
user: barbieswimcrew
repo: ddev-imgproxy
repo_id: 884199243
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2024-11-06
updated_at: 2025-04-10
stars: 0
---

# DDEV imgproxy

## What is this?

This repository allows you to quickly install an [imgproxy](https://imgproxy.net/) server into a [DDEV](https://ddev.readthedocs.io) project using the instructions below.

imgproxy is a standalone server for resizing, processing and converting images. With imgproxy you can quickly and easily edit images on the fly (resize, crop, watermark, filter,...), convert them to other file formats, optimise images and more.

In addition to the paid version, the free, self-hosted version of imgproxy (open source) can also be used without registration or further steps. Of course, there are differences between the open source and the pro version, which are described on the [imgproxy features](https://imgproxy.net/features/) page. 

## Installation

Uses [imgproxy official image](https://ghcr.io/imgproxy/imgproxy)


```sh
ddev add-on get barbieswimcrew/ddev-imgproxy && ddev restart
```

For earlier versions of DDEV run

```sh
ddev get barbieswimcrew/ddev-imgproxy && ddev restart
```


Navigating to your DDEV project's hostname:8081 (for example https://imgproxy.ddev.site:8081) should display the imgproxy default landingpage.

## Explanation

This recipe for [DDEV](https://ddev.readthedocs.io) installs a [.ddev/docker-compose.imgproxy.yaml](https://github.com/barbieswimcrew/ddev-imgproxy/blob/main/docker-compose.imgproxy.yaml) using the `ghcr.io/imgproxy/imgproxy` Docker image.

As soon as imgproxy is running successfully, requests can be sent to imgproxy to convert images from a source format to the desired target format. A prominent scenario for this is the generation of thumbnails:

```
ddev launch :8081/insecure/size:360:240/plain//web/path/to/image.png@avif
```

This request would convert/return the file `web/path/to/image.png` into an AVIF thumbnail with the desired dimensions 360x240px.

**Contributed and maintained by [@barbieswimcrew](https://github.com/barbieswimcrew)**
