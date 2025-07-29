package expansion

type Interface struct {
	Params func(name string, t string) (sql string, args []any, err error)
}

// Expander -
var Expander = func() *Interface {
	return &Interface{
		Params: func(name string, t string) (sql string, args []any, err error) {

			return "", nil, nil
		},
	}
}()
