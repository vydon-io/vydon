# EE Deletion Implementation Plan

This document is the executable plan for removing the inherited EE
(Enterprise Edition) code from the Vydon fork. Generated from the
inventory in [`ee-inventory.md`](ee-inventory.md). Execute in a dedicated
session with a clean working tree.

## Goal

Land a state where:

```
$ go build ./...
$ go test ./... -short
```

both pass, with **zero remaining imports** of `internal/ee/*`,
`backend/internal/ee/*`, or `worker/pkg/workflows/ee/*`.

## Strategy

The OSS-shipped Vydon codebase already contains OSS-compatible shims
inside the EE directories (most notably `internal/ee/rbac/allow_all_client.go`).
The migration is therefore:

1. **Extract** the OSS-compatible types and shims to non-`ee` locations.
2. **Stub** EE feature integrations (license check, slack notifications,
   account hooks, piidetect, presidio PII detection) as either no-ops or
   feature-disabled paths.
3. **Update imports** in the 49 OSS callers.
4. **Delete** the three `ee/` directories and their `LICENSE.md` files.
5. **Validate** with `go build ./...` and the existing test suite.

## Per-package migration

### `internal/ee/rbac` → `internal/rbac` (19 OSS importers)

**Extract to `internal/rbac/`:**

- `allow_all_client.go` (OSS shim — already implements the interface).
- `actions.go` (action enum strings).
- `entity.go` (Entity, EntityString, constructors).
- The `Interface` type definition from `client.go`.
- Wildcard constants from `entity.go`.

**Delete:**

- `internal/ee/rbac/client.go` (the Casbin-backed `Rbac` struct).
- `internal/ee/rbac/db.go` (policy persistence).
- `internal/ee/rbac/policy.go` (policy types).
- `internal/ee/rbac/roles.go` (role assignment).
- `internal/ee/rbac/enforcer/` (the Casbin enforcer).

**Caller migration:** replace
`"github.com/vydon-io/vydon/internal/ee/rbac"` with
`"github.com/vydon-io/vydon/internal/rbac"`. Any construction site
calling `rbac.New(enforcer)` must switch to `&rbac.AllowAllClient{}`.

**Drop dependency:** remove `github.com/casbin/casbin/v2` from `go.mod` if
no longer used anywhere.

### `internal/ee/license` → `internal/license` (17 OSS importers)

**Create OSS stub at `internal/license/license.go`:**

```go
package license

import "context"

// IsValid reports whether the license is valid. The OSS distribution
// always returns true; there is no proprietary feature to gate on.
func IsValid(_ context.Context) bool { return true }
```

**Caller migration:** replace each license check site with `license.IsValid(ctx)`
or delete the check entirely if it gated a feature being removed.

**Drop:** `EE_LICENSE` env var documentation, the bundled `.pem` keys
under `internal/ee/license/`.

### `internal/ee/cloud-license` (2 OSS importers)

The cloud license is a paid SaaS concept; it has no OSS analogue.

**Action:** delete entirely. Strip the 2 OSS callers' usage —
`backend/internal/cmds/mgmt/serve/connect/cmd.go` and one other —
replacing with hard-coded "not a cloud install" behavior.

### `internal/ee/events` (3 OSS importers)

Account hook event types.

**Action:** delete. The account-hook feature itself is removed in the
`backend/internal/ee/hooks` step below, so the event types lose their
only callers.

### `internal/ee/slack` (3 OSS importers)

Slack notifications for account hooks.

**Action:** delete. Same rationale as events — the feature gating slack
notifications is removed.

### `internal/ee/presidio` (10 OSS importers)

Microsoft Presidio API client for PII detection.

**Action:** delete. The PII detection feature is removed entirely from
the Vydon OSS surface. Callers in `internal/json-anonymizer/`,
`worker/pkg/benthos/transformer_executor/`, and
`backend/services/mgmt/v1alpha1/transformers-service/` must be edited
to either drop the PII-detection codepath or panic with a clear "not
supported in OSS" error.

### `internal/ee/mssql-manager` (2 OSS importers)

EE-only MSSQL transformer flavors.

**Action:** delete. The standard `backend/pkg/sqlmanager/mssql/` covers
the OSS MSSQL needs. Callers in `internal/schema-manager/mssql/` must be
edited to use the OSS mssql-manager.

### `internal/ee/transformers` (4 OSS importers)

EE-only transformer functions.

**Action:** delete. Strip the four OSS callers' usage.

### `backend/internal/ee/hooks` (4 OSS importers)

Account and job hook services.

**Action:** delete. The hooks feature is dropped in OSS. Remove all
references in `backend/internal/cmds/mgmt/serve/connect/cmd.go` and
the related service implementations.

### `worker/pkg/workflows/ee/account_hooks` (6 OSS importers)

Account hooks workflow definitions.

**Action:** delete. Strip references in
`worker/pkg/workflows/datasync/workflow/`, `worker/internal/cmds/worker/serve/serve.go`,
and the integration tests.

### `worker/pkg/workflows/ee/piidetect` (4 OSS importers)

PII detection workflows.

**Action:** delete. Strip references in the four worker files.

## Execution order

Run in this order to keep `go build` green at each commit:

1. Add OSS types to `internal/rbac/` (additive, breaks nothing).
2. Add OSS license stub at `internal/license/` (additive).
3. Migrate `rbac` callers and `license` callers (19 + 17 = 36 file edits).
4. Run `go build ./...`. Should still work because EE packages remain.
5. Strip `presidio` callers' PII-detection codepaths.
6. Strip `events`, `slack`, `cloud-license`, `mssql-manager`,
   `transformers` callers.
7. Strip `backend/internal/ee/hooks` callers (account hooks service path).
8. Strip `worker/pkg/workflows/ee/{account_hooks,piidetect}` callers.
9. Run `go build ./...` — should now pass with EE packages still present
   but unused.
10. Delete the three `ee/` directories.
11. Run `go build ./...` and `go test ./... -short`. Tag `phase-1-done`.

## Risk register

- **Casbin dependency:** if any non-EE code in the repo also uses Casbin,
  do not drop it from `go.mod`. Audit with `grep -r "casbin"` first.
- **Integration tests** under `internal/integration-tests/` exercise EE
  features (RBAC, account hooks). They will need to be deleted or
  rewritten alongside the feature removal.
- **Frontend protobufs** reference `account_hooks` types in
  `frontend/packages/sdk/`. Generated from `.proto` files. The proto
  removal is an EPIC-2 concern (rebranding + proto cleanup) — for now,
  the frontend will keep the types but no backend will respond.
- **Database migrations** for RBAC roles/policies may exist under
  `backend/internal/migrations/`. Leave the migrations in place; we do
  not roll back schemas, only stop writing to the related tables.

## Validation gates

Before merging the deletion PR:

- [ ] `go build ./...` exits 0
- [ ] `go vet ./...` exits 0
- [ ] `go test ./... -short` passes (excluding integration tests that
      gated EE features — those should be deleted, not made to pass)
- [ ] `grep -r "vydon-io/vydon/internal/ee\|vydon-io/vydon/worker/pkg/workflows/ee\|vydon-io/vydon/backend/internal/ee" --include="*.go" .` returns no matches
- [ ] `find . -type d -name ee -not -path "*/node_modules/*"` returns no matches
- [ ] CodeQL passes
- [ ] OSV-Scanner passes
- [ ] `helm lint charts/vydon` (no regression — the chart is still
      named vydon until EPIC-2 renames)
