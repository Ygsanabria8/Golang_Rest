package utils

import (
	"log"
	models "modules/src/models"
	"os"

	"gopkg.in/yaml.v2"
)

const configPath = "configuration/local.config.yaml"

var Config *models.Config

func LoadEnv() {
	f, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&Config)

	if err != nil {
		log.Fatal(err.Error())
	}
}
