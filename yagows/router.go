package yagows

type Router struct {
	routes map[string]RequestHandler
}

func NewRouter() *Router {
	return &Router{routes: map[string]RequestHandler{}}
}

func (r *Router) Get(path string, handler RequestHandler) {
	r.routes[path] = handler
}

func (r *Router) Match(method string, path string) RequestHandler {
	return r.routes[path]
}
