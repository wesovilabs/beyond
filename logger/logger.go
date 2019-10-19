package logger

import (
	"github.com/google/logger"
	"io/ioutil"
)

var (
	log = logger.Init("GoaLogger", false, false, ioutil.Discard)
)

// Enable enabling the logger
func Enable() {
	log = logger.Init("GoaLogger", true, false, ioutil.Discard)
	log.Info("log is enabled")
}

// Close close logger
func Close() {
	log.Close()
}

// Infof the messages
func Infof(text string, args ...interface{}) {
	log.Infof(text, args...)
}

// Info the messages
func Info(text string) {
	log.Info(text)
}

// Error log error messages
func Error(text string) {
	log.Error(text)
}

// Errorf log error messages
func Errorf(text string, args ...interface{}) {
	log.Errorf(text, args...)
}

// Fatal log fatal messages
func Fatal(text string, args ...interface{}) {
	log.Fatalf(text, args...)
}
