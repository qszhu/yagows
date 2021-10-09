package yagows

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
)

type App struct {
	middlewares []RequestHandler
	config      map[string]string
}

func NewApp() *App {
	return &App{middlewares: []RequestHandler{}, config: map[string]string{}}
}

func (a *App) Set(name string, value string) {
	a.config[name] = value
}

func (a *App) Get(name string) string {
	return a.config[name]
}

func (a *App) Use(middlewares ...RequestHandler) {
	a.middlewares = append(a.middlewares, middlewares...)
}

func (a *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := NewContext(a, req)

	ctx.handlers = []RequestHandler{}
	ctx.handlers = append(ctx.handlers, a.middlewares...)

	// kick off handlers
	ctx.Next()

	// writing response
	for name, headers := range ctx.Response.headers {
		for _, header := range headers {
			w.Header().Set(name, header)
		}
	}
	w.WriteHeader(ctx.Response.StatusCode)

	_, err := w.Write(ctx.Response.body)
	// TODO: better error handling?
	if err != nil {
		log.Printf("%v %s\n", err, debug.Stack())
		w.WriteHeader(HttpInternalError)
	}
}

func (a *App) Listen(bindAddress string, port int) {
	server := &http.Server{
		Handler: a,
		Addr:    fmt.Sprintf("%s:%d", bindAddress, port),
	}

	done := make(chan struct{})
	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		sig := <-quit

		log.Printf("Received %v, Shutting down...", sig)
		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(done)
	}()

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
	<-done
}
