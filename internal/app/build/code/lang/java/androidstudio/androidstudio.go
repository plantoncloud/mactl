package androidstudio

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/plantoncloud/mactl/internal/bundle/tool"
	"github.com/plantoncloud/mactl/internal/installer/brew"
	"os"
	"path/filepath"
)

const (
	BrewPkgSdk    = "android-sdk"
	BrewPkgStudio = "android-studio"
	LaunchAlias   = "aaa"
	LaunchScript  = `
#!/bin/sh

IDEA="/Applications/IntelliJ IDEA.app"

# were we given a file?
if [ -f "$1" ]; then
  open -a "$IDEA" "$1"
  exit 0
fi

if [ -f build.gradle ]; then
  open -a "$IDEA" "build.gradle"
  exit 0
fi

if [ -f pom.xml ]; then
  open -a "$IDEA" "pom.xml"
  exit 0
fi
open -a "$IDEA" .
`
)

func Setup() error {
	log.Info("installing android -tudio")
	if err := brew.Install(BrewPkgStudio); err != nil {
		return errors.Wrap(err, "failed to install android-studio")
	}
	log.Info("installed android-studio")
	log.Info("installing android-sdk")
	if err := brew.Install(BrewPkgSdk); err != nil {
		return errors.Wrap(err, "failed to install android-sdk")
	}
	log.Info("installed android-sdk")
	log.Infof("ensuring launch alias")
	if err := setupLaunchAlias(); err != nil {
		return errors.Wrapf(err, "failed to ensure launch alias")
	}
	log.Infof("ensured launch alias")
	return nil
}

func Upgrade() error {
	log.Info("upgrading android-studio")
	if err := brew.Upgrade(BrewPkgStudio); err != nil {
		return errors.Wrap(err, "failed to install android-studio")
	}
	log.Info("upgraded android-studio")
	log.Info("upgrading android-sdk")
	if err := brew.Upgrade(BrewPkgSdk); err != nil {
		return errors.Wrap(err, "failed to install android-sdk")
	}
	log.Info("upgraded android-sdk")
	return nil
}

func setupLaunchAlias() error {
	binDir, err := tool.GetBinDir()
	if err != nil {
		return errors.Wrapf(err, "failed to get bin dir")
	}
	if err := os.WriteFile(filepath.Join(binDir, LaunchAlias), []byte(LaunchScript), os.ModePerm); err != nil {
		return errors.Wrapf(err, "failed to write %s file", filepath.Join(binDir, LaunchScript))
	}
	return nil
}
