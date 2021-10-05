package main

import (
	"log"
	"time"
	. "yagows"
	. "yagows/middleware"
)

const BindAddress = "localhost"
const Port = 8090

func rootHandler(c *Context) {
	c.Response.WriteStringBody("ok")
}

func timeoutHandler(c *Context) {
	time.Sleep(10 * time.Second)
	c.Response.WriteStringBody("ok")
}

func panicHandler(_ *Context) {
	panic("error")
}

func main() {
	app := NewApp()

	app.Use(Logging())
	app.Use(Recovery())

	app.Router.Get("/", Timeout(1000), rootHandler)
	app.Router.Get("/timeout", Timeout(1000), timeoutHandler)
	app.Router.Get("/panic", Timeout(1000), panicHandler)

	log.Printf("Listening on %s:%d...", BindAddress, Port)
	app.Listen(BindAddress, Port)
}
