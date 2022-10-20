package key

import (
	"fmt"
	"github.com/leftbin/go-util/pkg/file"
	lbnos "github.com/leftbin/go-util/pkg/os"
	"github.com/pkg/errors"
	"github.com/plantoncloud/mactl/internal/bundle/build/scm"
	"github.com/plantoncloud/mactl/internal/git/config"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

const (
	DefaultPvtKeyName       = "id_rsa"
	BrowserOpenDelaySeconds = 10
)

// Cre a new ssh-key and return path
func Cre(host, workspace string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrap(err, "failed to get user home dir")
	}
	sshKeyPath := filepath.Join(homeDir, "scm", host)
	if workspace != "" {
		sshKeyPath = filepath.Join(homeDir, "scm", host, workspace)
	}
	pvtKeyFile := filepath.Join(sshKeyPath, DefaultPvtKeyName)
	if file.IsFileExists(pvtKeyFile) {
		isReplaceKey, err := isReplaceKey()
		if err != nil {
			return "", errors.Wrap(err, "failed to confirm replace key")
		}
		log.Debugf("replace key response recorded as %s", isReplaceKey)
		if isReplaceKey == "n" {
			return pvtKeyFile, nil
		}
	}
	if err := newKey(sshKeyPath, DefaultPvtKeyName); err != nil {
		return "", errors.Wrap(err, "failed to cre new ssh key")
	}
	return pvtKeyFile, nil
}

// Get returns the contents of the pub key file
func Get(host, workspace string) (string, error) {
	cfg, err := config.Get(host, workspace)
	if err != nil {
		return "", errors.Wrap(err, "failed to get config")
	}
	sshPubKeyFile := fmt.Sprintf("%s.pub", cfg.SshKeyPath)
	if !file.IsFileExists(sshPubKeyFile) {
		return "", errors.Errorf("pub key not found at %s", sshPubKeyFile)
	}
	pubKeyBytes, err := os.ReadFile(sshPubKeyFile)
	if err != nil {
		return "", errors.Wrap(err, "failed to read pub key contents")
	}
	return string(pubKeyBytes), nil
}

// Use copies the private of the selected host and workspace to the default scm host key
func Use(host, workspace string) error {
	cfg, err := config.Get(host, workspace)
	if err != nil {
		return errors.Wrap(err, "failed to get config")
	}
	sshKeyBytes, err := os.ReadFile(string(cfg.SshKeyPath))
	if err != nil {
		return errors.Wrapf(err, "failed to read %s file", cfg.SshKeyPath)
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get home dir")
	}
	scmSshDir := filepath.Join(homeDir, ".ssh", "scm")
	if err := os.MkdirAll(scmSshDir, 0744); err != nil {
		return errors.Wrapf(err, "failed to ensure %s dir", scmSshDir)
	}
	defaultScmSshKeyPath := filepath.Join(scmSshDir, host)
	if file.IsFileExists(defaultScmSshKeyPath) {
		if err := os.Remove(defaultScmSshKeyPath); err != nil {
			return errors.Wrapf(err, "failed to remove %s file", defaultScmSshKeyPath)
		}
	}
	if err := os.WriteFile(defaultScmSshKeyPath, sshKeyBytes, 0400); err != nil {
		return errors.Wrapf(err, "failed to write %s file", defaultScmSshKeyPath)
	}
	return nil
}

func HandlePubKeyUpdOnScm(host string) {
	fmt.Println("\n\n*** note: public key copied to clipboard. use cmd+v to paste ***")
	var url string
	switch scm.Provider(host) {
	case scm.ProviderGitlab:
		url = scm.ProviderGitlabSshKeyConfigurePageUrl
		break
	case scm.ProviderGitHub:
		url = scm.ProviderGitGithubSshKeyConfigurePageUrl
		break
	}
	if url == "" {
		return
	}
	fmt.Printf("\n\n*** note: in few seconds %s url will open in the browser. paste the copied ssh key ***\n\n", url)
	time.Sleep(BrowserOpenDelaySeconds * time.Second)
	lbnos.OpenBrowser(url, false)
}

func isReplaceKey() (string, error) {
	var choice string
	fmt.Print("\nkey already exists. replace? (y/n): ")
	if _, err := fmt.Scan(&choice); err != nil {
		return "", errors.Wrapf(err, "failed to read confirmation input")
	}
	return choice, nil
}
