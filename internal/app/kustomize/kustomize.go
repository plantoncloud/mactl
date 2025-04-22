package kustomize

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/plantoncloud/mactl/internal/installer/brew"
)

const (
	BrewPkg = "kustomize"
)

func Setup() error {
	log.Info("installing kustomize")
	if err := brew.Install(BrewPkg); err != nil {
		return errors.Wrap(err, "failed to install kustomize")
	}
	log.Info("installed kustomize")
	return nil
}
