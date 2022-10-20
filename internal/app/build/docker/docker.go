package docker

import (
	"github.com/leftbin/go-util/pkg/file"
	"github.com/pkg/errors"
	"github.com/plantoncloud/mactl/internal/installer/brew"
	"github.com/plantoncloud/mactl/internal/installer/macapp"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

const (
	BrewCaskPkg           = "docker"
	MacAppName            = "Docker.app"
	ConfigFilePermissions = 0644
	VmDiskSizeInGb        = 200
	VmMemoryInGb          = 10
	VmCpu                 = 4
)

type VmConfigTemplateInput struct {
	HomeDir      string
	Cpu          int
	MemoryInMb   int
	DiskSizeInMb int
}

func Setup() error {
	log.Infof("installing %s", BrewCaskPkg)
	if err := brew.InstallCask(BrewCaskPkg); err != nil {
		return errors.Wrapf(err, "failed to install %s pkg using brew", BrewCaskPkg)
	}
	log.Infof("installed %s", BrewCaskPkg)
	return nil
}

func Upgrade() error {
	log.Infof("upgrading %s", BrewCaskPkg)
	if err := brew.UpgradeCask(BrewCaskPkg); err != nil {
		return errors.Wrapf(err, "failed to upgrade %s pkg using brew", BrewCaskPkg)
	}
	log.Infof("upgraded %s", BrewCaskPkg)
	return nil
}

func Configure() error {
	log.Infof("configuring docker daemon")
	if err := configureDaemon(); err != nil {
		return errors.Wrap(err, "failed to configure docker daemon")
	}
	log.Infof("configured docker daemon")
	log.Infof("configuring docker vm")
	if err := configureVm(); err != nil {
		return errors.Wrap(err, "failed to configure docker vm")
	}
	log.Infof("configured docker vm")
	return nil
}

func configureDaemon() error {
	const FileLookupDelaySeconds = 5
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get home dir")
	}
	configFile := filepath.Join(homeDir, ".docker", "daemon.json")
	for {
		if file.IsFileExists(configFile) {
			break
		}
		log.Infof("%s file not found sleeping for %d seconds", configFile, FileLookupDelaySeconds)
		time.Sleep(FileLookupDelaySeconds * time.Second)
	}
	if err := os.WriteFile(configFile, []byte(DaemonConfig), ConfigFilePermissions); err != nil {
		return errors.Wrapf(err, "faiedl to write %s file", configFile)
	}
	return nil
}

func configureVm() error {
	const FileLookupDelaySeconds = 5
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get home dir")
	}
	renderedBytes, err := file.RenderTmplt(getVmConfigTemplateInput(homeDir), VmConfigTemplate)
	if err != nil {
		return errors.Wrap(err, "failed to render configure docker vm")
	}
	vmConfigFileLoc := filepath.Join(homeDir, "Library/Group Containers/group.com.docker/settings.json")
	for {
		if file.IsFileExists(vmConfigFileLoc) {
			break
		}
		log.Infof("%s file not found sleeping for %d seconds", vmConfigFileLoc, FileLookupDelaySeconds)
		time.Sleep(FileLookupDelaySeconds * time.Second)
	}
	log.Infof("creating %s docker vm config file", vmConfigFileLoc)
	if err := os.WriteFile(vmConfigFileLoc, renderedBytes, ConfigFilePermissions); err != nil {
		return errors.Wrapf(err, "failed to write %s file", vmConfigFileLoc)
	}
	log.Infof("created %s docker vm config file", vmConfigFileLoc)
	return nil
}

func getVmConfigTemplateInput(homeDir string) *VmConfigTemplateInput {
	return &VmConfigTemplateInput{
		HomeDir:      homeDir,
		Cpu:          VmCpu,
		MemoryInMb:   VmMemoryInGb * 1024,
		DiskSizeInMb: VmDiskSizeInGb * 1024,
	}
}

func Open() error {
	if err := macapp.Open(MacAppName); err != nil {
		return errors.Wrapf(err, "failed to open %s", MacAppName)
	}
	return nil
}
