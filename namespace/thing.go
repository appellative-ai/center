package namespace

import (
	"context"
	"errors"
	"github.com/appellative-ai/center/template"
	"github.com/appellative-ai/core/jsonx"
	"github.com/appellative-ai/postgres/request"
	"net/http"
)

const (
	nameName   = "name"
	cNameName  = "cname"
	authorName = "author"
	thing1Name = "thing1"
	thing2Name = "thing2"
)

type tagThing struct {
	Name   string `json:"name"`
	CName  string `json:"cname"`
	Author string `json:"author"`
}

func thingRequest(ctx context.Context, requester *request.Interface, processor template.Agent, r *http.Request) (request.Result, error) {
	if r == nil {
		return request.Result{}, errors.New("request is nil")
	}
	t, err := jsonx.New[tagThing](r.Body, nil)
	if err != nil {
		return request.Result{}, err
	}
	res, err1 := processor.Build(t.Name, []template.Arg{
		{Name: nameName, Value: t.Name},
		{Name: cNameName, Value: t.CName},
		{Name: authorName, Value: t.Author},
	})
	if err1 != nil {
		return request.Result{}, err
	}
	return requester.Execute(ctx, t.Name, res.Sql, res.Args)
}
