package config

import "fmt"

type DBConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Database  string `yaml:"database"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Migration bool   `yaml:"migration"`
}

func (conf *DBConfig) ConnectionString() string { // DSN
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		conf.Host, conf.User, conf.Password, conf.Database, conf.Port,
	)
}
