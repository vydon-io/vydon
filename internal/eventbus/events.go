// Package eventbus is a no-op stub for the upstream Account Hook event
// types. These types are emitted by the worker on job run lifecycle
// transitions and consumed by the (now-disabled) account hooks system.
// Reintroduced as a native event bus in a future release.
package eventbus

type Event struct {
	Name    string
	Payload map[string]any
}

func NewEvent_JobRunCreated(jobID, runID, accountID string) *Event {
	return &Event{Name: "job.run.created", Payload: map[string]any{
		"jobId": jobID, "runId": runID, "accountId": accountID,
	}}
}

func NewEvent_JobRunSucceeded(jobID, runID, accountID string) *Event {
	return &Event{Name: "job.run.succeeded", Payload: map[string]any{
		"jobId": jobID, "runId": runID, "accountId": accountID,
	}}
}

func NewEvent_JobRunFailed(jobID, runID, accountID string) *Event {
	return &Event{Name: "job.run.failed", Payload: map[string]any{
		"jobId": jobID, "runId": runID, "accountId": accountID,
	}}
}
