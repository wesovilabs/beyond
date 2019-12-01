package helper

import (
	"github.com/wesovilabs/beyond/logger"
	"os"
)

func closeFile(f *os.File) {
	if err := f.Close(); err != nil {
		logger.Errorf("error while closing file: '%v'", err)
	}
}
