package yagows

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	Router      *Router
	middlewares []Middleware
	config      map[string]string
}

func NewApp() *App {
	return &App{Router: NewRouter(), middlewares: []Middleware{}, config: map[string]string{}}
}

func (a *App) Set(name string, value string) {
	a.config[name] = value
}

func (a *App) Get(name string) string {
	return a.config[name]
}

func (a *App) Use(middlewares ...Middleware) {
	a.middlewares = middlewares
}

func (a *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			w.WriteHeader(HttpInternalError)
		}
	}()

	ctx := NewContext(a, req)

	handler := a.Router.Match(req.Method, req.URL.Path)
	if handler == nil {
		w.WriteHeader(HttpNotFound)
		return
	}

	for _, m := range a.middlewares {
		m.PreRequest(ctx)
	}
	handler(ctx)
	for i := len(a.middlewares) - 1; i >= 0; i-- {
		a.middlewares[i].PostRequest(ctx)
	}

	for name, headers := range ctx.Response.headers {
		for _, header := range headers {
			w.Header().Set(name, header)
		}
	}
	w.WriteHeader(ctx.Response.StatusCode)
	_, err := w.Write(ctx.Response.body)
	if err != nil {
		w.WriteHeader(HttpInternalError)
	}
}

func (a *App) Listen(bindAddress string, port int) {
	server := &http.Server{
		Handler: a,
		Addr:    fmt.Sprintf("%s:%d", bindAddress, port),
	}

	done := make(chan interface{})
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