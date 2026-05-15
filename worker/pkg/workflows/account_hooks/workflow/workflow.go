// Package workflow is a no-op stub for the upstream account-hooks
// Temporal workflow. Reintroduced natively in a future release.
package workflow

import (
	"go.temporal.io/sdk/workflow"
)

type ProcessAccountHookRequest struct {
	Event any
}

type ProcessAccountHookResponse struct{}

// ProcessAccountHook is the Temporal workflow entrypoint. With the
// account-hooks feature disabled in OSS, the workflow returns
// immediately.
func ProcessAccountHook(_ workflow.Context, _ *ProcessAccountHookRequest) (*ProcessAccountHookResponse, error) {
	return &ProcessAccountHookResponse{}, nil
}
