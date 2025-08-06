package templatetest

import (
	"errors"
	"fmt"
	"github.com/appellative-ai/center/template"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/std"
)

const (
	AgentName = "common:core:agent/template/center/test"
)

type agentT struct {
	cache *std.MapT[string, template.Entry]
}

func NewAgent() template.Agent {
	return newAgent()
}

func newAgent() *agentT {
	a := new(agentT)
	a.cache = std.NewSyncMap[string, template.Entry]()
	return a
}

// String - identity
func (a *agentT) String() string { return a.Name() }

// Name - agent identifier
func (a *agentT) Name() string { return AgentName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
}

func (a *agentT) Add(entry template.Entry) {
	if entry.Name == "" {
		return
	}
	a.cache.Store(entry.Name, entry)
}

func (a *agentT) Build(name string, args []template.Arg) (template.Result, error) {
	if name == "" {
		return template.Result{}, errors.New("name is empty")
	}
	if len(args) == 0 {
		return template.Result{}, errors.New("arguments are empty")
	}
	t := a.cache.Load(name)
	if t.Name == "" {
		return template.Result{}, errors.New(fmt.Sprintf("template [%v] not found", name))
	}
	newArgs, err := template.Build(args, t.Params)
	if err != nil {
		return template.Result{}, err
	}
	return template.Result{Sql: t.Sql, Args: newArgs}, nil
}
