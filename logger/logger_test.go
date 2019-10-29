package logger

import "testing"

func TestLogger(t *testing.T) {
	Enable()
	defer Close()
	Info("Info message")
	Infof("Info message %s", "hey")
	Error("Error message")
	Errorf("Error message %s", "hey")
}
