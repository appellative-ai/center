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

func ExampleLink() {
	r := tagLink{
		Name:   "test:agent",
		CName:  "headers",
		Author: "core:person/bob",
		Thing1: "core:thing/first",
		Thing2: "core:thing/second",
	}

	buf, err := json.Marshal(r)

	fmt.Printf("test: Link() -> [%v] [err:%v]\n", string(buf), err)

	//Output:
	//test: Link() -> [{"name":"test:agent","cname":"headers","thing1":"core:thing/first","thing2":"core:thing/second","author":"core:person/bob"}] [err:<nil>]

}

func Test_linkRequest(t *testing.T) {
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
			got, err := linkRequest(tt.args.ctx, tt.args.requester, tt.args.processor, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("linkRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("linkRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}
