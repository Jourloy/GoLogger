package GoLogger

import (
	"fmt"
	"time"

	cl "github.com/Jourloy/Colored"
)

type Logger struct {
	l int
	n string
}

// Create logger
//
// Levels
//
// 1 - Error
//
// 2 - Warn
//
// 3 - Log
//
// 4 - Info
//
// 5 - Debug
//
// Example
//
//     // For all levels
//     logger := GoLogger.New(5, "Logger")
//     // For only important levels
//     logger := GoLogger.New(3, "Logger")
func New(level int, name string) Logger {
	logger := Logger{}
	logger.l = level
	logger.n = name
	return logger
}

// You can increase or degrease level for your logger
//
// Example
//
//     // Increase level on 2
//     logger = logger.Level(2)
//     // Degrease level on 1
//     logger = logger.Level(-1)
func (l Logger) Level(level int) Logger {
	l.l += level
	if l.l < 1 {
		l.l = 1
	} else if l.l > 5 {
		l.l = 5
	}
	return l
}

// Important red messages when something go wrong
//
// Level: 1
//
// Example
//
//     logger.Error("User not found")
func (l Logger) Error(t string) bool {
	if l.l < 1 {
		return false
	}
	timeNow := time.Now()
	log := timeNow.Format(time.UnixDate) + "  " + l.n + " [ERROR] - " + t
	fmt.Println(cl.ChangeFG(log, "red"))
	return true
}

// Orange messages when something need your attention
//
// Level: 2
//
// Example
//
//     logger.Warn("Duplicated")
func (l Logger) Warn(t string) bool {
	if l.l < 2 {
		return false
	}
	timeNow := time.Now()
	log := timeNow.Format(time.UnixDate) + "  " + l.n + "  [WARN] - " + t
	fmt.Println(cl.ChangeFG(log, "yellow"))
	return true
}

// Green messages for success or just useful moments
//
// Level: 3
//
// Example
//
//     logger.Log("Success")
func (l Logger) Log(t string) bool {
	if l.l < 3 {
		return false
	}
	timeNow := time.Now()
	log := timeNow.Format(time.UnixDate) + "  " + l.n + "   [LOG] - " + t
	fmt.Println(cl.ChangeFG(log, "green"))
	return true
}

// Grey messages just for moments when you want get some information
//
// Level: 4
//
// Example
//
//     logger.Info(strconv.Itoa(len(members)))
func (l Logger) Info(t string) bool {
	if l.l < 4 {
		return false
	}
	timeNow := time.Now()
	log := timeNow.Format(time.UnixDate) + "  " + l.n + "  [INFO] - " + t
	fmt.Println(cl.ChangeFG(log, "grey"))
	return true
}

// Purple messages for debugging
//
// Level: 5
//
// Example
//
//     logger.Debug("1")
func (l Logger) Debug(t string) bool {
	if l.l < 5 {
		return false
	}
	timeNow := time.Now()
	log := timeNow.Format(time.UnixDate) + "  " + l.n + " [DEBUG] - " + t
	fmt.Println(cl.ChangeFG(log, "magenta"))
	return true
}
