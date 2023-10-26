package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

type AuthConfig struct {
	Auth struct {
		Username string
		Password string
	}
}

func ReadAuthConfig(configPath string) (*AuthConfig, error) {
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config AuthConfig
	if _, err := toml.Decode(string(content), &config); err != nil {
		fmt.Println("Error decoding config content:", err)
		return nil, err
	}

	return &config, nil
}
