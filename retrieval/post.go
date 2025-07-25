package retrieval

import (
	"bytes"
	"context"
	"net/http"
)

func post(agent *agentT, ctx context.Context, req *http.Request) (bytes.Buffer, error) {
	var buf bytes.Buffer
	sql, args, err1 := agent.expander.Params("", req.URL.Query())
	if err1 != nil {
		return buf, err1
	}
	return agent.retriever.Marshal(ctx, "", sql, args)
}
