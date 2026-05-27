# Vydon

Open-source data anonymization and synthetic data orchestration. Anonymize
PII, generate realistic synthetic data, and sync environments for safer
testing, debugging, and developer experience.

[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat)](https://makeapullrequest.com)
[![License: MIT](https://img.shields.io/github/license/vydon-io/vydon)](./LICENSE.md)
[![Go Tests](https://github.com/vydon-io/vydon/actions/workflows/go.yml/badge.svg)](https://github.com/vydon-io/vydon/actions/workflows/go.yml)

## Introduction

Vydon is a developer-first toolkit for handling sensitive data in
non-production environments. Teams use it to:

1. **Test code safely against production-shaped data** — anonymize
   sensitive production records and use them locally.
2. **Reproduce production bugs locally** — anonymize and subset the
   production database into a representative slice.
3. **Hydrate staging and QA** with high-quality, production-like data
   to catch bugs earlier.
4. **Reduce compliance scope** — meet GDPR, DPDP, FERPA, HIPAA and
   similar requirements by removing PII before it leaves production.
5. **Seed development databases** with deterministic synthetic data for
   unit tests and demos.

## Features

- Generate synthetic data based on your schema
- Anonymize existing production data
- Subset a production database with any SQL query
- Asynchronous pipeline with automatic retries, failure handling and
  event-sourced playback
- Referential integrity preserved automatically across subsetting and
  anonymization
- Declarative, GitOps-friendly configuration to hydrate CI databases
- Pre-built transformers for every common data type
- Custom transformers in JavaScript or via LLMs
- Built-in integrations with Postgres, MySQL, MS SQL Server, MongoDB,
  DynamoDB, S3 and Google Cloud Storage

## Getting started

Pick the path that matches your workflow. The full developer-environment
guide lives in [docs/docs/guides/vydon-local-dev.md](./docs/docs/guides/vydon-local-dev.md).

### Docker Compose (default, < 2 min)

The fastest way to get the full stack running. Requires Docker with the
modern `docker compose` plugin.

```sh
git clone https://github.com/vydon-io/vydon.git
cd vydon
make dev
```

`make dev` builds the local images and brings up Postgres, Redis,
Temporal, the API, the worker and the frontend with hot reload enabled.
Once it returns, the stack is reachable at:

| Service     | URL                                |
| ----------- | ---------------------------------- |
| Frontend    | <http://localhost:3000>            |
| API         | <http://localhost:8080>            |
| Temporal UI | <http://localhost:8233>            |

Useful follow-up commands:

```sh
make dev/logs    # tail logs from every container
make dev/down    # stop the stack (keeps volumes)
make dev/clean   # stop the stack and wipe all volumes (destructive)
```

If a previous run left the Postgres data volume in a bad state and the
API logs report `database "vydon" does not exist`, run `make dev/clean`
once to recreate it.

### Tilt on Kubernetes (kind or OrbStack)

For contributors who want a Kubernetes-shaped environment closer to
production. Requires [Tilt](https://tilt.dev) and either
[kind](https://kind.sigs.k8s.io) or [OrbStack](https://orbstack.dev)
with Kubernetes enabled.

```sh
# kind
kind create cluster --name vydon-dev
tilt up

# or, with OrbStack Kubernetes
orb start k8s
kubectl config use-context orbstack
tilt up
```

Full instructions, including the Compose authentication overlay, are in
[docs/docs/guides/vydon-local-dev.md](./docs/docs/guides/vydon-local-dev.md).

## Documentation

The full documentation source lives under [docs/](./docs/). Each topic
has a Markdown file under `docs/docs/`:

- [Local development](./docs/docs/guides/vydon-local-dev.md)
- [Deploy](./docs/docs/deploy/)
- [Connections](./docs/docs/connections/)
- [Transformers](./docs/docs/transformers/)
- [CLI](./docs/docs/cli/)

## Contributing

Contributions of every size are welcome. Start by reading
[CONTRIBUTING.md](./CONTRIBUTING.md) and opening a draft PR.

- File a [feature request](https://github.com/vydon-io/vydon/issues/new?labels=enhancement) or [bug report](https://github.com/vydon-io/vydon/issues/new?labels=bug)
- Disclose security issues privately — see [SECURITY.md](./SECURITY.md)

## License

Vydon is distributed under the [MIT Expat license](./LICENSE.md).
