package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	configFile := Config{}

	// Get the config file path
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	// Read the file on path
	byteValue, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	// Unmarshal bytes using the struct
	err = json.Unmarshal(byteValue, &configFile)
	if err != nil {
		return Config{}, err
	}

	return configFile, nil

}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName
	err := write(*c)
	if err != nil {
		return err
	}

	return nil

}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, configFileName), nil

}

func write(cfg Config) error {

	// Convert data to byte[]
	jsonData, err := json.MarshalIndent(cfg, "", "	")
	if err != nil {
		return err
	}

	// Get file path
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	// Write data to file
	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil

}
