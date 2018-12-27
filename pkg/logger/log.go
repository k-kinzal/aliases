package logger

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

const (
	DebugLevel = iota
	InfoLevel
	WarnLevel
	FatalLevel
)

var (
	logLevel  = InfoLevel
	logOutput = os.Stdout
)

func SetLogLevel(level int) {
	logLevel = level
}

func SetOutput(output *os.File) {
	logOutput = output
}

func log(level int, msg string) {
	if level < logLevel {
		return
	}
	fmt.Fprintln(logOutput, msg)
}

func Debug(a ...interface{}) {
	log(DebugLevel, color.HiBlackString(fmt.Sprint(a...)))
}

func Debugf(format string, a ...interface{}) {
	log(DebugLevel, color.HiBlackString(format, a...))
}

func Info(a ...interface{}) {
	log(InfoLevel, color.WhiteString(fmt.Sprint(a...)))
}

func Infof(format string, a ...interface{}) {
	log(InfoLevel, color.WhiteString(format, a...))
}

func Warn(a ...interface{}) {
	log(WarnLevel, color.YellowString(fmt.Sprint(a...)))
}

func Warnf(format string, a ...interface{}) {
	log(WarnLevel, color.YellowString(format, a...))
}

func Fatal(a ...interface{}) {
	log(FatalLevel, color.RedString(fmt.Sprint(a...)))
}

func Fatalf(format string, a ...interface{}) {
	log(FatalLevel, color.RedString(format, a...))
}
