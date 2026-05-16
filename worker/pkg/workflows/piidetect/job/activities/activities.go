package activities

import (
	mgmtv1alpha1 "github.com/vydon-io/vydon/backend/gen/go/protos/mgmt/v1alpha1"
)

type JobPiiDetectReport struct {
	JobID                  string
	SuccessfulTableReports []TableReportRef
	FailedTableReports     []TableReportRef
}

type TableReportRef struct {
	TableSchema string
	TableName   string
	ReportKey   *mgmtv1alpha1.RunContextKey
}

const JobReportSuffix = "pii-job-report"

func BuildJobReportExternalId(jobID string) string {
	return jobID + "/" + JobReportSuffix
}
