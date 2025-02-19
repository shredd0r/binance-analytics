package utils

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ReadYamlFile(pathToFile string, out interface{}) error {
	fileBytes, err := os.ReadFile(pathToFile)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(fileBytes, out)

	if err != nil {
		return err
	}

	return nil
}
