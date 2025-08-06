package namespace

import (
	"context"
	"errors"
	"github.com/appellative-ai/center/template"
	"github.com/appellative-ai/core/jsonx"
	"github.com/appellative-ai/postgres/request"
	"net/http"
)

type tagLink struct {
	Name   string `json:"name"`
	CName  string `json:"cname"`
	Thing1 string `json:"thing1"`
	Thing2 string `json:"thing2"`
	Author string `json:"author"`
}

func linkRequest(ctx context.Context, requester *request.Interface, processor template.Agent, r *http.Request) (request.Result, error) {
	if r == nil {
		return request.Result{}, errors.New("request is nil")
	}
	t, err := jsonx.New[tagLink](r.Body, nil)
	if err != nil {
		return request.Result{}, err
	}
	res, err1 := processor.Build(t.Name, []template.Arg{
		{Name: nameName, Value: t.Name},
		{Name: cNameName, Value: t.CName},
		{Name: thing1Name, Value: t.Thing1},
		{Name: thing2Name, Value: t.Thing2},
		{Name: authorName, Value: t.Author},
	})
	if err1 != nil {
		return request.Result{}, err
	}
	return requester.Execute(ctx, t.Name, res.Sql, res.Args)
}
