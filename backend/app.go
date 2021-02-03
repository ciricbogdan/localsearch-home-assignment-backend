package backend

import (
	"github.com/ciricbogdan/localsearch-home-assignment-backend/http/client"
	"github.com/ciricbogdan/localsearch-home-assignment-backend/http/server"
	"github.com/ciricbogdan/localsearch-home-assignment-backend/services"
	"log"
	"os"
)

// App defines the Backend application
type App struct {
	Server *server.Server
	Client *client.Client
}

// New retunes a new instance of the Backend application
func New() (*App, error) {

	// get address from env
	addr := os.Getenv("host") + ":" + os.Getenv("port")

	// init http server
	s, err := server.New(server.WithAddr(addr))
	if err != nil {
		return nil, err
	}

	// init http client
	c, err := client.New(os.Getenv("places_url"))
	if err != nil {
		return nil, err
	}

	// define app
	app := &App{
		Server: s,
		Client: c,
	}

	// define and register services
	places := services.Places{
		Path:   "/places",
		Client: c,
	}
	app.Server.RegisterService(&places)

	return app, nil
}

// Run starts the Backend application
func (a *App) Run() error {

	// run http server
	err := a.Server.Run()
	if err != nil {
		return err
	}

	return nil
}

// Stop stops the Backend application
func (a *App) Stop() error {

	// stop the http server
	err := a.Server.Stop()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
