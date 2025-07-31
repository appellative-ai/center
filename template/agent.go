package template

import (
	"errors"
	"fmt"
	"github.com/appellative-ai/center/exchange"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/postgres/request"
	"net/http"
	"time"
)

const (
	NamespaceName = "common:core:agent/namespace/templatecenter"
	duration      = time.Second * 30
	timeout       = time.Second * 4

	defaultGetName = "common:core:namespace/query"
)

// TODO: need to integrate a namespace that has an implied name, getDefaultName, as the GET arguments
//       are in the query string, not the request body

var (
	agent *agentT
)

type agentT struct {
	running bool
	timeout time.Duration

	ticker   *messaging.Ticker
	emissary *messaging.Channel

	requester *request.Interface
}

// init - register an agent constructor
func init() {
	exchange.RegisterConstructor(NamespaceName, func() messaging.Agent {
		agent = newAgent(request.Requester)
		return agent
	})
}

func newAgent(requester *request.Interface) *agentT {
	a := new(agentT)
	a.timeout = timeout
	a.ticker = messaging.NewTicker(messaging.ChannelEmissary, duration)
	a.emissary = messaging.NewEmissaryChannel()

	a.requester = requester
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
