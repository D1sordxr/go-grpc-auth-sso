package config

import (
	db "src/internal/db/config"
	api "src/internal/http/config"
	log "src/internal/logger/config"
)

type Config struct {
	AppConfig        `yaml:"app"`
	log.LoggerConfig `yaml:"logger"`
	db.DBConfig      `yaml:"db"`
	api.APIConfig    `yaml:"http_server"`
}

type AppConfig struct {
	Mode string `yaml:"mode" env-default:"local"`
}
