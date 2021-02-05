package services

import (
	"encoding/json"
	"errors"
	"github.com/ciricbogdan/localsearch-home-assignment-backend/http/client"
	"github.com/ciricbogdan/localsearch-home-assignment-backend/http/server"
	izap "github.com/ciricbogdan/localsearch-home-assignment-backend/infra/zap"
	"github.com/ciricbogdan/localsearch-home-assignment-backend/model/uimodel"
	"github.com/ciricbogdan/localsearch-home-assignment-backend/model/upstreamAPI"
	"go.uber.org/zap"
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

		// Get the response from the upstream API
		resp, err := client.Get(ctx.Params().ByName("id"))
		if err != nil {
			izap.Logger.Error("get place: ", zap.Error(err))
			return err
		}

		// Respond according to the upstream API response
		switch resp.StatusCode {
		case http.StatusNotFound:

			err = errors.New("place not found")
			izap.Logger.Error("Place not found: ", zap.Error(err))
			return ctx.Error(http.StatusNotFound, err, "Place not found")
		case http.StatusOK:
			// Decode place
			var place upstreamAPI.Place
			json.NewDecoder(resp.Body).Decode(&place)

			// Send Response
			return ctx.Encode(uimodel.PlaceFromUpstreamAPI(place))
		default:

			izap.Logger.Error("Unknown  error occurred")
			return ctx.Error(http.StatusInternalServerError, errors.New("unkown error"), "Unknown error")
		}
	}
}
