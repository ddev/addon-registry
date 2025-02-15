---
title: nickchomey/ddev-cloudflare
github_url: https://github.com/nickchomey/ddev-cloudflare
description: "Cloudflare Tunnels for DDEV"
user: nickchomey
repo: ddev-cloudflare
repo_id: 829567361
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2024-07-16
updated_at: 2024-11-03
stars: 0
---

## DDEV-Cloudflare
This ddev add-on helps you easily serve your ddev projects with real public subdomains via Cloudflare Tunnels

## Why does this exist?
### Why Not Ngrok?
DDEV already has an [integration](https://ddev.readthedocs.io/en/stable/users/topics/sharing/) with [ngrok](https://duckduckgo.com/?q=ngrok&ia=web), which allows for accessing your DDEV project via a public url, [for free](https://ngrok.com/pricing). However, this has a lot of complications and limitations, including getting a completely random domain name.

### Cloudflare Tunnels

[Cloudflare Tunnels](https://developers.cloudflare.com/cloudflare-one/connections/connect-networks/), on the other hand, allow you to create a secure (you don't need to have any ports open) tunnel between your DDEV server and Cloudflare, allowing you to publicly serve your DDEV Projects on *actual* subdomains on your *actual* domain. All while leveraging the full power of Cloudflare's CDN, WAF, and other Edge services. All for free. Each person on your team could have a `name-dev.yourproductiondomain.com` subdomain for easily and securely sharing what they are working on!

> "this is the production-quality service that gets the closest to achieving the dream" -- Github [Awesome Tunneling](https://github.com/anderspitman/awesome-tunneling?tab=readme-ov-file#recommendations) repo

And this DDEV Add-on makes it even easier to use Cloudflare Tunnels. A few CLI commands and prompts are all that are needed to seamlessly  install, configure and serve your DDEV Projects via Cloudflare Tunnels. No futzing with config.yaml files is required!

Some relevant reading on Cloudflare Tunnels:
* https://blog.cloudflare.com/tunnel-for-everyone/
* https://blog.cloudflare.com/ridiculously-easy-to-use-tunnels/

## How to use this
### Requirements
* A Cloudflare account with at least one domain name/"zone"
* Create an API Token that can edit Zone DNS. [Docs here](https://developers.cloudflare.com/fundamentals/api/get-started/create-token/). You can specify which zones/domains you want it to have access to, or allow it to access all of your zones. If you have access to multiple accounts (e.g. personal, professional, client etc...), you may want to include access to only specific ones.


### Install Add-on

From within a DDEV project:

For DDEV v1.23.5 or above run

```sh
ddev add-on get github.com/nickchomey/ddev-cloudflare
```

For earlier versions of DDEV run

```sh
ddev get github.com/nickchomey/ddev-cloudflare
```

**HOWEVER**, the addon is installed globally, so this only needs to be done once. The following commands will be available from any project.

### Commands
* `ddev cloudflare install` - installs, if necessary, the appropriate `cloudflared` and `flarectl` tools for your OS and CPU Architecture. These are needed to automatically manage your Cloudflare Tunnels from the command line.
    * `cloudflared` is for creating, managing and communicating over the actual tunnels. A system service will be created for it.
    * `flarectl` will create, update and delete the DNS records in your Cloudflare account, which are used to route the traffic to the appropriate tunnel (and, consequently, server(s))
* `ddev cloudflare connect` - (re)connects a Cloudflare Tunnel to the local server.
    * This only needs to be run once - as all traffic for all projects and domains will go through a single tunnel.
    * It is run automaticallly when you initially `install`, but can be re-run to change the configuration
* `ddev cloudflare serve [list|add|remove]` - Run this from within a DDEV project to manage tunnel routes for that project's fqdns
    * list - Lists the fqdns for the current project and whether they are being served by Cloudflare Tunnel
    * add - Adds new Cloudflare Tunnel Routes for this project. It will prompt you for fqdns(s) to link to the DDEV project. It automatically sets the `additional_fqdns` field in the project's `config.yaml`, and also sets up any required DNS records. It will then restart the project.
    * remove - Removes existing Cloudflare Tunnel Routes for this project. It will prompt you for fqdns(s) to remove from Cloudflare Tunnel and DNS records. It will not remove them from the  `additional_fqdns` field in the project's `config.yaml`
    

## Help Needed
This has only been tested on Ubuntu 24.04 in WSL2. However, code is in place to install and configure everything for other Linux Distros, macOS, and Windows. I assume that it does not work perfectly (or perhaps even at all) - in particular setting up `cloudflared` to run as a [system service](https://developers.cloudflare.com/cloudflare-one/connections/connect-networks/configure-tunnels/local-management/as-a-service/).

Likewise, the workflow and integration with DDEV is probably not ideal.

So, testing, debugging and contributions are very welcome.

Check the issues for outstanding tasks, or feel free to submit new ones.

# Reference Notes from the Add-on Template for outstanding tasks
## Components of the repository

* The fundamental contents of the add-on service or other component. For example, in this template there is a [docker-compose.ddev-cloudflare.yaml](https://github.com/nickchomey/ddev-cloudflare/blob/main/docker-compose.ddev-cloudflare.yaml) file.
* An [install.yaml](https://github.com/nickchomey/ddev-cloudflare/blob/main/install.yaml) file that describes how to install the service or other component.
* A test suite in [test.bats](https://github.com/nickchomey/ddev-cloudflare/blob/main/tests/test.bats) that makes sure the service continues to work as expected.
* [Github actions setup](https://github.com/nickchomey/ddev-cloudflare/blob/main/.github/workflows/tests.yml) so that the tests run automatically when you push to the repository.

## Getting started

6. Update `tests/test.bats` to provide a reasonable test for your repository. Tests are triggered either by manually executing `bats ./tests/test.bats`, automatically on every push to the repository, or periodically each night. Please make sure to attend to test failures when they happen. Others will be depending on you. Bats is a simple testing framework that just uses Bash. To run a Bats test locally, you have to [install bats-core](https://bats-core.readthedocs.io/en/stable/installation.html) first. Then you download your add-on, and finally run `bats ./tests/test.bats` within the root of the uncompressed directory. To learn more about Bats see the [documentation](https://bats-core.readthedocs.io/en/stable/).
7. Create a [release](https://docs.github.com/en/repositories/releasing-projects-on-github/managing-releases-in-a-repository) on GitHub.
8. Test manually with `ddev add-on get <owner/repo>` (or `ddev get <owner/repo>` for older versions of DDEV)
9.  You can test PRs with `ddev add-on get https://github.com/<user>/<repo>/tarball/<branch>` (or `ddev get https://github.com/<user>/<repo>/tarball/<branch>` for older versions of DDEV)


Add-ons were covered in [DDEV Add-ons: Creating, maintaining, testing](https://www.dropbox.com/scl/fi/bnvlv7zswxwm8ix1s5u4t/2023-11-07_DDEV_Add-ons.mp4?rlkey=5cma8s11pscxq0skawsoqrscp&dl=0) (part of the [DDEV Contributor Live Training](https://ddev.com/blog/contributor-training)).

Note that more advanced techniques are discussed in [DDEV docs](https://ddev.readthedocs.io/en/latest/users/extend/additional-services/#additional-service-configurations-and-add-ons-for-ddev).

## How to debug tests (Github Actions)

1. You need an SSH-key registered with GitHub. You either pick the key you have already used with `github.com` or you create a dedicated new one with `ssh-keygen -t ed25519 -a 64 -f tmate_ed25519 -C "$(date +'%d-%m-%Y')"` and add it at `https://github.com/settings/keys`.

2. Add the following snippet to `~/.ssh/config`:

```
Host *.tmate.io
    User git
    AddKeysToAgent yes
    UseKeychain yes
    PreferredAuthentications publickey
    IdentitiesOnly yes
    IdentityFile ~/.ssh/tmate_ed25519
```
3. Go to `https://github.com/<user>/<repo>/actions/workflows/tests.yml`.

4. Click the `Run workflow` button and you will have the option to select the branch to run the workflow from and activate `tmate` by checking the `Debug with tmate` checkbox for this run.

![tmate](https://raw.githubusercontent.com/nickchomey/ddev-cloudflare/main/images/gh-tmate.jpg)

5. After the `workflow_dispatch` event was triggered, click the `All workflows` link in the sidebar and then click the `tests` action in progress workflow.

7. Pick one of the jobs in progress in the sidebar.

8. Wait until the current task list reaches the `tmate debugging session` section and the output shows something like:

```
106 SSH: ssh PRbaS7SLVxbXImhjUqydQBgDL@nyc1.tmate.io
107 or: ssh -i <path-to-private-SSH-key> PRbaS7SLVxbXImhjUqydQBgDL@nyc1.tmate.io
108 SSH: ssh PRbaS7SLVxbXImhjUqydQBgDL@nyc1.tmate.io
109 or: ssh -i <path-to-private-SSH-key> PRbaS7SLVxbXImhjUqydQBgDL@nyc1.tmate.io
```

9. Copy and execute the first option `ssh PRbaS7SLVxbXImhjUqydQBgDL@nyc1.tmate.io` in the terminal and continue by pressing either <kbd>q</kbd> or <kbd>Ctrl</kbd> + <kbd>c</kbd>.

10. Start the Bats test with `bats ./tests/test.bats`.

For a more detailed documentation about `tmate` see [Debug your GitHub Actions by using tmate](https://mxschmitt.github.io/action-tmate/).

**Contributed and maintained by [@CONTRIBUTOR](https://github.com/CONTRIBUTOR)**
