package namespace

import (
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strconv"
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

func build(args []arg, params []param) ([]any, error) {
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
	var ix = 0
	var temp []any

	for _, a := range args {
		p := params[ix]
		switch strings.Compare(a.Name, p.Name) {
		case 0:
			v, err := createValue(a, p)
			if err != nil {
				return nil, err
			}
			temp = append(temp, v)
			ix++
		case -1:
			return nil, errors.New(fmt.Sprintf("argument name [%v] is not in the parameter list", a.Name))
		case 1:
			if !p.Nullable {
				return nil, errors.New(fmt.Sprintf("parameter [%v] does not allow nulls", p.Name))
			}
			ix++
		}
		if ix >= len(params) {
			return nil, errors.New(fmt.Sprintf("argument name [%v] is not in the parameter list", a.Name))
			//break
		}
	}
	return temp, nil
}

func createValue(a arg, p param) (any, error) {
	switch p.Type {
	case "string":
		return a.Value, nil
	case "int":
		i, err := strconv.Atoi(a.Value)
		return i, err
	}
	return nil, errors.New(fmt.Sprintf("parameter [%v] type [%v] is not supported", p.Name, p.Type))
}

func expand(r *http.Request) (response, error) {

	return response{}, nil
}
