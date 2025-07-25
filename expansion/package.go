package expansion

type Interface struct {
	Params func(name string, t any) (string, []any, error)
}

// Expander -
var Expander = func() *Interface {
	return &Interface{
		Params: func(name string, t any) (string, []any, error) {

			return "", nil, nil
		},
	}
}()
