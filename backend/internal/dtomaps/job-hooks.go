package dtomaps

import (
	db_queries "github.com/vydon-io/vydon/backend/gen/go/db"
	mgmtv1alpha1 "github.com/vydon-io/vydon/backend/gen/go/protos/mgmt/v1alpha1"
	"github.com/vydon-io/vydon/internal/vydondb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToJobHookDto(
	input *db_queries.NeosyncApiJobHook,
) (*mgmtv1alpha1.JobHook, error) {
	if input == nil {
		input = &db_queries.NeosyncApiJobHook{}
	}
	priority := uint32(0)
	if input.Priority > 0 {
		priority = uint32(input.Priority)
	}

	config := &mgmtv1alpha1.JobHookConfig{}
	err := config.UnmarshalJSON(input.Config)
	if err != nil {
		return nil, err
	}

	output := &mgmtv1alpha1.JobHook{
		Id:              vydondb.UUIDString(input.ID),
		Name:            input.Name,
		Description:     input.Description,
		JobId:           vydondb.UUIDString(input.JobID),
		CreatedByUserId: vydondb.UUIDString(input.CreatedByUserID),
		CreatedAt:       timestamppb.New(input.CreatedAt.Time),
		UpdatedByUserId: vydondb.UUIDString(input.UpdatedByUserID),
		UpdatedAt:       timestamppb.New(input.UpdatedAt.Time),
		Enabled:         input.Enabled,
		Priority:        priority,
		Config:          config,
	}

	return output, nil
}

func ToJobHooksDto(
	input []db_queries.NeosyncApiJobHook,
) ([]*mgmtv1alpha1.JobHook, error) {
	dtos := make([]*mgmtv1alpha1.JobHook, len(input))
	for idx := range input {
		hook := input[idx]
		dto, err := ToJobHookDto(&hook)
		if err != nil {
			return nil, err
		}
		dtos[idx] = dto
	}
	return dtos, nil
}
