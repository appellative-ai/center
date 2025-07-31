package namespace

import (
	"encoding/json"
	"fmt"
)

func ExampleRetrieval() {
	r := tagRetrieval{
		Name: "test:agent",

		Args: []arg{
			{Name: "kind", Value: "aspect"},
			{Name: "count", Value: "25"},
		},
	}

	buf, err := json.Marshal(r)

	fmt.Printf("test: Retrieval() -> [%v] [err:%v]\n", string(buf), err)

	//Output:
	//test: Retrieval() -> [{"name":"test:agent","args":[{"name":"kind","value":"aspect"},{"name":"count","value":"25"}]}] [err:<nil>]

}
