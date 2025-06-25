package logger

type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarning
	LevelError
)

type Logger interface {
	Log(level LogLevel, msg string)
	Info(message string)
	Warning(message string)
	Debug(message string)
	Error(message string)
}
