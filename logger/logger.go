package logger

import (
	"fmt"
	"github.com/google/logger"
	"io/ioutil"
)

var (
	beyondLog = logger.Init("BeyondErrorLogger", true, false, ioutil.Discard)
	debugLog  = logger.Init("BeyondDebugLogger", false, false, ioutil.Discard)
)

// Enable enabling the logger
func Enable() {
	debugLog = logger.Init("BeyondLogger", true, false, ioutil.Discard)
}

// Close close loggers
func Close() {
	debugLog.Close()
	beyondLog.Close()
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
	beyondLog.ErrorDepth(1, text)
}

// Errorf log error messages
func Errorf(text string, args ...interface{}) {
	beyondLog.ErrorDepth(1, fmt.Sprintf(text, args...))
}
