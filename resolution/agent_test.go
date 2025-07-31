package resolution

import "fmt"

func ExampleNewAgent() {
	a := newAgent(nil, nil, nil)

	fmt.Printf("test: newAgent() -> [%v]\n", a)

	//Output:
	//test: newAgent() -> [common:core:agent/namespace/center]

}
