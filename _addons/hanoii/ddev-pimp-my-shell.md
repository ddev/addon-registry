---
title: hanoii/ddev-pimp-my-shell
github_url: https://github.com/hanoii/ddev-pimp-my-shell
description: "A collection of tools I like to have on ddev."
user: hanoii
repo: ddev-pimp-my-shell
repo_id: 777249734
ddev_version_constraint: ""
dependencies: []
type: contrib
created_at: 2024-03-25
updated_at: 2025-07-31
workflow_status: success
stars: 4
---

[![tests](https://github.com/hanoii/ddev-pimp-my-shell/actions/workflows/tests.yml/badge.svg)](https://github.com/hanoii/ddev-pimp-my-shell/actions/workflows/tests.yml)
![project is maintained](https://img.shields.io/maintenance/yes/2025.svg)

# ddev-pimp-my-shell <!-- omit in toc -->

This a an addon that adds several nice command line utitlites I use across all
projects, making it easier/nicer to work inside the container with some initial
defaults.

<!-- toc -->

- [Install](#install)
- [Features](#features)
- [Tweaks](#tweaks)
  * [Xdebug command inside the container](#xdebug-command-inside-the-container)
  * [terminal window title](#terminal-window-title)
  * [ls replacement](#ls-replacement)
  * [ahoy](#ahoy)
  * [starship](#starship)
  * [fish](#fish)
  * [lazygit](#lazygit)
  * [go](#go)
  * [delta](#delta)
  * [DDEV `post-import-db` hook](#ddev-post-import-db-hook)

<!-- tocstop -->

## Install

You can always install the latest code:

For DDEV v1.23.5 or above run

```sh
ddev add-on get https://github.com/hanoii/ddev-pimp-my-shell/tarball/main
```

For earlier versions of DDEV run

```sh
ddev get https://github.com/hanoii/ddev-pimp-my-shell/tarball/main
```

**Contributed and maintained by [@hanoii](https://github.com/hanoii)**

## Features

- https://github.com/ahoy-cli/ahoy
- https://github.com/skywind3000/z.lua (working on both bash and fish)
- https://github.com/junegunn/fzf
- https://github.com/junegunn/fzf-git.sh
- https://starship.rs/
- https://github.com/charmbracelet/gum
- https://github.com/fish-shell/fish-shell / https://fishshell.com/
  - https://github.com/jorgebucaran/fisher
  - https://github.com/IlanCosman/tide
  - https://github.com/edc/bass
- https://chrisbuilds.github.io/terminaltexteffects/
- https://bun.sh/
- https://github.com/eza-community/eza/
- https://github.com/jesseduffield/lazygit/
- https://github.com/noborus/ov
- https://go.dev
- https://github.com/skx/sysbox
- https://github.com/dbohdan/recur
- https://www.rust-lang.org/
- https://github.com/samwho/spacer
- https://dandavison.github.io/delta/
- https://github.com/sharkdp/bat
- https://github.com/sharkdp/fd
- Some useful [scripts](https://github.com/hanoii/ddev-pimp-my-shell/blob/main/pimp-my-shell/scripts)
- https://github.com/cli/cli

## Tweaks

### Xdebug command inside the container

While you can enable Xdebug by following the [DDEV's guide][ddev-xdebug], I work
a lot inside the container.

I found that you can use DDEV's own `enable_xdebug` and `disable_xdebug` inside
the container but I wanted a shortcut to that so this add-on
[ships a small script](https://github.com/hanoii/ddev-pimp-my-shell/blob/main/homeadditions/.local/bin/xdebug) that simply toggles
Xdebug state by running `xdebug` inside the container.

[ddev-xdebug]:
  https://ddev.readthedocs.io/en/stable/users/debugging-profiling/step-debugging/

### terminal window title

Both for bash and fish, this add-on changes the window terminal with a better
default if working inside the container:

- Adds the DDEV project name
- shows that is a DDEV project
- includes the current $PWD shortenning it.
- shows the currently running command

i.e. for this add-on `sleep 10 - ddev-pimp-my-shell/ddev: /v/w/html`

### ls replacement

It aliases `ls` and `ll` to use [eza](https://github.com/eza-community/eza/)
instead of ls.

### ahoy

It bundles autocomplete for both bash and fish.

### starship

If you want to use a [Nerd font](https://starship.rs/presets/nerd-font), there
are many ways you can achieve that. I am personally editing my global
`~/.ddev/global_config.yaml` on the host and adding a `STARSHIP_CONFIG`
environment variable pointing to an alternative config file which I also added
to `~/.ddev/homeadditions/.config` on the host with the content of this add-on's
[`starship.toml`](https://github.com/hanoii/ddev-pimp-my-shell/blob/main/homeadditions/.config/starship.toml) and merging it with the
output of (`[directory]` is in both):

```
starship preset nerd-font-symbols
```

### fish

For tide, it uses the default configuration, if you want to have your own on all
of your DDEV projects you can create/edit
`~/.ddev/homeadditions/.config/fish/conf.d/mytide.fish` on the host with
something like the following:

```fish
# Doing it as a fish_prompt event to make sure it is shown the first time it's
# run. Not sure if this is the best workaround, but otherwise the first `ddev fish`
# would show no prompt.
function mytide --on-event fish_prompt
  if not test -f ~/.config/fish/conf.d/.mytide && tide --version > /dev/null
    tide configure --auto --style=Classic --prompt_colors='True color' --classic_prompt_color=Light --show_time=No --classic_prompt_separators=Angled --powerline_prompt_heads=Sharp --powerline_prompt_tails=Flat --powerline_prompt_style='One line' --prompt_spacing=Compact --icons='Many icons' --transient=Yes
    touch ~/.config/fish/conf.d/.mytide
  end
endkkk
```

### lazygit

It adds a [config.yml](https://github.com/hanoii/ddev-pimp-my-shell/blob/main/homeadditions/.config/lazygit/config.yml) diabling the
startup popups.

### go

Other than adding itself to the path and installing some packages I am unsetting
`GOARCH` and `GOOS` for both `ddev ssh` and `ddev fish` so that are not in
conflict with the container OS as
[they are added by DDEV from the host](https://github.com/ddev/ddev/issues/6748).

### delta

Having found [delta](https://dandavison.github.io/delta/) and wanting to give
this a try I am auto-configuring this for git. I am doing so in the shell
startup scripts so that it's picked up after any global homeaddition.

### DDEV `post-import-db` hook

I generally have import scripts that I ship either from other add-ons or
specific to projects and those scripts usually download the database from within
the container. Rather than downloading the database and then running
`ddev import-db --file=` I wanted to provide a way so that either way runs all
of my post-import-db commands both with the ddev import command as well as from
any helper script from within the project.

So the way this works is that you can run
`/var/www/html/.ddev/pimp-my-shell/hooks/post-import-db.sh` from any script that
will be run in the container and it will take care of running all scripts on
`.ddev/pimp-my-shell/post-import-db.d/`.

This also allows for other add-ons to add scripts to this directory so that they
will be run.

Example on how I am using it:

- https://github.com/hanoii/ddev-platformsh-lite/blob/ce4b95d578e82f942b75ec9fa6fa60f63473c90b/platformsh-lite/scripts/db-pull.sh#L91
- https://github.com/hanoii/ddev-platformsh-lite/blob/ce4b95d578e82f942b75ec9fa6fa60f63473c90b/pimp-my-shell/hooks/post-import-db.d/00-drupal.sh
