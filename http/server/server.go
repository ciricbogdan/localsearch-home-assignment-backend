package server

import (
	"context"
	izap "github.com/ciricbogdan/localsearch-home-assignment-backend/infra/zap"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"go.uber.org/zap"
	"net/http"
)

// Server defines a wrapper around http.Server
type Server struct {
	srv    http.Server
	router *httprouter.Router
}

// New returns an instance of the Server
func New(opts ...Option) (*Server, error) {

	// init server
	server := &Server{
		router: httprouter.New(),
	}

	// init cors
	cors := cors.AllowAll()

	// add middleware for all handlers
	server.srv.Handler = cors.Handler(WithLog(server.router))

	// append options of the server
	for _, opt := range opts {
		opt(server)
	}

	// add healthcheck endpoint
	server.Handle(http.MethodGet, "/", healthcheck)

	return server, nil
}

// Handle registers a handler with the passed method and path
func (s *Server) Handle(method, path string, h Handle) {

	s.router.Handle(
		method,
		path,
		s.handle(h),
	)
}

func (s *Server) handle(h Handle) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ctx := Context{
			ResponseWriter: w,
			Context:        r.Context(),
			request:        r,
			params:         p,
		}
		err := h(&ctx)
		if err != nil {
			izap.Logger.Error("request errored", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

// RegisterService registers a starts the self registration of the service on the server
func (s *Server) RegisterService(service Service) {
	service.Register(s)
}

// Run starts the server
func (s *Server) Run() error {

	go func() {
		izap.Logger.Info("Starting http server", zap.String("address", s.srv.Addr))
		err := s.srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			s.Stop()
		}
	}()

	return nil
}

// Stop stops the server gracefully
func (s *Server) Stop() error {

	izap.Logger.Info("Stopping http server", zap.String("address", s.srv.Addr))
	return s.srv.Shutdown(context.Background())
}
