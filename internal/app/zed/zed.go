package zed

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/mactl/internal/installer/brew"
	"github.com/plantoncloud/mactl/internal/installer/macapp"
	log "github.com/sirupsen/logrus"
)

const (
	AppName        = "zed"
	BrewPkg        = "zed"
	MacAppFileName = "Zed.app"
)

func Setup() error {
	log.Infof("installing %s", BrewPkg)
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to install %s pkg using brew", BrewPkg)
	}
	log.Infof("installed %s", BrewPkg)
	return nil
}

func Upgrade() error {
	log.Infof("upgrading %s", BrewPkg)
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrapf(err, "failed to upgrade %s pkg using brew", BrewPkg)
	}
	log.Infof("upgraded %s", BrewPkg)
	return nil
}

func GetPath() string {
	return macapp.GetPath(MacAppFileName)
}
