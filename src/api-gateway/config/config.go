package config

import (
	"api-gateway/helper"
	"gopkg.in/yaml.v3"
	"os"
)

var Config Configuration

func InitConfig() {
	content, err := os.ReadFile("./config.yaml")
	helper.LogFatalIfError(err)
	err = yaml.Unmarshal(content, &Config)
	helper.LogFatalIfError(err)
}
