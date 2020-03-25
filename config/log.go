package config

import (
	"os"

	"github.com/imjuanleonard/palu-covid/pkg/logger"
)

var Log *logger.LogConfig

func initLogConfig() {
	Log = &logger.LogConfig{
		Level:  mustGetString("LOG_LEVEL"),
		Format: mustGetString("LOG_FORMAT"),
		Out:    os.Stderr,
	}
}
