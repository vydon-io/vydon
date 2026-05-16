<p align="center">
  <!-- <img alt="vydonbanner" src="https://assets.vydon.io/vydon/docs/vydon-header.svg" > -->
  <picture>
  <source
    srcset="https://assets.vydon.io/vydon/docs/vydon-header.svg"
    media="(prefers-color-scheme: light)"
  />
  <source
    srcset="https://assets.vydon.io/vydon/docs/vydon-header-dark.svg"
    media="(prefers-color-scheme: dark), (prefers-color-scheme: no-preference)"
  />
  <img src="https://github-readme-stats.vercel.app/api?username=anuraghazra&show_icons=true" />
</picture>
</p>

<p align="center" style="font-size: 24px;font-weight: 500;">
Open Source Data Anonymization and Synthetic Data Orchestration
<p>

<div align='center'>
 | <a href="https://vydon.io">Website</a>
 | <a href="https://docs.vydon.io">Docs</a>
 | <a href="https://discord.com/invite/MFAMgnp4HF">Discord</a>
 | <a href="https://vydon.io/blog">Blog</a>
 | <a href="https://docs.vydon.io/changelog">Changelog</a>
 | <a href="https://vydon-io.productlane.com/roadmap">Roadmap</a>
</div>

 <br>

<div align="center">
  <a href='https://makeapullrequest.com'>
    <img alt='PRs Welcome' src='https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=shields'/>
  </a>
  <a href="https://github.com/vydon-io/vydon/blob/main/LICENSE.md">
    <img alt="License: MIT" src="https://img.shields.io/github/license/vydon-io/vydon"/>
  </a>
  <a href="https://github.com/vydon-io/vydon/actions/workflows/go.yml/">
    <img alt="Go Tests" src="https://github.com/vydon-io/vydon/actions/workflows/go.yml/badge.svg"/>
  </a>
  <a href="https://artifacthub.io/packages/search?repo=vydon">
    <img alt="ArtifactHub Vydon" src="https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/vydon" />
  </a>
</div>

## Introduction

[Vydon](https://vydon.io) is an open-source, developer-first way to anonymize PII, generate synthetic data and sync environments for better testing, debugging and developer experience.

Companies use Vydon to:

1. **Safely test code against production data** - Anonymize sensitive production data in order to safely use it locally for a better testing and developer experience
2. **Easily reproduce production bugs locally** - Anonymize and subset production data to get a safe, representative data set that you can use to locally reproduce production bugs quickly and efficiently
3. **High quality data for lower-level environments** - Catch bugs before they hit production when you hydrate your staging and QA environments with production-like data
4. **Solve GDPR, DPDP, FERPA, HIPAA and more** - Use anonymized and synthetic data to reduce your compliance scope and easily comply with laws like HIPAA, GDPR, and DPDP
5. **Seed development databases** - Easily seed development databases with synthetic data for unit testing, demos and more

## Features

- **Generate synthetic data** based on your schema
- **Anonymize existing production-data** for a better developer experience
- **Subset your production database** for local and CI testing using any SQL query
- **Complete async pipeline** that automatically handles job retries, failures and playback using an event-sourcing model
- **Referential integrity** for your data automatically
- **Declarative, GitOps based configs** as a step in your CI pipeline to hydrate your CI DB
- **Pre-built data transformers** for all major data types
- **Custom data transformers** using javascript or LLMs
- **Pre-built integrations** with Postgres, Mysql, S3

## Getting started

Pick the path that matches your workflow.

### Docker Compose (default, < 2 min)

The fastest way to get the full stack running. Requires Docker with the
modern `docker compose` plugin.

```sh
git clone https://github.com/vydon-io/vydon.git
cd vydon
make dev
```

`make dev` builds the local images and brings up Postgres, Redis, Temporal,
the API, the worker and the frontend with hot reload enabled. Once it
returns, the stack is reachable at:

| Service         | URL                                                 |
| --------------- | --------------------------------------------------- |
| Frontend        | <http://localhost:3000>                             |
| API healthcheck | <http://localhost:8080/healthz>                     |
| Temporal UI     | <http://localhost:8233>                             |

Useful follow-up commands:

```sh
make dev/logs    # tail logs from every container
make dev/down    # stop the stack (keeps volumes)
make dev/clean   # stop the stack and wipe all volumes (destructive)
```

If you previously ran an older revision and the API logs report
`database "vydon" does not exist`, run `make dev/clean` once to recreate
the Postgres data volume.

### Tilt + kind (Kubernetes path, optional)

For contributors who want a Kubernetes-shaped environment closer to
production. Requires [Tilt](https://tilt.dev) and a running
[kind](https://kind.sigs.k8s.io) cluster named `vydon-dev`.

```sh
kind create cluster --name vydon-dev
tilt up
```

Tilt watches the source tree and live-syncs Go and frontend changes into
the cluster. The Tilt UI exposes per-resource logs and health.

### Production-style compose

The root [compose.yml](./compose.yml) pulls published images from GHCR
(no local build) and seeds demo connections and jobs.

```sh
docker compose up -d    # start
docker compose down     # stop
```

This path is useful for demos and CI smoke tests, not for active
development — there is no hot reload.

## Kubernetes, Auth Mode and more

For more in-depth details on environment variables, Kubernetes deployments, and a production-ready guide, check out the [Deploy Vydon](https://docs.vydon.io/deploy/introduction) section of our Docs.

## Resources

Some resources to help you along the way:

- [Docs](https://docs.vydon.io) for comprehensive documentation and guides
- [Discord](https://discord.com/invite/MFAMgnp4HF) for discussion with the community and Vydon team
- [X](https://x.com/vydon) for the latest updates

## Contributing

We love contributions big and small. Here are just a few ways that you can contribute to Vydon.

- Join our [Discord](https://discord.com/invite/MFAMgnp4HF) channel and ask us any questions there
- Open a PR (see our instructions on [developing with Vydon locally](https://docs.vydon.io/guides/vydon-local-dev))
- Submit a [feature request](https://github.com/vydon-io/vydon/issues/new?assignees=&labels=enhancement%2C+feature&template=feature_request.md) or [bug report](https://github.com/vydon-io/vydon/issues/new?assignees=&labels=bug&template=bug_report.md)

## Licensing

We strongly believe in free and open source software and make this repo is available under the [MIT expat license](./LICENSE.md).
