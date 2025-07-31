package namespace

import (
	"bytes"
	"context"
	"net/http"
)

func (a *agentT) post(ctx context.Context, r *http.Request) (bytes.Buffer, error) {
	var buf bytes.Buffer
	resp, err := a.expand(r)
	if err != nil {
		return buf, err
	}
	return a.retriever.Marshal(ctx, resp.Name, resp.Sql, resp.Args)
}
