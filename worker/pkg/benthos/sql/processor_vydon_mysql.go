package vydon_benthos_sql

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/redpanda-data/benthos/v4/public/service"
	"github.com/vydon-io/vydon/internal/gotypeutil"
	vydontypes "github.com/vydon-io/vydon/internal/vydon-types"
	vydon_benthos "github.com/vydon-io/vydon/worker/pkg/benthos"
)

func vydonToMysqlProcessorConfig() *service.ConfigSpec {
	return service.NewConfigSpec().
		Field(service.NewStringListField("columns")).
		Field(service.NewStringMapField("column_data_types")).
		Field(service.NewAnyMapField("column_default_properties"))
}

func RegisterVydonToMysqlProcessor(env *service.Environment) error {
	return env.RegisterBatchProcessor(
		"vydon_to_mysql",
		vydonToMysqlProcessorConfig(),
		func(conf *service.ParsedConfig, mgr *service.Resources) (service.BatchProcessor, error) {
			proc, err := newVydonToMysqlProcessor(conf, mgr)
			if err != nil {
				return nil, err
			}
			return proc, nil
		})
}

type vydonToMysqlProcessor struct {
	logger                  *service.Logger
	columns                 []string
	columnDataTypes         map[string]string
	columnDefaultProperties map[string]*vydon_benthos.ColumnDefaultProperties
}

func newVydonToMysqlProcessor(
	conf *service.ParsedConfig,
	mgr *service.Resources,
) (*vydonToMysqlProcessor, error) {
	columns, err := conf.FieldStringList("columns")
	if err != nil {
		return nil, err
	}

	columnDataTypes, err := conf.FieldStringMap("column_data_types")
	if err != nil {
		return nil, err
	}

	columnDefaultPropertiesConfig, err := conf.FieldAnyMap("column_default_properties")
	if err != nil {
		return nil, err
	}

	columnDefaultProperties, err := getColumnDefaultProperties(columnDefaultPropertiesConfig)
	if err != nil {
		return nil, err
	}

	return &vydonToMysqlProcessor{
		logger:                  mgr.Logger(),
		columns:                 columns,
		columnDataTypes:         columnDataTypes,
		columnDefaultProperties: columnDefaultProperties,
	}, nil
}

func (p *vydonToMysqlProcessor) ProcessBatch(
	ctx context.Context,
	batch service.MessageBatch,
) ([]service.MessageBatch, error) {
	newBatch := make(service.MessageBatch, 0, len(batch))
	for _, msg := range batch {
		root, err := msg.AsStructuredMut()
		if err != nil {
			return nil, err
		}
		newRoot, err := transformVydonToMysql(
			root,
			p.columns,
			p.columnDataTypes,
			p.columnDefaultProperties,
		)
		if err != nil {
			return nil, err
		}
		newMsg := msg.Copy()
		newMsg.SetStructured(newRoot)
		newBatch = append(newBatch, newMsg)
	}

	if len(newBatch) == 0 {
		return nil, nil
	}
	return []service.MessageBatch{newBatch}, nil
}

func (m *vydonToMysqlProcessor) Close(context.Context) error {
	return nil
}

func transformVydonToMysql(
	root any,
	columns []string,
	columnDataTypes map[string]string,
	columnDefaultProperties map[string]*vydon_benthos.ColumnDefaultProperties,
) (map[string]any, error) {
	rootMap, ok := root.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("root value must be a map[string]any")
	}

	newMap := make(map[string]any)
	for col, val := range rootMap {
		// Skip values that aren't in the column list to handle circular references
		if !isColumnInList(col, columns) {
			continue
		}
		colDefaults := columnDefaultProperties[col]
		datatype := columnDataTypes[col]
		newVal, err := getMysqlValue(val, colDefaults, datatype)
		if err != nil {
			return nil, fmt.Errorf("failed to get MySQL value for column %s: %w", col, err)
		}
		newMap[col] = newVal
	}

	return newMap, nil
}

func getMysqlValue(
	value any,
	colDefaults *vydon_benthos.ColumnDefaultProperties,
	datatype string,
) (any, error) {
	if colDefaults != nil && colDefaults.HasDefaultTransformer {
		return goqu.Default(), nil
	}

	if value == nil {
		return nil, nil
	}

	value, isVydonValue, err := getMysqlVydonValue(value)
	if err != nil {
		return nil, fmt.Errorf("unable to get MySQL value from vydon value: %w", err)
	}
	if isVydonValue {
		return value, nil
	}

	if datatype == "json" {
		if v, ok := value.([]byte); ok {
			validJson, err := getValidJson(v)
			if err != nil {
				return nil, fmt.Errorf("unable to get valid json: %w", err)
			}
			return validJson, nil
		}
		if value == "null" {
			return value, nil
		}
		bits, err := json.Marshal(value)
		if err != nil {
			return nil, fmt.Errorf("unable to marshal mysql json to bits: %w", err)
		}
		return bits, nil
	}

	if gotypeutil.IsMap(value) {
		bits, err := json.Marshal(value)
		if err != nil {
			return nil, fmt.Errorf("unable to marshal go map to json bits: %w", err)
		}
		return bits, nil
	}

	return value, nil
}

func getMysqlVydonValue(root any) (value any, isVydonValue bool, err error) {
	if valuer, ok := root.(vydontypes.VydonMysqlValuer); ok {
		value, err := valuer.ValueMysql()
		if err != nil {
			return nil, false, fmt.Errorf(
				"unable to get MYSQL value from VydonMysqlValuer: %w",
				err,
			)
		}
		return value, true, nil
	}
	return root, false, nil
}
