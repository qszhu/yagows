package middleware

import (
	"log"
	"time"
	"yagows"
)

type LogMiddleware struct {
	startTime time.Time
}

func NewLogMiddleware() *LogMiddleware {
	return &LogMiddleware{}
}

func (m *LogMiddleware) PreRequest(c *yagows.Context) {
	m.startTime = time.Now()
	log.Printf("%s %s <- %s", c.Request.Method(), c.Request.Path(), c.Request.RemoteAddr())
}

func (m *LogMiddleware) PostRequest(c *yagows.Context) {
	elapsed := time.Now().Sub(m.startTime)
	log.Printf("-> %s: %d, %dms", c.Request.RemoteAddr(), c.Response.StatusCode, elapsed.Milliseconds())
}
