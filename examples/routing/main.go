package main

import (
	"fmt"
	"log"
	. "yagows"
	. "yagows/middleware"
	. "yagows/router"
)

const BindAddress = "localhost"
const Port = 8090

func main() {
	app := NewApp()

	app.Use(Logging())
	app.Use(Recovery())

	router := NewRouter()
	router.Post("/articles/", func(c *Context) {
		c.Response.WriteStringBody("new article")
	})
	router.Get("/articles/", func(c *Context) {
		c.Response.WriteStringBody("list articles")
	})
	router.Get("/article/:id", func(c *Context) {
		id := c.Params["id"]
		c.Response.WriteStringBody(fmt.Sprintf("get article %s details", id))
	})
	router.Put("/article/:id", func(c *Context) {
		id := c.Params["id"]
		c.Response.WriteStringBody(fmt.Sprintf("modify article %s", id))
	})
	router.Delete("/article/:id", func(c *Context) {
		id := c.Params["id"]
		c.Response.WriteStringBody(fmt.Sprintf("delete article %s", id))
	})
	app.Use(router.Routes())

	log.Printf("Listening on %s:%d...", BindAddress, Port)
	app.Listen(BindAddress, Port)
}
