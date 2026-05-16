package tablesync_workflow_register

import (
	"github.com/redis/go-redis/v9"
	"github.com/vydon-io/vydon/backend/gen/go/protos/mgmt/v1alpha1/mgmtv1alpha1connect"
	benthosstream "github.com/vydon-io/vydon/internal/benthos-stream"
	connectionmanager "github.com/vydon-io/vydon/internal/connection-manager"
	vydon_benthos_mongodb "github.com/vydon-io/vydon/worker/pkg/benthos/mongodb"
	vydon_benthos_sql "github.com/vydon-io/vydon/worker/pkg/benthos/sql"
	sync_activity "github.com/vydon-io/vydon/worker/pkg/workflows/tablesync/activities/sync"
	tablesync_workflow "github.com/vydon-io/vydon/worker/pkg/workflows/tablesync/workflow"
	"go.opentelemetry.io/otel/metric"
	"go.temporal.io/sdk/client"
)

type Worker interface {
	RegisterWorkflow(workflow any)
	RegisterActivity(activity any)
}

func Register(
	w Worker,
	connclient mgmtv1alpha1connect.ConnectionServiceClient,
	jobclient mgmtv1alpha1connect.JobServiceClient,
	sqlconnmanager connectionmanager.Interface[vydon_benthos_sql.SqlDbtx],
	mongoconnmanager connectionmanager.Interface[vydon_benthos_mongodb.MongoClient],
	meter metric.Meter, // optional
	benthosStreamManager benthosstream.BenthosStreamManagerClient,
	temporalclient client.Client,
	maxIterations int,
	anonymizationClient mgmtv1alpha1connect.AnonymizationServiceClient,
	redisclient redis.UniversalClient,
) {
	tsWf := tablesync_workflow.New(maxIterations)
	w.RegisterWorkflow(tsWf.TableSync)

	syncActivity := sync_activity.New(
		connclient,
		jobclient,
		sqlconnmanager,
		mongoconnmanager,
		meter,
		benthosStreamManager,
		temporalclient,
		anonymizationClient,
		redisclient,
	)

	w.RegisterActivity(syncActivity.Sync)
	w.RegisterActivity(syncActivity.SyncTable)
}
