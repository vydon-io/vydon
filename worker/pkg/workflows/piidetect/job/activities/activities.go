package activities

type JobPiiDetectReport struct {
	JobID                  string
	SuccessfulTableReports []TableReportRef
	FailedTableReports     []TableReportRef
}

type TableReportRef struct {
	Schema    string
	Table     string
	ReportKey string
}

const JobReportSuffix = "pii-job-report"

func BuildJobReportExternalId(jobID string) string {
	return jobID + "/" + JobReportSuffix
}
