package config

import "time"

type TokenConfig struct {
	Key      string        `yaml:"key" env-required:"true"`
	TokenTTL time.Duration `yaml:"token_ttl" env-required:"true"`
}
