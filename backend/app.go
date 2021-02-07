package backend

import (
	"fmt"
	"github.com/ciricbogdan/localsearch-home-assignment-backend/http/client"
	"github.com/ciricbogdan/localsearch-home-assignment-backend/http/server"
	izap "github.com/ciricbogdan/localsearch-home-assignment-backend/infra/zap"
	"github.com/ciricbogdan/localsearch-home-assignment-backend/services"
	"go.uber.org/zap"
	"os"
)

// App defines the Backend application
type App struct {
	Server *server.Server
	Client *client.Client
}

// New retunes a new instance of the Backend application
func New() (*App, error) {

	// get address from env file
	addr := fmt.Sprintf("%v:%v", os.Getenv("host"), os.Getenv("port"))

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
		izap.Logger.Fatal("Backend http server stop", zap.Error(err))
	}

	return nil
}
