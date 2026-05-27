---
title: View Job Run Logs
description: Learn how to view job run logs to assist with debugging or to view more details of the job run
id: viewing-job-run-logs
hide_title: false
slug: /guides/viewing-job-run-logs
# cSpell:words LOKICONFIG Promtail
---

## Job Run Logs

This section details the variety of ways that job run logs can be accessed depending on your Vydon environment.

![Job Run Logs](/img/runlogs.png)

## Docker Compose

If you're running the docker compose setup or just trying out Vydon locally, there is currently no option within the UI to view logs.

To see these logs, you'll need to tail the running Vydon `worker` container.

If you ran `make dev` from the root, you can run the following command in your terminal:

```console
docker compose -f compose.dev.yml logs -f worker
```

Or use `make dev/logs` to stream logs from every container in the stack.

An alternative is to use the `docker` command directly, or navigate to the vydon worker container in Docker UI.

```console
docker logs vydon-worker -f
```

## Kubernetes

If you're running in a Kubernetes environment, there are multiple ways to view worker logs, along with support for showing them natively in Vydon's UI.

### kubectl

The standard way of viewing the live pod logs:

```console
kubectl logs -n vydon deployment/vydon-worker -f
```

### Vydon UI

The API can surface pod logs directly in the Vydon job-run view when
`RUN_LOGS_ENABLED=true` and either a Kubernetes pod-log config or a
Loki config is supplied. See [api env vars](../deploy/environment-variables.md#backend-api)
for the full list of `RUN_LOGS_*` settings.

## Persistence with Loki

Vydon ships a first-class Loki integration. Point the API at a Loki
endpoint via the `RUN_LOGS_LOKICONFIG_*` env vars (URL, optional
tenant ID, optional credentials) and the job-run view will query
historical logs from Loki instead of the live pod. The same env vars
also drive any retention or label scoping you may need.

Shipping logs into Loki itself remains your choice — Promtail, Vector,
or the Docker logging driver all work.
