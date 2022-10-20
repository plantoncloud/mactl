package extra

import (
	"bytes"
	"github.com/ghodss/yaml"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/leftbin/go-util/pkg/file"
	lbnutil "github.com/leftbin/go-util/pkg/table"
	"github.com/pkg/errors"
	"github.com/plantoncloud/mactl/internal/keyboard/keys"
	"github.com/plantoncloud/mactl/internal/lib/shell"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	FileName    = "shortcuts.extra.yaml"
	InitContent = `
#- appName: iterm
#   key: x
#   appFilePath: /Applications/Some App.app

`
)

// Edit opens extra file in vs-code
func Edit() error {
	extraFile, err := getExtraShortcutFilePath()
	if err != nil {
		return errors.Wrapf(err, "failed to get extra file path")
	}
	if !file.IsFileExists(extraFile) {
		if err := Init(); err != nil {
			return errors.Wrapf(err, "failed to initialize extra")
		}
	}
	if err := shell.RunCmd(exec.Command("code", extraFile)); err != nil {
		return errors.Wrapf(err, "failed to run command to open cache loc %s in vs code. is vs code installed?", extraFile)
	}
	return nil
}

// Del deletes extra file in vs-code
func Del() error {
	extraFile, err := getExtraShortcutFilePath()
	if err != nil {
		return errors.Wrapf(err, "failed to get extra file path")
	}
	if !file.IsFileExists(extraFile) {
		return nil
	}
	if err := os.Remove(extraFile); err != nil {
		return errors.Wrapf(err, "failed to delete %s file", extraFile)
	}
	return nil
}

// Init creates empty extra shortcuts file
func Init() error {
	extraFile, err := getExtraShortcutFilePath()
	if err != nil {
		return errors.Wrapf(err, "failed to get extra file path")
	}
	if !file.IsDirExists(filepath.Dir(extraFile)) {
		if err := os.Mkdir(filepath.Dir(extraFile), os.ModePerm); err != nil {
			return errors.Wrapf(err, "failed to create %s dir", filepath.Dir(extraFile))
		}
	}
	if err := os.WriteFile(extraFile, []byte(InitContent), os.ModePerm); err != nil {
		return errors.Wrapf(err, "failed to write %s file", extraFile)
	}
	return nil
}

func List() ([]*keys.AppShortcut, error) {
	es := make([]*keys.AppShortcut, 0)
	resp := make([]*keys.AppShortcut, 0)
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get home dir")
	}
	extraFile := filepath.Join(homeDir, ".mactl", "keyboard", FileName)
	if !file.IsFileExists(extraFile) {
		return resp, nil
	}
	extraShortCutsFileContent, err := os.ReadFile(extraFile)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read %s file", extraFile)
	}
	if bytes.Equal(extraShortCutsFileContent, []byte(InitContent)) {
		return resp, nil
	}
	if err := yaml.Unmarshal(extraShortCutsFileContent, &es); err != nil {
		return nil, errors.Wrapf(err, "failed to yaml unmarshal extra shortcuts file")
	}
	for _, b := range es {
		resp = append(resp, b)
	}
	return resp, nil
}

func PrintList(shortcuts []*keys.AppShortcut) {
	header := table.Row{"key", "app name", "app path"}
	rows := make([]table.Row, 0)
	for _, r := range shortcuts {
		rows = append(rows, table.Row{r.Key, r.AppName, r.AppFilePath})
	}
	lbnutil.PrintTable(header, rows)
}

func getExtraShortcutFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrapf(err, "failed to get home dir")
	}
	return filepath.Join(homeDir, ".mactl", "keyboard", FileName), nil
}
