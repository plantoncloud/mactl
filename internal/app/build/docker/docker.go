package docker

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/mactl/internal/installer/brew"
	"github.com/plantoncloud/mactl/internal/installer/macapp"
	log "github.com/sirupsen/logrus"
)

const (
	BrewCaskPkg = "docker"
	MacAppName  = "Docker.app"
)

type VmConfigTemplateInput struct {
	HomeDir      string
	Cpu          int
	MemoryInMb   int
	DiskSizeInMb int
}

func Setup() error {
	log.Infof("installing %s", BrewCaskPkg)
	if err := brew.InstallCask(BrewCaskPkg); err != nil {
		return errors.Wrapf(err, "failed to install %s pkg using brew", BrewCaskPkg)
	}
	log.Infof("installed %s", BrewCaskPkg)
	return nil
}

func Upgrade() error {
	log.Infof("upgrading %s", BrewCaskPkg)
	if err := brew.UpgradeCask(BrewCaskPkg); err != nil {
		return errors.Wrapf(err, "failed to upgrade %s pkg using brew", BrewCaskPkg)
	}
	log.Infof("upgraded %s", BrewCaskPkg)
	return nil
}

func Open() error {
	if err := macapp.Open(MacAppName); err != nil {
		return errors.Wrapf(err, "failed to open %s", MacAppName)
	}
	return nil
}
