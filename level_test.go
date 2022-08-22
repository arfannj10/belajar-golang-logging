package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()

	logger.Trace("This is trace")
	logger.Debug("this is debug")
	logger.Info("this is info")
	logger.Warn("this is warn")
	logger.Error("this is Error")
}

func TestLoggingLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("This is trace")
	logger.Debug("this is debug")
	logger.Info("this is info")
	logger.Warn("this is warn")
	logger.Error("this is Error")
}