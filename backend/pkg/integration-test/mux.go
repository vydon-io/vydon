package integrationtests_test

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"connectrpc.com/connect"
	db_queries "github.com/vydon-io/vydon/backend/gen/go/db"
	mysql_queries "github.com/vydon-io/vydon/backend/gen/go/db/dbschemas/mysql"
	pg_queries "github.com/vydon-io/vydon/backend/gen/go/db/dbschemas/postgresql"
	"github.com/vydon-io/vydon/backend/gen/go/protos/mgmt/v1alpha1/mgmtv1alpha1connect"
	auth_apikey "github.com/vydon-io/vydon/backend/internal/auth/apikey"
	auth_jwt "github.com/vydon-io/vydon/backend/internal/auth/jwt"
	auth_interceptor "github.com/vydon-io/vydon/backend/internal/connect/interceptors/auth"
	accounthooks "github.com/vydon-io/vydon/backend/internal/hooks/accounts"
	jobhooks "github.com/vydon-io/vydon/backend/internal/hooks/jobs"
	"github.com/vydon-io/vydon/backend/internal/userdata"
	"github.com/vydon-io/vydon/backend/internal/utils"
	"github.com/vydon-io/vydon/backend/pkg/mongoconnect"
	"github.com/vydon-io/vydon/backend/pkg/sqlconnect"
	v1alpha1_accounthookservice "github.com/vydon-io/vydon/backend/services/mgmt/v1alpha1/account-hooks-service"
	v1alpha_anonymizationservice "github.com/vydon-io/vydon/backend/services/mgmt/v1alpha1/anonymization-service"
	v1alpha1_connectiondataservice "github.com/vydon-io/vydon/backend/services/mgmt/v1alpha1/connection-data-service"
	v1alpha1_connectionservice "github.com/vydon-io/vydon/backend/services/mgmt/v1alpha1/connection-service"
	v1alpha1_jobservice "github.com/vydon-io/vydon/backend/services/mgmt/v1alpha1/job-service"
	v1alpha1_transformersservice "github.com/vydon-io/vydon/backend/services/mgmt/v1alpha1/transformers-service"
	v1alpha1_useraccountservice "github.com/vydon-io/vydon/backend/services/mgmt/v1alpha1/user-account-service"
	"github.com/vydon-io/vydon/internal/apikey"
	"github.com/vydon-io/vydon/internal/authmgmt"
	awsmanager "github.com/vydon-io/vydon/internal/aws"
	"github.com/vydon-io/vydon/internal/billing"
	"github.com/vydon-io/vydon/internal/connectiondata"
	presidioapi "github.com/vydon-io/vydon/internal/piidetect/presidio"
	"github.com/vydon-io/vydon/internal/rbac"
	vydon_gcp "github.com/vydon-io/vydon/internal/gcp"
	vydontypes "github.com/vydon-io/vydon/internal/vydon-types"
	"github.com/vydon-io/vydon/internal/vydondb"
	"github.com/vydon-io/vydon/internal/testutil"
	tcpostgres "github.com/vydon-io/vydon/internal/testutil/testcontainers/postgres"
)

var (
	validAuthUser = &authmgmt.User{Name: "foo", Email: "bar", Picture: "baz"}

	authinterceptor = auth_interceptor.NewInterceptor(
		func(ctx context.Context, header http.Header, spec connect.Spec) (context.Context, error) {
			// will need to further fill this out as the tests grow
			authuserid, err := utils.GetBearerTokenFromHeader(header, "Authorization")
			if err != nil {
				return nil, err
			}
			if apikey.IsValidV1WorkerKey(authuserid) {
				return auth_apikey.SetTokenData(ctx, &auth_apikey.TokenContextData{
					RawToken:   authuserid,
					ApiKey:     nil,
					ApiKeyType: apikey.WorkerApiKey,
				}), nil
			}
			return auth_jwt.SetTokenData(ctx, &auth_jwt.TokenContextData{
				AuthUserId: authuserid,
				Claims:     &auth_jwt.CustomClaims{Email: &validAuthUser.Email},
			}), nil
		},
	)
)

const (
	// OSS, Unauthenticated, Licensed
	openSourceUnauthenticatedLicensedPostfix = "/oss-unauthenticated-licensed"
	// OSS, Authenticated, Licensed
	openSourceAuthenticatedLicensedPostfix = "/oss-authenticated-licensed"
	// OSS, Unauthenticated, Unlicensed
	openSourceUnauthenticatedUnlicensedPostfix = "/oss-unauthenticated-unlicensed"
	// NeoCloud, Licensed, Authenticated
	neoCloudAuthenticatedLicensedPostfix = "/vydon-authenticated"
)

func (s *VydonApiTestClient) setupOssUnauthenticatedLicensedMux(
	ctx context.Context,
	pgcontainer *tcpostgres.PostgresTestContainer,
	logger *slog.Logger,
) (*http.ServeMux, error) {
	isLicensed := true
	isAuthEnabled := false
	isVydonCloud := false
	enforcedRbacClient, err := s.getEnforcedRbacClient(ctx, pgcontainer)
	if err != nil {
		return nil, fmt.Errorf("unable to get enforced rbac client: %w", err)
	}
	return s.setupMux(
		pgcontainer,
		isAuthEnabled,
		isLicensed,
		isVydonCloud,
		enforcedRbacClient,
		logger,
	)
}

func (s *VydonApiTestClient) setupOssLicensedAuthMux(
	ctx context.Context,
	pgcontainer *tcpostgres.PostgresTestContainer,
	logger *slog.Logger,
) (*http.ServeMux, error) {
	isLicensed := true
	isAuthEnabled := true
	isVydonCloud := false
	enforcedRbacClient, err := s.getEnforcedRbacClient(ctx, pgcontainer)
	if err != nil {
		return nil, fmt.Errorf("unable to get enforced rbac client: %w", err)
	}
	return s.setupMux(
		pgcontainer,
		isAuthEnabled,
		isLicensed,
		isVydonCloud,
		enforcedRbacClient,
		logger,
	)
}

func (s *VydonApiTestClient) setupOssUnlicensedMux(
	pgcontainer *tcpostgres.PostgresTestContainer,
	logger *slog.Logger,
) (*http.ServeMux, error) {
	isLicensed := false
	isAuthEnabled := false
	isVydonCloud := false
	permissiveRbacClient := rbac.NewAllowAllClient()
	return s.setupMux(
		pgcontainer,
		isAuthEnabled,
		isLicensed,
		isVydonCloud,
		permissiveRbacClient,
		logger,
	)
}

func (s *VydonApiTestClient) setupNeoCloudMux(
	ctx context.Context,
	pgcontainer *tcpostgres.PostgresTestContainer,
	logger *slog.Logger,
) (*http.ServeMux, error) {
	isLicensed := true
	isAuthEnabled := true
	isVydonCloud := true
	enforcedRbacClient, err := s.getEnforcedRbacClient(ctx, pgcontainer)
	if err != nil {
		return nil, fmt.Errorf("unable to get enforced rbac client: %w", err)
	}
	return s.setupMux(
		pgcontainer,
		isAuthEnabled,
		isLicensed,
		isVydonCloud,
		enforcedRbacClient,
		logger,
	)
}

func (s *VydonApiTestClient) setupMux(
	pgcontainer *tcpostgres.PostgresTestContainer,
	isAuthEnabled bool,
	isLicensed bool,
	isVydonCloud bool,
	rbacClient rbac.Interface,
	logger *slog.Logger,
) (*http.ServeMux, error) {
	isPresidioEnabled := isLicensed || isVydonCloud

	maxAllowed := int64(10000)
	var license *testutil.FakeEELicense
	if isLicensed {
		license = testutil.NewFakeEELicense(testutil.WithIsValid())
	} else {
		license = testutil.NewFakeEELicense()
	}

	vydonDb := vydondb.New(pgcontainer.DB, db_queries.New())

	var billingclient billing.Interface
	if isVydonCloud {
		billingclient = s.Mocks.Billingclient
	} else {
		billingclient = nil
	}

	userService := v1alpha1_useraccountservice.New(
		&v1alpha1_useraccountservice.Config{
			IsAuthEnabled:            isAuthEnabled,
			IsVydonCloud:           isVydonCloud,
			DefaultMaxAllowedRecords: &maxAllowed,
		},
		vydondb.New(pgcontainer.DB, db_queries.New()),
		s.Mocks.TemporalConfigProvider,
		s.Mocks.Authclient,
		s.Mocks.Authmanagerclient,
		billingclient,
		rbacClient, // rbac client
		license,
	)
	userclient := userdata.NewClient(userService, rbacClient, license)

	transformerService := v1alpha1_transformersservice.New(
		&v1alpha1_transformersservice.Config{
			IsPresidioEnabled: isPresidioEnabled,
		},
		vydondb.New(pgcontainer.DB, db_queries.New()),
		s.Mocks.Presidio.Entities,
		userclient,
		license,
	)

	sqlmanagerclient := NewTestSqlManagerClient()

	connectionService := v1alpha1_connectionservice.New(
		&v1alpha1_connectionservice.Config{IsVydonCloud: isVydonCloud},
		vydonDb,
		userclient,
		mongoconnect.NewConnector(),
		awsmanager.New(),
		sqlmanagerclient,
		&sqlconnect.SqlOpenConnector{},
	)

	var jobhookService *jobhooks.Service
	if isLicensed {
		jobhookService = jobhooks.New(
			vydonDb,
			userclient,
			jobhooks.WithEnabled(),
		)
	} else {
		jobhookService = jobhooks.New(
			vydonDb,
			userclient,
		)
	}

	awsManager := awsmanager.New()
	sqlConnector := &sqlconnect.SqlOpenConnector{}
	pgquerier := pg_queries.New()
	mysqlquerier := mysql_queries.New()
	mongoconnector := mongoconnect.NewConnector()
	sqlmanager := sqlmanagerclient
	gcpmanager := vydon_gcp.NewManager()
	vydontyperegistry := vydontypes.NewTypeRegistry(logger)

	connectiondatabuilder := connectiondata.NewConnectionDataBuilder(
		sqlConnector,
		sqlmanager,
		pgquerier,
		mysqlquerier,
		awsManager,
		gcpmanager,
		mongoconnector,
		vydontyperegistry,
	)

	jobService := v1alpha1_jobservice.New(
		&v1alpha1_jobservice.Config{IsAuthEnabled: isAuthEnabled, IsVydonCloud: isVydonCloud},
		vydonDb,
		s.Mocks.TemporalClientManager,
		connectionService,
		sqlmanagerclient,
		jobhookService,
		userclient,
		connectiondatabuilder,
	)

	var presAnalyzeClient presidioapi.AnalyzeInterface
	var presAnonClient presidioapi.AnonymizeInterface

	anonymizationService := v1alpha_anonymizationservice.New(
		&v1alpha_anonymizationservice.Config{
			IsPresidioEnabled: isPresidioEnabled,
			IsAuthEnabled:     isAuthEnabled,
			IsVydonCloud:    isVydonCloud,
		},
		nil, // meter
		userclient,
		userService,
		transformerService,
		presAnalyzeClient,
		presAnonClient,
		vydonDb,
		license,
	)

	connectionDataService := v1alpha1_connectiondataservice.New(
		&v1alpha1_connectiondataservice.Config{},
		connectionService,
		connectiondatabuilder,
	)

	accountHookService := v1alpha1_accounthookservice.New(
		accounthooks.New(
			vydonDb,
			userclient,
			accounthooks.WithSlackClient(s.Mocks.Slackclient),
		),
	)

	mux := http.NewServeMux()

	interceptors := []connect.Interceptor{}

	if isAuthEnabled {
		interceptors = append(interceptors, authinterceptor)
	}

	mux.Handle(mgmtv1alpha1connect.NewUserAccountServiceHandler(
		userService,
		connect.WithInterceptors(interceptors...),
	))
	mux.Handle(mgmtv1alpha1connect.NewTransformersServiceHandler(
		transformerService,
		connect.WithInterceptors(interceptors...),
	))
	mux.Handle(mgmtv1alpha1connect.NewConnectionServiceHandler(
		connectionService,
		connect.WithInterceptors(interceptors...),
	))
	mux.Handle(mgmtv1alpha1connect.NewJobServiceHandler(
		jobService,
		connect.WithInterceptors(interceptors...),
	))
	mux.Handle(mgmtv1alpha1connect.NewAnonymizationServiceHandler(
		anonymizationService,
		connect.WithInterceptors(interceptors...),
	))
	mux.Handle(mgmtv1alpha1connect.NewConnectionDataServiceHandler(
		connectionDataService,
		connect.WithInterceptors(interceptors...),
	))

	if isLicensed {
		mux.Handle(mgmtv1alpha1connect.NewAccountHookServiceHandler(
			accountHookService,
			connect.WithInterceptors(interceptors...),
		))
	} else {
		mux.Handle(mgmtv1alpha1connect.NewAccountHookServiceHandler(
			mgmtv1alpha1connect.UnimplementedAccountHookServiceHandler{},
			connect.WithInterceptors(interceptors...),
		))
	}

	return mux, nil
}

func (s *VydonApiTestClient) getEnforcedRbacClient(
	_ context.Context,
	_ *tcpostgres.PostgresTestContainer,
) (rbac.Interface, error) {
	return rbac.NewAllowAllClient(), nil
}
