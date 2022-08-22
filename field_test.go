package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

//Singular
func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "arfan").Info("Hello loging")
	logger.WithField("username", "empang").
		WithField("name", "Arfan H").
		Info("Hello loging")

}


//Plural
func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username" : "arfan",
		"name" : "Arfan H",
	}).Info("Hello loging")

}