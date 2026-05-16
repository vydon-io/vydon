package gob

import (
	"encoding/gob"
	"time"

	dynamotypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	vydontypes "github.com/vydon-io/vydon/internal/vydon-types"
)

// need to register all the types that are used in the connection data service
// because we use interfaces for the types
func RegisterGobTypes() {
	gob.Register(map[string]any{})
	gob.Register([]any{})
	gob.Register(time.Time{})
	gob.Register([][]uint8{})
	gob.RegisterName("vydontypes.VydonDateTime", &vydontypes.VydonDateTime{})
	gob.RegisterName("vydontypes.Bits", &vydontypes.Bits{})
	gob.RegisterName("vydontypes.Binary", &vydontypes.Binary{})
	gob.RegisterName("vydontypes.VydonArray", &vydontypes.VydonArray{})
	gob.RegisterName("vydontypes.Interval", &vydontypes.Interval{})
	gob.RegisterName("dynamodb.AttributeValueMemberB", &dynamotypes.AttributeValueMemberB{})
	gob.RegisterName("dynamodb.AttributeValueMemberBOOL", &dynamotypes.AttributeValueMemberBOOL{})
	gob.RegisterName("dynamodb.AttributeValueMemberBS", &dynamotypes.AttributeValueMemberBS{})
	gob.RegisterName("dynamodb.AttributeValueMemberL", &dynamotypes.AttributeValueMemberL{})
	gob.RegisterName("dynamodb.AttributeValueMemberM", &dynamotypes.AttributeValueMemberM{})
	gob.RegisterName("dynamodb.AttributeValueMemberN", &dynamotypes.AttributeValueMemberN{})
	gob.RegisterName("dynamodb.AttributeValueMemberNS", &dynamotypes.AttributeValueMemberNS{})
	gob.RegisterName("dynamodb.AttributeValueMemberNULL", &dynamotypes.AttributeValueMemberNULL{})
	gob.RegisterName("dynamodb.AttributeValueMemberS", &dynamotypes.AttributeValueMemberS{})
	gob.RegisterName("dynamodb.AttributeValueMemberSS", &dynamotypes.AttributeValueMemberSS{})
}
