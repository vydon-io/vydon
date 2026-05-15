package v1alpha1_connectionservice

import (
	"github.com/vydon-io/vydon/backend/internal/userdata"
	"github.com/vydon-io/vydon/backend/pkg/mongoconnect"
	"github.com/vydon-io/vydon/backend/pkg/sqlconnect"
	sql_manager "github.com/vydon-io/vydon/backend/pkg/sqlmanager"
	awsmanager "github.com/vydon-io/vydon/internal/aws"
	"github.com/vydon-io/vydon/internal/neosyncdb"
)

type Service struct {
	cfg            *Config
	db             *neosyncdb.NeosyncDb
	userclient     userdata.Interface
	sqlConnector   sqlconnect.SqlConnector
	sqlmanager     sql_manager.SqlManagerClient
	mongoconnector mongoconnect.Interface
	awsManager     awsmanager.NeosyncAwsManagerClient
}

type Config struct {
	IsNeosyncCloud bool
}

func New(
	cfg *Config,
	db *neosyncdb.NeosyncDb,
	userclient userdata.Interface,
	mongoconnector mongoconnect.Interface,
	awsManager awsmanager.NeosyncAwsManagerClient,
	sqlmanager sql_manager.SqlManagerClient,
	sqlconnector sqlconnect.SqlConnector,
) *Service {
	return &Service{
		cfg:            cfg,
		db:             db,
		userclient:     userclient,
		sqlmanager:     sqlmanager,
		mongoconnector: mongoconnector,
		awsManager:     awsManager,
		sqlConnector:   sqlconnector,
	}
}
