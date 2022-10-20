package postgresql

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/plantoncloud/mactl/internal/installer/brew"
)

const (
	BrewPkg = "postgresql"
)

func Setup() error {
	log.Info("installing postgresql")
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to install postgresql")
	}
	log.Info("installed postgresql")
	return nil
}

func Upgrade() error {
	log.Info("upgrading postgresql")
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to upgrade postgresql")
	}
	log.Info("upgraded postgresql")
	return nil
}
