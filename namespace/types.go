package namespace

type arg struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
}

type tagRelation struct {
	Name string `json:"name"`
	Args []arg  `json:"args"`
}

type tagRetrieval struct {
	Name string `json:"name"`
	Args []arg  `json:"args"`
}

type tagThing struct {
	Name   string `json:"name"`
	CName  string `json:"cname"`
	Author string `json:"author"`
}

type tagLink struct {
	Name   string `json:"name"`
	CName  string `json:"cname"`
	Thing1 string `json:"thing1"`
	Thing2 string `json:"thing2"`
	Author string `json:"author"`
}
