package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/ciricbogdan/localsearch-home-assignment-backend/backend"
	izap "github.com/ciricbogdan/localsearch-home-assignment-backend/infra/zap"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var envFile = flag.String("env", "", "mandatory flag with the .env file in the current directory")

func main() {

	flag.Parse()

	izap.Logger.Info("Backend is starting", zap.String("env file", *envFile))

	if *envFile != "" {
		err := godotenv.Load(*envFile)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(errors.New("-env flag is mandatory"))
	}

	// Init the app
	app, err := backend.New()
	if err != nil {
		log.Fatal(err)
	}

	// Start the app
	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}

	// If the app is shut down manually we catch it to gracefully stop the app
	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	<-signals

	app.Stop()
	fmt.Println("Stopped")
}
