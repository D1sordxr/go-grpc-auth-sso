package zerolog

import (
	"github.com/D1sordxr/aviasales/src/internal/logger/config"
	"github.com/rs/zerolog"
	"os"
	"time"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

type Logger struct {
	zerolog.Logger
}

func NewLogger(cfg config.LoggerConfig) *Logger {
	var logger zerolog.Logger

	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().UTC()
	}

	switch cfg.Mode {
	case envLocal:
		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
		logger = zerolog.New(output).With().Timestamp().Logger().Level(zerolog.DebugLevel)
	case envDev:
		logger = zerolog.New(os.Stdout).With().Timestamp().Logger().Level(zerolog.DebugLevel)
	case envProd:
		logger = zerolog.New(os.Stdout).With().Timestamp().Logger().Level(zerolog.InfoLevel)
	default:
		logger = zerolog.New(os.Stdout).With().Timestamp().Logger().Level(zerolog.DebugLevel)
	}

	return &Logger{Logger: logger}
}
