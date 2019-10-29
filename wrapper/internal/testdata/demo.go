package testdata

import (
	"encoding/json"
	"runtime"
	tt "time"
)

func test() (tt.Time, runtime.Error) {
	return tt.Now(), nil
}

func convert() *json.Encoder {
	return nil
}
