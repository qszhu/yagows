package yagows

import "net/http"

type Context struct {
	Request  *Request
	Response *Response
	App      *App
}

func NewContext(app *App, req *http.Request) *Context {
	return &Context{
		Request:  &Request{req},
		Response: &Response{StatusCode: HttpOk, headers: map[string][]string{}, body: []byte{}},
		App:      app,
	}
}
