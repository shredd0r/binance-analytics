package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		Credentials Credentials `yaml:"credentials"`
	}

	Sqlite struct {
		FilePath string `yaml:"file-path"`
	}

	Credentials struct {
		ApiKey    string `yaml:"api-key"`
		SecretKey string `yaml:"secret-key"`
	}
)

func Read() (*Config, error) {
	cfg := &Config{}
	bytesConfig, err := os.ReadFile("config.yml")

	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(bytesConfig, cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
