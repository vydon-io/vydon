package sqlmanager

import (
	"testing"

	"github.com/stretchr/testify/require"
	mysql_queries "github.com/vydon-io/vydon/backend/gen/go/db/dbschemas/mysql"
	pg_queries "github.com/vydon-io/vydon/backend/gen/go/db/dbschemas/postgresql"
	mssql_queries "github.com/vydon-io/vydon/backend/pkg/mssql-querier"
	"github.com/vydon-io/vydon/backend/pkg/sqlconnect"
	connectionmanager "github.com/vydon-io/vydon/internal/connection-manager"
	"github.com/vydon-io/vydon/internal/connection-manager/providers/sqlprovider"
)

func Test_NewSqlManager(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		mgr := NewSqlManager()
		require.NotNil(t, mgr)
	})
	t.Run("new_with_opts", func(t *testing.T) {
		mgr := NewSqlManager(
			WithConnectionManager(connectionmanager.NewConnectionManager(sqlprovider.NewProvider(&sqlconnect.SqlOpenConnector{}))),
			WithPostgresQuerier(pg_queries.New()),
			WithMssqlQuerier(mssql_queries.New()),
			WithMysqlQuerier(mysql_queries.New()),
		)
		require.NotNil(t, mgr)
	})
}
