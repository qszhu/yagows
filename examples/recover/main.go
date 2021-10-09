package main

import (
	"log"
	. "yagows"
	. "yagows/middleware"
	. "yagows/router"
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

	router := NewRouter()
	router.Get("/", rootHandler)
	app.Use(router.Routes())

	log.Printf("Listening on %s:%d...", BindAddress, Port)
	app.Listen(BindAddress, Port)
}
