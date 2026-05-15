// Package functions stubs the upstream Enterprise transformer functions
// (PII text rewrite via Presidio). The OSS implementation returns the
// input unchanged with a clear error so callers can detect the missing
// feature.
package functions

import (
	"context"
	"errors"
	"log/slog"

	mgmtv1alpha1 "github.com/vydon-io/vydon/backend/gen/go/protos/mgmt/v1alpha1"
	"github.com/vydon-io/vydon/internal/piidetect/presidio"
	"github.com/stretchr/testify/mock"
)

var ErrUnsupported = errors.New("pii text transformer is not supported in the open-source vydon distribution")

type NeosyncOperatorApi interface {
	Transform(ctx context.Context, config *mgmtv1alpha1.TransformerConfig, value string) (string, error)
}

func TransformPiiText(
	_ context.Context,
	_ presidio.AnalyzeInterface,
	_ presidio.AnonymizeInterface,
	_ NeosyncOperatorApi,
	_ *mgmtv1alpha1.TransformPiiText,
	value string,
	_ *slog.Logger,
) (string, error) {
	return value, ErrUnsupported
}

type MockNeosyncOperatorApi struct {
	mock.Mock
}

type mockT interface {
	mock.TestingT
	Cleanup(func())
}

func NewMockNeosyncOperatorApi(t mockT) *MockNeosyncOperatorApi {
	m := &MockNeosyncOperatorApi{}
	m.Mock.Test(t)
	t.Cleanup(func() { m.AssertExpectations(t) })
	return m
}

func (m *MockNeosyncOperatorApi) Transform(ctx context.Context, config *mgmtv1alpha1.TransformerConfig, value string) (string, error) {
	args := m.Called(ctx, config, value)
	return args.String(0), args.Error(1)
}
