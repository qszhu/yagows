package main

import (
	"log"
	"os"
	"time"
	. "yagows"
	. "yagows/middleware"
)

const BindAddress = "localhost"
const Port = 8090

func rootHandler(c *Context) {
	time.Sleep(5 * time.Second)
	c.Response.WriteStringBody("ok")
}

func main() {
	app := NewApp()

	app.Use(Logging())

	app.Router.Get("/", rootHandler)

	log.Printf("pid: %d", os.Getpid())
	log.Printf("Listening on %s:%d...", BindAddress, Port)
	app.Listen(BindAddress, Port)
}
