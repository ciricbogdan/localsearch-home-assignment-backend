package server

import (
	"fmt"
)

func healthcheck(ctx *Context) error {
	msg := fmt.Sprintf("Server is running")
	return ctx.Encode(msg)
}
