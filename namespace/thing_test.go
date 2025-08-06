package namespace

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/appellative-ai/center/template"
	"github.com/appellative-ai/postgres/request"
	"net/http"
	"reflect"
	"testing"
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

func Test_thingRequest(t *testing.T) {
	type args struct {
		ctx       context.Context
		requester *request.Interface
		processor template.Agent
		r         *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    request.Result
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := thingRequest(tt.args.ctx, tt.args.requester, tt.args.processor, tt.args.r)
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
