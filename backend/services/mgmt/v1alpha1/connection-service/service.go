package v1alpha1_connectionservice

import (
	"github.com/vydon-io/vydon/backend/internal/userdata"
	"github.com/vydon-io/vydon/backend/pkg/mongoconnect"
	"github.com/vydon-io/vydon/backend/pkg/sqlconnect"
	sql_manager "github.com/vydon-io/vydon/backend/pkg/sqlmanager"
	awsmanager "github.com/vydon-io/vydon/internal/aws"
	"github.com/vydon-io/vydon/internal/vydondb"
)

type Service struct {
	cfg            *Config
	db             *vydondb.VydonDb
	userclient     userdata.Interface
	sqlConnector   sqlconnect.SqlConnector
	sqlmanager     sql_manager.SqlManagerClient
	mongoconnector mongoconnect.Interface
	awsManager     awsmanager.VydonAwsManagerClient
}

type Config struct {
	IsVydonCloud bool
}

func New(
	cfg *Config,
	db *vydondb.VydonDb,
	userclient userdata.Interface,
	mongoconnector mongoconnect.Interface,
	awsManager awsmanager.VydonAwsManagerClient,
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
