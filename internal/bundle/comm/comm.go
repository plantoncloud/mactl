package comm

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/mactl/internal/app/comm/discord"
	"github.com/plantoncloud/mactl/internal/app/comm/slack"
	"github.com/plantoncloud/mactl/internal/app/comm/telegram"
	"github.com/plantoncloud/mactl/internal/app/comm/whatsapp"
	log "github.com/sirupsen/logrus"
)

func Setup() error {
	log.Info("ensuring slack")
	if err := slack.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure slack")
	}
	log.Info("ensured slack")
	log.Info("ensuring whatsapp")
	if err := whatsapp.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure whatsapp")
	}
	log.Info("ensured whatsapp")
	log.Info("ensuring telegram")
	if err := telegram.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure telegram")
	}
	log.Info("ensured telegram")
	//todo: failed installing.
	//log.Info("ensuring gitter")
	//if err := gitter.Setup(); err != nil {
	//	return errors.Wrap(err, "failed to ensure gitter")
	//}
	//log.Info("ensured gitter")
	log.Info("ensuring discord")
	if err := discord.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure discord")
	}
	log.Info("ensured discord")
	return nil
}
