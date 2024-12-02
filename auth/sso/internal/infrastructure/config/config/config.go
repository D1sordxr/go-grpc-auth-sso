package config

import (
	db "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/infrastructure/db/config"
	logger "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/infrastructure/logger/config"
	grpc "github.com/D1sordxr/go-grpc-auth-sso/auth/sso/internal/presentation/grpc/config"
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
