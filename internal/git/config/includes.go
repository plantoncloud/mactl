package config

import (
	"github.com/leftbin/go-util/pkg/file"
	"github.com/pkg/errors"
	"os"
	"strings"
)

var ErrGitConfigNotFoundForIncludeDir = errors.New("git config include not found for include dir")

func addIncludes(includesDir, path string) error {
	includes, err := getIncludes()
	if err != nil {
		return errors.Wrap(err, "failed to get includes")
	}
	includes[includesDir] = path
	if err := updGlobalConfigIncludes(includes); err != nil {
		return errors.Wrap(err, "failed to update global includes")
	}
	return nil
}

func delIncludes(includesDir string) error {
	includes, err := getIncludes()
	if err != nil {
		return errors.Wrap(err, "failed to get includes")
	}
	delete(includes, includesDir)
	if err := updGlobalConfigIncludes(includes); err != nil {
		return errors.Wrap(err, "failed to update global includes")
	}
	return nil
}

func getIncludes() (map[string]string, error) {
	globalConfig, err := parseGlobalConfig()
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse global git config")
	}
	includesMap := make(map[string]string, 0)
	const IncludeIfPrefix = "includeif.gitdir:"
	for k, v := range globalConfig {
		if !strings.HasPrefix(k, IncludeIfPrefix) {
			continue
		}
		dir := strings.TrimPrefix(k, IncludeIfPrefix)
		dir = strings.TrimSuffix(dir, "/.path")
		includesMap[dir] = v
	}
	return includesMap, nil
}

func getGitConfigForIncludeDir(includeDir string) (string, error) {
	includes, err := getIncludes()
	if err != nil {
		return "", errors.Wrap(err, "failed to get global git includes")
	}
	if includes[includeDir] == "" {
		return "", ErrGitConfigNotFoundForIncludeDir
	}
	return includes[includeDir], nil
}

func updGlobalConfigIncludes(includes map[string]string) error {
	renderedBytes, err := file.RenderTmplt(includes, GlobalGitConfigTemplate)
	if err != nil {
		return errors.Wrap(err, "failed to render global git config")
	}
	gitConfigPath, err := getGlobalConfigPath()
	if err != nil {
		return errors.Wrap(err, "failed to get global git config path")
	}
	if err := os.WriteFile(gitConfigPath, renderedBytes, 0644); err != nil {
		return errors.Wrapf(err, "failed to write %s file", gitConfigPath)
	}
	return nil
}
