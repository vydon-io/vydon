package mongodb

import (
	"math/big"
	"testing"

	vydon_types "github.com/vydon-io/vydon/internal/types"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_UnmarshalPrimitives(t *testing.T) {
	input := map[string]any{
		"string": "test",
		"int":    42,
		"bool":   true,
	}
	expectedOutput := map[string]any{
		"string": "test",
		"int":    42,
		"bool":   true,
	}
	expectedKTM := map[string]vydon_types.KeyType{}

	mapper := NewMongoBuilder()

	t.Run("Basic types", func(t *testing.T) {
		output, ktm, err := mapper.MapRecordWithKeyType(input)
		require.NoError(t, err)
		require.Equal(t, expectedOutput, output)
		require.Equal(t, expectedKTM, ktm)
	})

	dec128 := primitive.NewDecimal128(3, 14159)
	objectId := primitive.NewObjectID()
	input = map[string]any{
		"decimal":   dec128,
		"binary":    primitive.Binary{Data: []byte("test")},
		"objectID":  objectId,
		"timestamp": primitive.Timestamp{T: 1, I: 1},
	}
	expectedOutput = map[string]any{
		"decimal":   getBigFloat(dec128.String()),
		"binary":    primitive.Binary{Data: []byte("test")},
		"objectID":  objectId,
		"timestamp": primitive.Timestamp{T: 1, I: 1},
	}
	expectedKTM = map[string]vydon_types.KeyType{
		"decimal":   vydon_types.Decimal128,
		"binary":    vydon_types.Binary,
		"objectID":  vydon_types.ObjectID,
		"timestamp": vydon_types.Timestamp,
	}

	t.Run("BSON types", func(t *testing.T) {
		output, ktm, err := mapper.MapRecordWithKeyType(input)
		require.NoError(t, err)
		require.Equal(t, expectedOutput, output)
		require.Equal(t, expectedKTM, ktm)
	})
}

func getBigFloat(v string) *big.Float {
	f, _, _ := big.ParseFloat(v, 10, 128, big.ToNearestEven)
	return f
}

func Test_ParsePrimitives(t *testing.T) {
	objectId := primitive.NewObjectID()
	dec128 := primitive.NewDecimal128(3, 14159)
	testCases := []struct {
		name        string
		key         string
		value       any
		expectedKTM map[string]vydon_types.KeyType
		expected    any
	}{
		{
			name:        "Decimal128",
			key:         "decimal",
			value:       dec128,
			expectedKTM: map[string]vydon_types.KeyType{"decimal": vydon_types.Decimal128},
			expected:    getBigFloat(dec128.String()),
		},
		{
			name:        "Binary",
			key:         "binary",
			value:       primitive.Binary{Data: []byte("test")},
			expectedKTM: map[string]vydon_types.KeyType{"binary": vydon_types.Binary},
			expected:    primitive.Binary{Data: []byte("test")},
		},
		{
			name:        "ObjectID",
			key:         "objectID",
			value:       objectId,
			expectedKTM: map[string]vydon_types.KeyType{"objectID": vydon_types.ObjectID},
			expected:    objectId,
		},
		{
			name:        "Timestamp",
			key:         "timestamp",
			value:       primitive.Timestamp{T: 1, I: 1},
			expectedKTM: map[string]vydon_types.KeyType{"timestamp": vydon_types.Timestamp},
			expected:    primitive.Timestamp{T: 1, I: 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ktm := make(map[string]vydon_types.KeyType)
			result, err := parsePrimitives(tc.key, tc.value, ktm)
			require.NoError(t, err)
			require.Equal(t, tc.expectedKTM, ktm)
			require.Equal(t, tc.expected, result)
		})
	}
}
