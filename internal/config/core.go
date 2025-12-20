package config

import (
	"encoding/json"
	"io"
	"os"
	"strings"
)

func (c *Config) SetUser(username string) error {
	c.Username = username
	err := c.write()
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) write() error {
	var fp *os.File
	var configPath string

	raw, err := json.Marshal(c)
	if err != nil {
		return err
	}

	configPath, err = getConfigFilePath()
	if err != nil {
		return err
	}

	fp, err = os.OpenFile(configPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write(raw)
	if err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return strings.Join([]string{homePath, configFileName}, "/"), nil
}

func Read() (Config, error) {
	var fp *os.File
	var raw []byte
	var config Config

	configPath, err := getConfigFilePath()
	if err != nil {
		return config, err
	}

	fp, err = os.Open(configPath)
	if err != nil {
		return Config{}, err
	}
	defer fp.Close()

	raw, err = io.ReadAll(fp)
	if err != nil {
		return Config{}, err
	}

	err = json.Unmarshal(raw, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
