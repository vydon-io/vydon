package job

import (
	"go.temporal.io/sdk/workflow"
)

type PiiDetectRequest struct {
	JobId     string
	AccountID string
}

type PiiDetectResponse struct{}

type Workflow struct{}

// JobPiiDetect is the Temporal workflow entrypoint. With the PII
// detection feature disabled in OSS, the workflow returns immediately.
func (*Workflow) JobPiiDetect(_ workflow.Context, _ *PiiDetectRequest) (*PiiDetectResponse, error) {
	return &PiiDetectResponse{}, nil
}
