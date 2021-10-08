package yagows

import (
	"strings"
)

const GET = "get"
const POST = "post"
const DELETE = "delete"
const PUT = "put"
const PATCH = "patch"

type Router struct {
	routes map[string]map[string][]RequestHandler
}

func NewRouter() *Router {
	return &Router{
		routes: map[string]map[string][]RequestHandler{
			GET:    {},
			POST:   {},
			DELETE: {},
			PUT:    {},
			PATCH:  {},
		},
	}
}

func (r *Router) Get(path string, handlers ...RequestHandler) {
	r.routes[GET][path] = handlers
}

func (r *Router) Post(path string, handlers ...RequestHandler) {
	r.routes[GET][path] = handlers
}

func (r *Router) Delete(path string, handlers ...RequestHandler) {
	r.routes[GET][path] = handlers
}

func (r *Router) Put(path string, handlers ...RequestHandler) {
	r.routes[GET][path] = handlers
}

func (r *Router) Path(path string, handlers ...RequestHandler) {
	r.routes[GET][path] = handlers
}

func (r *Router) Match(method string, path string) []RequestHandler {
	return r.routes[strings.ToLower(method)][path]
}

func defaultNotFoundHandler(c *Context) {
	c.Response.StatusCode = HttpNotFound
}

func (r *Router) Routes() RequestHandler {
	return func(c *Context) {
		req := c.Request.request
		handlers := r.Match(req.Method, req.URL.Path)

		if handlers == nil {
			handlers = []RequestHandler{defaultNotFoundHandler}
		}

		c.use(handlers...)
		c.Next()
	}
}
