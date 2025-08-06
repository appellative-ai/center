package namespace

import (
	"encoding/json"
	"fmt"
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
