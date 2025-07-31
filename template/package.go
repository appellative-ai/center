package template

type Arg struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Result struct {
	Sql  string
	Args []any
}
