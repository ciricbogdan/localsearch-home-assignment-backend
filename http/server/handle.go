package server

// Handle handles HTTP Request
type Handle func(ctx *Context) error
