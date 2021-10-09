package main

import (
	"log"
	"os"
	. "yagows"
	. "yagows/middleware"
	. "yagows/router"
)

const BindAddress = "localhost"
const Port = 8090
const KeyVersion = "VERSION"

func rootHandler(c *Context) {
	for name, headers := range c.Request.Headers() {
		for _, header := range headers {
			c.Response.WriteHeader(name, header)
		}
	}

	c.Response.WriteHeader(KeyVersion, c.App.Get(KeyVersion))

	c.Response.StatusCode = 200
}

func main() {
	app := NewApp()

	app.Set(KeyVersion, os.Getenv(KeyVersion))

	app.Use(Logging())

	router := NewRouter()
	router.Get("/", rootHandler)
	router.Get("/healthz", func(*Context) {})
	app.Use(router.Routes())

	log.Printf("Listening on %s:%d...", BindAddress, Port)
	app.Listen(BindAddress, Port)
}
