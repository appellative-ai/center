package namespace

import (
	"encoding/json"
	"fmt"
)

func ExampleArgs() {
	a := template{Name: "CALL dbo.QueryNamespace($1,$2,$3)", Params: []param{
		{Name: "name", Nullable: true, Type: "string", SqlType: ""},
		{Name: "count", Nullable: false, Type: "int", SqlType: ""},
		{Name: "createDate", Nullable: false, Type: "string", SqlType: "DateTime"},
	},
	}

	fmt.Printf("test: template() -> [sql:%v] [args:%v]\n", a.Name, a.Params)
	buf, err := json.Marshal(a)
	fmt.Printf("test: template() -> [%v] [err:%v]\n", string(buf), err)

	//Output:
	//test: template() -> [sql:CALL dbo.QueryNamespace($1,$2,$3)] [args:[{name true string } {count false int } {createDate false string DateTime}]]
	//test: template() -> [{"name":"CALL dbo.QueryNamespace($1,$2,$3)","args":[{"name":"name","nullable":true,"type":"string","sql-type":""},{"name":"count","nullable":false,"type":"int","sql-type":""},{"name":"createDate","nullable":false,"type":"string","sql-type":"DateTime"}]}] [err:<nil>]

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

func ExampleBuild_Sort() {
	args := []arg{
		{Name: "test5", Value: "string"},
		{Name: "TESt1", Value: "string"},
	}
	params := []param{
		{Name: "test5"},
		{Name: "TESt4"},
	}

	build(args, params)

	fmt.Printf("test: build() [args:%v] [params:%v]\n", args, params)

	//Output:
	//test: build() [args:[{TESt1 string} {test5 string}]] [params:[{TESt4 false  } {test5 false  }]]

}

func ExampleSlice() {

	params := []param{
		{Name: "test5"},
		{Name: "TESt4"},
	}

	fmt.Printf("test: build() [params:%v]\n", params)
	appendValue(params)
	fmt.Printf("test: build() [params:%v]\n", params)
	appendPointer(&params)
	fmt.Printf("test: build() [params:%v]\n", params)

	//Output:
	//test: build() [params:[{test5 false  } {TESt4 false  }]]
	//test: build() [params:[{test5 false  } {TESt4 false  }]]
	//test: build() [params:[{test5 false  } {TESt4 false  } {new item false  }]]
	
}

func appendValue(params []param) {
	params = append(params, param{Name: "new item"})
}

func appendPointer(params *[]param) {
	*params = append(*params, param{Name: "new item"})
}
