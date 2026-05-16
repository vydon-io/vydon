// Package mssqlmanager stubs the upstream EE MSSQL manager. The OSS
// MSSQL implementation embeds Manager for shared state; the embedded
// struct is empty in OSS.
package mssqlmanager

import (
	"log/slog"
)

const (
	SchemasLabel        = "schemas"
	TableIndexLabel     = "table-index"
	ViewsFunctionsLabel = "views-functions"
)

type Manager struct{}

// NewManager accepts the original argument list and returns an empty
// Manager. Callers compose it via struct embedding for backwards
// compatibility.
func NewManager(_, _, _, _ any) *Manager {
	_ = slog.Default()
	return &Manager{}
}
