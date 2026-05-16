package benthosbuilder_builders

import (
	"context"
	"fmt"

	"github.com/vydon-io/vydon/backend/gen/go/protos/mgmt/v1alpha1/mgmtv1alpha1connect"
	"github.com/vydon-io/vydon/backend/pkg/metrics"
	sqlmanager_shared "github.com/vydon-io/vydon/backend/pkg/sqlmanager/shared"
	bb_internal "github.com/vydon-io/vydon/internal/benthos/benthos-builder/internal"
	bb_shared "github.com/vydon-io/vydon/internal/benthos/benthos-builder/shared"
	"github.com/vydon-io/vydon/internal/runconfigs"
	vydon_benthos "github.com/vydon-io/vydon/worker/pkg/benthos"
)

type mongodbSyncBuilder struct {
	transformerclient mgmtv1alpha1connect.TransformersServiceClient
}

func NewMongoDbSyncBuilder(
	transformerclient mgmtv1alpha1connect.TransformersServiceClient,
) bb_internal.BenthosBuilder {
	return &mongodbSyncBuilder{
		transformerclient: transformerclient,
	}
}

func (b *mongodbSyncBuilder) BuildSourceConfigs(
	ctx context.Context,
	params *bb_internal.SourceParams,
) ([]*bb_internal.BenthosSourceConfig, error) {
	sourceConnection := params.SourceConnection
	job := params.Job
	groupedMappings := groupMappingsByTable(job.GetMappings())

	benthosConfigs := []*bb_internal.BenthosSourceConfig{}
	for _, tableMapping := range groupedMappings {
		bc := &vydon_benthos.BenthosConfig{
			StreamConfig: vydon_benthos.StreamConfig{
				Input: &vydon_benthos.InputConfig{
					Inputs: vydon_benthos.Inputs{
						PooledMongoDB: &vydon_benthos.InputMongoDb{
							ConnectionId: params.SourceConnection.GetId(),
							Database:     tableMapping.Schema,
							Collection:   tableMapping.Table,
							Query:        "root = this",
						},
					},
				},
				Pipeline: &vydon_benthos.PipelineConfig{
					Threads:    -1,
					Processors: []vydon_benthos.ProcessorConfig{},
				},
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

		columns := []string{}
		for _, jm := range tableMapping.Mappings {
			columns = append(columns, jm.Column)
		}

		schemaTable := sqlmanager_shared.SchemaTable{
			Schema: tableMapping.Schema,
			Table:  tableMapping.Table,
		}
		runconfigType := runconfigs.RunTypeInsert
		runconfigId := fmt.Sprintf("%s.%s", schemaTable.String(), runconfigType)
		splitColumnPaths := true
		processorConfigs, err := buildProcessorConfigsByRunType(
			ctx,
			b.transformerclient,
			runconfigs.NewRunConfig(
				runconfigId,
				schemaTable,
				runconfigType,
				[]string{},
				nil,
				columns,
				columns,
				nil,
				splitColumnPaths,
			),
			map[string][]*bb_internal.ReferenceKey{},
			map[string][]*bb_internal.ReferenceKey{},
			params.Job.Id,
			params.JobRunId,
			tableMapping.Mappings,
			map[string]*sqlmanager_shared.DatabaseSchemaRow{},
			job.GetSource().GetOptions(),
			columns,
		)
		if err != nil {
			return nil, err
		}
		for _, pc := range processorConfigs {
			bc.Pipeline.Processors = append(bc.Pipeline.Processors, *pc)
		}

		benthosConfigs = append(benthosConfigs, &bb_internal.BenthosSourceConfig{
			Config:      bc,
			Name:        fmt.Sprintf("%s.%s", tableMapping.Schema, tableMapping.Table), // todo
			TableSchema: tableMapping.Schema,
			TableName:   tableMapping.Table,
			RunType:     runconfigs.RunTypeInsert,
			DependsOn:   []*runconfigs.DependsOn{},
			Columns:     columns,
			BenthosDsns: []*bb_shared.BenthosDsn{{ConnectionId: sourceConnection.GetId()}},

			Metriclabels: metrics.MetricLabels{
				metrics.NewEqLabel(metrics.TableSchemaLabel, tableMapping.Schema),
				metrics.NewEqLabel(metrics.TableNameLabel, tableMapping.Table),
				metrics.NewEqLabel(metrics.JobTypeLabel, "sync"),
			},
		})
	}

	return benthosConfigs, nil
}

func (b *mongodbSyncBuilder) BuildDestinationConfig(
	ctx context.Context,
	params *bb_internal.DestinationParams,
) (*bb_internal.BenthosDestinationConfig, error) {
	config := &bb_internal.BenthosDestinationConfig{}

	benthosConfig := params.SourceConfig
	config.BenthosDsns = append(
		config.BenthosDsns,
		&bb_shared.BenthosDsn{ConnectionId: params.DestConnection.GetId()},
	)
	config.Outputs = append(
		config.Outputs,
		vydon_benthos.Outputs{PooledMongoDB: &vydon_benthos.OutputMongoDb{
			ConnectionId: params.DestConnection.GetId(),

			Database:   benthosConfig.TableSchema,
			Collection: benthosConfig.TableName,
			Operation:  "update-one",
			Upsert:     true,
			DocumentMap: `
			root = {
				"$set": this
			}
		`,
			FilterMap: `
			root._id = this._id
		`,
			WriteConcern: &vydon_benthos.MongoWriteConcern{
				W: "1",
			},
		},
		},
	)
	return config, nil
}
