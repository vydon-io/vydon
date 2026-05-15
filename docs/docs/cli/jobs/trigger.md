---
title: Trigger
description: Learn how to trigger a Vydon job with the vydon jobs trigger command.
id: trigger
hide_title: false
slug: /cli/jobs/trigger
---

## Overview

Learn how to trigger a Vydon job with the vydon jobs trigger command.

The `vydon jobs trigger` command is used to trigger an execution of a Vydon job.
This is useful if a Job is configured but is not running on a schedule, or it's desired to trigger a job outside of the normal scheduled flow.

## Usage

```bash
vydon jobs trigger <job-id>
```

### Argument: job-id

A job-id must be provided as the first command-line argument. This is required and will fail otherwise.
This job-id is used to trigger a workflow execution of the relevant Vydon Job.
