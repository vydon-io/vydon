---
title: Switch
description: Learn how to switch accounts with the vydon accounts switch command.
id: switch
hide_title: false
slug: /cli/accounts/switch
---

# vydon switch

## Overview

Learn how to switch accounts with the vydon accounts switch command.

The `vydon accounts switch` command is used to switch the account in the cli context.
This is useful if you want to run cli commands for a different account.

## Usage

```bash
vydon switch <name | id>
```

### Argument: name | id

An account-name or account-id can be provided as the first command-line argument. If omitted
it will start an interactive selection mode.
