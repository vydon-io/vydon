package vydontypes

import (
	"fmt"

	"github.com/lib/pq"
	"github.com/vydon-io/vydon/internal/gotypeutil"
)

type VydonArray struct {
	BaseType `                 json:",inline"`
	Elements []VydonAdapter `json:"elements"`
}

func NewVydonArray(
	elements []VydonAdapter,
	opts ...VydonTypeOption,
) (*VydonArray, error) {
	pgArray := &VydonArray{
		Elements: elements,
	}
	pgArray.Vydon.TypeId = VydonArrayId
	pgArray.setVersion(LatestVersion)

	if err := applyOptions(pgArray, opts...); err != nil {
		return nil, err
	}

	return pgArray, nil
}

func (a *VydonArray) setVersion(v Version) {
	a.Vydon.Version = v
}

func (a *VydonArray) GetVersion() Version {
	return a.Vydon.Version
}

func (a *VydonArray) ScanPgx(value any) error {
	valueSlice, err := gotypeutil.ParseSlice(value)
	if err != nil {
		return err
	}
	if len(valueSlice) != len(a.Elements) {
		return fmt.Errorf(
			"length mismatch: got %d elements, expected %d",
			len(valueSlice),
			len(a.Elements),
		)
	}
	for i, v := range valueSlice {
		if err := a.Elements[i].ScanPgx(v); err != nil {
			return fmt.Errorf("scanning element %d: %w", i, err)
		}
	}
	return nil
}

func (a *VydonArray) ValuePgx() (any, error) {
	values := make([]any, len(a.Elements))
	for i, e := range a.Elements {
		v, err := e.ValuePgx()
		if err != nil {
			return nil, fmt.Errorf("getting value for element %d: %w", i, err)
		}
		values[i] = v
	}
	return pq.Array(values), nil
}

func (a *VydonArray) ScanJson(value any) error {
	valueSlice, err := gotypeutil.ParseSlice(value)
	if err != nil {
		return err
	}
	if len(valueSlice) != len(a.Elements) {
		return fmt.Errorf(
			"length mismatch: got %d elements, expected %d",
			len(valueSlice),
			len(a.Elements),
		)
	}
	for i, v := range valueSlice {
		if err := a.Elements[i].ScanJson(v); err != nil {
			return fmt.Errorf("scanning element %d: %w", i, err)
		}
	}
	return nil
}

func (a *VydonArray) ValueJson() (any, error) {
	values := make([]any, len(a.Elements))
	for i, e := range a.Elements {
		v, err := e.ValueJson()
		if err != nil {
			return nil, fmt.Errorf("getting value for element %d: %w", i, err)
		}
		values[i] = v
	}
	return values, nil
}

func (a *VydonArray) ScanMysql(value any) error {
	valueSlice, err := gotypeutil.ParseSlice(value)
	if err != nil {
		return err
	}
	if len(valueSlice) != len(a.Elements) {
		return fmt.Errorf(
			"length mismatch: got %d elements, expected %d",
			len(valueSlice),
			len(a.Elements),
		)
	}
	for i, v := range valueSlice {
		if err := a.Elements[i].ScanMysql(v); err != nil {
			return fmt.Errorf("scanning element %d: %w", i, err)
		}
	}
	return nil
}

func (a *VydonArray) ValueMysql() (any, error) {
	values := make([]any, len(a.Elements))
	for i, e := range a.Elements {
		v, err := e.ValueMysql()
		if err != nil {
			return nil, fmt.Errorf("getting value for element %d: %w", i, err)
		}
		values[i] = v
	}
	return values, nil
}

func (a *VydonArray) ScanMssql(value any) error {
	return a.ScanMysql(value)
}

func (a *VydonArray) ValueMssql() (any, error) {
	return a.ValueMysql()
}
