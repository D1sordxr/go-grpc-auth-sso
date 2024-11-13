package config

import (
	"errors"
	"os"
)

func NewConfig() (*Config, error) {
	var cfg Config

	cfg.AppConfig.Mode = os.Getenv("APP_MODE")
	if cfg.AppConfig.Mode == "" {
		return nil, errors.New("APP_MODE is not set")
	}

	cfg.DBConfig.DSN = os.Getenv("DB_STRING")
	if cfg.DBConfig.DSN == "" {
		return nil, errors.New("DB_STRING is not set")

	}

	cfg.APIConfig.Host = os.Getenv("HOST")
	if cfg.APIConfig.Host == "" {
		cfg.APIConfig.Host = "localhost"
	}

	cfg.APIConfig.Port = os.Getenv("PORT")
	if cfg.APIConfig.Port == "" {
		return nil, errors.New("PORT is not set")
	}

	return &cfg, nil
}
