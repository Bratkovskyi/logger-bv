package main

import (
	"log"
	"logger-bv/logger"
	"time"
)

type App struct {
	logger logger.Logger
	// spinner *logger.Spinner
}

func (a *App) Run() {
	startApp := logger.NewSpinner("Starting app")
	a.logger.Info("Starting app")
	startApp.Start()

	time.Sleep(500 * time.Millisecond)
	startApp.Stop()

	debugStart := logger.NewSpinner("Debug starting app")
	a.logger.Debug("Debugging something...")
	debugStart.Start()
	time.Sleep(500 * time.Millisecond)
	debugStart.Stop()

	a.logger.Warning("Low memory")
	time.Sleep(500 * time.Millisecond)

	a.logger.Error("Something went wrong!")

}

// ---------------------------------------------------------------

func main() {
	//-------------------------------------------------------------
	// 1) Console Logger
	//-------------------------------------------------------------
	// consoleLogger := logger.ConsoleLogger{}
	// app := App{logger: consoleLogger}
	// app.Run()

	//-------------------------------------------------------------
	// 2) Silent Logger
	//-------------------------------------------------------------
	// silentLogger := logger.SilentLogger{}
	// app := App{logger: silentLogger}
	// app.Run()

	//-------------------------------------------------------------
	// 3) File Logger
	//-------------------------------------------------------------
	// file, err := logger.NewFileLogger("app.log")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// app := App{logger: file}
	// app.Run()

	//-------------------------------------------------------------
	// 4) MultiLogger  (Console + File)
	//-------------------------------------------------------------
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
	// app.Run()

	//-------------------------------------------------------------
	// 5) MultiLogger + Spinner
	//-------------------------------------------------------------
	// spin := logger.NewSpinner("Loading big task")
	// consoleLogger := logger.ConsoleLogger{Spinner: spin}
	consoleLogger := logger.ConsoleLogger{}

	file, err := logger.NewFileLogger("app.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	multi := logger.MultiLogger{
		Loggers: []logger.Logger{consoleLogger, file},
	}

	// spin.Start()
	// time.Sleep(2 * time.Second)

	// spin.Stop()

	app := App{logger: multi}
	app.Run()
}
