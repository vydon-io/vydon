---
title: OpenAI
description: Learn how to use OpenAI as a synthetic data generator inside Vydon
id: openai
hide_title: false
slug: /connections/openai
---

## Introduction

Vydon supports OpenAI-compatible APIs as a synthetic data generator.
The connection is used by the `Generate AI` transformer to produce
realistic, schema-aware values for columns where a static distribution
or hand-written JavaScript would feel too rigid.

The connection is OpenAI-compatible: any API exposing the OpenAI
chat-completions endpoint (Azure OpenAI, Anyscale, local llama.cpp
servers, etc.) can be used by overriding the URL.

## Configuring an OpenAI connection

Open the Vydon UI, navigate to **Connections**, and create a new
connection of type **OpenAI**.

**Connection Name**: a unique, human-readable label.

**SDK URL**: the base URL of the chat-completions API. Defaults to
`https://api.openai.com/v1`. Override for Azure OpenAI or a
self-hosted compatible server.

**API Key**: the secret key used to authenticate against the API. Vydon
stores the key encrypted at rest and only decrypts it when issuing a
request from the worker.

## Using the connection in a job

In the job configuration, pick **Generate AI** as the transformer for
any column you want to synthesise with an LLM. The transformer asks
for:

- **Connection**: the OpenAI connection you just created.
- **Model**: the model identifier to request (for example `gpt-4o-mini`
  or any name accepted by your endpoint).
- **Prompt**: the instruction passed to the model. Vydon injects the
  column name, type and sample of source values to keep the output
  shape consistent.

## Operational notes

- Requests are issued from the **worker** process. Make sure the
  worker has network reachability to your chosen endpoint.
- LLM generation costs and rate limits apply per request; pair the
  `Generate AI` transformer with batching options on your destination
  to avoid exhausting your provider quota.
- The connection currently exposes only the chat-completions surface.
  Embeddings, fine-tuning and assistants APIs are not yet wired up.
