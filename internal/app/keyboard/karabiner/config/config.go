package config

import (
	"encoding/json"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

const (
	File = "karabiner.json"
)

var (
	CapsLockMofifierKeys = []string{"left_shift", "left_control", "left_option", "left_command"}
)

func Setup() error {
	configFile, err := getFilePath()
	if err != nil {
		return errors.Wrapf(err, "failed to get path of the config file")
	}
	if err := os.MkdirAll(filepath.Dir(configFile), 0644); err != nil {
		return errors.Wrapf(err, "failed to ensure %s dir", filepath.Dir(configFile))
	}
	c, err := GetDefault()
	if err != nil {
		return errors.Wrapf(err, "failed to get default config")
	}

	configBytes, err := json.Marshal(c)
	if err != nil {
		return errors.Wrapf(err, "failed to json marshal config")
	}
	if err := os.WriteFile(configFile, configBytes, 0644); err != nil {
		return errors.Wrapf(err, "failed to write %s file", configFile)
	}
	log.Infof("karabiner config file %s created", configFile)
	return nil
}

func GetDefault() (*Config, error) {
	c := new(Config)
	if err := json.Unmarshal([]byte(Default), c); err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal default config")
	}
	return c, nil
}

func (c *Config) GetDefaultProfile() (*Profile, error) {
	if len(c.Profiles) <= 0 {
		return nil, errors.New("no profile found")
	}
	return c.Profiles[0], nil
}

func getFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Wrap(err, "failed to get home dir")
	}
	return filepath.Join(homeDir, ".config", "karabiner", File), nil
}

func (c *Config) Save() error {
	karabinerCfgFile, err := getFilePath()
	if err != nil {
		return errors.Wrapf(err, "failed to get path of the config file")
	}
	cfgJson, err := json.Marshal(c)
	if err != nil {
		return errors.Wrapf(err, "failed to json unmarshal karabiner config")
	}
	if err := os.WriteFile(karabinerCfgFile, cfgJson, 0644); err != nil {
		return errors.Wrapf(err, "failed to write %s file", karabinerCfgFile)
	}
	return nil
}
