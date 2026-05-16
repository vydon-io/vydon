// Package functions stubs the upstream Enterprise transformer functions
// (PII text rewrite via Presidio). The OSS implementation returns the
// input unchanged with a clear error so callers can detect the missing
// feature.
package functions

import (
	"context"
	"errors"
	"log/slog"

	"github.com/stretchr/testify/mock"
	mgmtv1alpha1 "github.com/vydon-io/vydon/backend/gen/go/protos/mgmt/v1alpha1"
	"github.com/vydon-io/vydon/internal/piidetect/presidio"
)

var ErrUnsupported = errors.New("pii text transformer is not supported in the open-source vydon distribution")

type VydonOperatorApi interface {
	Transform(ctx context.Context, config *mgmtv1alpha1.TransformerConfig, value string) (string, error)
}

func TransformPiiText(
	_ context.Context,
	_ presidio.AnalyzeInterface,
	_ presidio.AnonymizeInterface,
	_ VydonOperatorApi,
	_ *mgmtv1alpha1.TransformPiiText,
	value string,
	_ *slog.Logger,
) (string, error) {
	return value, ErrUnsupported
}

type MockVydonOperatorApi struct {
	mock.Mock
}

type mockT interface {
	mock.TestingT
	Cleanup(func())
}

func NewMockVydonOperatorApi(t mockT) *MockVydonOperatorApi {
	m := &MockVydonOperatorApi{}
	m.Test(t)
	t.Cleanup(func() { m.AssertExpectations(t) })
	return m
}

func (m *MockVydonOperatorApi) Transform(ctx context.Context, config *mgmtv1alpha1.TransformerConfig, value string) (string, error) {
	args := m.Called(ctx, config, value)
	return args.String(0), args.Error(1)
}
