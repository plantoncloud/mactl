package cache

import (
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

func GetLoc() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrap(err, "failed to get home dir")
	}
	cacheLoc := filepath.Join(homeDir, ".mactl", "cache")
	if err := os.MkdirAll(cacheLoc, os.ModePerm); err != nil {
		return "", errors.Wrapf(err, "failed to ensure %s dir", cacheLoc)
	}
	return cacheLoc, nil
}
