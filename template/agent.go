package template

import (
	"errors"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/std"
	"github.com/appellative-ai/postgres/retrieval"
	"time"
)

const (
	NamespaceName = "common:core:agent/template/center"
	timeout       = time.Second * 4
	defaultSql    = "CALL dbo.Representation($1)"
)

type Agent interface {
	messaging.Agent
	Build(name string, args []Arg) (Result, error)
}

var (
	agent *agentT
)

type agentT struct {
	timeout   time.Duration
	cache     *std.MapT[string, template]
	retriever *retrieval.Interface
}

func NewAgent(retriever *retrieval.Interface) Agent {
	a := new(agentT)
	a.timeout = timeout
	a.cache = std.NewSyncMap[string, template]()
	a.retriever = retriever
	return a
}

// String - identity
func (a *agentT) String() string { return a.Name() }

// Name - agent identifier
func (a *agentT) Name() string { return NamespaceName }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	switch m.Name {
	case messaging.ConfigEvent:
		messaging.UpdateContent[time.Duration](&a.timeout, m)
		return
	}
}

func (a *agentT) Build(name string, args []Arg) (Result, error) {
	if name == "" {
		return Result{}, errors.New("name is empty")
	}
	if len(args) == 0 {
		return Result{}, errors.New("arguments are empty")
	}
	t := a.cache.Load(name)
	if t.Name == "" {
		var err error
		t, err = a.add(name)
		if err != nil {
			return Result{}, err
		}
	}
	newArgs, err := build(args, t.Params)
	if err != nil {
		return Result{}, err
	}
	return Result{Sql: t.Sql, Args: newArgs}, nil
}
