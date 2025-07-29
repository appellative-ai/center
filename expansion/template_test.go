package expansion

import (
	"encoding/json"
	"fmt"
)

func ExampleArgs() {
	a := template{Name: "CALL dbo.QueryNamespace($1,$2,$3)", Args: map[string]arg{
		"name":       {Nullable: true, Type: "string", SqlType: ""},
		"count":      {Nullable: false, Type: "int", SqlType: ""},
		"createDate": {Nullable: false, Type: "string", SqlType: "DateTime"},
	}}

	fmt.Printf("test: template() -> [sql:%v] [args:%v]\n", a.Name, a.Args)
	buf, err := json.Marshal(a)
	fmt.Printf("test: template() -> [%v] [err:%v]\n", string(buf), err)

	//Output:
	//fail
}

/*
func ExampleArgs2() {
	a := template{Sql: "sp.QueryNamespace", Args: []arg{
		{Name: "name", Value: "test:agent"},
		{Name: "count", Value: 123},
	},
	}
	fmt.Printf("test: template() -> [sql:%v] [args:%v]\n", a.Sql, a.Args)
	buf, err := json.Marshal(a)
	fmt.Printf("test: template() -> [%v] [err:%v]\n", string(buf), err)

	//Output:
	//fail
}


*/
