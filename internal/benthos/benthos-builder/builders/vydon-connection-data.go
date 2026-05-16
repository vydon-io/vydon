package benthosbuilder_builders

import (
	"context"
	"errors"

	mgmtv1alpha1 "github.com/vydon-io/vydon/backend/gen/go/protos/mgmt/v1alpha1"
	"github.com/vydon-io/vydon/backend/gen/go/protos/mgmt/v1alpha1/mgmtv1alpha1connect"
	"github.com/vydon-io/vydon/backend/pkg/sqlmanager"
	sqlmanager_shared "github.com/vydon-io/vydon/backend/pkg/sqlmanager/shared"
	bb_internal "github.com/vydon-io/vydon/internal/benthos/benthos-builder/internal"
	bb_shared "github.com/vydon-io/vydon/internal/benthos/benthos-builder/shared"
	"github.com/vydon-io/vydon/internal/runconfigs"
	vydon_benthos "github.com/vydon-io/vydon/worker/pkg/benthos"
)

type vydonConnectionDataBuilder struct {
	connectiondataclient  mgmtv1alpha1connect.ConnectionDataServiceClient
	sqlmanagerclient      sqlmanager.SqlManagerClient
	sourceJobRunId        *string
	syncConfigs           []*runconfigs.RunConfig
	destinationConnection *mgmtv1alpha1.Connection
	sourceConnectionType  bb_shared.ConnectionType
}

func NewVydonConnectionDataSyncBuilder(
	connectiondataclient mgmtv1alpha1connect.ConnectionDataServiceClient,
	sqlmanagerclient sqlmanager.SqlManagerClient,
	sourceJobRunId *string,
	syncConfigs []*runconfigs.RunConfig,
	destinationConnection *mgmtv1alpha1.Connection,
	sourceConnectionType bb_shared.ConnectionType,
) bb_internal.BenthosBuilder {
	return &vydonConnectionDataBuilder{
		connectiondataclient:  connectiondataclient,
		sqlmanagerclient:      sqlmanagerclient,
		sourceJobRunId:        sourceJobRunId,
		syncConfigs:           syncConfigs,
		destinationConnection: destinationConnection,
		sourceConnectionType:  sourceConnectionType,
	}
}

func (b *vydonConnectionDataBuilder) BuildSourceConfigs(
	ctx context.Context,
	params *bb_internal.SourceParams,
) ([]*bb_internal.BenthosSourceConfig, error) {
	sourceConnection := params.SourceConnection
	job := params.Job
	configs := []*bb_internal.BenthosSourceConfig{}

	for _, config := range b.syncConfigs {
		schema, table := sqlmanager_shared.SplitTableKey(config.Table())

		bc := &vydon_benthos.BenthosConfig{
			StreamConfig: vydon_benthos.StreamConfig{
				Logger: &vydon_benthos.LoggerConfig{
					Level:        "ERROR",
					AddTimestamp: true,
				},
				Input: &vydon_benthos.InputConfig{
					Inputs: vydon_benthos.Inputs{
						VydonConnectionData: &vydon_benthos.VydonConnectionData{
							ConnectionId:   sourceConnection.GetId(),
							ConnectionType: string(b.sourceConnectionType),
							JobId:          &job.Id,
							JobRunId:       b.sourceJobRunId,
							Schema:         schema,
							Table:          table,
						},
					},
				},
				Pipeline: &vydon_benthos.PipelineConfig{},
				Output: &vydon_benthos.OutputConfig{
					Outputs: vydon_benthos.Outputs{
						Broker: &vydon_benthos.OutputBrokerConfig{
							Pattern: "fan_out",
							Outputs: []vydon_benthos.Outputs{},
						},
					},
				},
			},
		}
		configs = append(configs, &bb_internal.BenthosSourceConfig{
			Name:      config.Id(),
			Config:    bc,
			DependsOn: config.DependsOn(),
			RunType:   config.RunType(),

			BenthosDsns: []*bb_shared.BenthosDsn{{ConnectionId: sourceConnection.Id}},

			TableSchema: schema,
			TableName:   table,
			Columns:     config.InsertColumns(),
			PrimaryKeys: config.PrimaryKeys(),
		})
	}

	return configs, nil
}

func (b *vydonConnectionDataBuilder) BuildDestinationConfig(
	ctx context.Context,
	params *bb_internal.DestinationParams,
) (*bb_internal.BenthosDestinationConfig, error) {
	return nil, errors.ErrUnsupported
}
