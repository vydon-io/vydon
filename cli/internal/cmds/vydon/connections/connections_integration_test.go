package connections_cmd

import (
	"context"
	"testing"

	mgmtv1alpha1 "github.com/vydon-io/vydon/backend/gen/go/protos/mgmt/v1alpha1"
	integrationtests_test "github.com/vydon-io/vydon/backend/pkg/integration-test"
	tcvydonapi "github.com/vydon-io/vydon/backend/pkg/integration-test"
	"github.com/vydon-io/vydon/internal/testutil"
	"github.com/stretchr/testify/require"
)

const vydonDbMigrationsPath = "../../../../../backend/sql/postgresql/schema"

func Test_Connections(t *testing.T) {
	t.Parallel()
	ok := testutil.ShouldRunCLIIntegrationTest()
	if !ok {
		return
	}
	ctx := context.Background()

	vydonApi, err := tcvydonapi.NewVydonApiTestClient(ctx, t, tcvydonapi.WithMigrationsDirectory(vydonDbMigrationsPath))
	if err != nil {
		panic(err)
	}
	postgresUrl := "postgresql://postgres:foofar@localhost:5434/vydon"

	t.Run("list_unauthed", func(t *testing.T) {
		accountId := tcvydonapi.CreatePersonalAccount(ctx, t, vydonApi.OSSUnauthenticatedLicensedClients.Users())
		conn1 := tcvydonapi.CreatePostgresConnection(ctx, t, vydonApi.OSSUnauthenticatedLicensedClients.Connections(), accountId, "conn1", postgresUrl)
		conn2 := tcvydonapi.CreatePostgresConnection(ctx, t, vydonApi.OSSUnauthenticatedLicensedClients.Connections(), accountId, "conn2", postgresUrl)
		conns := []*mgmtv1alpha1.Connection{conn1, conn2}
		connections, err := getConnections(ctx, vydonApi.OSSUnauthenticatedLicensedClients.Connections(), accountId)
		require.NoError(t, err)
		require.Len(t, connections, len(conns))
	})

	t.Run("list_auth", func(t *testing.T) {
		testAuthUserId := "c3b32842-9b70-4f4e-ad45-9cab26c6f2f1"
		userclient := vydonApi.OSSAuthenticatedLicensedClients.Users(integrationtests_test.WithUserId(testAuthUserId))
		connclient := vydonApi.OSSAuthenticatedLicensedClients.Connections(integrationtests_test.WithUserId(testAuthUserId))
		tcvydonapi.SetUser(ctx, t, userclient)
		accountId := tcvydonapi.CreatePersonalAccount(ctx, t, userclient)
		conn1 := tcvydonapi.CreatePostgresConnection(ctx, t, connclient, accountId, "conn1", postgresUrl)
		conn2 := tcvydonapi.CreatePostgresConnection(ctx, t, connclient, accountId, "conn2", postgresUrl)
		conns := []*mgmtv1alpha1.Connection{conn1, conn2}
		connections, err := getConnections(ctx, connclient, accountId)
		require.NoError(t, err)
		require.Len(t, connections, len(conns))
	})

	t.Run("list_cloud", func(t *testing.T) {
		testAuthUserId := "34f3e404-c995-452b-89e4-9c486b491dab"
		userclient := vydonApi.VydonCloudAuthenticatedLicensedClients.Users(integrationtests_test.WithUserId(testAuthUserId))
		connclient := vydonApi.VydonCloudAuthenticatedLicensedClients.Connections(integrationtests_test.WithUserId(testAuthUserId))
		tcvydonapi.SetUser(ctx, t, userclient)
		accountId := tcvydonapi.CreatePersonalAccount(ctx, t, userclient)
		conn1 := tcvydonapi.CreatePostgresConnection(ctx, t, connclient, accountId, "conn1", postgresUrl)
		conn2 := tcvydonapi.CreatePostgresConnection(ctx, t, connclient, accountId, "conn2", postgresUrl)
		conns := []*mgmtv1alpha1.Connection{conn1, conn2}
		connections, err := getConnections(ctx, connclient, accountId)
		require.NoError(t, err)
		require.Len(t, connections, len(conns))
	})

	err = vydonApi.TearDown(ctx)
	if err != nil {
		panic(err)
	}
}
