package template

import "fmt"

const (
	fileName = "file://[cwd]/templatetest/template.json"
	name     = "common:core:retrieval/test"
)

func ExampleAddEntry() {
	agent := newAgent(nil)

	err := AddEntry(agent, fileName)
	fmt.Printf("test: AddEntry() -> [err:%v]\n", err)

	t := agent.cache.Load(name)
	fmt.Printf("test: Entry() -> [%v]\n", t)

	//Output:
	//test: AddEntry() -> [err:<nil>]
	//test: Entry() -> [{common:core:retrieval/test CALL dbo.QueryNamespace($1,$2,$3) [{name true string } {count false int } {createDate false string DateTime}]}]

}
