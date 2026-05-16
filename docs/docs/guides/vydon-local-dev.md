---
title: Developing Vydon Locally
description: Learn how to develop with Vydon Open Source locally in order to get up to speed with how Vydon works
id: vydon-local-dev
hide_title: false
slug: /guides/vydon-local-dev
---

## Introduction

This section goes into detail on each tool that is used for developing with Vydon locally.

## Setup with Compose

### Pre-requisites

- Go >= 1.24 (see [go.mod](https://github.com/vydon-io/vydon/blob/main/go.mod) for the exact minimum)
- Docker Compose >= 2.26

### Setup

#### Buf Login

Vydon uses Buf to generate code from our proto files. This is possible to do unauthenticated, but if done often (more than 10 requests in an hour), you will be rate limited. To combat this, you must login to the [BSR](https://buf.build) and create a user token.

Afterwards, drop your token in the `backend/.env.dev.secrets` file.

```console
echo "BUF_TOKEN=<token>" >> ./backend/.env.dev.secrets
```

The docker compose environment runs entirely by itself.

To start:

```console
make dev
```

> Note: The `backend` and `worker` containers will start but may take some time to do their initial build.

Once they have a build cache, they will come online and re-build much faster!

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

### Running Compose with Authentication

The repository ships a `compose.auth.yml` overlay that stands up Keycloak with a pre-configured realm so you can sign in with a standard username and password, completely offline.

```console
docker compose -f compose.yml -f compose.auth.yml up -d
```

To stop:

```console
docker compose -f compose.yml -f compose.auth.yml down
```

## Setup with Tilt

Developing with Kubernetes via Tilt is an alternative path. It is heavier than the Compose flow but reproduces a Kubernetes environment closer to production.

> The default Vydon contributor workflow is Compose. Tilt is kept for contributors who want a kind-based cluster.

### Docker Desktop

If using Docker Desktop, the host file path to the `.data` folder will need to be added to the File Sharing tab.

The allow list can be found by first opening Docker Desktop. `Settings -> Resources -> File Sharing` and add the path to the Vydon repository.

If you don't want to do this, the volume mappings can be removed by removing the PVC for Tilt.
This comes at a negative of the local database not surviving restarts.

### Cluster Setup

Create a `kind` cluster named `vydon-dev` (the cluster name expected by the top-level [Tiltfile](https://github.com/vydon-io/vydon/blob/main/Tiltfile)):

```console
kind create cluster --name vydon-dev
```

If you prefer a declarative cluster spec, the project ships one at [tilt/kind/cluster.yaml](https://github.com/vydon-io/vydon/blob/main/tilt/kind/cluster.yaml) which you can pass via `kind create cluster --config tilt/kind/cluster.yaml`.

After the cluster is up, run `tilt up`. Each dependency in the `vydon` repo has its own sub-Tiltfile so it can be enabled in isolation:

```console
tilt up                # everything
tilt up backend        # backend only
tilt up frontend       # frontend (also brings backend)
```

Once everything is up and running, the app can be accessed locally at [http://localhost:3000](http://localhost:3000).

## Developing on Bare Metal

You can develop Vydon totally on bare metal. Every service supports a .env file along with environment specific .env overrides.
This way of developing isn't really used today as we've invested heavily in developing within containerized environments to be more closely aligned with a production environment.

## Tools

This section contains a flat list of the tools that are used to develop Vydon and why.

Detailed below are the main dependencies are descriptions of how they are utilized:

### Kubernetes

If you're choosing to develop in a Tilt environment, this section is more important as it contains all of the K8s focused tooling.

Tilt is a great tool that is used to automate the setup of a Kubernetes cluster. There are multiple `Tiltfile`'s throughout the code, along with a top-level one that is used to inject all of the K8s manifests to setup Vydon inside of a K8s cluster.

This enables fast development, locally, while closely mimicking a real production environment.

- [kind](https://github.com/kubernetes-sigs/kind)
  - Kubernetes in Docker. We use this to spin up a slim kubernetes cluster that deploys all of the `vydon` resources.
- [tilt](https://github.com/tilt-dev/tilt)
  - Allows us to define our development environment as code.
- [ctlptl](https://github.com/tilt-dev/ctlptl)
  - CLI provided by the Tilt-team to make it easy to declaratively define the kind cluster that is used for development
- [kubectl](https://github.com/kubernetes/kubectl)
  - Allows for observability and management into the spun-up kind cluster.
- [kustomize](https://github.com/kubernetes-sigs/kustomize)
  - yaml template tool for ad-hoc patches to kubernetes configurations
- [helm](https://github.com/helm/helm)
  - Kubernetes package manager. All of our app deployables come with a helm-chart for easy installation into kubernetes
- [helmfile](https://github.com/helmfile/helmfile)
  - Declaratively define a helmfile in code! We have all of our dev charts defined as a helmfile, of which Tilt points directly to.

### Go + Protobuf

- [Go](https://go.dev/)
  - The language of choice for our backend and worker packages
- [sqlc](https://github.com/sqlc-dev/sqlc)
  - Our tool of choice for the data-layer. This lets us write pure SQL and let sqlc generate the rest.
- [buf](https://github.com/bufbuild/buf)
  - Our tool of choice for interfacing with protobuf
- [golangci-ci](https://github.com/golangci/golangci-lint)
  - The golang linter of choice
- [migrate](https://github.com/golang-migrate/migrate)
  - Golang Migrate is the tool that is used to run DB Migrations for the API.

### Npm/Nodejs

- [Node/Npm](https://nodejs.org/en)
  - Used to run the app, along with Nextjs.

All of these tools can be easily installed with `brew` if on a Mac.
Today, `sqlc` and `buf` don't need to be installed locally as we exec docker images for running them.
This lets us declare the versions in code and docker takes care of the rest.

## Brew Install

Each tool above can be straightforwardly installed with brew if on Linux/MacOS. The Compose path only needs Go, Node and Docker; the Tilt path needs the full list.

```console
# Compose path (default)
brew install go node

# Tilt + kind path (optional)
brew install kind tilt-dev/tap/tilt tilt-dev/tap/ctlptl kubernetes-cli kustomize helm helmfile go sqlc buf golangci-lint node
```
