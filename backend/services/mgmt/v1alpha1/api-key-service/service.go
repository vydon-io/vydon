package v1alpha1_apikeyservice

import (
	"github.com/vydon-io/vydon/backend/internal/userdata"
	"github.com/vydon-io/vydon/internal/neosyncdb"
)

type Service struct {
	cfg            *Config
	db             *neosyncdb.NeosyncDb
	userdataclient userdata.Interface
}

type Config struct {
	IsAuthEnabled bool
}

func New(
	cfg *Config,
	db *neosyncdb.NeosyncDb,
	userdataclient userdata.Interface,
) *Service {
	return &Service{cfg: cfg, db: db, userdataclient: userdataclient}
}
