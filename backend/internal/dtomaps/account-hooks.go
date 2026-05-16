package dtomaps

import (
	"encoding/json"

	db_queries "github.com/vydon-io/vydon/backend/gen/go/db"
	mgmtv1alpha1 "github.com/vydon-io/vydon/backend/gen/go/protos/mgmt/v1alpha1"
	"github.com/vydon-io/vydon/internal/vydondb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToAccountHookDto(
	input *db_queries.VydonApiAccountHook,
) (*mgmtv1alpha1.AccountHook, error) {
	if input == nil {
		input = &db_queries.VydonApiAccountHook{}
	}

	config := &mgmtv1alpha1.AccountHookConfig{}
	err := json.Unmarshal(input.Config, config)
	if err != nil {
		return nil, err
	}

	output := &mgmtv1alpha1.AccountHook{
		Id:              vydondb.UUIDString(input.ID),
		Name:            input.Name,
		Description:     input.Description,
		AccountId:       vydondb.UUIDString(input.AccountID),
		CreatedByUserId: vydondb.UUIDString(input.CreatedByUserID),
		CreatedAt:       timestamppb.New(input.CreatedAt.Time),
		UpdatedByUserId: vydondb.UUIDString(input.UpdatedByUserID),
		UpdatedAt:       timestamppb.New(input.UpdatedAt.Time),
		Enabled:         input.Enabled,
		Config:          config,
		Events:          toAccountHookEvents(input.Events),
	}

	return output, nil
}

func ToAccountHooksDto(
	input []db_queries.VydonApiAccountHook,
) ([]*mgmtv1alpha1.AccountHook, error) {
	dtos := make([]*mgmtv1alpha1.AccountHook, len(input))
	for idx := range input {
		hook := input[idx]
		dto, err := ToAccountHookDto(&hook)
		if err != nil {
			return nil, err
		}
		dtos[idx] = dto
	}
	return dtos, nil
}

func toAccountHookEvents(events []int32) []mgmtv1alpha1.AccountHookEvent {
	output := make([]mgmtv1alpha1.AccountHookEvent, 0, len(events))
	for _, event := range events {
		if _, ok := mgmtv1alpha1.AccountHookEvent_name[event]; ok {
			output = append(output, mgmtv1alpha1.AccountHookEvent(event))
		}
	}
	return output
}
