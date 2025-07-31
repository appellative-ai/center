package namespace

import (
	"bytes"
	"context"
	"net/http"
)

type tagThing struct {
	Name   string `json:"name"`
	CName  string `json:"cname"`
	Author string `json:"author"`
}

type tagLink struct {
	Name   string `json:"name"`
	CName  string `json:"cname"`
	Thing1 string `json:"thing1"`
	Thing2 string `json:"thing2"`
	Author string `json:"author"`
}

func (a *agentT) linkRequest(ctx context.Context, r *http.Request) (bytes.Buffer, error) {
	var buf bytes.Buffer
	name := ""
	res, err := a.processor.Build(name, nil)
	if err != nil {
		return buf, err
	}
	return a.retriever.Marshal(ctx, name, res.Sql, res.Args)
}

func (a *agentT) thingRequest(ctx context.Context, r *http.Request) (bytes.Buffer, error) {
	var buf bytes.Buffer
	name := ""
	res, err := a.processor.Build(name, nil)
	if err != nil {
		return buf, err
	}
	return a.retriever.Marshal(ctx, name, res.Sql, res.Args)
}
