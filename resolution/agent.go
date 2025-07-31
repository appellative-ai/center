package resolution

import (
	"fmt"
	"github.com/appellative-ai/center/exchange"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
	"github.com/appellative-ai/postgres/request"
	"github.com/appellative-ai/postgres/retrieval"
	"net/http"
	"time"
)

const (
	NamespaceName = "common:core:agent/resolution/center"
	duration      = time.Second * 30
	timeout       = time.Second * 4
)

var (
	agent *agentT
)

type agentT struct {
	running bool
	timeout time.Duration

	ticker   *messaging.Ticker
	emissary *messaging.Channel

	retriever *retrieval.Interface
	requester *request.Interface
}

// init - register an agent constructor
func init() {
	exchange.RegisterConstructor(NamespaceName, func() messaging.Agent {
		agent = newAgent(retrieval.Retriever, request.Requester)
		return agent
	})
}

func newAgent(retriever *retrieval.Interface, requester *request.Interface) *agentT {
	a := new(agentT)
	a.timeout = timeout
	a.ticker = messaging.NewTicker(messaging.ChannelEmissary, duration)
	a.emissary = messaging.NewEmissaryChannel()
	a.retriever = retriever
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
		messaging.UpdateContent[time.Duration](&a.timeout, m)
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
	//go emissaryAttend(a)
}

func (a *agentT) emissaryFinalize() {
	a.emissary.Close()
	a.ticker.Stop()
}

// Link - chainable exchange
func (a *agentT) Link(next rest.Exchange) rest.Exchange {
	return func(req *http.Request) (resp *http.Response, err error) {
		//ctx, cancel := httpx.NewContext(nil, a.timeout)
		//defer cancel()

		/*
			buf, err1 := a.retriever.Marshal(ctx, ""thingQueryName, "select * from thing", nil)
			if err1 != nil {
				return httpx.NewResponse(messaging.StatusExecError, nil, err1), err1
			}
			return httpx.NewResponse(http.StatusOK, nil, buf), nil

		*/
		return
	}
}
