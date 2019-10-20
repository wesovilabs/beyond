package logger

import (
	"fmt"
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
func Infof(text string,args ...interface{}) {
	log.InfoDepth(1,fmt.Sprintf(text,args...))
}

// Info the messages
func Info(text string) {
	log.InfoDepth(1, text)
}

// Error log error messages
func Error(text string) {
	log.ErrorDepth(1, text)
}

// Errorf log error messages
func Errorf(text string, args ...interface{}) {
	log.Errorf(text, args...)
}

// Fatal log fatal messages
func Fatal(text string, args ...interface{}) {
	log.Fatalf(text, args...)
}
