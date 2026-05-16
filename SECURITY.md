# Security Policy

## Supported versions

Vydon is pre-1.0. We support the latest minor release. Older releases receive
security patches only at maintainer discretion.

## Reporting a vulnerability

**Do not open a public issue, pull request, or discussion for a security
vulnerability.**

Use GitHub's [private vulnerability reporting][gh-pvr] for this repository:

[gh-pvr]: https://github.com/vydon-io/vydon/security/advisories/new

You will get an acknowledgment within **72 hours** and a substantive response
within **7 business days**. We aim to ship a fix within **30 days** of
confirmation for critical and high severity issues.

Please include:

- a description of the issue and its impact
- a minimal reproduction (proof-of-concept, payload, or affected code path)
- the version, deployment mode (local, Helm, Docker), and OS
- any mitigations you are aware of

## Scope

In scope:

- the `vydon` core services (backend, worker, frontend)
- the `vydon` CLI and SDKs
- the Helm charts and container images we publish
- the connectors and transformers shipped in this repository

Out of scope:

- third-party dependencies (report upstream; CC us if helpful)
- self-hosted deployments running on unsupported infrastructure
- denial-of-service via resource exhaustion on default limits
- vulnerabilities requiring physical access or compromised credentials
- social engineering of contributors

## Coordinated disclosure

We follow a 90-day coordinated disclosure window. Credit is given to reporters
in the published advisory unless you opt out.

## Supply chain

All container images, Helm charts, and SDK packages are signed with Sigstore
keyless OIDC and ship SLSA build provenance plus SBOM attestations. See
[docs/security/SIGNING.md](docs/security/SIGNING.md) for verification recipes.
