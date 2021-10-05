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
	Router      *Router
	middlewares []RequestHandler
	config      map[string]string
}

func NewApp() *App {
	return &App{Router: NewRouter(), middlewares: []RequestHandler{}, config: map[string]string{}}
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

func notFoundHandler(c *Context) {
	c.Response.StatusCode = HttpNotFound
}

func (a *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := NewContext(a, req)

	handlers := a.Router.Match(req.Method, req.URL.Path)
	if handlers == nil {
		handlers = []RequestHandler{notFoundHandler}
	}

	ctx.handlers = []RequestHandler{}
	ctx.handlers = append(ctx.handlers, a.middlewares...)
	ctx.handlers = append(ctx.handlers, handlers...)

	ctx.Next()

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
