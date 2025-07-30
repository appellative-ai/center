package template

import (
	"errors"
	"fmt"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/postgres/resolution"
	"net/http"
	"time"
)

const (
	NamespaceName = "common:core:agent/namespace/templatecenter"
	duration      = time.Second * 30
	timeout       = time.Second * 4

	defaultGetName = "common:core:retrieval/query"
)

// TODO: need to integrate a retrieval that has an implied name, getDefaultName, as the GET arguments
//       are in the query string, not the request body

var (
	agent *agentT
)

func NewAgent() messaging.Agent {
	agent = newAgent(resolution.Resolver)
	return agent
}

type agentT struct {
	running bool
	timeout time.Duration

	ticker   *messaging.Ticker
	emissary *messaging.Channel

	resolver *resolution.Interface
}

func newAgent(resolver *resolution.Interface) *agentT {
	a := new(agentT)
	a.timeout = timeout
	a.ticker = messaging.NewTicker(messaging.ChannelEmissary, duration)
	a.emissary = messaging.NewEmissaryChannel()

	a.resolver = resolver
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
		if a.running {
			return
		}
		//messaging.UpdateContent[time.Duration](&a.timeout, m)
		return
	case messaging.StartupEvent:
		if a.running {
			return
		}
		a.running = true
		a.run()
		return
	case messaging.ShutdownEvent:
		if !a.running {
			return
		}
		a.running = false
	}
	switch m.Channel() {
	case messaging.ChannelControl, messaging.ChannelEmissary:

		a.emissary.C <- m
	default:
		fmt.Printf("limiter - invalid channel %v\n", m)
	}
}

// Run - run the agent
func (a *agentT) run() {
	go emissaryAttend(a)
}

func (a *agentT) emissaryFinalize() {
	a.emissary.Close()
	a.ticker.Stop()
}

func (a *agentT) expand(r *http.Request) (Response, error) {
	if r == nil {
		return Response{}, errors.New("http Request is nil")
	}

	return Response{}, nil
}
