package main

import (
	"log"
	"net/http"
	"os"
	"yagows"
	"yagows/middleware"
)

const KeyVersion = "VERSION"

func rootHandler(c *yagows.Context) {
	for name, headers := range c.Request.Headers() {
		for _, header := range headers {
			c.Response.WriteHeader(name, header)
		}
	}

	c.Response.WriteHeader(KeyVersion, c.App.Get(KeyVersion))

	c.Response.StatusCode = 200
}

func main() {
	app := yagows.NewApp()

	app.Set(KeyVersion, os.Getenv(KeyVersion))

	app.Use(middleware.NewLogMiddleware())

	app.Router.Get("/", rootHandler)
	app.Router.Get("/healthz", func(*yagows.Context) {})

	server := &http.Server{
		Handler: app,
		Addr:    ":8090",
	}
	log.Fatal(server.ListenAndServe())
}
