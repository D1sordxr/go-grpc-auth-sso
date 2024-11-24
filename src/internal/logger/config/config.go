package config

type LoggerConfig struct {
	Mode      string `yaml:"mode"`
	LogOutput string `yaml:"log_output"`
	LogLevel  string `yaml:"log_level"`
}
