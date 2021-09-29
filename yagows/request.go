package yagows

import (
	"net/http"
	"strings"
)

type Request struct {
	request *http.Request
}

func (r *Request) RemoteAddr() string {
	return r.request.RemoteAddr
}

func (r *Request) Headers() http.Header {
	return r.request.Header
}

func (r *Request) Method() string {
	return strings.ToUpper(r.request.Method)
}

func (r *Request) Path() string {
	return r.request.URL.Path
}
