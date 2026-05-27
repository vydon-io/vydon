---
title: GCS
description: Learn how to use Vydon to sync data from a database to Google Cloud Storage
id: gcs
hide_title: false
slug: /connections/gcs
---

## Introduction

Vydon supports Google Cloud Storage (GCS) as a destination connection.
Like the AWS S3 destination, GCS receives anonymized or synthetic
records produced by a Vydon job and stores them as objects in a bucket
of your choosing.

## Configuring GCS

Open the Vydon UI, navigate to **Connections** and create a new
connection of type **GCP Cloud Storage**.

**Connection Name**: a unique, human-readable label.

**Bucket**: the name of the GCS bucket Vydon writes to.

**Path Prefix**: optional path prefix within the bucket. Useful to scope
exports per environment (for example `staging/exports`).

**Service Account JSON**: paste the contents of a Google Cloud service
account key file. The service account needs `storage.objects.create`
and `storage.objects.list` on the bucket. If you would rather not paste
a key, you can leave this empty and rely on Application Default
Credentials picked up from the worker's runtime environment
(`GOOGLE_APPLICATION_CREDENTIALS` or workload identity on GKE).

## Permissions

The minimum IAM permissions required on the bucket are:

- `storage.objects.create`
- `storage.objects.list`
- `storage.objects.get` (only required if you want to re-sync from a
  GCS-backed job run)

The `Storage Object User` predefined role grants this combination on a
single bucket. Apply it at the bucket level rather than at the project
level to scope access tightly.

## Use as a source

GCS can also be used as a source for an existing job run, in the same
way as S3: provide the `job-id` or `job-run-id` of the job whose
outputs you want to replay. See the
[CLI sync](../cli/sync.md) reference for the corresponding flags.
