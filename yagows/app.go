package yagows

import "net/http"

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
