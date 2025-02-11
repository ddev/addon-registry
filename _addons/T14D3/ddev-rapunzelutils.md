---
title: T14D3/ddev-rapunzelutils
github_url: https://github.com/T14D3/ddev-rapunzelutils
description: ""
user: T14D3
repo: ddev-rapunzelutils
categories:
  - community
created_at: 2025-02-10
updated_at: 2025-02-10
stars: 1
---

# ddev-rapunzelutils

## Using Custom Aliases

This add-on allows you to create short custom DDEV aliases. To add an alias, you can either manually edit the `.ddev/aliases.yaml` file or use the `ddev alias` command.

To create a new alias, use the `--add` or `--create` flag:

```sh
ddev alias --add <alias_name> <container> <command>
```

For example:

```sh
ddev alias --add myalias web "echo 'This is my custom alias'"
```

### Listing Aliases

To list all available aliases, use the `--list` or `-l` flag:

```sh
ddev alias --list
```

### Removing an Alias

To remove an alias, use the `--remove` or `--delete` flag:

```sh
ddev alias --remove <alias_name>
```

For example:

```sh
ddev alias --remove myalias
```

### Using Aliases

You can then use your aliases with the `ddev alias` command. For example, if we want to create an alias for clearing 
the cache of a Symfony project, we can use the following command:

```sh
ddev alias --add cc web "php bin/console cache:clear"
```
And then we can use the alias like this:

```sh
ddev alias cc
```
The output will look something like this:

```sh
Executing in "web": "php bin/console cache:clear"

 // Clearing the cache for the dev environment with debug true

                                                                                                                        
 [OK] Cache for the "dev" environment (debug=true) was successfully cleared.                                            
                                                                                                                        
```


