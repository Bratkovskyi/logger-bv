package logger

import (
	"fmt"
	"runtime"
	"time"
)

// --- ANSI Colors
const (
	ColorReset  = "\033[0m"
	ColorGreen  = "\033[32m" // Info
	ColorBlue   = "\033[34m" // Debug
	ColorYellow = "\033[33m" // Warning
	ColorRed    = "\033[31m" // Error
)

func (c ConsoleLogger) Log(level LogLevel, msg string) {
	var color, prefix string

	switch level {
	case LevelDebug:
		color = ColorBlue
		prefix = "[DEBUG]"
	case LevelInfo:
		color = ColorGreen
		prefix = "[INFO]"
	case LevelWarning:
		color = ColorYellow
		prefix = "[WARN]"
	case LevelError:
		color = ColorRed
		prefix = "[ERROR]"
	default:
		color = ColorReset
		prefix = "[LOG]"
	}

	timestamp := time.Now().Format(time.RFC3339Nano)

	_, file, line, ok := runtime.Caller(2)
	location := ""
	if ok {
		location = fmt.Sprintf("%s:%d", file, line)
	} else {
		location = "unknown"
	}

	fmt.Printf("%s%s %s [%s] %s%s\n", color, prefix, timestamp, location, msg, ColorReset)
}

type ConsoleLogger struct{}

func (c ConsoleLogger) Info(msg string)    { c.Log(LevelInfo, msg) }
func (c ConsoleLogger) Debug(msg string)   { c.Log(LevelDebug, msg) }
func (c ConsoleLogger) Warning(msg string) { c.Log(LevelWarning, msg) }
func (c ConsoleLogger) Error(msg string)   { c.Log(LevelError, msg) }
