package gcloud

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/mactl/internal/installer/brew"
	log "github.com/sirupsen/logrus"
)

const (
	BrewPkg = "google-cloud-sdk"
)

func Setup() error {
	log.Info("installing gcloud")
	err := brew.Install(BrewPkg)
	if err != nil {
		return errors.Wrapf(err, "failed to install gcloud")
	}
	log.Info("installed gcloud")
	return nil
}

func Upgrade() error {
	log.Info("upgrading gcloud")
	err := brew.Upgrade(BrewPkg)
	if err != nil {
		return errors.Wrapf(err, "failed to upgrade gcloud")
	}
	log.Info("upgraded gcloud")
	return nil
}
