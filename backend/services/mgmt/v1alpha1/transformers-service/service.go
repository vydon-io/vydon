package v1alpha1_transformersservice

import (
	"github.com/vydon-io/vydon/backend/internal/userdata"
	"github.com/vydon-io/vydon/internal/license"
	presidioapi "github.com/vydon-io/vydon/internal/piidetect/presidio"
	"github.com/vydon-io/vydon/internal/neosyncdb"
)

type Service struct {
	cfg            *Config
	db             *neosyncdb.NeosyncDb
	entityclient   presidioapi.EntityInterface
	userdataclient userdata.Interface
	license        license.EEInterface
}

type Config struct {
	IsPresidioEnabled bool
}

func New(
	cfg *Config,
	db *neosyncdb.NeosyncDb,
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
