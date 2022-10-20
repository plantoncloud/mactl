package bootstrap

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/mactl/internal/app/browser/chrome"
	"github.com/plantoncloud/mactl/internal/app/build/terminal/iterm"
	"github.com/plantoncloud/mactl/internal/app/installer/mas"
	"github.com/plantoncloud/mactl/internal/app/keyboard/rectangle"
	"github.com/plantoncloud/mactl/internal/bundle/build"
	"github.com/plantoncloud/mactl/internal/bundle/hotkey"
	"github.com/plantoncloud/mactl/internal/bundle/tool"
	"github.com/plantoncloud/mactl/internal/optimize/dock"
	log "github.com/sirupsen/logrus"
)

const (
	UpgradeOptionalCommApps = true
)

func Upgrade() error {
	log.Info("bootstrap fulfill started")

	done := make(chan bool)
	fatalErrors := make(chan error)

	if err := fulfill(fatalErrors); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to ensure fulfill all bootstrap components")
	}

	if err := open(fatalErrors); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to open all apps installed as part of bootstrap fulfill")
	}

	if err := postOpen(fatalErrors); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to run post open steps as part of bootstrap fulfill")
	}

	close(done)
	select {
	case <-done:
		break
	case err := <-fatalErrors:
		close(fatalErrors)
		log.Warnf("Error encountered : %v", err.Error())
		return err
	}

	log.Info("bootstrap fulfill completed")
	return nil
}

func fulfill(fatalErrors chan error) error {
	log.Info("optimizing dock")
	if err := dock.Optimize(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to optimize dock")
	}
	log.Info("optimized dock")

	log.Info("upgrading installers")
	if err := mas.UpgradeMas(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to fulfill installers")
	}
	log.Info("upgraded installers")

	log.Info("upgrading hotkey")
	if err := hotkey.Upgrade(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to fulfill hotkey")
	}
	log.Info("upgraded hotkey")

	log.Info("upgrading window-management")
	if err := rectangle.Upgrade(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to fulfill window management using rectangle app")
	}
	log.Info("upgraded window-management")

	log.Info("upgrading browser")
	if err := chrome.Upgrade(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to fulfill browser")
	}
	log.Info("upgraded browser")

	log.Info("upgrading terminal")
	if err := iterm.Upgrade(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to fulfill scm")
	}
	log.Info("upgraded terminal")

	log.Info("upgrading tool bundle")
	if err := tool.Upgrade(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to fulfill tool bundle")
	}
	log.Info("upgraded tool bundle")

	log.Info("upgrading build bundle")
	if err := build.Upgrade(); err != nil {
		fatalErrors <- errors.Wrap(err, "failed to fulfill ides")
	}
	log.Info("upgraded build bundle")
	return nil
}
