package namespace

import (
	"context"
	"fmt"
	"github.com/appellative-ai/core/std"
	"github.com/appellative-ai/postgres/request"
	"github.com/appellative-ai/postgres/request/requesttest"
	"net/url"
	"reflect"
	"testing"
)

func ExampleCreateThing() {
	values := make(url.Values)

	name, args, err := createThingArgs(values)
	fmt.Printf("test: Thing() -> [name:%v] [args:%v] [err:%v]\n", name, args, err)

	values.Add(nameName, "common:resiliency:agent/rate-limiting/request/http")
	name, args, err = createThingArgs(values)
	fmt.Printf("test: Thing() -> [name:%v] [args:%v] [err:%v]\n", name, args, err)

	values.Add(authorName, "bobs uncle")
	name, args, err = createThingArgs(values)
	fmt.Printf("test: Thing() -> [name:%v] [args:%v] [err:%v]\n", name, args, err)

	//Output:
	//test: Thing() -> [name:] [args:[]] [err:name is empty]
	//test: Thing() -> [name:] [args:[]] [err:author is empty]
	//test: Thing() -> [name:common:resiliency:agent/rate-limiting/request/http] [args:[common:resiliency:agent/rate-limiting/request/http bobs uncle common resiliency agent /rate-limiting/request/http]] [err:<nil>]

}

func ExampleThingRequest() {
	name := "common:resiliency:agent/rate-limiting/request/http"

	r, err := thingRequest(nil, nil, nil)
	fmt.Printf("test: thingRequest() -> [result:%v] [err:%v]\n", r, err)

	m := std.NewSyncMap[string, any]()
	requester := requesttest.NewRequester(m)
	values := make(url.Values)
	values.Add(nameName, name)
	values.Add(authorName, "bobs uncle")

	r, err = thingRequest(nil, requester, values)
	fmt.Printf("test: thingRequest() -> [result:%v] [err:%v]\n", r, err)

	t, ok := m.Load(name)
	fmt.Printf("test: Load() -> [%v] [ok:%v]\n", t, ok)

	//Output:
	//test: thingRequest() -> [result:{ 0 false false false false}] [err:query values are nil]
	//test: thingRequest() -> [result:{ 1 true false false false}] [err:<nil>]
	//test: Load() -> [[common:resiliency:agent/rate-limiting/request/http bobs uncle common resiliency agent /rate-limiting/request/http]] [ok:true]

}

func Test_thingRequest(t *testing.T) {
	type args struct {
		ctx       context.Context
		requester *request.Interface
		values    url.Values
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
			got, err := thingRequest(tt.args.ctx, tt.args.requester, tt.args.values)
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
