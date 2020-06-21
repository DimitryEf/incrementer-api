package config

import (
	"github.com/sirupsen/logrus"
	"os"
)

// Logger - структура с логером
type Logger struct {
	Log *logrus.Logger
}

// Конструктор логера, по умолчанию пишет логи в стандартный вывод
func NewLogger() *Logger {
	log := logrus.New()
	log.SetOutput(os.Stdout)

	return &Logger{
		Log: log,
	}
}
