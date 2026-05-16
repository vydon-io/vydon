// Package register is a no-op stub for the upstream piidetect Temporal
// workflow registration. The OSS distribution does not run PII
// detection workflows; Register accepts the same arguments for binary
// compatibility but does not wire any workflow.
package register

// Register accepts the upstream argument list and is a no-op. The first
// argument is the Temporal worker; the remaining arguments are clients
// and config used by the now-disabled PII detection feature.
func Register(_ ...any) {
	// no-op
}
