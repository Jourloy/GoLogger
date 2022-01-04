package logger

import (
	"fmt"
	"time"

	cl "github.com/Jourloy/Colored"
)

var n string
var l int

func Log(t string) {
	timeNow := time.Now()
	log := timeNow.Format(time.UnixDate) + " " + n + "   [LOG] - " + t
	fmt.Println(cl.ChangefG(log, "green"))
}

func Warn(t string) {
	timeNow := time.Now()
	log := timeNow.Format(time.UnixDate) + " " + n + "  [WARN] - " + t
	fmt.Println(cl.ChangefG(log, "yellow"))
}

func Info(t string) {
	if l == 1 {
		return
	}
	timeNow := time.Now()
	log := timeNow.Format(time.UnixDate) + " " + n + "  [INFO] - " + t
	fmt.Println(cl.ChangefG(log, "grey"))
}

func Error(t string) {
	timeNow := time.Now()
	log := timeNow.Format(time.UnixDate) + " " + n + " [Error] - " + t
	fmt.Println(cl.ChangefG(log, "red"))
}

func Debug(t string) {
	if l == 1 {
		return
	}
	timeNow := time.Now()
	log := timeNow.Format(time.UnixDate) + " " + n + " [Debug] - " + t
	fmt.Println(cl.ChangefG(log, "magenta"))
}

// Create logger
//
// level - logger level. Prod mode show only LOG, WARN and ERROR levels
//     Possible: all | prod
// name - name of logger
func New(level string, name string) {
	if level == "all" {
		l = 0
	} else {
		l = 1
	}
	n = name
}
