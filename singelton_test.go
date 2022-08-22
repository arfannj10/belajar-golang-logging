package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingleton(t *testing.T) {
	logrus.Info("Hello info")
	logrus.Warn("Hello Warn")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Hello info")
	logrus.Warn("Hello Warn")

}