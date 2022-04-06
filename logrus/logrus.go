package logrus

import (
	"github.com/sirupsen/logrus"
	"go.clever-cloud.dev/app"
)

func SetupLogger(logger *logrus.Logger) {
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		DisableTimestamp: true,
	})

	if app.IsDebugEnabled() {
		logger.SetLevel(logrus.DebugLevel)
	}
}

func SetupStandardLogger() {
	SetupLogger(logrus.StandardLogger())
}
