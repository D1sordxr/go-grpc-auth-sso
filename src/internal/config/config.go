package config

import (
	db "aviasales/src/internal/db/config"
	api "aviasales/src/internal/http/config"
)

type Config struct {
	AppConfig
	db.DBConfig
	api.APIConfig
}

type AppConfig struct {
	Mode string
}
