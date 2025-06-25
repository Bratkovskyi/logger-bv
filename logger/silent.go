package logger

type SilentLogger struct{}

func (s SilentLogger) Log(level LogLevel, msg string) {}
func (s SilentLogger) Info(msg string)                {}
func (s SilentLogger) Debug(msg string)               {}
func (s SilentLogger) Warning(msg string)             {}
func (s SilentLogger) Error(msg string)               {}
