package zshrc

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
)

func Get() (string, error) {
	var zb strings.Builder
	zb.WriteString("\n\n# java env vars begin\n\n")
	envVars, err := getEnvVars()
	if err != nil {
		return "", errors.Wrapf(err, "failed to get env vars")
	}
	zb.WriteString(envVars)
	zb.WriteString("\n\n# java env vars end\n\n")
	return zb.String(), nil
}

func getEnvVars() (string, error) {
	var vb strings.Builder
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrap(err, "failed to get home dir")
	}
	vb.WriteString(fmt.Sprintf("export PATH=%s/.jenv/shims:${PATH}", homeDir))
	vb.WriteString(`
export JENV_SHELL=zsh
export JENV_LOADED=1`)
	return vb.String(), nil
}
