package belajar_golang_logging

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

type SampleHook struct {
}

func (s *SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel} // level apa aja yg mau dieksekusi
}

func (s *SampleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("sample hook", entry.Level, entry.Message) // mau eksekusi hook kemana aja email atau tele
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(&SampleHook{})

	logger.Info("Hello Info")
	logger.Error("Hello Error")
	logger.Warn("Hello Warn")


}