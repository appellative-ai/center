package expansion

type arg struct {
	//Name     string `json:"name"`
	Nullable bool
	Type     string `json:"type"`
	SqlType  string `json:"sql-type"`
}
type template struct {
	Name string         `json:"name"`
	Args map[string]arg `json:"args"`
}

type inputArg struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
}

type template2 struct {
	Sql  string `json:"sql"`
	Args []arg  `json:"args"`
}
