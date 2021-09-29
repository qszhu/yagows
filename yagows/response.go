package yagows

type Response struct {
	StatusCode int
	headers    map[string][]string
	body       []byte
}

func (r *Response) WriteHeader(name string, value string) {
	r.headers[name] = append(r.headers[name], value)
}
