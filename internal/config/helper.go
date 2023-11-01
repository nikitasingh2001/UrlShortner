package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func loadConfigFromFile(basePath string, filename string, env string) {
	path := getFilePath(basePath, filename, env)
	_, err := os.Stat(path)
	if err != nil {
		logrus.WithError(err).Error(configFileNotFound)
	}
	content := readConfigFile(path)
	content = os.ExpandEnv(string(content))
	if _, err := toml.Decode(content, &config); err != nil {
		logrus.WithError(err).Panic(invalidConfigType)
	}
}

func readConfigFile(path string) string {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		logrus.WithError(err).Error(configFileNotFound)
		return ""
	}

	return string(data)
}

func getFilePath(basePath string, fileName string, env string) string {
	if env != "" {
		fileName = fmt.Sprintf(fileName, env)
	}

	path := fmt.Sprintf(FilePath, basePath, fileName)

	return path
}
