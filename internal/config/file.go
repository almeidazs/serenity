package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/almeidazs/gowther/internal"
)

var configFile = "gowther.json"

func GetPath() (path string, err error) {
	path, err = os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error to get user working directory: %w", err)
	}

	return path, nil
}

func CreateConfigFile(path string) error {
	_, err := os.Create(filepath.Join(path, configFile))
	if err != nil {
		return fmt.Errorf("error creating config file: %w", err)
	}
	return nil
}

func CheckHasConfigFile(path string) (bool, error) {
	configFile := filepath.Join(path, configFile)
	if _, err := os.Stat(configFile); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		// TODO: criar write (escrever dentro do arquivo o JSON)

		return false, fmt.Errorf("error to get file info: %w", err)
	}

	return true, nil
}

func ReadConfigFile(path string) (*internal.Config, error) {
	configPath := filepath.Join(path, configFile)
	fileBytes, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("config file '%s' not found", configPath)
		}
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config internal.Config

	if err := json.Unmarshal(fileBytes, &config); err != nil {
		return nil, fmt.Errorf("error parsing config file json: %w", err)
	}

	return &config, nil
}
