package logger

type MultiLogger struct {
	Loggers []Logger
}

func (m MultiLogger) Log(level LogLevel, msg string) {
	for _, l := range m.Loggers {
		l.Log(level, msg)
	}
}

func (m MultiLogger) Info(msg string)    { m.Log(LevelInfo, msg) }
func (m MultiLogger) Debug(msg string)   { m.Log(LevelDebug, msg) }
func (m MultiLogger) Warning(msg string) { m.Log(LevelWarning, msg) }
func (m MultiLogger) Error(msg string)   { m.Log(LevelError, msg) }
