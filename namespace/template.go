package namespace

import (
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strings"
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

func build(args []arg, params []param) (t []any, err error) {
	if len(args) == 0 {
		return nil, errors.New("input args are empty")
	}
	if len(params) == 0 {
		return nil, errors.New("input parameters are empty")
	}
	slices.SortFunc(params, func(a, b param) int {
		return strings.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
	})
	slices.SortFunc(args, func(a, b arg) int {
		return strings.Compare(strings.ToLower(a.Name), strings.ToLower(b.Name))
	})
	// var argIndex = 0
	var ix = 0

	for _, a := range args {
		p := params[ix]
		if a.Name == params[ix].Name {
			err = validateType(a, p)
			if err != nil {
				return
			}

		}

	}
	return
}

func validateType(a arg, p param) error {
	if a.Value == nil && !p.Nullable {
		return errors.New(fmt.Sprintf("parameter %v does not allow nulls", p.Name))
	}

	return nil
}

func expand(r *http.Request) (response, error) {

	return response{}, nil
}
