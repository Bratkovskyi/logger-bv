package logger

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

type FileLogger struct {
	file     *os.File
	MinLevel LogLevel
}

func (f *FileLogger) Close() error {
	return f.file.Close()
}

func NewFileLogger(path string, minLevel ...LogLevel) (*FileLogger, error) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	logger := &FileLogger{file: f}

	if len(minLevel) > 0 {
		logger.MinLevel = minLevel[0]
	} else {
		logger.MinLevel = LevelDebug
	}

	return logger, nil
}

func (f *FileLogger) Log(level LogLevel, msg string) {
	if level < f.MinLevel {
		return
	}

	timestamp := time.Now().Format(time.RFC3339Nano)

	var prefix string
	switch level {
	case LevelDebug:
		prefix = "[DEBUG]"
	case LevelInfo:
		prefix = "[INFO]"
	case LevelWarning:
		prefix = "[WARN]"
	case LevelError:
		prefix = "[ERROR]"
	default:
		prefix = "[LOG]"
	}

	_, file, line, ok := runtime.Caller(2)
	location := ""
	if ok {
		location = fmt.Sprintf("%s:%d", file, line)
	} else {
		location = "unknown"
	}

	fmt.Fprintf(f.file, "%s %s [%s] %s\n", prefix, timestamp, location, msg)
}

func (f *FileLogger) Info(msg string)    { f.Log(LevelInfo, msg) }
func (f *FileLogger) Debug(msg string)   { f.Log(LevelDebug, msg) }
func (f *FileLogger) Warning(msg string) { f.Log(LevelWarning, msg) }
func (f *FileLogger) Error(msg string)   { f.Log(LevelError, msg) }
