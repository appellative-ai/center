package templatetest

import "fmt"

func ExampleNewAgent() {
	a := NewAgent()

	fmt.Printf("test: newAgent() -> [%v]\n", a)

	//Output:
	//test: newAgent() -> [common:core:agent/template/center/test]

}
