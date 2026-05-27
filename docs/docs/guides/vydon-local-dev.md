---
title: Developing Vydon Locally
description: Learn how to develop with Vydon Open Source locally in order to get up to speed with how Vydon works
id: vydon-local-dev
hide_title: false
slug: /guides/vydon-local-dev
---

## Introduction

Vydon offers three local development paths, in increasing order of complexity:

1. **Docker Compose** — the default contributor flow. Fastest to start, mirrors production services.
2. **Tilt on kind** — Kubernetes via a local `kind` cluster, useful for chart and manifest work.
3. **Tilt on OrbStack Kubernetes** — same Tilt path, targeting OrbStack's built-in Kubernetes (no `kind` install required).

A Bare-Metal path is also supported but rarely used.

## Setup with Compose

### Pre-requisites

- Go matching `go.mod` (currently 1.26.3)
- Docker Compose >= 2.26

### Setup

The Docker Compose environment is fully self-contained. No external login is required; protobuf code is committed under `backend/gen/`.

To start:

```console
make dev
```

> Note: The `backend` and `worker` containers will start but may take some time to do their initial build. Subsequent rebuilds are fast thanks to the build cache.

To stream logs from every container:

```console
make dev/logs
```

To stop the stack (keeps volumes):

```console
make dev/down
```

To stop and wipe all volumes (destructive — useful when the Postgres data dir holds a stale database name):

```console
make dev/clean
```

Once everything is up and running, the app can be accessed locally at [http://localhost:3000](http://localhost:3000), the API at [http://localhost:8080](http://localhost:8080), and the Temporal UI at [http://localhost:8233](http://localhost:8233).

#### Regenerating protobuf code (optional)

If you change a `.proto` file you need to regenerate the bindings via Buf. The Buf CLI is invoked through Docker, so no local install is required — but anonymous calls to the Buf Schema Registry are rate-limited to 10 requests per hour. To raise the limit, [create a BSR token](https://buf.build/) and export it before running `make generate`:

```console
export BUF_TOKEN=<token>
make generate
```

### Running Compose with Authentication

The repository ships a `compose.auth.yml` overlay that stands up Keycloak with a pre-configured realm so you can sign in offline.

```console
docker compose -f compose.dev.yml -f compose.auth.yml up -d
```

To stop:

```console
docker compose -f compose.dev.yml -f compose.auth.yml down
```

Keycloak is exposed at [http://localhost:8083](http://localhost:8083). The realm is `vydon`. The realm has self-registration enabled — open the app's sign-in page and choose **Register** to create a user. The Keycloak admin console is reachable with `admin` / `change_me`.

## Setup with Tilt

Developing on Kubernetes via Tilt reproduces an environment closer to production. The same Tilt setup works against either a local `kind` cluster or OrbStack's built-in Kubernetes.

### Pre-requisites

- Docker Compose >= 2.26 (Tilt invokes the Docker daemon for builds)
- [tilt](https://tilt.dev/), [kubectl](https://kubernetes.io/docs/reference/kubectl/), [helm](https://helm.sh/), [helmfile](https://github.com/helmfile/helmfile)
- For the kind path: [kind](https://kind.sigs.k8s.io/) and [ctlptl](https://github.com/tilt-dev/ctlptl)
- For the OrbStack path: [OrbStack](https://orbstack.dev/) with Kubernetes enabled in its settings

### Cluster Setup — kind

Create a kind cluster named `vydon-dev` (the cluster name expected by the top-level Tiltfile):

```console
kind create cluster --name vydon-dev
```

Alternatively, the project ships a declarative `ctlptl` spec (kind cluster + local registry):

```console
ctlptl apply -f tilt/kind/cluster.yaml
```

Verify the active context:

```console
kubectl config use-context kind-vydon-dev
```

### Cluster Setup — OrbStack Kubernetes

Enable Kubernetes in OrbStack (**OrbStack → Settings → Kubernetes → Enable**). OrbStack registers itself as the `orbstack` kubectl context:

```console
kubectl config use-context orbstack
```

No additional cluster bootstrap is required — OrbStack provides networking, storage, and an ingress out of the box.

### Running Tilt

After the cluster context is selected, run `tilt up`. Each component has its own sub-Tiltfile so it can be enabled in isolation:

```console
tilt up                # everything
tilt up backend        # backend only
tilt up frontend       # frontend (also brings backend)
```

The app, API and Temporal UI are port-forwarded to [http://localhost:3000](http://localhost:3000), [http://localhost:8080](http://localhost:8080), and [http://localhost:8233](http://localhost:8233) respectively.

To tear everything down:

```console
tilt down
```

## Developing on Bare Metal

You can develop Vydon totally on bare metal. Every service supports a `.env` file along with environment-specific `.env` overrides. This way of developing isn't really used today as we've invested heavily in containerized environments to stay closer to production.

## Tools

This section contains a flat list of the tools used to develop Vydon.

### Kubernetes

If you're choosing to develop in a Tilt environment, this section is more important as it contains all of the K8s-focused tooling.

- [kind](https://github.com/kubernetes-sigs/kind) — Kubernetes in Docker. Used to spin up a slim Kubernetes cluster.
- [OrbStack](https://orbstack.dev/) — Alternative to kind on macOS, ships with a built-in Kubernetes cluster.
- [tilt](https://github.com/tilt-dev/tilt) — Defines our development environment as code.
- [ctlptl](https://github.com/tilt-dev/ctlptl) — CLI from the Tilt team to declaratively define kind clusters and registries.
- [kubectl](https://github.com/kubernetes/kubectl) — Observability and management of the local cluster.
- [kustomize](https://github.com/kubernetes-sigs/kustomize) — YAML template tool for ad-hoc patches.
- [helm](https://github.com/helm/helm) — Kubernetes package manager. All app deployables ship a Helm chart.
- [helmfile](https://github.com/helmfile/helmfile) — Declarative helmfile descriptors. Tilt points directly at them.

### Go + Protobuf

- [Go](https://go.dev/) — Language used for the backend and worker.
- [sqlc](https://github.com/sqlc-dev/sqlc) — Generates Go data-layer code from pure SQL.
- [buf](https://github.com/bufbuild/buf) — Tooling for protobuf.
- [golangci-lint](https://github.com/golangci/golangci-lint) — Go linter.
- [migrate](https://github.com/golang-migrate/migrate) — Runs DB migrations for the API.

### Npm/Node.js

- [Node/Npm](https://nodejs.org/en) — Runs the Next.js app.

All tools can be installed via `brew` on macOS or Linux. `sqlc` and `buf` are invoked through Docker images, so they don't need to be installed locally — the versions are pinned in code and Docker handles the rest.

## Brew Install

The Compose path only needs Go, Node and Docker. The Tilt path needs the full list.

```console
# Compose path (default)
brew install go node

# Tilt + kind path
brew install kind tilt-dev/tap/tilt tilt-dev/tap/ctlptl kubernetes-cli kustomize helm helmfile go sqlc buf golangci-lint node

# Tilt + OrbStack path (no kind/ctlptl needed)
brew install --cask orbstack
brew install tilt-dev/tap/tilt kubernetes-cli kustomize helm helmfile go sqlc buf golangci-lint node
```
