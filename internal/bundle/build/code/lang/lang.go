package lang

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/mactl/internal/app/build/code/lang/golang"
	"github.com/plantoncloud/mactl/internal/app/build/code/lang/java"
	"github.com/plantoncloud/mactl/internal/app/build/code/lang/javascript"
	"github.com/plantoncloud/mactl/internal/app/build/code/lang/javascript/nodejs"
	"github.com/plantoncloud/mactl/internal/app/build/code/lang/python"
	"github.com/plantoncloud/mactl/internal/app/build/code/lang/sql"
	"github.com/plantoncloud/mactl/internal/app/build/scm/battenberg"
	log "github.com/sirupsen/logrus"
)

func Setup() error {
	log.Info("ensuring golang")
	if err := golang.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure golang")
	}
	log.Info("ensured golang")
	log.Info("ensuring javascript")
	if err := javascript.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure javascript")
	}
	log.Info("ensured javascript")
	log.Info("ensuring python")
	if err := python.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure python")
	}
	log.Info("ensured python")
	log.Info("ensuring battenberg")
	if err := battenberg.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure battenberg")
	}
	log.Info("ensured battenberg")
	log.Info("ensuring java")
	if err := java.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure java")
	}
	log.Info("ensured java")
	log.Info("ensuring sql")
	if err := sql.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure sql")
	}
	log.Info("ensured sql")
	return nil
}

func Upgrade() error {
	log.Info("upgrading golang")
	if err := golang.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to fulfill golang")
	}
	log.Info("upgraded golang")

	log.Info("upgrading nodejs")
	if err := nodejs.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to fulfill nodejs")
	}
	log.Info("upgraded nodejs")

	log.Info("upgrading python")
	if err := python.Upgrade(); err != nil {
		return errors.Wrap(err, "failed to fulfill python")
	}
	log.Info("upgraded python")

	// java is not being upgraded. need to discuss and will come up with strategy to upgrade java version
	return nil
}
