---
title: Security Overview
description: How Vydon handles security for self-hosted deployments
id: cloud-security-overview
hide_title: false
slug: /cloud-security-overview
---

Vydon is shipped as OSS for self-hosting. The notes below describe the
security posture of the software itself; the security of any deployment
is the operator's responsibility.

## Code

All Vydon code is open source and lives on [GitHub](https://github.com/vydon-io/vydon).
The container images and Helm charts published by GoReleaser are built
from the same source you can read in the repo.

If you find a security vulnerability, please follow the disclosure
process in [SECURITY.md](https://github.com/vydon-io/vydon/blob/main/SECURITY.md)
— open a private security advisory on the repository.

## Connecting a production database

We do not recommend pointing Vydon at a production database directly.
Beyond the obvious security exposure, Vydon can put noticeable load on
the source database during a sync. The recommended pattern is to
restore a snapshot of the production database into a separate instance
and point Vydon at that copy.

## Network and credentials

For self-hosted deployments we recommend:

- Running Vydon in a private network and not exposing it to the public
  internet without an authenticating proxy or [Vydon Auth Mode](/deploy/authentication).
- Reaching source databases via a private peering, VPN, or
  [Bastion Host](/guides/connect-private-postgres-via-bastion-host).
- Rotating API keys regularly. API keys created in Vydon expire after
  at most one year, and once revealed at creation time they cannot be
  retrieved again from the API.
