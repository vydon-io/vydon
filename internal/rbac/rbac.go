package rbac

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgtype"
	mgmtv1alpha1 "github.com/vydon-io/vydon/backend/gen/go/protos/mgmt/v1alpha1"
	"github.com/vydon-io/vydon/internal/vydondb"
)

const Wildcard = "*"

var (
	JobWildcard        = NewEntity("jobs", Wildcard)
	ConnectionWildcard = NewEntity("connections", Wildcard)
)

type Entity struct {
	prefix string
	value  string
}

type EntityString interface {
	String() string
}

func NewEntity(prefix, value string) *Entity {
	return &Entity{prefix: prefix, value: value}
}

func (e *Entity) String() string {
	return fmt.Sprintf("%s/%s", e.prefix, e.value)
}

func NewAccountIdEntity(value string) *Entity    { return NewEntity("accounts", value) }
func NewJobIdEntity(value string) *Entity        { return NewEntity("jobs", value) }
func NewUserIdEntity(value string) *Entity       { return NewEntity("users", value) }
func NewConnectionIdEntity(value string) *Entity { return NewEntity("connections", value) }
func NewPgUserIdEntity(value pgtype.UUID) *Entity {
	return NewUserIdEntity(vydondb.UUIDString(value))
}

type AccountAction string

const (
	AccountAction_Create AccountAction = "create"
	AccountAction_Delete AccountAction = "delete"
	AccountAction_View   AccountAction = "view"
	AccountAction_Edit   AccountAction = "edit"
)

func (a AccountAction) String() string { return string(a) }

type ConnectionAction string

const (
	ConnectionAction_Create        ConnectionAction = "create"
	ConnectionAction_Delete        ConnectionAction = "delete"
	ConnectionAction_View          ConnectionAction = "view"
	ConnectionAction_ViewSensitive ConnectionAction = "view_sensitive"
	ConnectionAction_Edit          ConnectionAction = "edit"
)

func (c ConnectionAction) String() string { return string(c) }

type JobAction string

const (
	JobAction_Create  JobAction = "create"
	JobAction_Delete  JobAction = "delete"
	JobAction_Execute JobAction = "execute"
	JobAction_View    JobAction = "view"
	JobAction_Edit    JobAction = "edit"
)

func (a JobAction) String() string { return string(a) }

type Role string

const (
	Role_AccountAdmin Role = "account_admin"
	Role_JobDeveloper Role = "job_developer"
	Role_JobExecutor  Role = "job_executor"
	Role_JobViewer    Role = "job_viewer"
)

func (r Role) String() string { return string(r) }

func (r Role) ToDto() mgmtv1alpha1.AccountRole {
	switch r {
	case Role_AccountAdmin:
		return mgmtv1alpha1.AccountRole_ACCOUNT_ROLE_ADMIN
	case Role_JobDeveloper:
		return mgmtv1alpha1.AccountRole_ACCOUNT_ROLE_JOB_DEVELOPER
	case Role_JobExecutor:
		return mgmtv1alpha1.AccountRole_ACCOUNT_ROLE_JOB_EXECUTOR
	case Role_JobViewer:
		return mgmtv1alpha1.AccountRole_ACCOUNT_ROLE_JOB_VIEWER
	default:
		return mgmtv1alpha1.AccountRole_ACCOUNT_ROLE_UNSPECIFIED
	}
}

type EntityEnforcer interface {
	Job(ctx context.Context, user, account, job EntityString, action JobAction) (bool, error)
	EnforceJob(ctx context.Context, user, account, job EntityString, action JobAction) error
	Connection(ctx context.Context, user, account, connection EntityString, action ConnectionAction) (bool, error)
	EnforceConnection(ctx context.Context, user, account, connection EntityString, action ConnectionAction) error
	Account(ctx context.Context, user, account EntityString, action AccountAction) (bool, error)
	EnforceAccount(ctx context.Context, user, account EntityString, action AccountAction) error
}

type RoleAdmin interface {
	SetAccountRole(ctx context.Context, user, account EntityString, role mgmtv1alpha1.AccountRole) error
	RemoveAccountRole(ctx context.Context, user, account EntityString, role mgmtv1alpha1.AccountRole) error
	RemoveAccountUser(ctx context.Context, user, account EntityString) error
	SetupNewAccount(ctx context.Context, accountId string, logger *slog.Logger) error
	GetUserRoles(ctx context.Context, users []EntityString, account EntityString, logger *slog.Logger) map[string]Role
}

type Interface interface {
	EntityEnforcer
	RoleAdmin
}

type AllowAllClient struct{}

var _ Interface = (*AllowAllClient)(nil)

func NewAllowAllClient() *AllowAllClient { return &AllowAllClient{} }

func (a *AllowAllClient) Job(_ context.Context, _, _, _ EntityString, _ JobAction) (bool, error) {
	return true, nil
}

func (a *AllowAllClient) Connection(_ context.Context, _, _, _ EntityString, _ ConnectionAction) (bool, error) {
	return true, nil
}

func (a *AllowAllClient) Account(_ context.Context, _, _ EntityString, _ AccountAction) (bool, error) {
	return true, nil
}

func (a *AllowAllClient) EnforceJob(_ context.Context, _, _, _ EntityString, _ JobAction) error {
	return nil
}

func (a *AllowAllClient) EnforceConnection(_ context.Context, _, _, _ EntityString, _ ConnectionAction) error {
	return nil
}

func (a *AllowAllClient) EnforceAccount(_ context.Context, _, _ EntityString, _ AccountAction) error {
	return nil
}

func (a *AllowAllClient) SetAccountRole(_ context.Context, _, _ EntityString, _ mgmtv1alpha1.AccountRole) error {
	return nil
}

func (a *AllowAllClient) RemoveAccountRole(_ context.Context, _, _ EntityString, _ mgmtv1alpha1.AccountRole) error {
	return nil
}

func (a *AllowAllClient) RemoveAccountUser(_ context.Context, _, _ EntityString) error {
	return nil
}

func (a *AllowAllClient) SetupNewAccount(_ context.Context, _ string, _ *slog.Logger) error {
	return nil
}

func (a *AllowAllClient) GetUserRoles(_ context.Context, _ []EntityString, _ EntityString, _ *slog.Logger) map[string]Role {
	return map[string]Role{}
}
