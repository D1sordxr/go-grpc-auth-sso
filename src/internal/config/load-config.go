package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

const BasicConfigPath = "./src/configs/app/local.yaml"

func NewConfig() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig(BasicConfigPath, &cfg); err != nil {
		log.Fatalf("failed to read config: %v", err.Error())
	}

	return &cfg, nil
}
