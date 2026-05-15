package integrationtest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	tcvydonapi "github.com/vydon-io/vydon/backend/pkg/integration-test"
	"github.com/vydon-io/vydon/internal/testutil"
	tcdynamodb "github.com/vydon-io/vydon/internal/testutil/testcontainers/dynamodb"
	tcmongodb "github.com/vydon-io/vydon/internal/testutil/testcontainers/mongodb"
	tcmysql "github.com/vydon-io/vydon/internal/testutil/testcontainers/mysql"
	tcpostgres "github.com/vydon-io/vydon/internal/testutil/testcontainers/postgres"
	tcredis "github.com/vydon-io/vydon/internal/testutil/testcontainers/redis"
	tcmssql "github.com/vydon-io/vydon/internal/testutil/testcontainers/sqlserver"
)

const vydonDbMigrationsPath = "../../../../backend/sql/postgresql/schema"

func Test_Workflow(t *testing.T) {
	t.Parallel()
	ok := testutil.ShouldRunWorkerIntegrationTest()
	if !ok {
		return
	}
	ctx := context.Background()

	vydonApi, err := tcvydonapi.NewVydonApiTestClient(ctx, t, tcvydonapi.WithMigrationsDirectory(vydonDbMigrationsPath))
	if err != nil {
		t.Fatal(err)
	}

	connclient := vydonApi.OSSUnauthenticatedLicensedClients.Connections()
	accountId := tcvydonapi.CreatePersonalAccount(ctx, t, vydonApi.OSSUnauthenticatedLicensedClients.Users())
	dbManagers := NewTestDatabaseManagers(t)

	t.Run("postgres", func(t *testing.T) {
		t.Log("Starting postgres tests")
		t.Parallel()
		postgres, err := tcpostgres.NewPostgresTestSyncContainer(ctx, []tcpostgres.Option{}, []tcpostgres.Option{})
		if err != nil {
			t.Fatal(err)
		}
		sourceConn := tcvydonapi.CreatePostgresConnection(ctx, t, connclient, accountId, "postgres-source", postgres.Source.URL)
		destConn := tcvydonapi.CreatePostgresConnection(ctx, t, connclient, accountId, "postgres-dest", postgres.Target.URL)

		_, err = postgres.Source.DB.Exec(ctx, `CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
		require.NoError(t, err)

		// Sync workflow tests
		t.Run("types", func(t *testing.T) {
			t.Parallel()
			test_postgres_types(t, ctx, postgres, vydonApi, dbManagers, accountId, sourceConn, destConn)
		})

		t.Run("edgecases", func(t *testing.T) {
			t.Parallel()
			test_postgres_edgecases(t, ctx, postgres, vydonApi, dbManagers, accountId, sourceConn, destConn)
		})

		t.Run("virtual_foreign_keys", func(t *testing.T) {
			t.Parallel()
			test_postgres_virtual_foreign_keys(t, ctx, postgres, vydonApi, dbManagers, accountId, sourceConn, destConn)
		})

		t.Run("javascript_transformers", func(t *testing.T) {
			t.Parallel()
			test_postgres_javascript_transformers(t, ctx, postgres, vydonApi, dbManagers, accountId, sourceConn, destConn)
		})

		t.Run("skip_foreign_keys_violations", func(t *testing.T) {
			t.Parallel()
			test_postgres_skip_foreign_keys_violations(t, ctx, postgres, vydonApi, dbManagers, accountId, sourceConn, destConn)
		})

		t.Run("foreign_keys_violations_error", func(t *testing.T) {
			t.Parallel()
			test_postgres_foreign_keys_violations_error(t, ctx, postgres, vydonApi, dbManagers, accountId, sourceConn, destConn)
		})

		t.Run("subsetting", func(t *testing.T) {
			t.Parallel()
			test_postgres_subsetting(t, ctx, postgres, vydonApi, dbManagers, accountId, sourceConn, destConn)
		})

		t.Run("primary_key_transformations", func(t *testing.T) {
			t.Parallel()
			redis, err := tcredis.NewRedisTestContainer(ctx)
			require.NoError(t, err)

			test_postgres_primary_key_transformations(t, ctx, postgres, redis, vydonApi, dbManagers, accountId, sourceConn, destConn)

			t.Cleanup(func() {
				err := redis.TearDown(ctx)
				require.NoError(t, err)
			})
		})

		t.Run("small_batch_size", func(t *testing.T) {
			t.Parallel()
			test_postgres_small_batch_size(t, ctx, postgres, vydonApi, dbManagers, accountId, sourceConn, destConn)
		})

		t.Run("complex", func(t *testing.T) {
			t.Parallel()
			test_postgres_complex(t, ctx, postgres, vydonApi, dbManagers, accountId, sourceConn, destConn)
		})

		t.Run("passthrough_on_new_column_addition", func(t *testing.T) {
			t.Parallel()
			test_postgres_passthrough_on_new_column_addition(t, ctx, postgres, vydonApi, dbManagers, accountId, sourceConn, destConn)
		})

		t.Run("schema_reconciliation", func(t *testing.T) {
			t.Parallel()
			t.Run("truncate", func(t *testing.T) {
				t.Parallel()
				test_postgres_schema_reconciliation(t, ctx, postgres, vydonApi, dbManagers, accountId, sourceConn, destConn, true)
			})
			t.Run("retain_data", func(t *testing.T) {
				t.Parallel()
				test_postgres_schema_reconciliation(t, ctx, postgres, vydonApi, dbManagers, accountId, sourceConn, destConn, false)
			})
		})

		// Generate workflow tests
		t.Run("generate", func(t *testing.T) {
			t.Parallel()
			test_postgres_generate_workflow(t, ctx, postgres, vydonApi, dbManagers, accountId, destConn)
		})

		t.Cleanup(func() {
			err := postgres.TearDown(ctx)
			if err != nil {
				t.Fatal(err)
			}
		})
	})

	t.Run("mysql", func(t *testing.T) {
		t.Log("Starting mysql tests")
		t.Parallel()
		mysql, err := tcmysql.NewMysqlTestSyncContainer(ctx, []tcmysql.Option{}, []tcmysql.Option{})
		if err != nil {
			t.Fatal(err)
		}
		sourceConn := tcvydonapi.CreateMysqlConnection(ctx, t, connclient, accountId, "mysql-source", mysql.Source.URL)
		destConn := tcvydonapi.CreateMysqlConnection(ctx, t, connclient, accountId, "mysql-dest", mysql.Target.URL)

		t.Run("types", func(t *testing.T) {
			t.Parallel()
			test_mysql_types(t, ctx, mysql, vydonApi, dbManagers, accountId, sourceConn, destConn)
		})

		t.Run("edgecases", func(t *testing.T) {
			t.Parallel()
			test_mysql_edgecases(t, ctx, mysql, vydonApi, dbManagers, accountId, sourceConn, destConn)
		})

		t.Run("composite_keys", func(t *testing.T) {
			t.Parallel()
			test_mysql_composite_keys(t, ctx, mysql, vydonApi, dbManagers, accountId, sourceConn, destConn)
		})
		t.Run("on_conflict_do_update", func(t *testing.T) {
			t.Parallel()
			test_mysql_on_conflict_do_update(t, ctx, mysql, vydonApi, dbManagers, accountId, sourceConn, destConn)
		})

		t.Run("schema_reconciliation", func(t *testing.T) {
			t.Parallel()
			t.Run("truncate", func(t *testing.T) {
				t.Parallel()
				test_mysql_schema_reconciliation(t, ctx, mysql, vydonApi, dbManagers, accountId, sourceConn, destConn, true)
			})
			t.Run("retain_data", func(t *testing.T) {
				t.Parallel()
				test_mysql_schema_reconciliation(t, ctx, mysql, vydonApi, dbManagers, accountId, sourceConn, destConn, false)
			})
		})

		t.Run("complex", func(t *testing.T) {
			t.Parallel()
			test_mysql_complex(t, ctx, mysql, vydonApi, dbManagers, accountId, sourceConn, destConn)
		})

		t.Cleanup(func() {
			err := mysql.TearDown(ctx)
			if err != nil {
				t.Fatal(err)
			}
		})
	})

	t.Run("mssql", func(t *testing.T) {
		// MSSQL schema bootstrap relied on the upstream EE sqlmanager (Get*Init*Statements,
		// GetColumnsByTables, GetTableConstraintsByTables, GetFunctionsBySchemas). The OSS
		// distribution stubs those methods until a native MSSQL provider is implemented;
		// the worker integration tests below cannot exercise the real path yet.
		t.Skip("mssql worker integration tests are paused until the native mssql provider lands")
		t.Log("Starting mssql tests")
		t.Parallel()
		mssql, err := tcmssql.NewMssqlTestSyncContainer(ctx, []tcmssql.Option{}, []tcmssql.Option{})
		if err != nil {
			t.Fatal(err)
		}
		sourceConn := tcvydonapi.CreateMssqlConnection(ctx, t, connclient, accountId, "mssql-source", mssql.Source.URL)
		destConn := tcvydonapi.CreateMssqlConnection(ctx, t, connclient, accountId, "mssql-dest", mssql.Target.URL)

		t.Run("types", func(t *testing.T) {
			t.Parallel()
			test_mssql_types(t, ctx, mssql, vydonApi, dbManagers, accountId, sourceConn, destConn)
		})

		t.Run("cross_schema_foreign_keys", func(t *testing.T) {
			t.Parallel()
			test_mssql_cross_schema_foreign_keys(t, ctx, mssql, vydonApi, dbManagers, accountId, sourceConn, destConn)
		})

		t.Run("subset", func(t *testing.T) {
			t.Parallel()
			test_mssql_subset(t, ctx, mssql, vydonApi, dbManagers, accountId, sourceConn, destConn)
		})

		t.Run("identity_columns", func(t *testing.T) {
			t.Parallel()
			test_mssql_identity_columns(t, ctx, mssql, vydonApi, dbManagers, accountId, sourceConn, destConn)
		})

		t.Cleanup(func() {
			err := mssql.TearDown(ctx)
			if err != nil {
				t.Fatal(err)
			}
		})
	})

	t.Run("dynamodb", func(t *testing.T) {
		t.Log("Starting dynamodb tests")
		t.Parallel()
		dynamo, err := tcdynamodb.NewDynamoDBTestSyncContainer(ctx, t, []tcdynamodb.Option{}, []tcdynamodb.Option{})
		if err != nil {
			t.Fatal(err)
		}
		sourceConn := tcvydonapi.CreateDynamoDBConnection(
			ctx,
			t,
			connclient,
			accountId,
			"dynamo-source",
			dynamo.Source.URL,
			dynamo.Source.Credentials,
		)
		destConn := tcvydonapi.CreateDynamoDBConnection(
			ctx,
			t,
			connclient,
			accountId,
			"dynamo-dest",
			dynamo.Target.URL,
			dynamo.Target.Credentials,
		)

		t.Run("types", func(t *testing.T) {
			t.Parallel()
			test_dynamodb_alltypes(t, ctx, dynamo, vydonApi, dbManagers, accountId, sourceConn, destConn)
		})

		t.Run("subset", func(t *testing.T) {
			t.Parallel()
			test_dynamodb_subset(t, ctx, dynamo, vydonApi, dbManagers, accountId, sourceConn, destConn)
		})

		t.Run("default_transformers", func(t *testing.T) {
			t.Parallel()
			test_dynamodb_default_transformers(t, ctx, dynamo, vydonApi, dbManagers, accountId, sourceConn, destConn)
		})

		t.Cleanup(func() {
			err := dynamo.TearDown(ctx)
			if err != nil {
				t.Fatal(err)
			}
		})
	})

	t.Run("mongodb", func(t *testing.T) {
		t.Log("Starting mongodb tests")
		t.Parallel()
		mongodb, err := tcmongodb.NewMongoDBTestSyncContainer(ctx, t)
		if err != nil {
			t.Fatal(err)
		}

		sourceConn := tcvydonapi.CreateMongodbConnection(ctx, t, connclient, accountId, "mongodb-source", mongodb.Source.URL)
		destConn := tcvydonapi.CreateMongodbConnection(ctx, t, connclient, accountId, "mongodb-dest", mongodb.Target.URL)

		t.Run("types", func(t *testing.T) {
			t.Parallel()
			test_mongodb_alltypes(t, ctx, mongodb, vydonApi, dbManagers, accountId, sourceConn, destConn)
		})

		t.Run("transform", func(t *testing.T) {
			t.Parallel()
			test_mongodb_transform(t, ctx, mongodb, vydonApi, dbManagers, accountId, sourceConn, destConn)
		})

		t.Cleanup(func() {
			err := mongodb.TearDown(ctx)
			if err != nil {
				t.Fatal(err)
			}
		})
	})

	t.Cleanup(func() {
		err = vydonApi.TearDown(ctx)
		if err != nil {
			panic(err)
		}
	})
}
