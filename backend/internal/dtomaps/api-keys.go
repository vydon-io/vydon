package dtomaps

import (
	db_queries "github.com/vydon-io/vydon/backend/gen/go/db"
	mgmtv1alpha1 "github.com/vydon-io/vydon/backend/gen/go/protos/mgmt/v1alpha1"
	"github.com/vydon-io/vydon/internal/vydondb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToAccountApiKeyDto(
	input *db_queries.NeosyncApiAccountApiKey,
	cleartextKeyValue *string,
) *mgmtv1alpha1.AccountApiKey {
	return &mgmtv1alpha1.AccountApiKey{
		Id:          vydondb.UUIDString(input.ID),
		Name:        input.KeyName,
		AccountId:   vydondb.UUIDString(input.AccountID),
		CreatedById: vydondb.UUIDString(input.CreatedByID),
		CreatedAt:   timestamppb.New(input.CreatedAt.Time),
		UpdatedById: vydondb.UUIDString(input.UpdatedByID),
		UpdatedAt:   timestamppb.New(input.UpdatedAt.Time),
		KeyValue:    cleartextKeyValue,
		UserId:      vydondb.UUIDString(input.UserID),
		ExpiresAt:   timestamppb.New(input.ExpiresAt.Time),
	}
}
