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

	router := NewRouter()
	router.Get("/", Timeout(1000), rootHandler)
	router.Get("/timeout", Timeout(1000), timeoutHandler)
	router.Get("/panic", Timeout(1000), panicHandler)
	app.Use(router.Routes())

	log.Printf("Listening on %s:%d...", BindAddress, Port)
	app.Listen(BindAddress, Port)
}
