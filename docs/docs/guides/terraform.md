---
title: Configuring Vydon with Terraform
description: Learn how to use Vydon's Terraform provider within your GitOps flow to manage your Terraform infrastructure
id: terraform
hide_title: false
slug: /guides/terraform
---

## Introduction

Vydon ships with an official Terraform provider that can be used to create, read, update, and delete supported Vydon resources.

## Setup and Configuration

Before configuring the Vydon provider, you'll need to know a few pieces of config data so that you can properly configure the provider.

### Endpoint Url

The url to your Vydon API must be provided. It can be set either directly on the provider as a configuration parameter, or via the `VYDON_ENDPOINT` environment variable. This is detailed in the Terraform Registry docs as well. For a local stack started with `make dev`, the API listens on [http://localhost:8080](http://localhost:8080).

### API Key

Next, you'll need to generate an API key that the Terraform provider can use to act on the behalf of your account.
If you haven't configured one, you can do so by heading over to the api key page in the settings for your specific account and creating one.

If the self-hosted instance is running without authentication enabled, this API Key is not utilized, but an account-id must be provided.

The API Key may be input as a variable to the provider, or provided in the environment through the `VYDON_API_TOKEN` environment variable.

### Account Id

Generally, this option is omitted as it is inferred through the API Key.
If self-hosting Vydon and running without authentication, or simply wanting to be redundant, provide the account id to the provider or via the `VYDON_ACCOUNT_ID` environment variable to explicitly tell the provider which account id to use.

## Provider documentation

Until the provider is published to the Terraform Registry, the full
resource and data-source reference lives directly in the
[vydon-io/terraform-provider-vydon](https://github.com/vydon-io/terraform-provider-vydon)
repository under `docs/`. Each resource and data source ships an HCL
example under `examples/`.

## Bugs or Features

If there is an issue with the provider, or there is a feature that is
missing, please [open an issue](https://github.com/vydon-io/terraform-provider-vydon/issues/new)
on the provider repository.
