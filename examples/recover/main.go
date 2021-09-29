package main

import (
	"log"
	"yagows"
	"yagows/middleware"
)

const BindAddress = "localhost"
const Port = 8090

func rootHandler(c *yagows.Context) {
	c.Response.WriteStringBody("ok")

	panic("something wrong")
}

func main() {
	app := yagows.NewApp()

	app.Use(middleware.NewLogMiddleware())

	app.Router.Get("/", rootHandler)

	log.Printf("Listening on %s:%d...", BindAddress, Port)
	app.Listen(BindAddress, Port)
}
