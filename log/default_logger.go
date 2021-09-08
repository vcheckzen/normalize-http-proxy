package log

import (
	"fmt"
	"log"
	"os"
)

// https://blog.csdn.net/books1958/article/details/22795203
type defaultLogger struct {
	verbose bool
	*log.Logger
}

func (l *defaultLogger) out(lvl string, v ...interface{}) {
	l.Output(calldepth, prefix(lvl, fmt.Sprint(v...)))

}

func (l *defaultLogger) outf(lvl string, format string, v ...interface{}) {
	l.Output(calldepth, prefix(lvl, fmt.Sprintf(format, v...)))
}

func (l *defaultLogger) SetVerbose(verbose bool) {
	l.verbose = verbose
}

func (l *defaultLogger) Info(v ...interface{}) {
	if l.verbose {
		l.out("INFO", v...)
	}
}

func (l *defaultLogger) Infof(format string, v ...interface{}) {
	if l.verbose {
		l.outf("INFO", format, v...)
	}
}

func (l *defaultLogger) Warn(v ...interface{}) {
	l.out("WARNING", v...)
}

func (l *defaultLogger) Warnf(format string, v ...interface{}) {
	l.outf("WARNING", format, v...)
}

func (l *defaultLogger) Error(v ...interface{}) {
	l.out("ERROR", v...)
}

func (l *defaultLogger) Errorf(format string, v ...interface{}) {
	l.outf("ERROR", format, v...)
}

func (l *defaultLogger) Fatal(v ...interface{}) {
	l.out("FATAL", v...)
	os.Exit(1)
}

func (l *defaultLogger) Fatalf(format string, v ...interface{}) {
	l.outf("FATAL", format, v...)
	os.Exit(1)
}

func prefix(lvl, msg string) string {
	return fmt.Sprintf("%s: %s", lvl, msg)
}
