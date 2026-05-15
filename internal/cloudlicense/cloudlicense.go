// Package cloudlicense is a no-op replacement for the upstream
// cloud-license gate. Vydon has no proprietary "cloud install" flavor;
// IsValid always returns false so callers do not enable cloud-only
// code paths.
package cloudlicense

import "time"

type Interface interface {
	IsValid() bool
	ExpiresAt() time.Time
}

type CloudLicense struct{}

func NewFromEnv() (*CloudLicense, error)        { return &CloudLicense{}, nil }
func (c *CloudLicense) IsValid() bool           { return false }
func (c *CloudLicense) ExpiresAt() time.Time    { return time.Time{} }
