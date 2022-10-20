package config

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	osint "github.com/plantoncloud/mactl/internal/lib/os"
	"github.com/plantoncloud/mactl/internal/lib/shell"
	"os/exec"
)

const (
	// PlistFileIntel todo: hardcoded version should be replaced
	PlistFileIntel   = "/usr/local/Cellar/dnsmasq/2.87/homebrew.mxcl.dnsmasq.plist"
	PlistFileSilicon = "/opt/homebrew/Cellar/dnsmasq/2.87/homebrew.mxcl.dnsmasq.plist"
	LaunchDaemonLoc  = "/Library/LaunchDaemons/homebrew.mxcl.dnsmasq.plist"
)

func Setup() error {
	if err := setupLaunchDaemon(); err != nil {
		return errors.Wrapf(err, "failed to write config file")
	}
	return nil
}

func setupLaunchDaemon() error {
	log.Infof("ensuring launch daemon")
	if err := shell.RunCmd(exec.Command("sudo", "cp", getPlistPath(), LaunchDaemonLoc)); err != nil {
		return errors.Wrapf(err, "failed to copy launch daemon file from %s to %s", getPlistPath(), LaunchDaemonLoc)
	}
	if err := shell.RunCmd(exec.Command("sudo", "launchctl", "load", LaunchDaemonLoc)); err != nil {
		return errors.Wrapf(err, "failed to copy launch daemon file from %s to %s", getPlistPath(), LaunchDaemonLoc)
	}
	log.Infof("ensured launch daemon")
	return nil
}

func getPlistPath() string {
	if osint.GetArch() == osint.ARM64 {
		return PlistFileSilicon
	}
	return PlistFileIntel
}
