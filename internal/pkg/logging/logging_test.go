package logging

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	err := os.Setenv("LOG_LEVEL", "debug")
	if err != nil {
		t.Error(err)
	}

	logger := newLogger()

	if logger.Level != logrus.DebugLevel {
		t.Errorf("wrong logger level: %d != %d", logger.Level, logrus.DebugLevel)
	}
}

func TestNewWrongLevel(t *testing.T) {
	os.Setenv("LOG_LEVEL", "foobar")

	logger := newLogger()

	if logger.Level != defaultLevel {
		t.Error("wrong logger default level")
	}
}

func TestNewWithoutLevel(t *testing.T) {
	os.Unsetenv("LOG_LEVEL")

	logger := newLogger()

	if logger.Level != defaultLevel {
		t.Error("wrong logger default level")
	}
}

func TestLoggerOnce(t *testing.T) {
	logger1 := New()
	logger2 := New()

	if logger1 != logger2 {
		t.Error("different loggers")
	}
}
