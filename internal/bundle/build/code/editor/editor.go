package editor

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/plantoncloud/mactl/internal/app/build/code/vscode"
)

func Setup() error {
	log.Info("ensuring editors")
	if err := vscode.Setup(); err != nil {
		return errors.Wrap(err, "failed to ensure golang")
	}
	log.Info("ensured editors")
	return nil
}
