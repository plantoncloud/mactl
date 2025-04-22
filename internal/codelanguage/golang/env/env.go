package env

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

const (
	GoPathEnvVarKey = "GOPATH"
	GoPathDirName   = "gopa"
)

// GetGoPathLoc returns the complete path of GOPATH
func GetGoPathLoc() (string, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrap(err, "failed to get user home directory")
	}
	return fmt.Sprintf("%s/%s", userHomeDir, GoPathDirName), nil
}
