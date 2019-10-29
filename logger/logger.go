// +build !test

package logger

import (
	"fmt"
	"github.com/google/logger"
	"io/ioutil"
)

var (
	log      = logger.Init("GoaErrorLogger", true, false, ioutil.Discard)
	debugLog = logger.Init("GoaDebugLogger", false, false, ioutil.Discard)
)

// Enable enabling the logger
func Enable() {
	debugLog = logger.Init("GoaLogger", true, false, ioutil.Discard)
}

// Close close loggers
func Close() {
	debugLog.Close()
	log.Close()
}

// Infof the messages
func Infof(text string, args ...interface{}) {
	debugLog.InfoDepth(1, fmt.Sprintf(text, args...))
}

// Info the messages
func Info(text string) {
	debugLog.InfoDepth(1, text)
}

// Error log error messages
func Error(text string) {
	log.ErrorDepth(1, text)
}

// Errorf log error messages
func Errorf(text string, args ...interface{}) {
	log.ErrorDepth(1, fmt.Sprintf(text, args...))
}

// Fatal log fatal messages
func Fatal(text string, args ...interface{}) {
	log.Fatalf(text, args...)
}
