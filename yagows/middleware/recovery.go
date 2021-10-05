package middleware

import (
	"log"
	"runtime/debug"
	. "yagows"
)

func Recovery() RequestHandler {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("%v %s\n", err, debug.Stack())
				c.Response.WriteStringBody("")
				c.Response.StatusCode = HttpInternalError
			}
		}()

		c.Next()
	}
}
