package nodejs

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/mactl/internal/installer/brew"
	log "github.com/sirupsen/logrus"
)

const (
	BrewPkgNvm  = "nvm"
	BrewPkgYarn = "yarn"
)

func Setup() error {
	log.Info("installing nvm")
	err := brew.Install(BrewPkgNvm)
	if err != nil {
		return errors.Wrapf(err, "failed to install nvm")
	}
	log.Info("installed nvm")
	log.Info("installing yarn")
	if err := brew.Install(BrewPkgYarn); err != nil {
		return errors.Wrapf(err, "failed to install yarn")
	}
	log.Info("installed yarn")
	return nil
}

func Upgrade() error {
	log.Info("upgrading nvm")
	err := brew.Upgrade(BrewPkgNvm)
	if err != nil {
		return errors.Wrapf(err, "failed to upgrade nvm")
	}
	log.Info("upgraded nvm")

	log.Info("upgrading yarn")
	if err := brew.Upgrade(BrewPkgYarn); err != nil {
		return errors.Wrapf(err, "failed to upgrade yarn")
	}
	log.Info("upgraded yarn")
	return nil
}
