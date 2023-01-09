package logger

import (
	"fmt"
	"io"
	"os"
	"time"
)

func init() {
	defaultLogger.output = os.Stderr
}

type LogLevel int

const (
	Error LogLevel = iota
	Warn
	Info
	Debug
)

var defaultLogger logger

type logger struct {
	level  LogLevel
	output io.Writer
}

func Default() logger {
	return defaultLogger
}

func SetLevel(l LogLevel) {
	defaultLogger.level = l
}

func (l logger) Debug(s ...string) {
	if l.level >= Debug {
		write(l.output, "DEBUG", s...)
	}
}

func (l logger) Info(s ...string) {
	if l.level >= Info {
		write(l.output, "INFO", s...)
	}
}

func (l logger) Warn(s ...string) {
	if l.level >= Warn {
		write(l.output, "WARN", s...)
	}
}

func (l logger) Error(s ...string) {
	if l.level >= Error {
		write(l.output, "ERROR", s...)
	}
}

func write(w io.Writer, level string, s ...string) {
	fmt.Fprint(w, time.Now().Format(time.RFC3339), " ", level, " ")
	for _, str := range s {
		fmt.Fprint(w, str, " ")
	}
	fmt.Fprintln(w)
}
