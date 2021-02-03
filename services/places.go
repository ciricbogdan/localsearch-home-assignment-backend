package services

import (
	"encoding/json"
	"fmt"
	"github.com/ciricbogdan/localsearch-home-assignment-backend/http/client"
	"github.com/ciricbogdan/localsearch-home-assignment-backend/http/server"
	"github.com/ciricbogdan/localsearch-home-assignment-backend/model"
	"net/http"
)

// Places defines a service for places
type Places struct {
	Path   string
	Client *client.Client
}

// Register is part of the services interface to register the places service into the passed http server
func (p *Places) Register(s *server.Server) {
	s.Handle(http.MethodGet, p.Path+"/:id", placesByID(p.Client))
}

func placesByID(client *client.Client) server.Handle {

	return func(ctx *server.Context) error {

		resp, err := client.Get(ctx.Params().ByName("id"))
		if err != nil {
			fmt.Println("error occurred")
		}

		// Decode place
		var place model.Place
		json.NewDecoder(resp.Body).Decode(&place)

		// Send Response
		return ctx.Encode(place)
	}
}
