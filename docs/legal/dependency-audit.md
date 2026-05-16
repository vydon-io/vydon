# Dependency License Audit

Audit performed on the Vydon fork of Vydon. Scope: all production and
development dependencies in Go, Node, and Python. Tool: [Trivy][trivy]
`v0.70.0` filesystem scanner.

[trivy]: https://github.com/aquasecurity/trivy

## Compatibility policy

Vydon ships under MIT. Dependencies are accepted if they fall under one
of:

- MIT
- Apache-2.0
- BSD-2-Clause, BSD-3-Clause
- MPL-2.0
- ISC
- 0BSD
- CC0-1.0
- BSL-1.0 (Boost — permissive)
- Zlib
- OpenSSL
- Public domain

Rejected: GPL, LGPL, AGPL, SSPL, BUSL, Elastic License, Server Side
Public License, anything custom-restrictive.

## Summary

| Ecosystem                              | Total deps       | Compatible | Flagged | Critical |
| -------------------------------------- | ---------------- | ---------- | ------- | -------- |
| Go (`go.mod`)                          | 286              | 286        | 0       | 0        |
| Node (`frontend/package-lock.json`)    | 173              | 160        | 13      | 12       |
| Node (`docs/package-lock.json`)        | 37               | 37         | 0       | 0        |
| Node (root `package-lock.json`)        | 0 (tooling-only) | n/a        | 0       | 0        |
| Python (`python/pyproject.toml`)       | 2                | 2          | 0       | 0        |
| Python (`worker/scripts/name_dataset`) | 1                | 1          | 0       | 0        |
| **Total**                              | **499**          | **486**    | **13**  | **12**   |

## Go (`go.mod`)

**Verdict: clean.** 286 dependencies, all permissive.

| License                                   | Count |
| ----------------------------------------- | ----- |
| Apache-2.0                                | 118   |
| MIT                                       | 105   |
| BSD-3-Clause                              | 38    |
| BSD-2-Clause                              | 10    |
| MPL-2.0                                   | 8     |
| BSL-1.0                                   | 1     |
| CC0-1.0                                   | 1     |
| ISC                                       | 1     |
| OpenSSL                                   | 1     |
| Zlib                                      | 1     |
| (Apache 2.0 with extra copyright headers) | 2     |

**Trivy flags clarified:**

- `github.com/apache/arrow/go/v15` — Apache 2.0 with embedded `LicenseRef-C-Ares`
  attribution. Apache Arrow ships Apache 2.0; the additional `c-ares`
  reference is upstream attribution. Permissive, no action needed.
- 8× MPL-2.0 packages (`hashicorp/*`, `go-sql-driver/mysql`,
  `shoenig/go-m1cpu`) — accepted per policy.

## Node (`frontend/package-lock.json`)

**Verdict: 12 packages need attention.** All are the `sharp` family (image
optimization used by Next.js).

| Package                              | License                                               | Status           |
| ------------------------------------ | ----------------------------------------------------- | ---------------- |
| `@img/sharp-libvips-darwin-arm64`    | LGPL-3.0-or-later                                     | flagged          |
| `@img/sharp-libvips-darwin-x64`      | LGPL-3.0-or-later                                     | flagged          |
| `@img/sharp-libvips-linux-arm`       | LGPL-3.0-or-later                                     | flagged          |
| `@img/sharp-libvips-linux-arm64`     | LGPL-3.0-or-later                                     | flagged          |
| `@img/sharp-libvips-linux-ppc64`     | LGPL-3.0-or-later                                     | flagged          |
| `@img/sharp-libvips-linux-s390x`     | LGPL-3.0-or-later                                     | flagged          |
| `@img/sharp-libvips-linux-x64`       | LGPL-3.0-or-later                                     | flagged          |
| `@img/sharp-libvips-linuxmusl-arm64` | LGPL-3.0-or-later                                     | flagged          |
| `@img/sharp-libvips-linuxmusl-x64`   | LGPL-3.0-or-later                                     | flagged          |
| `@img/sharp-wasm32`                  | Apache-2.0 AND LGPL-3.0-or-later AND MIT              | flagged          |
| `@img/sharp-win32-ia32`              | Apache-2.0 AND LGPL-3.0-or-later                      | flagged          |
| `@img/sharp-win32-x64`               | Apache-2.0 AND LGPL-3.0-or-later                      | flagged          |
| `posthog-js`                         | "SEE LICENSE IN LICENSE" (manually verified: **MIT**) | clarify metadata |

### `sharp` decision

`sharp` is the de-facto image optimization library for Node and ships as a
transitive of `next/image`. The native bindings wrap `libvips`, which is
LGPL-3.0. Two paths:

1. **Accept LGPL-3.0 (build-time + runtime as dynamically linked library).**
   This matches the position taken by Vercel itself when shipping Next.js
   and is the consensus in the Node ecosystem. Note: LGPL is not in our
   accepted list above, so accepting it requires an explicit carve-out
   documented in this file.

2. **Replace `next/image` with alternatives** (custom `<img>`,
   `@unpic/next`, Cloudflare Images, etc.). Loses optimization automation;
   forces hand-tuning of every image.

**Recommendation:** explicit carve-out for `sharp` (option 1). Add a note
to the policy. Tracked separately.

### `posthog-js` decision

Manual verification of the upstream
[`PostHog/posthog-js`](https://github.com/PostHog/posthog-js/blob/main/LICENSE)
LICENSE file shows MIT. The package's `license` field in `package.json`
uses the non-standard string `"SEE LICENSE IN LICENSE"`, which Trivy
correctly fails to parse. No action required beyond noting the
discrepancy.

## Node (`docs/package-lock.json`)

**Verdict: clean.** 37 dependencies, all MIT or Apache-2.0.

## Python

`python/pyproject.toml`:

| Package    | Version | License      | Status |
| ---------- | ------- | ------------ | ------ |
| `grpcio`   | 1.69.\* | Apache-2.0   | ok     |
| `protobuf` | 5.\*    | BSD-3-Clause | ok     |

`worker/scripts/name_dataset/requirements.txt`:

| Package         | Version | License    | Status |
| --------------- | ------- | ---------- | ------ |
| `names-dataset` | 3.1.0   | Apache-2.0 | ok     |

**Verdict: clean.**

## Reproduce

```bash
trivy fs --scanners license --license-full --format table go.mod
trivy fs --scanners license --license-full --format table frontend/package-lock.json
trivy fs --scanners license --license-full --format table docs/package-lock.json
trivy fs --scanners license --license-full --format table python/pyproject.toml
```

Trivy required: `>=v0.55.0`.

## Follow-ups

- **LGPL carve-out:** add `LGPL-3.0-or-later` allowance scoped strictly
  to `sharp`/`libvips` build-time bindings; document the rationale (Next.js
  ecosystem standard, dynamically linked, no user-facing code under LGPL).
- **Upstream fix (optional):** open a PR against `posthog-js` to fix the
  `package.json` license field to `"MIT"`.
- **Re-audit cadence:** rerun this audit on every dependency bump that
  changes more than 10 packages, and at minimum quarterly.
