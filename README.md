# 📦 Logger BV

A simple, extensible logger in Go with support for:

- Log levels (Debug, Info, Warning, Error)
- Colorful console output
- File logging with level filtering
- Silent mode (`SilentLogger`)
- Multi-logger output (`MultiLogger`)
- Timestamp, caller location, and level info in every log

---

## 🔧 Installation

Add the package to your project:

```bash
go mod init my-project
go get github.com/yourname/logger-bv
```

---

## 🛠 Features

### ✅ Log Levels (enum-like style)

```go
const (
    LevelDebug LogLevel = iota
    LevelInfo
    LevelWarning
    LevelError
)
```

### ✅ Logger Interface

```go
type Logger interface {
    Log(level LogLevel, msg string)
    Info(message string)
    Warning(message string)
    Debug(message string)
    Error(message string)
}
```

---

## 🔎 Usage

### 🖥 ConsoleLogger

```go
logger := logger.ConsoleLogger{}
logger.Info("App started")
```

✅ Colored output + timestamp + caller location

---

### 🗃 FileLogger

```go
fileLogger, err := logger.NewFileLogger("app.log", logger.LevelInfo)
if err != nil {
    log.Fatal(err)
}
defer fileLogger.Close()

fileLogger.Debug("This won't be logged — below LevelInfo")
fileLogger.Warning("This will be written to the file")
```

✅ Logs to file  
✅ Level filtering  
✅ Includes timestamp and caller

---

### 🤫 SilentLogger

```go
logger := logger.SilentLogger{}
logger.Info("This won't appear")
```

✅ Suppresses all log output

---

### 🧩 MultiLogger

```go
fileLogger, _ := logger.NewFileLogger("app.log")
multi := logger.MultiLogger{
    Loggers: []logger.Logger{
        logger.ConsoleLogger{},
        fileLogger,
    },
}
multi.Info("Logs to both console and file")
```

✅ Dispatches to multiple loggers

---

## 📁 Example (from `example/loggerDemo.go`)

```go
package main

import (
    "log"
    "logger-bv/logger"
)

func main() {
    fileLogger, err := logger.NewFileLogger("app.log")
    if err != nil {
        log.Fatal(err)
    }
    defer fileLogger.Close()

    app := App{
        logger: logger.MultiLogger{
            Loggers: []logger.Logger{
                logger.ConsoleLogger{},
                fileLogger,
            },
        },
    }
    app.Run()
}
```

---

## 📌 TODO (future improvements)

- [ ] Support for stdout/stderr targeting
- [ ] JSON log format
- [ ] Configurable log format
- [ ] Environment variable level control (`LOG_LEVEL`)

---

## 🧑‍💻 Author

[Vladyslav Bratkovskyi](https://github.com/Bratkovskyi) — project made for learning purposes.