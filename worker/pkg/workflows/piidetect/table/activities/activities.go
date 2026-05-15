// Package activities holds the PII detect table-level Temporal activity
// types referenced by the OSS job-service. With the PII detection
// feature disabled, the types remain for binary compatibility but no
// runtime work is performed.
package activities

const (
	PiiTableReportSuffix = "pii-table-report"
)

type PiiCategory string

const PiiCategoryPersonal PiiCategory = "personal"

func BuildTableReportExternalId(jobID, schema, table string) string {
	return jobID + "/" + schema + "/" + table + "/" + PiiTableReportSuffix
}

type RegexPiiDetectReport struct {
	Category PiiCategory
}

func (c PiiCategory) String() string { return string(c) }

type LLMPiiDetectReport struct {
	Category   PiiCategory
	Confidence float64
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
	Schema        string
	Table         string
	ColumnReports []*ColumnReport
	ReportKey     string
}
