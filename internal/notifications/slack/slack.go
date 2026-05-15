// Package slack is a stub for the upstream Slack-connector integration.
// All operations return an "unsupported" error; the feature will be
// reimplemented natively in a future release.
package slack

import (
	"context"
	"errors"
	"net/http"
)

var ErrUnsupported = errors.New("slack notifications are not supported in the open-source vydon distribution")

type Interface interface {
	ExchangeOAuthCode(ctx context.Context, code, redirectURI string) ([]byte, error)
	GetOauthState(ctx context.Context) (*OauthState, error)
	SendMessage(ctx context.Context, accessToken, channel, text string) error
	TestConnection(ctx context.Context, accessToken string) error
}

type OauthState struct {
	State     string
	URL       string
	AccountId string
	UserId    string
	Timestamp int64
}

type Encryptor interface {
	Encrypt(plaintext string) (string, error)
	Decrypt(ciphertext string) (string, error)
}

type Option func(*Client)

func WithAuthClientCreds(_, _ string) Option { return func(_ *Client) {} }
func WithScope(_ string) Option              { return func(_ *Client) {} }
func WithRedirectUrl(_ string) Option        { return func(_ *Client) {} }
func WithHTTPClient(_ *http.Client) Option   { return func(_ *Client) {} }

type Client struct{}

func NewClient(_ Encryptor, _ ...Option) *Client { return &Client{} }

var _ Interface = (*Client)(nil)

func (c *Client) ExchangeOAuthCode(_ context.Context, _, _ string) ([]byte, error) {
	return nil, ErrUnsupported
}
func (c *Client) GetOauthState(_ context.Context) (*OauthState, error) {
	return nil, ErrUnsupported
}
func (c *Client) SendMessage(_ context.Context, _, _, _ string) error { return ErrUnsupported }
func (c *Client) TestConnection(_ context.Context, _ string) error    { return ErrUnsupported }

// MockInterface satisfies the slack.Interface for tests.
type MockInterface struct{}

func NewMockInterface(_ any) *MockInterface { return &MockInterface{} }

var _ Interface = (*MockInterface)(nil)

func (m *MockInterface) ExchangeOAuthCode(_ context.Context, _, _ string) ([]byte, error) {
	return nil, ErrUnsupported
}
func (m *MockInterface) GetOauthState(_ context.Context) (*OauthState, error) {
	return nil, ErrUnsupported
}
func (m *MockInterface) SendMessage(_ context.Context, _, _, _ string) error { return nil }
func (m *MockInterface) TestConnection(_ context.Context, _ string) error    { return nil }
