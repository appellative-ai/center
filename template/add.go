package template

import "errors"

func (a *agentT) add(name string) (template, error) {
	if name == "" {
		return template{}, errors.New("name is empty")
	}
	t, err := a.get(name)
	if err != nil {
		return template{}, err
	}
	a.cache.Store(name, t)
	return t, nil
}
