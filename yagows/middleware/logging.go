package middleware

import (
	"log"
	"time"
	. "yagows"
)

func Logging() RequestHandler {
	return func(c *Context) {
		startTime := time.Now()
		log.Printf("%s %s <- %s", c.Request.Method(), c.Request.Path(), c.Request.RemoteAddr())

		c.Next()

		elapsed := time.Now().Sub(startTime)
		log.Printf("-> %s: %d, %dms", c.Request.RemoteAddr(), c.Response.StatusCode, elapsed.Milliseconds())
	}
}
