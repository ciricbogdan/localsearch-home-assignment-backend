package server

import (
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Context describes HTTP request context. Provides access to application
// resources.
type Context struct {
	http.ResponseWriter
	context.Context
	request *http.Request
	params  httprouter.Params
}

// Request returns an HTTP request.
func (c *Context) Request() *http.Request {
	return c.request
}

// Params returns the request params
func (c *Context) Params() httprouter.Params {
	return c.params
}

// DecodeBody decodes request body
func (c *Context) DecodeBody(dst interface{}) error {
	return json.NewDecoder(c.request.Body).Decode(dst)
}

// Encode encodes and writes the response body.
func (c *Context) Encode(src interface{}) error {
	return json.NewEncoder(c.ResponseWriter).Encode(ResponseFromData(src))
}

// Error is a standardized error response that uses application's default codec!
func (c *Context) Error(status int, err error, code string) error {
	c.WriteHeader(status)
	return json.NewEncoder(c.ResponseWriter).Encode(ErrorResponse(err, code))
}
