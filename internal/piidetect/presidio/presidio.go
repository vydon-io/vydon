// Package presidio is a stub for the Microsoft Presidio PII-detection
// integration that the upstream fork wrapped as an Enterprise feature.
// Every method on every interface here is a no-op or returns an "unsupported
// in OSS" error so the rest of the codebase compiles. The feature itself
// will be reintroduced as a native module calling Presidio directly.
package presidio

import (
	"context"
	"errors"
	"net/http"
)

var ErrUnsupported = errors.New("pii detection is not supported in the open-source vydon distribution")

type ClientOption func(*ClientWithResponses) error

func WithHTTPClient(_ *http.Client) ClientOption {
	return func(_ *ClientWithResponses) error { return nil }
}

type ClientWithResponses struct{}

func NewClientWithResponses(_ string, _ ...ClientOption) (*ClientWithResponses, error) {
	return &ClientWithResponses{}, nil
}

type AnalyzeInterface interface {
	PostAnalyzeWithResponse(ctx context.Context, body interface{}) (*PostAnalyzeResponse, error)
}

type AnonymizeInterface interface {
	PostAnonymizeWithResponse(ctx context.Context, body interface{}) (*PostAnonymizeResponse, error)
}

type EntityInterface interface {
	GetSupportedentitiesWithResponse(ctx context.Context, params *GetSupportedentitiesParams) (*GetSupportedentitiesResponse, error)
}

func (c *ClientWithResponses) PostAnalyzeWithResponse(_ context.Context, _ interface{}) (*PostAnalyzeResponse, error) {
	return nil, ErrUnsupported
}

func (c *ClientWithResponses) PostAnonymizeWithResponse(_ context.Context, _ interface{}) (*PostAnonymizeResponse, error) {
	return nil, ErrUnsupported
}

func (c *ClientWithResponses) GetSupportedentitiesWithResponse(_ context.Context, _ *GetSupportedentitiesParams) (*GetSupportedentitiesResponse, error) {
	return nil, ErrUnsupported
}

type GetSupportedentitiesParams struct {
	Language *string
}

type GetSupportedentitiesResponse struct {
	JSON200    *[]string
	Body       []byte
	HTTPResponse *http.Response
}

func (r *GetSupportedentitiesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return ""
}

func (r *GetSupportedentitiesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type RecognizerResultWithAnaysisExplanation struct {
	EntityType *string
	Start      *int
	End        *int
	Score      *float32
}

type PostAnalyzeResponse struct {
	JSON200 *[]RecognizerResultWithAnaysisExplanation
}

type OperatorResult struct {
	Start    *int
	End      *int
	Text     *string
	Operator *string
}

type AnonymizeResponse struct {
	Text  *string
	Items *[]OperatorResult
}

type PostAnonymizeResponse struct {
	JSON200 *AnonymizeResponse
}
