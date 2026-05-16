package vydon_benthos_connectiondata

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	vydon_types "github.com/vydon-io/vydon/internal/types"
)

func Test_UnmarshalDynamoDBItem(t *testing.T) {
	tests := []struct {
		name            string
		input           map[string]any
		wantStandardMap map[string]any
		wantKeyTypeMap  map[string]vydon_types.KeyType
	}{
		{
			name: "Basic test",
			input: map[string]any{
				"StringKey": map[string]any{"S": "value"},
				"NumberKey": map[string]any{"N": "123"},
				"BoolKey":   map[string]any{"BOOL": true},
			},
			wantStandardMap: map[string]any{
				"StringKey": "value",
				"NumberKey": int64(123),
				"BoolKey":   true,
			},
			wantKeyTypeMap: map[string]vydon_types.KeyType{},
		},
		{
			name: "Complex test with sets",
			input: map[string]any{
				"StringSetKey": map[string]any{"SS": []any{"a", "b", "c"}},
				"NumberSetKey": map[string]any{"NS": []any{"1", "2", "3"}},
			},
			wantStandardMap: map[string]any{
				"StringSetKey": []string{"a", "b", "c"},
				"NumberSetKey": []any{int64(1), int64(2), int64(3)},
			},
			wantKeyTypeMap: map[string]vydon_types.KeyType{
				"StringSetKey": vydon_types.StringSet,
				"NumberSetKey": vydon_types.NumberSet,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStandardMap, gotKeyTypeMap := unmarshalDynamoDBItem(tt.input)
			require.Equal(t, tt.wantStandardMap, gotStandardMap)
			require.Equal(t, tt.wantKeyTypeMap, gotKeyTypeMap)
		})
	}
}

func Test_ParseDynamoDBAttributeValue(t *testing.T) {
	tests := []struct {
		name        string
		key         string
		value       any
		keyTypeMap  map[string]vydon_types.KeyType
		want        any
		wantKeyType vydon_types.KeyType
	}{
		{"String", "StrKey", map[string]any{"S": "value"}, map[string]vydon_types.KeyType{}, "value", 0},
		{"Number", "NumKey", map[string]any{"N": "123"}, map[string]vydon_types.KeyType{}, int64(123), 0},
		{"Boolean", "BoolKey", map[string]any{"BOOL": true}, map[string]vydon_types.KeyType{}, true, 0},
		{"Null", "NullKey", map[string]any{"NULL": true}, map[string]vydon_types.KeyType{}, nil, 0},
		{
			"StringSet",
			"SSKey",
			map[string]any{"SS": []any{"a", "b"}},
			map[string]vydon_types.KeyType{},
			[]string{"a", "b"},
			vydon_types.StringSet,
		},
		{
			"NumberSet",
			"NSKey",
			map[string]any{"NS": []any{"1", "2"}},
			map[string]vydon_types.KeyType{},
			[]any{int64(1), int64(2)},
			vydon_types.NumberSet,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseDynamoDBAttributeValue(tt.key, tt.value, tt.keyTypeMap)
			require.Equalf(t, tt.want, got, fmt.Sprintf("ParseDynamoDBAttributeValue() = %v, want %v", got, tt.want))
			if gotKeyType, ok := tt.keyTypeMap[tt.key]; ok {
				require.Equalf(
					t,
					tt.wantKeyType,
					gotKeyType,
					fmt.Sprintf("ParseDynamoDBAttributeValue() key type = %v, want %v", gotKeyType, tt.wantKeyType),
				)
			}
		})
	}
}
