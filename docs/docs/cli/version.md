---
title: Version
description: Learn how to view the Vydon CLI Version with the vydon version command and vydon --version flag.
id: version
hide_title: false
slug: /cli/version
---

# vydon version

## Overview

Learn how to view the Vydon CLI Version with the vydon version command and vydon --version flag.

## Command Usage

```bash
vydon version
```

The basic command will print out details such as the current Git tag, commit, build date, go version, and other relevant operating system information.

This can be augmented for systems by providing the `--output` flag to generate the version data in either `yaml` or `json`.

```bash
vydon version --output json
vydon version --output yaml
```

## Flag Usage

```bash
vydon --version
```

The top-level `--version` flag can be used to simply print out the Git tagged version of Vydon CLI.
For more details, it's recommended to use the `vydon version` command.
