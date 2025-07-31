package namespace

import (
	"encoding/json"
	"fmt"
)

func ExampleRelation() {
	r := tagRelation{
		Name: "test:agent",
		//Instance: "core:aspect/resiliency",
		//Pattern:  "core:aspect/expressive",
		Args: []arg{
			{Name: "instance", Value: "core:aspect/resiliency"},
			{Name: "pattern", Value: "core:aspect/expressive"},
			{Name: "kind", Value: "aspect"},
			{Name: "count", Value: "25"},
		},
	}

	buf, err := json.Marshal(r)

	fmt.Printf("test: Relation() -> [%v] [err:%v]\n", string(buf), err)

	//Output:
	//test: Relation() -> [{"name":"test:agent","instance":"core:aspect/resiliency","pattern":"core:aspect/expressive","args":[{"name":"kind","value":"aspect"},{"name":"count","value":25}]}] [err:<nil>]

}
