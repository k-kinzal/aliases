package logger

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// Log levels.
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

// SetLogLevel changes log level.
func SetLogLevel(level int) {
	logLevel = level
}

// SetOutput changes the output destination of log.
func SetOutput(output *os.File) {
	logOutput = output
}

// log outputs destination of log.
func log(level int, msg string) {
	if level < logLevel {
		return
	}
	fmt.Fprintln(logOutput, msg)
}

// Debug output the debug log.
func Debug(a ...interface{}) {
	log(DebugLevel, color.HiBlackString(fmt.Sprint(a...)))
}

// Debugf output the debug log with formatted.
func Debugf(format string, a ...interface{}) {
	log(DebugLevel, color.HiBlackString(format, a...))
}

// Info output the information log.
func Info(a ...interface{}) {
	log(InfoLevel, fmt.Sprint(a...))
}

// Infof output the information log with formatted.
func Infof(format string, a ...interface{}) {
	log(InfoLevel, fmt.Sprintf(format, a...))
}

// Warn output the warning log.
func Warn(a ...interface{}) {
	log(WarnLevel, "WARNING: "+fmt.Sprint(a...))
}

// Warnf output the warning log with formatted.
func Warnf(format string, a ...interface{}) {
	log(WarnLevel, "WARNING: "+fmt.Sprintf(format, a...))
}

// Fatal output the fatal log.
func Fatal(a ...interface{}) {
	log(FatalLevel, fmt.Sprint(a...))
}

// Fatalf output the fatal log with formatted.
func Fatalf(format string, a ...interface{}) {
	log(FatalLevel, fmt.Sprintf(format, a...))
}
