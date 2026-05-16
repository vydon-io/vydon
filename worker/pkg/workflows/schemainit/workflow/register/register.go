package schemainit_workflow_register

import (
	"github.com/vydon-io/vydon/backend/gen/go/protos/mgmt/v1alpha1/mgmtv1alpha1connect"
	sql_manager "github.com/vydon-io/vydon/backend/pkg/sqlmanager"
	"github.com/vydon-io/vydon/internal/license"
	initschema_activity "github.com/vydon-io/vydon/worker/pkg/workflows/schemainit/activities/init-schema"
	reconcileschema_activity "github.com/vydon-io/vydon/worker/pkg/workflows/schemainit/activities/reconcile-schema"
	schemainit_workflow "github.com/vydon-io/vydon/worker/pkg/workflows/schemainit/workflow"
)

type Worker interface {
	RegisterWorkflow(workflow any)
	RegisterActivity(activity any)
}

func Register(
	w Worker,
	jobclient mgmtv1alpha1connect.JobServiceClient,
	connclient mgmtv1alpha1connect.ConnectionServiceClient,
	sqlmanager *sql_manager.SqlManager,
	eelicense license.EEInterface,
) {
	runSqlInitTableStatements := initschema_activity.New(
		jobclient,
		connclient,
		sqlmanager,
		eelicense,
	)
	runReconcileSchema := reconcileschema_activity.New(jobclient, connclient, sqlmanager, eelicense)
	siWf := schemainit_workflow.New()
	w.RegisterWorkflow(siWf.SchemaInit)
	w.RegisterActivity(runSqlInitTableStatements.RunSqlInitTableStatements)
	w.RegisterActivity(runReconcileSchema.RunReconcileSchema)
}
