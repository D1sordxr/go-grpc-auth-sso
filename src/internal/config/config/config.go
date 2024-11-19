package config

import (
	db "github.com/D1sordxr/aviasales/src/internal/db/config"
	log "github.com/D1sordxr/aviasales/src/internal/logger/config"
	api "github.com/D1sordxr/aviasales/src/internal/presentation/api/config"
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
