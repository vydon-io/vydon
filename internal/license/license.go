// Package license provides a permissive license shim for the OSS
// distribution. The original Neosync fork supported a paid Enterprise
// Edition gated by a signed license file. Vydon ships under MIT with
// no proprietary gates, so every license check returns true.
//
// The API is kept intentionally close to the upstream shape (IsValid,
// ExpiresAt, EEInterface) to minimize churn in callers, but every
// implementation is a no-op.
package license

import (
	"time"
)

type EEInterface interface {
	IsValid() bool
	ExpiresAt() time.Time
}

type ValidLicense struct{}

func NewValidLicense() *ValidLicense                 { return &ValidLicense{} }
func (v *ValidLicense) IsValid() bool                { return true }
func (v *ValidLicense) ExpiresAt() time.Time         { return time.Time{} }

type EELicense = ValidLicense

func NewFromEnv() (*EELicense, error) { return NewValidLicense(), nil }

type CascadeLicense struct{}

func NewCascadeLicense(_ ...EEInterface) *CascadeLicense { return &CascadeLicense{} }
func (c *CascadeLicense) IsValid() bool                  { return true }
func (c *CascadeLicense) ExpiresAt() time.Time           { return time.Time{} }

// Signature is the deprecated stand-in for the upstream signed-license
// payload type. Kept so call sites that fish a Signature out of an
// EELicense compile without churn; never populated in Vydon.
type Signature = string
