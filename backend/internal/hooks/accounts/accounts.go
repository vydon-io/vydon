// Package accounts is a no-op stub for the upstream Account Hooks
// feature. All operations return an "unsupported" error so the OSS
// account-hooks management endpoints reply consistently. The feature
// will be reimplemented natively in a future release.
package accounts

import (
	"context"
	"errors"

	mgmtv1alpha1 "github.com/vydon-io/vydon/backend/gen/go/protos/mgmt/v1alpha1"
	"github.com/vydon-io/vydon/backend/internal/userdata"
	slack "github.com/vydon-io/vydon/internal/notifications/slack"
	"github.com/vydon-io/vydon/internal/vydondb"
)

var ErrUnsupported = errors.New("account hooks are not supported in the open-source vydon distribution")

type Interface interface {
	GetAccountHooks(ctx context.Context, req *mgmtv1alpha1.GetAccountHooksRequest) (*mgmtv1alpha1.GetAccountHooksResponse, error)
	GetAccountHook(ctx context.Context, req *mgmtv1alpha1.GetAccountHookRequest) (*mgmtv1alpha1.GetAccountHookResponse, error)
	CreateAccountHook(ctx context.Context, req *mgmtv1alpha1.CreateAccountHookRequest) (*mgmtv1alpha1.CreateAccountHookResponse, error)
	DeleteAccountHook(ctx context.Context, req *mgmtv1alpha1.DeleteAccountHookRequest) (*mgmtv1alpha1.DeleteAccountHookResponse, error)
	IsAccountHookNameAvailable(
		ctx context.Context,
		req *mgmtv1alpha1.IsAccountHookNameAvailableRequest,
	) (*mgmtv1alpha1.IsAccountHookNameAvailableResponse, error)
	UpdateAccountHook(ctx context.Context, req *mgmtv1alpha1.UpdateAccountHookRequest) (*mgmtv1alpha1.UpdateAccountHookResponse, error)
	SetAccountHookEnabled(
		ctx context.Context,
		req *mgmtv1alpha1.SetAccountHookEnabledRequest,
	) (*mgmtv1alpha1.SetAccountHookEnabledResponse, error)
	GetActiveAccountHooksByEvent(
		ctx context.Context,
		req *mgmtv1alpha1.GetActiveAccountHooksByEventRequest,
	) (*mgmtv1alpha1.GetActiveAccountHooksByEventResponse, error)
	GetSlackConnectionUrl(
		ctx context.Context,
		req *mgmtv1alpha1.GetSlackConnectionUrlRequest,
	) (*mgmtv1alpha1.GetSlackConnectionUrlResponse, error)
	HandleSlackOAuthCallback(
		ctx context.Context,
		req *mgmtv1alpha1.HandleSlackOAuthCallbackRequest,
	) (*mgmtv1alpha1.HandleSlackOAuthCallbackResponse, error)
	TestSlackConnection(
		ctx context.Context,
		req *mgmtv1alpha1.TestSlackConnectionRequest,
	) (*mgmtv1alpha1.TestSlackConnectionResponse, error)
	SendSlackMessage(ctx context.Context, req *mgmtv1alpha1.SendSlackMessageRequest) (*mgmtv1alpha1.SendSlackMessageResponse, error)
}

type Option func(*Service)

func WithAppBaseUrl(_ string) Option           { return func(_ *Service) {} }
func WithSlackClient(_ slack.Interface) Option { return func(_ *Service) {} }

type Service struct{}

var _ Interface = (*Service)(nil)

func New(_ *vydondb.VydonDb, _ userdata.Interface, _ ...Option) *Service {
	return &Service{}
}

func (s *Service) GetAccountHooks(
	_ context.Context,
	_ *mgmtv1alpha1.GetAccountHooksRequest,
) (*mgmtv1alpha1.GetAccountHooksResponse, error) {
	return &mgmtv1alpha1.GetAccountHooksResponse{}, nil
}
func (s *Service) GetAccountHook(_ context.Context, _ *mgmtv1alpha1.GetAccountHookRequest) (*mgmtv1alpha1.GetAccountHookResponse, error) {
	return nil, ErrUnsupported
}

func (s *Service) CreateAccountHook(
	_ context.Context,
	_ *mgmtv1alpha1.CreateAccountHookRequest,
) (*mgmtv1alpha1.CreateAccountHookResponse, error) {
	return nil, ErrUnsupported
}

func (s *Service) DeleteAccountHook(
	_ context.Context,
	_ *mgmtv1alpha1.DeleteAccountHookRequest,
) (*mgmtv1alpha1.DeleteAccountHookResponse, error) {
	return &mgmtv1alpha1.DeleteAccountHookResponse{}, nil
}

func (s *Service) IsAccountHookNameAvailable(
	_ context.Context,
	_ *mgmtv1alpha1.IsAccountHookNameAvailableRequest,
) (*mgmtv1alpha1.IsAccountHookNameAvailableResponse, error) {
	return &mgmtv1alpha1.IsAccountHookNameAvailableResponse{IsAvailable: true}, nil
}

func (s *Service) UpdateAccountHook(
	_ context.Context,
	_ *mgmtv1alpha1.UpdateAccountHookRequest,
) (*mgmtv1alpha1.UpdateAccountHookResponse, error) {
	return nil, ErrUnsupported
}

func (s *Service) SetAccountHookEnabled(
	_ context.Context,
	_ *mgmtv1alpha1.SetAccountHookEnabledRequest,
) (*mgmtv1alpha1.SetAccountHookEnabledResponse, error) {
	return nil, ErrUnsupported
}

func (s *Service) GetActiveAccountHooksByEvent(
	_ context.Context,
	_ *mgmtv1alpha1.GetActiveAccountHooksByEventRequest,
) (*mgmtv1alpha1.GetActiveAccountHooksByEventResponse, error) {
	return &mgmtv1alpha1.GetActiveAccountHooksByEventResponse{}, nil
}

func (s *Service) GetSlackConnectionUrl(
	_ context.Context,
	_ *mgmtv1alpha1.GetSlackConnectionUrlRequest,
) (*mgmtv1alpha1.GetSlackConnectionUrlResponse, error) {
	return nil, ErrUnsupported
}

func (s *Service) HandleSlackOAuthCallback(
	_ context.Context,
	_ *mgmtv1alpha1.HandleSlackOAuthCallbackRequest,
) (*mgmtv1alpha1.HandleSlackOAuthCallbackResponse, error) {
	return nil, ErrUnsupported
}

func (s *Service) TestSlackConnection(
	_ context.Context,
	_ *mgmtv1alpha1.TestSlackConnectionRequest,
) (*mgmtv1alpha1.TestSlackConnectionResponse, error) {
	return nil, ErrUnsupported
}

func (s *Service) SendSlackMessage(
	_ context.Context,
	_ *mgmtv1alpha1.SendSlackMessageRequest,
) (*mgmtv1alpha1.SendSlackMessageResponse, error) {
	return nil, ErrUnsupported
}
