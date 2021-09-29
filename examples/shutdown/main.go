package main

import (
	"log"
	"os"
	"time"
	"yagows"
	"yagows/middleware"
)

const BindAddress = "localhost"
const Port = 8090

func rootHandler(c *yagows.Context) {
	time.Sleep(5 * time.Second)
	c.Response.WriteStringBody("ok")
}

func main() {
	app := yagows.NewApp()

	app.Use(middleware.NewLogMiddleware())

	app.Router.Get("/", rootHandler)

	log.Printf("pid: %d", os.Getpid())
	log.Printf("Listening on %s:%d...", BindAddress, Port)
	app.Listen(BindAddress, Port)
}
