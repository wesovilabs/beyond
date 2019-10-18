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

// Log return an instance of logger
func Log() *logger.Logger {
	return log
}
