package job

type PiiDetectRequest struct {
	JobID     string
	AccountID string
}

// Workflow is the Temporal workflow type name. The PII detection
// workflow is disabled in OSS.
const Workflow = "PiiDetectJob"
