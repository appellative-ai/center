package namespace

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/appellative-ai/center/template"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
	"github.com/appellative-ai/postgres/request"
	"github.com/appellative-ai/postgres/retrieval"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func ExampleThing() {
	r := tagThing{
		Name:   "test:agent",
		CName:  "headers",
		Author: "core:person/bob",
	}

	buf, err := json.Marshal(r)

	fmt.Printf("test: Thing() -> [%v] [err:%v]\n", string(buf), err)

	//Output:
	//test: Thing() -> [{"name":"test:agent","cname":"headers","author":"core:person/bob"}] [err:<nil>]

}

func Test_agentT_Link(t *testing.T) {
	type fields struct {
		running   bool
		timeout   time.Duration
		ticker    *messaging.Ticker
		emissary  *messaging.Channel
		retriever *retrieval.Interface
		requester *request.Interface
		processor template.Agent
	}
	type args struct {
		next rest.Exchange
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   rest.Exchange
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &agentT{
				running:   tt.fields.running,
				timeout:   tt.fields.timeout,
				ticker:    tt.fields.ticker,
				emissary:  tt.fields.emissary,
				retriever: tt.fields.retriever,
				requester: tt.fields.requester,
				processor: tt.fields.processor,
			}
			if got := a.Link(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Link() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_agentT_Message(t *testing.T) {
	type fields struct {
		running   bool
		timeout   time.Duration
		ticker    *messaging.Ticker
		emissary  *messaging.Channel
		retriever *retrieval.Interface
		requester *request.Interface
		processor template.Agent
	}
	type args struct {
		m *messaging.Message
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &agentT{
				running:   tt.fields.running,
				timeout:   tt.fields.timeout,
				ticker:    tt.fields.ticker,
				emissary:  tt.fields.emissary,
				retriever: tt.fields.retriever,
				requester: tt.fields.requester,
				processor: tt.fields.processor,
			}
			a.Message(tt.args.m)
		})
	}
}

func Test_agentT_Name(t *testing.T) {
	type fields struct {
		running   bool
		timeout   time.Duration
		ticker    *messaging.Ticker
		emissary  *messaging.Channel
		retriever *retrieval.Interface
		requester *request.Interface
		processor template.Agent
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &agentT{
				running:   tt.fields.running,
				timeout:   tt.fields.timeout,
				ticker:    tt.fields.ticker,
				emissary:  tt.fields.emissary,
				retriever: tt.fields.retriever,
				requester: tt.fields.requester,
				processor: tt.fields.processor,
			}
			if got := a.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_agentT_String(t *testing.T) {
	type fields struct {
		running   bool
		timeout   time.Duration
		ticker    *messaging.Ticker
		emissary  *messaging.Channel
		retriever *retrieval.Interface
		requester *request.Interface
		processor template.Agent
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &agentT{
				running:   tt.fields.running,
				timeout:   tt.fields.timeout,
				ticker:    tt.fields.ticker,
				emissary:  tt.fields.emissary,
				retriever: tt.fields.retriever,
				requester: tt.fields.requester,
				processor: tt.fields.processor,
			}
			if got := a.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_agentT_emissaryFinalize(t *testing.T) {
	type fields struct {
		running   bool
		timeout   time.Duration
		ticker    *messaging.Ticker
		emissary  *messaging.Channel
		retriever *retrieval.Interface
		requester *request.Interface
		processor template.Agent
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &agentT{
				running:   tt.fields.running,
				timeout:   tt.fields.timeout,
				ticker:    tt.fields.ticker,
				emissary:  tt.fields.emissary,
				retriever: tt.fields.retriever,
				requester: tt.fields.requester,
				processor: tt.fields.processor,
			}
			a.emissaryFinalize()
		})
	}
}

func Test_agentT_relation(t *testing.T) {
	type fields struct {
		running   bool
		timeout   time.Duration
		ticker    *messaging.Ticker
		emissary  *messaging.Channel
		retriever *retrieval.Interface
		requester *request.Interface
		processor template.Agent
	}
	type args struct {
		ctx context.Context
		r   *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bytes.Buffer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &agentT{
				running:   tt.fields.running,
				timeout:   tt.fields.timeout,
				ticker:    tt.fields.ticker,
				emissary:  tt.fields.emissary,
				retriever: tt.fields.retriever,
				requester: tt.fields.requester,
				processor: tt.fields.processor,
			}
			got, err := a.relation(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("relation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("relation() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_agentT_retrieval(t *testing.T) {
	type fields struct {
		running   bool
		timeout   time.Duration
		ticker    *messaging.Ticker
		emissary  *messaging.Channel
		retriever *retrieval.Interface
		requester *request.Interface
		processor template.Agent
	}
	type args struct {
		ctx context.Context
		r   *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bytes.Buffer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &agentT{
				running:   tt.fields.running,
				timeout:   tt.fields.timeout,
				ticker:    tt.fields.ticker,
				emissary:  tt.fields.emissary,
				retriever: tt.fields.retriever,
				requester: tt.fields.requester,
				processor: tt.fields.processor,
			}
			got, err := a.retrieval(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("retrieval() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("retrieval() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_agentT_run(t *testing.T) {
	type fields struct {
		running   bool
		timeout   time.Duration
		ticker    *messaging.Ticker
		emissary  *messaging.Channel
		retriever *retrieval.Interface
		requester *request.Interface
		processor template.Agent
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &agentT{
				running:   tt.fields.running,
				timeout:   tt.fields.timeout,
				ticker:    tt.fields.ticker,
				emissary:  tt.fields.emissary,
				retriever: tt.fields.retriever,
				requester: tt.fields.requester,
				processor: tt.fields.processor,
			}
			a.run()
		})
	}
}

func Test_agentT_thingRequest(t *testing.T) {
	type fields struct {
		running   bool
		timeout   time.Duration
		ticker    *messaging.Ticker
		emissary  *messaging.Channel
		retriever *retrieval.Interface
		requester *request.Interface
		processor template.Agent
	}
	type args struct {
		ctx context.Context
		r   *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    request.Result
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &agentT{
				running:   tt.fields.running,
				timeout:   tt.fields.timeout,
				ticker:    tt.fields.ticker,
				emissary:  tt.fields.emissary,
				retriever: tt.fields.retriever,
				requester: tt.fields.requester,
				processor: tt.fields.processor,
			}
			got, err := a.thingRequest(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("thingRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("thingRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_emissaryAttend(t *testing.T) {
	type args struct {
		a *agentT
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			emissaryAttend(tt.args.a)
		})
	}
}

func Test_newAgent(t *testing.T) {
	tests := []struct {
		name string
		want *agentT
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newAgent(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newAgent() = %v, want %v", got, tt.want)
			}
		})
	}
}
