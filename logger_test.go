package logger_test

import (
	"testing"

	lg "github.com/Jourloy/GoLogger"
)

func TestAll(t *testing.T) {
	lg.New("all", "Test Logger")
	lg.Log("Test log")
	lg.Warn("Test warning")
	lg.Info("Test info")
	lg.Debug("Test debug")
	lg.Error("Test error")
}

func TestProd(t *testing.T) {
	lg.New("prod", "Test Logger")
	lg.Log("Test log")
	lg.Warn("Test warning")
	lg.Info("Test info")
	lg.Debug("Test debug")
	lg.Error("Test error")
}
