package retrieval

import (
	"bytes"
	"context"
	"net/http"
)

func (a *agentT) get(ctx context.Context, req *http.Request) (bytes.Buffer, error) {
	var buf bytes.Buffer
	sql, args, err1 := a.expander.Params("", req.URL.Query())
	if err1 != nil {
		return buf, err1
	}
	return a.retriever.Marshal(ctx, "", sql, args)
}
