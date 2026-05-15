---
title: Docker Compose
description: Learn how to deploy Vydon using Docker compose for a better local developer experience
id: docker-compose
hide_title: false
slug: /deploy/docker-compose
---

## Trying Vydon with Compose

A `compose.yml` file is provided at the root of the repository. This uses our pre-built docker images so no building is required.

This file includes two other compose files that are found in the main repository.

- Temporal Compose
- Test Databases

We split out the Temporal compose file to make it easier to include in other places, as well as to keep the main compose clean and to have a separate of concerns.

**This main compose.yml file is made to easily try Vydon and should not be used as-is for production deployments.**

To run this you can run one of the two following commands:

```console
make compose/up
docker compose up -d
```

## Deploying Vydon with Compose

If you wish to deploy Vydon to production with `compose.yml`, we don't currently offer a single `compose.yml` file to do this (yet.).
However, you can easily combine main `compose.yml` and the temporal `compose.yml` files to achieve this.

The main `compose.yml` includes an `api-seed` and `temporal-seed` that may not be necessary and require extra files, so those can be omitted for minimal dependencies.

Once all of the containers come online, the app is now routable via [http://localhost:3000](http://localhost:3000).

## Deploy with Docker Compose and Authentication

> **NB:** This requires a valid Vydon Enterprise license for OSS deployments. If you would like to try this out, please contact us.

Vydon provides an auth friendly compose file that will stand up Vydon in auth-mode with Keycloak.

```console
make compose/auth/up
docker compose -f compose.yml -f compose.auth.yml up -d
```

Keycloak comes default with two clients that allow the app and cli to login successfully.

On first boot up, Keycloak will assert itself with the provided realm Vydon realm.

When navigating to Vydon for the first time, you'll land on the Keycloak sign-in page. It is easy to create an account simply by going through the register flow.
This will persist restarts due to the postgres volume mapping. If you wish to start over, simply delete the vydon docker volume to reset your database to a fresh state.
