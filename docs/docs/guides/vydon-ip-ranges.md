---
title: Vydon IP Ranges
description: Historical IP ranges used by the former hosted Vydon offering
id: vydon-ip-ranges
hide_title: false
slug: /guides/vydon-ip-ranges
# cSpell:words Vydon
---

## Introduction

> The hosted "Vydon Cloud" SaaS referenced below is no longer a publicly available offering. If you are running Vydon self-hosted, no inbound allow-listing of the IPs below is needed — your instance reaches your databases from your own network. This page is kept for historical reference only.

When self-hosting, you control the egress IPs and can allow-list them on your database firewall directly.

If you need help configuring a Bastion Host, check out the [Connect Postgres via Bastion Host](/guides/connect-private-postgres-via-bastion-host) guide.

Generally, it's good practice to limit inbound connections into a private network.

The following IP addresses were used by the previous Vydon Cloud regions.

## AWS

### us-west-2

```
54.69.79.83
44.235.108.235
35.84.248.98
```

### eu-central-1

```
3.64.167.74
3.65.103.40
3.121.94.104
```
