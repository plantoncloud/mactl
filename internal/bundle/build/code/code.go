package code

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/plantoncloud/mactl/internal/bundle/build/code/editor"
	"github.com/plantoncloud/mactl/internal/bundle/build/code/lang"
)

func Setup() error {
	log.Info("ensuring code bundle")
	if err := lang.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure code bundle")
	}
	log.Info("ensured code bundle")
	log.Info("ensuring editor bundle")
	if err := editor.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure editor bundle")
	}
	log.Info("ensured editor bundle")
	return nil
}
