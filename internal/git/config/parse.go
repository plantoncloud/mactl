package config

import (
	"github.com/leftbin/go-util/pkg/file"
	"github.com/muja/goconfig"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"strings"
)

func parseConfig(gitConfigPath string) (*Config, error) {
	bytes, err := ioutil.ReadFile(gitConfigPath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %s file", gitConfigPath)
	}
	configContent, _, err := goconfig.Parse(bytes)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse git config from %s file", gitConfigPath)
	}
	sshKeyPath, err := parseSshKeyPath(configContent["core.sshcommand"])
	if err != nil {
		if err == ErrSshKeyPathNotFound {
			return &Config{
				Username:   configContent["user.name"],
				Email:      configContent["user.email"],
				SshKeyPath: "not set",
			}, nil
		}
		return nil, errors.Wrap(err, "failed to get ssh key path")
	}
	return &Config{
		Username:   configContent["user.name"],
		Email:      configContent["user.email"],
		SshKeyPath: sshKeyPath,
	}, nil
}

func parseGlobalConfig() (map[string]string, error) {
	gitConfigPath, err := getGlobalConfigPath()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get global git config path")
	}
	if !file.IsFileExists(gitConfigPath) {
		return make(map[string]string, 0), nil
	}
	bytes, err := ioutil.ReadFile(gitConfigPath)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %s file", gitConfigPath)
	}
	configContent, _, err := goconfig.Parse(bytes)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse git config from %s file", gitConfigPath)
	}
	return configContent, nil
}

func parseSshKeyPath(sshCmd string) (FilePath, error) {
	if sshCmd == "" {
		return "", ErrSshKeyPathNotFound
	}
	parts := strings.Split(sshCmd, " ")
	if len(parts) != 3 {
		return "", ErrSshKeyPathNotFound
	}
	sshKeyPath := parts[2]
	if strings.HasPrefix(sshKeyPath, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", errors.Wrap(err, "failed to get home dir")
		}
		sshKeyPath = strings.ReplaceAll(sshKeyPath, "~", homeDir)
	}
	return FilePath(sshKeyPath), nil
}
