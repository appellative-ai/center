package notification

import "fmt"

func ExampleNewAgent() {
	a := newAgent(nil, nil)

	fmt.Printf("test: newAgent() -> [%v]\n", a)

	//Output:
	//test: newAgent() -> [common:core:agent/notification/center]

}
