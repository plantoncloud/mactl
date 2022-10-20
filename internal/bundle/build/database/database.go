package database

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/plantoncloud/mactl/internal/app/build/database/postgresql"
)

func Setup() error {
	log.Info("ensuring postgresql")
	if err := postgresql.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure postgresql")
	}
	log.Info("ensured postgresql")
	return nil
}
