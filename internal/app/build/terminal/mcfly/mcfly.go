// Package mcfly installs https://github.com/cantino/mcfly
package mcfly

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/mactl/internal/installer/brew"
	log "github.com/sirupsen/logrus"
)

const (
	BrewPkg = "cantino/mcfly/mcfly"
)

func Setup() error {
	log.Info("installing mcfly")
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to install mcfly")
	}
	log.Info("installed mcfly")
	return nil
}

func Upgrade() error {
	log.Info("upgrading mcfly")
	if err := brew.Upgrade(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to upgrade mcfly")
	}
	log.Info("upgraded mcfly")
	return nil
}
