package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/pkg/errors"
)

type (
	Config struct {
		Credentials Credentials `envPrefix:"CREDENTIALS_"`
	}

	Credentials struct {
		ApiKey    string `env:"API_KEY"`
		SecretKey string `env:"SECRET_KEY"`
	}
)

func Read() (*Config, error) {
	var cfg *Config
	if err := env.Parse(cfg); err != nil {
		return nil, errors.Wrap(err, "error parsing config from env")
	}
	return cfg, nil
}
