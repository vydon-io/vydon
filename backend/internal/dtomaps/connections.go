package dtomaps

import (
	db_queries "github.com/vydon-io/vydon/backend/gen/go/db"
	mgmtv1alpha1 "github.com/vydon-io/vydon/backend/gen/go/protos/mgmt/v1alpha1"
	"github.com/vydon-io/vydon/internal/vydondb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToConnectionDto(
	input *db_queries.VydonApiConnection,
	canViewSensitive bool,
) (*mgmtv1alpha1.Connection, error) {
	ccDto, err := input.ConnectionConfig.ToDto(canViewSensitive)
	if err != nil {
		return nil, err
	}
	return &mgmtv1alpha1.Connection{
		Id:               vydondb.UUIDString(input.ID),
		Name:             input.Name,
		ConnectionConfig: ccDto,
		CreatedAt:        timestamppb.New(input.CreatedAt.Time),
		UpdatedAt:        timestamppb.New(input.UpdatedAt.Time),
		CreatedByUserId:  vydondb.UUIDString(input.CreatedByID),
		UpdatedByUserId:  vydondb.UUIDString(input.UpdatedByID),
		AccountId:        vydondb.UUIDString(input.AccountID),
	}, nil
}
