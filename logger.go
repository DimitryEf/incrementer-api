package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Logger struct {
	Log *logrus.Logger
}

func NewLogger() *Logger {
	log := logrus.New()
	log.SetOutput(os.Stdout)

	return &Logger{
		Log: log,
	}
}
