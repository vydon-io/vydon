---
title: login
description: Learn how to login to vydon via the vydon login command.
id: login
hide_title: false
slug: /cli/login
---

# vydon login

## Overview

Learn how to login to vydon via the vydon login command.

The `vydon login` command is used to login to vydon through the CLI to generate an access token.
This can then be utilized with other commands to perform authenticated requests to Vydon API.

## Usage

```bash
vydon login
```

Running this command will open up a browser window and direct the user through the login flow that has been configured by the API.

If on success, an access token and (optionally) a refresh token will be saved to the filesystem in `$VYDON_CONFIG_DIR`.

## Environment Variables

| Variable            | Description                                                                                          | Is Required | Default Value         |
| ------------------- | ---------------------------------------------------------------------------------------------------- | ----------- | --------------------- |
| VYDON_API_URL       | The base url of the Vydon API. This can be overridden to connect to different Vydon API environments | false       | http://localhost:8080 |
| VYDON_API_KEY       | The api key for Vydon API.                                                                           | false       |                       |
| LOGIN_HOST          | The http server that is booted up running `vydon login` via an oauth flow                            | false       | 127.0.0.1             |
| LOGIN_REDIRECT_HOST | The redirect host that is sent alongside the oauth flow when running `vydon login`                   | false       | 127.0.0.1             |
| LOGIN_PORT          | The port the http server runs on when running `vydon login`                                          | false       | 4242                  |
