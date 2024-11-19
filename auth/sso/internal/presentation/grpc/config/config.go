package config

import "time"

type GRPCConfig struct {
	Port     int           `yaml:"port"`
	Timeout  time.Duration `yaml:"timeout"`
	TokenTTL time.Duration `yaml:"token_ttl"`
}
