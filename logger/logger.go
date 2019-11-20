package logger

import (
	"fmt"
	"github.com/google/logger"
	"io/ioutil"
	"log"
)

var (
	goaLog   = logger.Init("GoaErrorLogger", true, false, ioutil.Discard)
	debugLog = logger.Init("GoaDebugLogger", false, false, ioutil.Discard)
)

// Enable enabling the logger
func Enable() {
	debugLog = logger.Init("GoaLogger", true, false, ioutil.Discard)
}

// Close close loggers
func Close() {
	debugLog.Close()
	goaLog.Close()
}

// Infof the messages
func Infof(text string, args ...interface{}) {
	debugLog.InfoDepth(0, fmt.Sprintf(text, args...))
}

// Info the messages
func Info(text string) {
	debugLog.InfoDepth(1, text)
}

// Error log error messages
func Error(text string) {
	goaLog.ErrorDepth(1, text)
}

// Errorf log error messages
func Errorf(text string, args ...interface{}) {
	goaLog.ErrorDepth(1, fmt.Sprintf(text, args...))
}

// Fatal log fatal messages
func Fatal(text string, args ...interface{}) {
	log.Fatalf(text, args...)
}
