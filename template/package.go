package template

import "net/http"

type Response struct {
	Name string
	Sql  string
	Args []any
}
type Interface struct {
	Expand func(r *http.Request) (Response, error)
}

// Processor -
var Processor = func() *Interface {
	return &Interface{
		Expand: func(r *http.Request) (Response, error) {
			return agent.expand(r)
		},
	}
}()
