package main

import (
	"log"
	. "yagows"
	. "yagows/middleware"
)

const BindAddress = "localhost"
const Port = 8090

func rootHandler(c *Context) {
	c.Response.WriteStringBody("ok")

	panic("something wrong")
}

func main() {
	app := NewApp()

	app.Use(Logging())
	app.Use(Recovery())

	app.Router.Get("/", rootHandler)

	log.Printf("Listening on %s:%d...", BindAddress, Port)
	app.Listen(BindAddress, Port)
}
