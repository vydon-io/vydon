package v1alpha1_metricsservice

import (
	promv1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/vydon-io/vydon/backend/gen/go/protos/mgmt/v1alpha1/mgmtv1alpha1connect"
	"github.com/vydon-io/vydon/backend/internal/userdata"
)

type Service struct {
	cfg *Config

	userdataclient   userdata.Interface
	jobservice       mgmtv1alpha1connect.JobServiceHandler
	prometheusclient promv1.API
}

type Config struct {
	IsAuthEnabled bool
}

func New(
	cfg *Config,
	userdataclient userdata.Interface,
	jobservice mgmtv1alpha1connect.JobServiceHandler,
	promclient promv1.API,
) *Service {
	return &Service{
		cfg:              cfg,
		userdataclient:   userdataclient,
		jobservice:       jobservice,
		prometheusclient: promclient,
	}
}
