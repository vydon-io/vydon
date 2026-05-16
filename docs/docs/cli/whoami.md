---
title: whoami
description: Learn how to display the currently logged in user with the vydon whoami CLI command.
id: whoami
hide_title: true
slug: /cli/whoami
---

# vydon whoami

## Overview

Learn how to display the currently logged in user with the vydon whoami CLI command.

The `vydon whoami` command is used to show the currently logged in user.

## Usage

```bash
vydon whoami
```

## Options

The following options can be passed using the `vydon whoami` command:

- `--api-key` - Vydon API Key. Takes precedence over `$VYDON_API_KEY`

## Environment Variables

| Variable      | Description                                                                                          | Is Required | Default Value         |
| ------------- | ---------------------------------------------------------------------------------------------------- | ----------- | --------------------- |
| VYDON_API_URL | The base url of the Vydon API. This can be overridden to connect to different Vydon API environments | false       | http://localhost:8080 |
| VYDON_API_KEY | The api key for Vydon API.                                                                           | false       |                       |
