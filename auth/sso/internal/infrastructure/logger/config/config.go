package config

type LoggerConfig struct {
	Mode      string `yaml:"mode"`
	LogLevel  string `yaml:"log_level"`
	LogOutput string `yaml:"log_output"`
}
