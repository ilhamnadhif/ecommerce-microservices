package config

import (
	"auth/helper"
	"os"

	"gopkg.in/yaml.v3"
)

var Config Configuration

func InitConfig() {
	content, err := os.ReadFile("./config.yaml")
	helper.LogFatalIfError(err)
	err = yaml.Unmarshal(content, &Config)
	helper.LogFatalIfError(err)
}
