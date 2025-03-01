package config

import (
	"github.com/ghodss/yaml"
	"github.com/leftbin/go-util/pkg/file"
	"github.com/pkg/errors"
	"github.com/plantoncloud/mactl/internal/defaulteditor"
	"github.com/plantoncloud/mactl/internal/keyboard/keys"
	_default "github.com/plantoncloud/mactl/internal/keyboard/shortcut/openapp/default"
	"github.com/plantoncloud/mactl/internal/lib/shell"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	FileName = "shortcuts.yaml"
)

// Edit opens config file in vs-code and waits until the file is closed.
func Edit() error {
	configFile, err := getShortcutConfigFilePath()
	if err != nil {
		return errors.Wrapf(err, "failed to get shortcut config file path")
	}
	if !file.IsFileExists(configFile) {
		if err := Write(_default.DefaultShortcuts); err != nil {
			return errors.Wrapf(err, "failed to initialize shortcut config file")
		}
	}
	for {
		duplicates := make([]string, 0)
		if err := shell.RunCmd(exec.Command(defaulteditor.DefaultEditor, "--wait", configFile)); err != nil {
			return errors.Wrapf(err, "failed to run command to open cache loc %s in vs code. is vs code installed?", configFile)
		}
		appShortcuts, err := List()
		if err != nil {
			return errors.Wrapf(err, "failed to list shortcuts")
		}
		duplicates = getDuplicates(appShortcuts)
		if len(duplicates) == 0 {
			break
		} else {
			log.Errorf("fix %v keys which may have either empty or duplicate shortcut mappings", duplicates)
		}
	}
	return nil
}

// getDuplicates iterates through the shortcuts and considers shortcuts that have empty app name to be a duplicate entry.
// note: the value could also be left empty intentionally.
func getDuplicates(shortcuts []*keys.AppShortcut) []string {
	duplicates := make([]string, 0)
	for _, s := range shortcuts {
		if s.AppName != "" {
			continue
		}
		duplicates = append(duplicates, string(s.Key))
	}
	return duplicates
}

// Write writes provided shortcuts to shortcuts config file
func Write(shortcuts []*keys.AppShortcut) error {
	shortcutConfigFilePath, err := getShortcutConfigFilePath()
	if err != nil {
		return errors.Wrapf(err, "failed to get config file path")
	}
	if !file.IsDirExists(filepath.Dir(shortcutConfigFilePath)) {
		if err := os.Mkdir(filepath.Dir(shortcutConfigFilePath), os.ModePerm); err != nil {
			return errors.Wrapf(err, "failed to create %s dir", filepath.Dir(shortcutConfigFilePath))
		}
	}
	defaultShortcutsYamlBytes, err := yaml.Marshal(shortcuts)
	if err != nil {
		return errors.Wrapf(err, "failed to initialize")
	}
	if err := os.WriteFile(shortcutConfigFilePath, defaultShortcutsYamlBytes, os.ModePerm); err != nil {
		return errors.Wrapf(err, "failed to write %s file", shortcutConfigFilePath)
	}
	return nil
}

// List returns list of app shortcuts from the shortcuts config file
func List() ([]*keys.AppShortcut, error) {
	appShortcuts := make([]*keys.AppShortcut, 0)
	shortcutConfigFile, err := getShortcutConfigFilePath()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get extra file path")
	}
	shortcutsFileContent, err := os.ReadFile(shortcutConfigFile)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %s file", shortcutConfigFile)
	}
	if err := yaml.Unmarshal(shortcutsFileContent, &appShortcuts); err != nil {
		return nil, errors.Wrapf(err, "failed to yaml unmarshal app shortcuts config file")
	}
	return appShortcuts, nil
}

// getShortcutConfigFilePath returns the location of the keyboard shortcut configs
func getShortcutConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrapf(err, "failed to get home dir")
	}
	return filepath.Join(homeDir, ".mactl", "keyboard", FileName), nil
}
