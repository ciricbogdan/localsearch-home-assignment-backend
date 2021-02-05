package server

import (
	izap "github.com/ciricbogdan/localsearch-home-assignment-backend/infra/zap"
	"go.uber.org/zap"
	"net/http"
)

// Middleware defines a function that wrapes around an http handelr
type Middleware func(http.Handler) http.Handler

// WithLog is a middleware which logs the request
func WithLog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		izap.Logger.Info("Request", zap.String("Method", r.Method), zap.String("URL", r.URL.Path))
		h.ServeHTTP(w, r)
	})
}
