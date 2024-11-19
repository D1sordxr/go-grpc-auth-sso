package config

import (
	db "sso/internal/infrastructure/db/config"
	logger "sso/internal/infrastructure/logger/config"
	grpc "sso/internal/presentation/grpc/config"
)

type Config struct {
	AppConfig           `yaml:"app"`
	logger.LoggerConfig `yaml:"logger"`
	db.DBConfig         `yaml:"db"`
	grpc.GRPCConfig     `yaml:"grpc"`
}

type AppConfig struct {
	Mode string `yaml:"mode" env-default:"local"`
}
