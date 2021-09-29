package main

import (
	"fmt"
	"log"
	"net/http"
	"yagows"
	"yagows/middleware"
)

const Port = 8090

func rootHandler(c *yagows.Context) {
	c.Response.WriteStringBody("ok")

	panic("something wrong")
}

func main() {
	app := yagows.NewApp()

	app.Use(middleware.NewLogMiddleware())

	app.Router.Get("/", rootHandler)

	server := &http.Server{
		Handler: app,
		Addr:    fmt.Sprintf(":%d", Port),
	}
	log.Printf("Listening on port %d...", Port)
	log.Fatal(server.ListenAndServe())
}
