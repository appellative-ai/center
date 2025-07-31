package namespace

import (
	"bytes"
	"context"
	"net/http"
)

type tagRetrieval struct {
	Name string `json:"name"`
	Args []arg  `json:"args"`
}

func (a *agentT) retrieval(ctx context.Context, r *http.Request) (bytes.Buffer, error) {
	var buf bytes.Buffer
	name := ""
	res, err := a.processor.Build(name, nil)
	if err != nil {
		return buf, err
	}
	return a.retriever.Marshal(ctx, name, res.Sql, res.Args)
}
