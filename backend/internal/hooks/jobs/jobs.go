// Package jobs is a no-op stub for the upstream Job Hooks feature.
// All operations return empty results so the OSS job management
// endpoints continue to function. The hook execution itself will be
// reimplemented natively in a future release.
package jobs

import (
	"context"
	"errors"

	mgmtv1alpha1 "github.com/vydon-io/vydon/backend/gen/go/protos/mgmt/v1alpha1"
	"github.com/vydon-io/vydon/backend/internal/userdata"
	"github.com/vydon-io/vydon/internal/vydondb"
)

var ErrUnsupported = errors.New("job hooks are not supported in the open-source vydon distribution")

type Interface interface {
	GetJobHooks(ctx context.Context, req *mgmtv1alpha1.GetJobHooksRequest) (*mgmtv1alpha1.GetJobHooksResponse, error)
	GetJobHook(ctx context.Context, req *mgmtv1alpha1.GetJobHookRequest) (*mgmtv1alpha1.GetJobHookResponse, error)
	CreateJobHook(ctx context.Context, req *mgmtv1alpha1.CreateJobHookRequest) (*mgmtv1alpha1.CreateJobHookResponse, error)
	DeleteJobHook(ctx context.Context, req *mgmtv1alpha1.DeleteJobHookRequest) (*mgmtv1alpha1.DeleteJobHookResponse, error)
	IsJobHookNameAvailable(ctx context.Context, req *mgmtv1alpha1.IsJobHookNameAvailableRequest) (*mgmtv1alpha1.IsJobHookNameAvailableResponse, error)
	UpdateJobHook(ctx context.Context, req *mgmtv1alpha1.UpdateJobHookRequest) (*mgmtv1alpha1.UpdateJobHookResponse, error)
	SetJobHookEnabled(ctx context.Context, req *mgmtv1alpha1.SetJobHookEnabledRequest) (*mgmtv1alpha1.SetJobHookEnabledResponse, error)
	GetActiveJobHooksByTiming(ctx context.Context, req *mgmtv1alpha1.GetActiveJobHooksByTimingRequest) (*mgmtv1alpha1.GetActiveJobHooksByTimingResponse, error)
}

type Option func(*Service)

func WithEnabled() Option { return func(_ *Service) {} }

type Service struct{}

var _ Interface = (*Service)(nil)

func New(_ *vydondb.VydonDb, _ userdata.Interface, _ ...Option) *Service {
	return &Service{}
}

func (s *Service) GetJobHooks(_ context.Context, _ *mgmtv1alpha1.GetJobHooksRequest) (*mgmtv1alpha1.GetJobHooksResponse, error) {
	return &mgmtv1alpha1.GetJobHooksResponse{}, nil
}
func (s *Service) GetJobHook(_ context.Context, _ *mgmtv1alpha1.GetJobHookRequest) (*mgmtv1alpha1.GetJobHookResponse, error) {
	return nil, ErrUnsupported
}
func (s *Service) CreateJobHook(_ context.Context, _ *mgmtv1alpha1.CreateJobHookRequest) (*mgmtv1alpha1.CreateJobHookResponse, error) {
	return nil, ErrUnsupported
}
func (s *Service) DeleteJobHook(_ context.Context, _ *mgmtv1alpha1.DeleteJobHookRequest) (*mgmtv1alpha1.DeleteJobHookResponse, error) {
	return &mgmtv1alpha1.DeleteJobHookResponse{}, nil
}
func (s *Service) IsJobHookNameAvailable(_ context.Context, _ *mgmtv1alpha1.IsJobHookNameAvailableRequest) (*mgmtv1alpha1.IsJobHookNameAvailableResponse, error) {
	return &mgmtv1alpha1.IsJobHookNameAvailableResponse{IsAvailable: true}, nil
}
func (s *Service) UpdateJobHook(_ context.Context, _ *mgmtv1alpha1.UpdateJobHookRequest) (*mgmtv1alpha1.UpdateJobHookResponse, error) {
	return nil, ErrUnsupported
}
func (s *Service) SetJobHookEnabled(_ context.Context, _ *mgmtv1alpha1.SetJobHookEnabledRequest) (*mgmtv1alpha1.SetJobHookEnabledResponse, error) {
	return nil, ErrUnsupported
}
func (s *Service) GetActiveJobHooksByTiming(_ context.Context, _ *mgmtv1alpha1.GetActiveJobHooksByTimingRequest) (*mgmtv1alpha1.GetActiveJobHooksByTimingResponse, error) {
	return &mgmtv1alpha1.GetActiveJobHooksByTimingResponse{}, nil
}
