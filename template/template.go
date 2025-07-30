package template

type arg struct {
	Name     string `json:"name"`
	Nullable bool   `json:"nullable"`
	Type     string `json:"type"`
	SqlType  string `json:"sql-type"`
}
type template struct {
	Name string `json:"name"`
	Args []arg  `json:"args"`
}
