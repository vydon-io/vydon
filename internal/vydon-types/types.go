package vydontypes

import (
	"encoding/json"
	"fmt"
)

type Version uint

const (
	V1            Version = iota + 1
	LatestVersion         = V1
)

const (
	VydonArrayId    = "VYDON_ARRAY"
	VydonBitsId     = "VYDON_BIT"
	VydonBinaryId   = "VYDON_BINARY"
	VydonDateTimeId = "VYDON_DATETIME"
	VydonIntervalId = "VYDON_INTERVAL"
)

type VydonAdapter interface {
	VydonMetadataType
	// Pgx
	ScanPgx(value any) error
	ValuePgx() (any, error)
	// Json
	ScanJson(value any) error
	ValueJson() (any, error)
	// Mysql
	ScanMysql(value any) error
	ValueMysql() (any, error)
	// Mssql
	ScanMssql(value any) error
	ValueMssql() (any, error)
}

type VydonPgxValuer interface {
	ValuePgx() (any, error)
}

type VydonMysqlValuer interface {
	ValueMysql() (any, error)
}

type VydonMssqlValuer interface {
	ValueMssql() (any, error)
}

type VydonJsonValuer interface {
	ValueJson() (any, error)
}

type Vydon struct {
	Version Version `json:"version"`
	TypeId  string  `json:"type_id"`
}
type BaseType struct {
	Vydon Vydon `json:"_vydon"`
}

type VydonMetadataType interface {
	setVersion(Version)
	GetVersion() Version
}

type VydonTypeOption func(VydonAdapter) error

func WithVersion(version Version) VydonTypeOption {
	return func(t VydonAdapter) error {
		if !IsValidVersion(version) {
			return fmt.Errorf("invalid Vydon Type version: %d", version)
		}
		if version == 0 {
			t.setVersion(LatestVersion)
			return nil
		}
		t.setVersion(version)
		return nil
	}
}

func applyOptions(t VydonAdapter, opts ...VydonTypeOption) error {
	for _, opt := range opts {
		if err := opt(t); err != nil {
			return err
		}
	}
	return nil
}

func IsValidVersion(ver Version) bool {
	return ver == V1 || ver == LatestVersion
}

type JsonScanner struct{}

func (js *JsonScanner) ScanJson(value, target any) error {
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, target)
	case string:
		return json.Unmarshal([]byte(v), target)
	default:
		return fmt.Errorf("unsupported scan type for Json: %T", value)
	}
}

func (js *JsonScanner) ValueJson(value any) (any, error) {
	return json.Marshal(value)
}
