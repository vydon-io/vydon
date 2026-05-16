---
title: Introduction
description: This section details Vydon CLI and all of its available commands.
id: intro
hide_title: false
slug: /cli/introduction
---

# CLI Overview

## Introduction

This section details Vydon CLI and all of its available commands.

```console
➜  ~ vydon
Terminal UI that interfaces with the Vydon system.

Usage:
  vydon [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  jobs        Parent command for jobs
  login       Login to Vydon
  sync        One off sync job to local resource
  version     Print the client version information
  whoami      Find out who you are

Flags:
      --api-key string   Vydon API Key. Takes precedence over $VYDON_API_KEY
      --config string    config file (default is $HOME/.vydon/vydon.yaml)
  -h, --help             help for vydon
  -v, --version          version for vydon

Use "vydon [command] --help" for more information about a command.
```

## Environment Variables

There are a few global environment variables that are available on every request.

### `VYDON_API_KEY`

Vydon API Key. Used if logging in via a system api key.

### `VYDON_API_URL`

The url of the Vydon API to direct the request to.

### Persisting CLI Environment Variables

Environment Variables for the CLI may be persisted by setting them in the config file.
By default this is located at `$HOME/.vydon/config.yaml`.
The CLI does respect `XDG_CONFIG_HOME` as well as a `VYDON_CONFIG_DIR` may be optionally set to override the default location.

Example of a config.yaml:

```yaml
VYDON_API_URL: 'http://localhost:8080'
```

The CLI uses [viper](https://github.com/spf13/viper) for environment management, and has various configuration options that come with it.
You can find the environment setup method [here](https://github.com/vydon-io/vydon/blob/main/cli/internal/cmds/vydon/vydon.go#L80).

### Full list

For a full list of environment variables and flags available, see the specific command you are running.
Otherwise, there is a top-level list of all environment variables spread across all commands available [here](../deploy/environment-variables.md#cli).

## Metadata

CLI metadata is appended to the outgoing gRPC context and HTTP Headers to provide tracking and metadata to the API.
This lets the API know which version the CLI is using when it invokes commands to better track CLI usage over time.

The following metadata is added to all CLI context:

- Git Version
- Git Commit
- OS Platform
