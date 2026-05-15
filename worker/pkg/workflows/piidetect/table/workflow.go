// Package table is a no-op stub for the upstream table-level PII
// detect workflow. Reintroduced natively in a future release.
package table

type TablePiiDetectRequest struct {
	JobID     string
	AccountID string
	Schema    string
	Table     string
}
