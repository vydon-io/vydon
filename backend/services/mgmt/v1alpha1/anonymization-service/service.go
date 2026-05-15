package v1alpha_anonymizationservice

import (
	"github.com/vydon-io/vydon/backend/gen/go/protos/mgmt/v1alpha1/mgmtv1alpha1connect"
	"github.com/vydon-io/vydon/backend/internal/userdata"
	"github.com/vydon-io/vydon/internal/license"
	presidioapi "github.com/vydon-io/vydon/internal/piidetect/presidio"
	"github.com/vydon-io/vydon/internal/vydondb"
	"go.opentelemetry.io/otel/metric"
)

type Service struct {
	cfg                *Config
	meter              metric.Meter // optional
	userdataclient     userdata.Interface
	useraccountService mgmtv1alpha1connect.UserAccountServiceClient
	transformerClient  mgmtv1alpha1connect.TransformersServiceClient
	analyze            presidioapi.AnalyzeInterface
	anonymize          presidioapi.AnonymizeInterface
	db                 *vydondb.VydonDb
	license            license.EEInterface
}

type Config struct {
	IsAuthEnabled           bool
	IsNeosyncCloud          bool
	IsPresidioEnabled       bool
	PresidioDefaultLanguage *string
}

func New(
	cfg *Config,
	meter metric.Meter,
	userdataclient userdata.Interface,
	useraccountService mgmtv1alpha1connect.UserAccountServiceClient,
	transformerClient mgmtv1alpha1connect.TransformersServiceClient,
	analyzeclient presidioapi.AnalyzeInterface,
	anonymizeclient presidioapi.AnonymizeInterface,
	db *vydondb.VydonDb,
	license license.EEInterface,
) *Service {
	return &Service{
		cfg:                cfg,
		meter:              meter,
		userdataclient:     userdataclient,
		useraccountService: useraccountService,
		transformerClient:  transformerClient,
		analyze:            analyzeclient,
		anonymize:          anonymizeclient,
		db:                 db,
		license:            license,
	}
}
