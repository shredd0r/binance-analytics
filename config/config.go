package config

import (
	"log/slog"

	"github.com/shredd0r/binance-analytics/utils"
)

type (
	Config struct {
		Logger      Logger      `yaml:"logger"`
		Credentials Credentials `yaml:"credentials"`
		Storage     Sqlite      `yaml:"storage"`
	}

	Sqlite struct {
		FilePath string `yaml:"file-path"`
	}

	Credentials struct {
		ApiKey    string `yaml:"api-key"`
		SecretKey string `yaml:"secret-key"`
	}

	Logger struct {
		Format string     `yaml:"format"`
		Level  slog.Level `yaml:"level"`
	}
)

const pathToConfig = "config.yml"

func Read() (*Config, error) {
	cfg := &Config{}

	utils.ReadYamlFile(pathToConfig, cfg)

	return cfg, nil
}
