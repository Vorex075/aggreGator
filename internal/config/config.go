package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Config struct {
	DbUrl       string `json:"db_url"`
	CurrentUser string `json:"current_user_name"`
	LastPost    struct {
		Publicated_at time.Time `json:"publicated_at"`
		Id            int       `json:"id"`
	} `json:"last_post"`
}

const configFileName = "/.gatorconfig.json"

// Read reads the config file and creates a Config instance.
func Read() (Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, fmt.Errorf("error while trying to read the config file: %v", err)
	}
	var conf Config
	err = json.Unmarshal(data, &conf)
	if err != nil {
		return Config{}, fmt.Errorf("error while trying to read json file: %v", err)
	}
	if conf.LastPost.Publicated_at.IsZero() {
		conf.LastPost.Publicated_at = time.Now()
	}
	return conf, nil
}

// SetUser modifies the struct to set the given username. This function also modifies the config file. If an error ocurrs, the user wont be modified.
func (c *Config) SetUser(user string) error {
	oldUser := c.CurrentUser
	c.CurrentUser = user
	err := Write(c)
	if err != nil {
		c.CurrentUser = oldUser
		return err
	}
	c.UpdateLastPost(time.Now(), 0)
	return nil
}

func Write(cfg *Config) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	jsonData, err := json.MarshalIndent(*cfg, "", "\t")
	if err != nil {
		return fmt.Errorf("error while trying to marshal the json config: %v", err)
	}
	err = os.WriteFile(configFilePath, jsonData, 700)
	if err != nil {
		return fmt.Errorf("error while trying to write the config into the file: %v", err)
	}
	return nil
}

func (c *Config) UpdateLastPost(pubTime time.Time, id int) {
	c.LastPost.Publicated_at = pubTime
	c.LastPost.Id = id
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error while trying to read home dir path: %v", err)
	}
	configFilePath := homePath + configFileName
	return configFilePath, nil
}
