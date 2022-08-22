package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello logging")
	logger.WithField("username", "arfan").Info("hello loging")

	entry := logrus.NewEntry(logger)
	entry.WithField("username", "khannedy")
	entry.Info("Hello Entry")
}