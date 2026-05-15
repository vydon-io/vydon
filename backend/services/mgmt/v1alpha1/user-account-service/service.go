package v1alpha1_useraccountservice

import (
	auth_client "github.com/vydon-io/vydon/backend/internal/auth/client"
	"github.com/vydon-io/vydon/backend/internal/userdata"
	"github.com/vydon-io/vydon/internal/authmgmt"
	"github.com/vydon-io/vydon/internal/billing"
	"github.com/vydon-io/vydon/internal/license"
	"github.com/vydon-io/vydon/internal/rbac"
	"github.com/vydon-io/vydon/internal/temporal/clientmanager"
	"github.com/vydon-io/vydon/internal/vydondb"
)

type Service struct {
	cfg                    *Config
	db                     *vydondb.VydonDb
	temporalConfigProvider clientmanager.ConfigProvider
	authclient             auth_client.Interface
	authadminclient        authmgmt.Interface
	billingclient          billing.Interface
	rbacClient             rbac.Interface
	licenseclient          license.EEInterface
}

type Config struct {
	IsAuthEnabled            bool
	IsVydonCloud             bool
	DefaultMaxAllowedRecords *int64
}

func New(
	cfg *Config,
	db *vydondb.VydonDb,
	temporalConfigProvider clientmanager.ConfigProvider,
	authclient auth_client.Interface,
	authadminclient authmgmt.Interface,
	billingclient billing.Interface,
	rbacClient rbac.Interface,
	licenseclient license.EEInterface,
) *Service {
	return &Service{
		cfg:                    cfg,
		db:                     db,
		temporalConfigProvider: temporalConfigProvider,
		authclient:             authclient,
		authadminclient:        authadminclient,
		billingclient:          billingclient,
		rbacClient:             rbacClient,
		licenseclient:          licenseclient,
	}
}

func (s *Service) UserDataClient() userdata.Interface {
	return userdata.NewClient(s, s.rbacClient, s.licenseclient)
}
