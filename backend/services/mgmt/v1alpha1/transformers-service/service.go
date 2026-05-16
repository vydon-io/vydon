package v1alpha1_transformersservice

import (
	"github.com/vydon-io/vydon/backend/internal/userdata"
	"github.com/vydon-io/vydon/internal/license"
	presidioapi "github.com/vydon-io/vydon/internal/piidetect/presidio"
	"github.com/vydon-io/vydon/internal/vydondb"
)

type Service struct {
	cfg            *Config
	db             *vydondb.VydonDb
	entityclient   presidioapi.EntityInterface
	userdataclient userdata.Interface
	license        license.EEInterface
}

type Config struct {
	IsPresidioEnabled bool
}

func New(
	cfg *Config,
	db *vydondb.VydonDb,
	recognizerclient presidioapi.EntityInterface,
	userdataclient userdata.Interface,
	license license.EEInterface,
) *Service {
	return &Service{
		cfg:            cfg,
		db:             db,
		entityclient:   recognizerclient,
		userdataclient: userdataclient,
		license:        license,
	}
}
