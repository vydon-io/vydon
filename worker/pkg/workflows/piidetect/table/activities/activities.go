// Package activities holds the PII detect table-level Temporal activity
// types referenced by the OSS job-service. With the PII detection
// feature disabled, the types remain for binary compatibility but no
// runtime work is performed.
package activities

import (
	mgmtv1alpha1 "github.com/vydon-io/vydon/backend/gen/go/protos/mgmt/v1alpha1"
)

const (
	PiiTableReportSuffix = "pii-table-report"
)

type PiiCategory string

const PiiCategoryPersonal PiiCategory = "personal"

func (c PiiCategory) String() string { return string(c) }

func BuildTableReportExternalId(args ...string) string {
	out := PiiTableReportSuffix
	for _, a := range args {
		out = a + "/" + out
	}
	return out
}

type RegexPiiDetectReport struct {
	Category PiiCategory
}

type LLMPiiDetectReport struct {
	Category   PiiCategory
	Confidence float32
}

type CombinedPiiDetectReport struct {
	Regex *RegexPiiDetectReport
	LLM   *LLMPiiDetectReport
}

type ColumnReport struct {
	ColumnName string
	Report     CombinedPiiDetectReport
}

type TableReport struct {
	TableSchema   string
	TableName     string
	ColumnReports []ColumnReport
	ReportKey     *mgmtv1alpha1.RunContextKey
}
