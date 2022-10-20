package config

import (
	"customer/helper"
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
