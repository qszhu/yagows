package yagows

import (
	"context"
	"net/http"
)

type Context struct {
	Request  *Request
	Response *Response
	App      *App
	handlers []RequestHandler
	idx      int
}

func NewContext(app *App, req *http.Request) *Context {
	return &Context{
		Request:  &Request{req},
		Response: &Response{StatusCode: HttpOk, headers: map[string][]string{}, body: []byte{}},
		App:      app,
		idx:      0,
	}
}

func (c *Context) Next() {
	if c.idx < len(c.handlers) {
		handler := c.handlers[c.idx]
		c.idx++
		handler(c)
	}
}

func (c *Context) BaseContext() context.Context {
	return c.Request.request.Context()
}

func (c *Context) use(handlers ...RequestHandler) {
	c.handlers = append(c.handlers, handlers...)
}