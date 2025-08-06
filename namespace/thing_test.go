package namespace

import (
	"encoding/json"
	"fmt"
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
