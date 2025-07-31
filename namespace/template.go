package namespace

import (
	"errors"
	"net/http"
)

type response struct {
	Name string
	Sql  string
	Args []any
}

type param struct {
	Name     string `json:"name"`
	Nullable bool   `json:"nullable"`
	Type     string `json:"type"`
	SqlType  string `json:"sql-type"`
}
type template struct {
	Name   string  `json:"name"`
	Params []param `json:"args"`
}

func expand(r *http.Request) (response, error) {
	if r == nil {
		return response{}, errors.New("http Request is nil")
	}

	return response{}, nil
}
