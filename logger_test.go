package GoLogger_test

import (
	"strconv"
	"testing"

	"github.com/Jourloy/GoLogger"
)

func print(l GoLogger.Logger, i int, t *testing.T) {
	if i == 1 {
		if b := l.Error(""); b == false {
			t.Error("Except 'Error'")
		}
		if b := l.Debug(""); b == true {
			t.Error("Except false for 'Debug'")
		}
	} else if i == 2 {
		if b := l.Warn(""); b == false {
			t.Error("Except 'Warn'")
		}
		if b := l.Debug(""); b == true {
			t.Error("Except false for 'Debug'")
		}
	} else if i == 3 {
		if b := l.Log(""); b == false {
			t.Error("Except 'Log'")
		}
		if b := l.Debug(""); b == true {
			t.Error("Except false for 'Debug'")
		}
	} else if i == 4 {
		if b := l.Info(""); b == false {
			t.Error("Except 'Info'")
		}
		if b := l.Debug(""); b == true {
			t.Error("Except false for 'Debug'")
		}
	} else if i == 5 {
		if b := l.Debug(""); b == false {
			t.Error("Except 'Debug'")
		}
	}
}

func TestProd(t *testing.T) {
	for i := 0; i <= 6; i++ {
		logger := GoLogger.New(i, "I: "+strconv.Itoa(i))
		print(logger, i, t)
	}
}
