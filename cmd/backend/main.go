package main

import (
	"flag"
	"github.com/ciricbogdan/localsearch-home-assignment-backend/backend"
	izap "github.com/ciricbogdan/localsearch-home-assignment-backend/infra/zap"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

var envFile = flag.String("env", "", "mandatory flag for the name of the env file")

func main() {

	flag.Parse()

	izap.Logger.Info("Backend is starting", zap.String("env file", *envFile))

	if *envFile != "" {
		err := godotenv.Load(*envFile)
		if err != nil {
			izap.Logger.Fatal("Env file error", zap.Error(err))
		}
	} else {
		izap.Logger.Fatal("Env file is mandatory")
	}

	// Init the app
	app, err := backend.New()
	if err != nil {
		izap.Logger.Fatal("Backend initialization", zap.Error(err))
	}

	// Start the app
	err = app.Run()
	if err != nil {
		izap.Logger.Fatal("Backend start", zap.Error(err))
	}

	// If the app is shut down manually we catch it to gracefully stop the app
	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	<-signals

	app.Stop()
	izap.Logger.Info("Backend has stoppend")
}
