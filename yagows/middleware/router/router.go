package router

import (
	"strings"
	. "yagows"
)

const GET = "get"
const POST = "post"
const DELETE = "delete"
const PUT = "put"
const PATCH = "patch"

type Router struct {
	routes map[string]*Trie
}

func NewRouter() *Router {
	return &Router{
		routes: map[string]*Trie{
			GET:    NewTrie(),
			POST:   NewTrie(),
			DELETE: NewTrie(),
			PUT:    NewTrie(),
			PATCH:  NewTrie(),
		},
	}
}

func toTrieData(handlers []RequestHandler) []interface{} {
	data := make([]interface{}, len(handlers))
	for i, h := range handlers {
		data[i] = h
	}
	return data
}

func (r *Router) Get(path string, handlers ...RequestHandler) {
	r.routes[GET].Add(path, toTrieData(handlers)...)
}

func (r *Router) Post(path string, handlers ...RequestHandler) {
	r.routes[POST].Add(path, toTrieData(handlers)...)
}

func (r *Router) Delete(path string, handlers ...RequestHandler) {
	r.routes[DELETE].Add(path, toTrieData(handlers)...)
}

func (r *Router) Put(path string, handlers ...RequestHandler) {
	r.routes[PUT].Add(path, toTrieData(handlers)...)
}

func (r *Router) Patch(path string, handlers ...RequestHandler) {
	r.routes[PATCH].Add(path, toTrieData(handlers)...)
}

func (r *Router) Match(method string, path string) *MatchResult {
	return r.routes[strings.ToLower(method)].Match(path)
}

func defaultNotFoundHandler(c *Context) {
	c.Response.StatusCode = HttpNotFound
}

func (r *Router) Routes() RequestHandler {
	return func(c *Context) {
		req := c.Request.RawRequest()
		res := r.Match(req.Method, req.URL.Path)

		var handlers []RequestHandler
		if res.data == nil {
			handlers = []RequestHandler{defaultNotFoundHandler}
		} else {
			handlers = make([]RequestHandler, len(res.data))
			for i, h := range res.data {
				handlers[i] = h.(RequestHandler)
			}
		}

		c.Use(handlers...)
		c.Params = res.params

		c.Next()
	}
}
