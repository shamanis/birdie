package logging

import (
	"github.com/sirupsen/logrus"
	"os"
)

var (
	logger *logrus.Logger
)

func init() {
	logger = newLogger()
}

func newLogger() *logrus.Logger {
	log := logrus.New()
	log.Out = os.Stdout

	lvl, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		lvl = logrus.InfoLevel // Default log level
	}
	log.SetLevel(lvl)
	log.SetFormatter(&logrus.JSONFormatter{})
	return log
}

func New() *logrus.Logger {
	return logger
}
