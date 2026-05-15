package presidio

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockAnalyzeInterface struct {
	mock.Mock
}

func NewMockAnalyzeInterface(t mockT) *MockAnalyzeInterface {
	m := &MockAnalyzeInterface{}
	m.Mock.Test(t)
	t.Cleanup(func() { m.AssertExpectations(t) })
	return m
}

func (m *MockAnalyzeInterface) PostAnalyzeWithResponse(ctx context.Context, body interface{}) (*PostAnalyzeResponse, error) {
	args := m.Called(ctx, body)
	var r0 *PostAnalyzeResponse
	if rf, ok := args.Get(0).(*PostAnalyzeResponse); ok {
		r0 = rf
	}
	return r0, args.Error(1)
}

type MockAnonymizeInterface struct {
	mock.Mock
}

func NewMockAnonymizeInterface(t mockT) *MockAnonymizeInterface {
	m := &MockAnonymizeInterface{}
	m.Mock.Test(t)
	t.Cleanup(func() { m.AssertExpectations(t) })
	return m
}

func (m *MockAnonymizeInterface) PostAnonymizeWithResponse(ctx context.Context, body interface{}) (*PostAnonymizeResponse, error) {
	args := m.Called(ctx, body)
	var r0 *PostAnonymizeResponse
	if rf, ok := args.Get(0).(*PostAnonymizeResponse); ok {
		r0 = rf
	}
	return r0, args.Error(1)
}

type MockEntityInterface struct {
	mock.Mock
}

func NewMockEntityInterface(t mockT) *MockEntityInterface {
	m := &MockEntityInterface{}
	m.Mock.Test(t)
	t.Cleanup(func() { m.AssertExpectations(t) })
	return m
}

func (m *MockEntityInterface) GetSupportedentitiesWithResponse(ctx context.Context, params *GetSupportedentitiesParams) (*GetSupportedentitiesResponse, error) {
	args := m.Called(ctx, params)
	var r0 *GetSupportedentitiesResponse
	if rf, ok := args.Get(0).(*GetSupportedentitiesResponse); ok {
		r0 = rf
	}
	return r0, args.Error(1)
}

// mockT is the minimum testify mock interface used to wire cleanup.
type mockT interface {
	mock.TestingT
	Cleanup(func())
}
