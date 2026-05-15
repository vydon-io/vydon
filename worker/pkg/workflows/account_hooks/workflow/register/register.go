package register

import (
	"go.temporal.io/sdk/worker"

	accounthook_workflow "github.com/vydon-io/vydon/worker/pkg/workflows/account_hooks/workflow"
)

// Register wires the account-hooks workflow on a Temporal worker. The
// workflow body is a no-op in the OSS distribution.
func Register(w worker.Worker) {
	w.RegisterWorkflow(accounthook_workflow.ProcessAccountHook)
}
