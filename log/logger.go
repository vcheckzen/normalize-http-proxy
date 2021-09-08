// https://github.com/smallnest/rpcxlite/tree/master/log
// https://qvault.io/golang/golang-logging-best-practices/
package log

import (
	"log"
	"os"
)

const (
	// https://chende.ren/2021/01/19131400-016-go-log.html
	calldepth = 3
)

var l Logger = &defaultLogger{true, log.New(os.Stderr, "[NP] ", log.Ldate|log.Ltime|log.Lshortfile|log.Lmsgprefix)}

type Logger interface {
	SetVerbose(verbose bool)

	Info(v ...interface{})
	Infof(format string, v ...interface{})

	Warn(v ...interface{})
	Warnf(format string, v ...interface{})

	Error(v ...interface{})
	Errorf(format string, v ...interface{})

	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
}

func SetLogger(logger Logger) {
	l = logger
}

func SetVerbose(verbose bool) {
	l.SetVerbose(verbose)
}

func Info(v ...interface{}) {
	l.Info(v...)
}
func Infof(format string, v ...interface{}) {
	l.Infof(format, v...)
}

func Warn(v ...interface{}) {
	l.Warn(v...)
}
func Warnf(format string, v ...interface{}) {
	l.Warnf(format, v...)
}

func Error(v ...interface{}) {
	l.Error(v...)
}
func Errorf(format string, v ...interface{}) {
	l.Errorf(format, v...)
}

func Fatal(v ...interface{}) {
	l.Fatal(v...)
}
func Fatalf(format string, v ...interface{}) {
	l.Fatalf(format, v...)
}
