# Enterprise Edition Inventory

Passive inventory of EE-licensed material inherited from the Vydon fork.
This document is the source-of-truth for the upcoming EE deletion work.
**No code is removed in this change.**

Scope: anything under `*/ee/` directories, files named `*_ee*` or `*-ee*`,
Go files with `//go:build ee` tags, and OSS source files importing EE
packages.

## Summary

| Category                          | Count                                        |
| --------------------------------- | -------------------------------------------- |
| Top-level EE directories          | 3                                            |
| EE Go source files                | 53                                           |
| EE LICENSE files                  | 3                                            |
| OSS files importing EE packages   | 49                                           |
| EE helper scripts                 | 2                                            |
| Go files with `//go:build ee` tag | 0                                            |
| Charts/Helm EE-only values        | 0                                            |
| Terraform EE modules              | 0 (no `terraform/` directory in this repo)   |
| Frontend EE pages/components      | 0 (UI is OSS-only; EE was server-side gated) |

## EE directories

```
internal/ee/                       # transformers, license, RBAC, presidio, slack, mssql-manager
backend/internal/ee/               # account/job hooks
worker/pkg/workflows/ee/           # account_hooks, piidetect workflows
```

## EE LICENSE files (preserved verbatim per MIT attribution)

```
internal/ee/LICENSE.md
backend/internal/ee/LICENSE.md
worker/pkg/workflows/ee/LICENSE.md
```

These will be deleted along with their parent directories. The upstream
license history is preserved in `git log`.

## EE Go source files

### `internal/ee/`

```
internal/ee/rbac/roles.go
internal/ee/rbac/allow_all_client.go
internal/ee/rbac/db.go
internal/ee/rbac/policy.go
internal/ee/rbac/actions.go
internal/ee/rbac/client.go
internal/ee/rbac/entity.go
internal/ee/rbac/enforcer/enforcer.go
internal/ee/rbac/enforcer/rbac_model.conf
internal/ee/rbac/enforcer/model.go
internal/ee/license/license.go
internal/ee/license/license_test.go
internal/ee/license/cascade.go
internal/ee/license/vydon_ee_pub.pem
internal/ee/cloud-license/license.go
internal/ee/cloud-license/license_test.go
internal/ee/cloud-license/vydon_cloud_pub.pem
internal/ee/mssql-manager/ee-mssql-manager.go
internal/ee/mssql-manager/ee-mssql-manager_test.go
internal/ee/mssql-manager/generate-sql.go
internal/ee/presidio/client.gen.go
internal/ee/presidio/generate.go
internal/ee/presidio/interface.go
internal/ee/presidio/util.go
internal/ee/presidio/mock_AnalyzeInterface.go
internal/ee/presidio/mock_AnonymizeInterface.go
internal/ee/presidio/mock_EntityInterface.go
internal/ee/presidio/mock_RecognizerInterface.go
internal/ee/presidio/config/api.yml
internal/ee/presidio/config/clientconfig.yml
internal/ee/slack/slack.go
internal/ee/slack/slack_test.go
internal/ee/slack/mock_Interface.go
internal/ee/transformers/transformers.go
internal/ee/transformers/functions/functions.go
internal/ee/transformers/functions/functions_test.go
internal/ee/transformers/functions/mock_VydonOperatorApi.go
internal/ee/events/events.go
```

### `backend/internal/ee/`

```
backend/internal/ee/hooks/accounts/service.go
backend/internal/ee/hooks/accounts/service_test.go
backend/internal/ee/hooks/jobs/service.go
```

### `worker/pkg/workflows/ee/`

```
worker/pkg/workflows/ee/account_hooks/workflow/workflow.go
worker/pkg/workflows/ee/account_hooks/workflow/workflow_test.go
worker/pkg/workflows/ee/account_hooks/workflow/register/register.go
worker/pkg/workflows/ee/account_hooks/activities/execute/activity.go
worker/pkg/workflows/ee/account_hooks/activities/execute/activity_test.go
worker/pkg/workflows/ee/account_hooks/activities/hooks-by-event/activity.go
worker/pkg/workflows/ee/account_hooks/activities/hooks-by-event/activity_test.go
worker/pkg/workflows/ee/piidetect/workflows/register/register.go
worker/pkg/workflows/ee/piidetect/workflows/job/workflow.go
worker/pkg/workflows/ee/piidetect/workflows/job/workflow_test.go
worker/pkg/workflows/ee/piidetect/workflows/job/activities/activities.go
worker/pkg/workflows/ee/piidetect/workflows/job/activities/activities_test.go
worker/pkg/workflows/ee/piidetect/workflows/table/workflow.go
worker/pkg/workflows/ee/piidetect/workflows/table/workflow_test.go
worker/pkg/workflows/ee/piidetect/workflows/table/activities/activities.go
worker/pkg/workflows/ee/piidetect/workflows/table/activities/activities_test.go
worker/pkg/workflows/ee/piidetect/workflows/table/activities/mock_OpenAiCompletionsClient.go
```

## EE helper scripts (to delete)

```
scripts/gen-license.md       # documents EE license generation
scripts/gen-cust-license.sh  # generates customer EE licenses
```

## OSS source files importing EE packages (49)

These OSS files reference EE packages and must be edited to either drop the
EE dependency or replace it with an OSS-only equivalent. Grouped by package:

### Backend services and userdata (`rbac`, `license`, `events`)

```
backend/internal/cmds/mgmt/serve/connect/cmd.go
backend/internal/userdata/client.go
backend/internal/userdata/entity.go
backend/internal/userdata/entity_enforcer.go
backend/internal/userdata/mock_EntityEnforcer.go
backend/internal/userdata/user.go
backend/services/mgmt/v1alpha1/account-hooks-service/service.go
backend/services/mgmt/v1alpha1/anonymization-service/service.go
backend/services/mgmt/v1alpha1/api-key-service/api-keys.go
backend/services/mgmt/v1alpha1/connection-service/connection.go
backend/services/mgmt/v1alpha1/job-service/jobs.go
backend/services/mgmt/v1alpha1/job-service/runs.go
backend/services/mgmt/v1alpha1/job-service/service.go
backend/services/mgmt/v1alpha1/metrics-service/metrics.go
backend/services/mgmt/v1alpha1/transformers-service/entities.go
backend/services/mgmt/v1alpha1/transformers-service/service.go
backend/services/mgmt/v1alpha1/transformers-service/system_transformers.go
backend/services/mgmt/v1alpha1/transformers-service/userdefined_transformers.go
backend/services/mgmt/v1alpha1/user-account-service/account-onboarding.go
backend/services/mgmt/v1alpha1/user-account-service/account-temporal-config.go
backend/services/mgmt/v1alpha1/user-account-service/billing.go
backend/services/mgmt/v1alpha1/user-account-service/service.go
backend/services/mgmt/v1alpha1/user-account-service/users.go
```

### Backend integration tests (`rbac`, `events`, `slack`, `presidio`)

```
backend/pkg/integration-test/integration-test.go
backend/pkg/integration-test/mux.go
backend/pkg/sqlmanager/mssql/mssql-manager.go
```

### Schema manager (`license`, `mssql-manager`)

```
internal/schema-manager/schema-manager.go
internal/schema-manager/mssql/mssql.go
internal/json-anonymizer/json-anonymizer.go
```

### Integration test harness (`rbac`, `events`, `slack`, `presidio`, `piidetect`, `account_hooks`)

```
internal/integration-tests/rbac/rbac_integration_test.go
internal/integration-tests/api/jobs-service_integration_test.go
internal/integration-tests/api/account-hooks-service_integration_test.go
internal/integration-tests/api/transformers-service_integration_test.go
internal/integration-tests/worker/workflow/datasync-workflow.go
internal/integration-tests/worker/workflow/process-account-hook-workflow-integration_test.go
```

### Worker (`license`, `presidio`, `account_hooks`, `piidetect`)

```
worker/internal/cmds/worker/serve/serve.go
worker/pkg/benthos/transformer_executor/executor.go
worker/pkg/benthos/transformer_executor/transform_pii_text_api.go
worker/pkg/workflows/datasync/workflow/register/register.go
worker/pkg/workflows/datasync/workflow/workflow.go
worker/pkg/workflows/schemainit/activities/init-schema/activity.go
worker/pkg/workflows/schemainit/activities/init-schema/init-schema.go
worker/pkg/workflows/schemainit/activities/reconcile-schema/activity.go
worker/pkg/workflows/schemainit/activities/reconcile-schema/reconcile-schema.go
worker/pkg/workflows/schemainit/workflow/register/register.go
```

## EE-related environment variables (to remove)

| Variable     | File                             | Action                         |
| ------------ | -------------------------------- | ------------------------------ |
| `EE_LICENSE` | `internal/ee/license/license.go` | Deleted with the EE directory. |

## Charts / Helm

No EE-specific values, templates, or feature toggles found in `charts/`. The
default Helm chart is OSS-clean today.

## Compose / Tilt / scripts

- `compose/compose-metrics.yml:57` references `grafana/grafana-enterprise`.
  This is a third-party Docker image, not an EE feature of Vydon — flagged
  but **not** subject to deletion in this cleanup. The image will be re-evaluated
  in EPIC-2 rebranding.

## Frontend

No EE pages, components, or feature flags found under `frontend/apps/web`
or `frontend/packages/*`. Server-side EE features were gated entirely on
the API; the frontend will respond gracefully once the EE endpoints are
removed (404 → fall back to OSS behavior).

The single string `enterprise` in
`frontend/packages/sdk/src/client/mgmt/v1alpha1/user_account_pb.ts` refers
to the protobuf concept of an "enterprise account" — an organizational
account type, not an EE feature. Keep.

## Documentation

Docs referencing EE features (cleanup pending):

```
docs/protos/home.md
docs/docs/deploy/kubernetes.md
docs/docs/guides/vydon-local-dev.md
frontend/apps/web/README.md
```

## Deletion plan

1. **Backend Go**: delete `internal/ee/` and `backend/internal/ee/`, edit
   the 33 OSS files importing them in `backend/` and `internal/`. Goal:
   `go build ./...` passes after the change.

2. **Worker**: delete `worker/pkg/workflows/ee/`, edit the 10 OSS worker
   files. Workflows still register and compile.

3. **Frontend**: no-op — no EE code in the frontend. Close with this
   inventory linked.

4. **Helm / Terraform / scripts**: no-op for Helm and Terraform.
   Delete `scripts/gen-license.md` and `scripts/gen-cust-license.sh`.

5. **Docs**: rewrite the four docs files that reference EE features to
   drop EE-only sections.
