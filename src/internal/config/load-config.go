package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"src/internal/config/config"
)

const BasicConfigPath = "./configs/app/local.yaml"

func NewConfig() (*config.Config, error) {
	var cfg config.Config

	if err := cleanenv.ReadConfig(BasicConfigPath, &cfg); err != nil {
		log.Fatalf("failed to read config: %v", err.Error())
	}

	return &cfg, nil
}
