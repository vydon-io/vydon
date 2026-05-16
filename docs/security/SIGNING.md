# Supply Chain Security

Every artefact published by this repository is signed with [Sigstore][sigstore]
keyless OIDC. There is no private key to lose, no PGP web of trust to manage:
signatures are bound to the GitHub Actions workflow identity (`repository_owner`,
`repository`, `ref`, `workflow`) and recorded in the [Rekor][rekor] public
transparency log.

[sigstore]: https://www.sigstore.dev
[rekor]: https://docs.sigstore.dev/logging/overview/

## What gets signed

| Artefact                                              | Tool                          | What proves authenticity                       |
| ----------------------------------------------------- | ----------------------------- | ---------------------------------------------- |
| Backend image `ghcr.io/vydon-io/vydon/api`            | `cosign` keyless              | Sigstore signature, SLSA provenance, SPDX SBOM |
| Frontend image `ghcr.io/vydon-io/vydon/app`           | `cosign` keyless              | Sigstore signature, SLSA provenance, SPDX SBOM |
| Worker image `ghcr.io/vydon-io/vydon/worker`          | `cosign` keyless              | Sigstore signature, SLSA provenance, SPDX SBOM |
| Helm OCI charts under `ghcr.io/vydon-io/vydon/helm/*` | `cosign` keyless              | Sigstore signature                             |
| Python SDK on PyPI                                    | `pypa/gh-action-pypi-publish` | [PEP 740][pep740] attestations                 |
| TypeScript SDK on npm                                 | `npm publish --provenance`    | [npm provenance][npmprov] (SLSA)               |

[pep740]: https://peps.python.org/pep-0740/
[npmprov]: https://docs.npmjs.com/generating-provenance-statements

## Verifying a container image

Install the [cosign CLI][cosign-install] (v2.x), then:

[cosign-install]: https://docs.sigstore.dev/cosign/system_config/installation/

```bash
IMAGE=ghcr.io/vydon-io/vydon/api:latest

cosign verify "$IMAGE" \
  --certificate-identity-regexp "^https://github.com/vydon-io/vydon/\.github/workflows/artifact-release\.yml@.+$" \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com
```

A successful run prints the signature, the OIDC identity that produced it, the
Rekor entry, and the inclusion proof. Any other output means the image was not
signed by the expected workflow and **must not be trusted**.

## Verifying the SLSA build provenance

Build provenance attestations live next to the image and can be inspected with
the GitHub CLI:

```bash
gh attestation verify oci://ghcr.io/vydon-io/vydon/api:latest \
  --owner vydon-io
```

`gh attestation` also accepts `--repo vydon-io/vydon` to pin verification to
this repository instead of the whole org.

## Extracting and verifying the SBOM

Each image ships with an SPDX-JSON SBOM as a signed attestation. Pull the
attestation and reuse it with any SBOM consumer (Grype, Trivy, Dependency-Track):

```bash
cosign download attestation \
  --predicate-type https://spdx.dev/Document \
  ghcr.io/vydon-io/vydon/api:latest \
  | jq -r '.payload' \
  | base64 -d \
  | jq '.predicate' \
  > sbom-api.spdx.json
```

Scan it with Grype:

```bash
grype sbom:./sbom-api.spdx.json
```

## Verifying a Helm chart

Helm OCI artefacts are signed the same way as container images.

```bash
CHART=ghcr.io/vydon-io/vydon/helm/vydon:0.4.0

cosign verify "$CHART" \
  --certificate-identity-regexp "^https://github.com/vydon-io/vydon/\.github/workflows/artifact-release\.yml@.+$" \
  --certificate-oidc-issuer https://token.actions.githubusercontent.com
```

Helm itself does not check OCI signatures yet ([open RFC][helm-rfc]). Run the
`cosign verify` step before `helm pull` / `helm install` in your release
pipeline.

[helm-rfc]: https://github.com/helm/community/blob/main/hips/hip-0007.md

## Verifying SDK packages

### Python (PyPI)

```bash
pip install vydon
python -m pip download vydon --no-deps -d /tmp/v
# pip verifies PEP 740 attestations automatically when supported
```

### TypeScript (npm)

```bash
npm install @vydon/sdk
npm view @vydon/sdk dist.attestations
```

## What "identity" means in keyless signing

Keyless signing pins the signature to:

- `--certificate-identity` — the GitHub Actions workflow URL
  (e.g. `https://github.com/vydon-io/vydon/.github/workflows/artifact-release.yml@refs/tags/v0.4.0`)
- `--certificate-oidc-issuer` — `https://token.actions.githubusercontent.com`

If either field does not match, the artefact was not produced by this
repository's release pipeline. Treat it as untrusted, regardless of digest or
tag.

## When verification fails

1. Confirm the artefact actually came from this repository (typos, supply-chain
   typo-squatting).
2. Make sure your `cosign` is at least v2.0 — earlier versions used the old
   flag names.
3. Check Rekor directly: <https://search.sigstore.dev/>
4. Open a security report via the process in [SECURITY.md](../../SECURITY.md).

## Reproducing the signatures locally

Signing happens in [`.github/workflows/artifact-release.yml`](../../.github/workflows/artifact-release.yml).
The relevant steps are:

- `sigstore/cosign-installer` — install cosign v2
- `cosign sign --yes <image>@<digest>` — Sigstore keyless signing
- `actions/attest-build-provenance@v2` — SLSA build provenance
- `actions/attest-sbom@v2` — SPDX SBOM attestation

The workflow holds no signing key. All trust derives from GitHub's OIDC
identity token, exchanged with Sigstore's Fulcio CA at signing time.
