package yagows

type Middleware interface {
	PreRequest(c *Context)
	PostRequest(c *Context)
}
