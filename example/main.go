package main

import (
	"logger-bv/logger"
)

type App struct {
	logger logger.Logger
}

func (a *App) Run() {
	a.logger.Info("Starting app")
	a.logger.Debug("Debugging something...")
	a.logger.Warning("Low memory")
	a.logger.Error("Something went wrong!")
}

func main() {
	// Coonsole Logger Example
	consoleLogger := logger.ConsoleLogger{}
	app := App{logger: consoleLogger}

	// Silent Logger Example
	// silentLogger := logger.SilentLogger{}
	// app := App{logger: silentLogger}

	// File Logger Example
	// file, err := logger.NewFileLogger("app.log")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// app := App{logger: file}

	// MultiLogger Example
	// consoleLogger := logger.ConsoleLogger{}
	// file, err := logger.NewFileLogger("app.log")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// multi := logger.MultiLogger{
	// 	Loggers: []logger.Logger{consoleLogger, file},
	// }
	// app := App{logger: multi}

	app.Run()
}
