package namespace

import (
	"context"
	"errors"
	"github.com/appellative-ai/core/std"
	"github.com/appellative-ai/postgres/request"
	"net/url"
)

const (
	requestLinkSql = "CALL dbo.InsertLink($1,$2,$3,$4,$5,$6,$7,$8,$9)"
)

func linkRequest(ctx context.Context, requester *request.Interface, values url.Values) (request.Result, error) {
	if values == nil {
		return request.Result{}, errors.New("query values are nil")
	}
	name, args, err := createLinkArgs(values)
	if err != nil {
		return request.Result{}, err
	}
	return requester.Execute(ctx, name, requestLinkSql, args)
}

func createLinkArgs(values url.Values) (string, []any, error) {
	name := values.Get(nameName)
	if name == "" {
		return "", nil, errors.New("name is empty")
	}
	author := values.Get(authorName)
	if author == "" {
		return "", nil, errors.New("author is empty")
	}
	thing1 := values.Get(thing1Name)
	if thing1 == "" {
		return "", nil, errors.New("thing1 is empty")
	}
	thing2 := values.Get(thing2Name)
	if thing2 == "" {
		return "", nil, errors.New("thing2 is empty")
	}
	n := std.NewName(name)
	var args []any

	args = append(args, name)
	cname := values.Get(cNameName)
	if cname != "" {
		args = append(args, cname)
	}
	args = append(args, thing1)
	args = append(args, thing2)
	args = append(args, author)
	args = append(args, n.Collective)
	args = append(args, n.Domain)
	args = append(args, n.Kind)
	args = append(args, n.Path)
	return name, args, nil
}
