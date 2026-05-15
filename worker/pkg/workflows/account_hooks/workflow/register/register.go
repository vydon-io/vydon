// Package register stubs the account-hooks workflow registration. The
// signature accepts variadic extras so it tolerates the upstream call
// sites that pass an AccountHookServiceClient or other dependencies that
// the OSS no-op workflow does not need.
package register

import (
	accounthook_workflow "github.com/vydon-io/vydon/worker/pkg/workflows/account_hooks/workflow"
)

// temporalWorker is the minimal Temporal worker contract used here. It
// accepts both the real worker.Worker and the testsuite environment
// callers pass in.
type temporalWorker interface {
	RegisterWorkflow(workflow any)
}

func Register(w temporalWorker, _ ...any) {
	w.RegisterWorkflow(accounthook_workflow.ProcessAccountHook)
}
