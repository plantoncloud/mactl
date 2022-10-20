package override

import (
	"github.com/leftbin/go-util/pkg/file"
	"github.com/pkg/errors"
	"github.com/plantoncloud/mactl/internal/lib/shell"
	"os"
	"os/exec"
	"path/filepath"
)

const FileName = ".zshrc.override"

var ErrNotFound = errors.New("override file not found")

func Get() ([]byte, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get home dir")
	}
	f := filepath.Join(homeDir, FileName)
	if !file.IsFileExists(f) {
		return nil, nil
	}
	fileContent, err := os.ReadFile(f)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read file contents")
	}
	return fileContent, nil
}

func Show() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get home dir")
	}
	f := filepath.Join(homeDir, FileName)
	if !file.IsFileExists(f) {
		return ErrNotFound
	}
	if err := shell.RunCmd(exec.Command("cat", f)); err != nil {
		return errors.Wrapf(err, "failed to run command to open %s in vs code. is vs code installed?", f)
	}
	return nil
}

func Edit() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return errors.Wrap(err, "failed to get home dir")
	}
	f := filepath.Join(homeDir, FileName)
	if err := shell.RunCmd(exec.Command("code", f)); err != nil {
		return errors.Wrapf(err, "failed to run command to open %s in vs code. is vs code installed?", f)
	}
	return nil
}
