package server

import (
	"fmt"
)

func healthcheck(ctx *Context) error {
	msg := fmt.Sprintf("Server is running")
	ctx.ResponseWriter.Write([]byte(msg))

	return nil
}
