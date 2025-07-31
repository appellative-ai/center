package namespace

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/appellative-ai/center/exchange"
	"github.com/appellative-ai/center/template"
	"github.com/appellative-ai/core/httpx"
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/rest"
	"github.com/appellative-ai/postgres/request"
	"github.com/appellative-ai/postgres/retrieval"
	"net/http"
	"time"
)

const (
	NamespaceName = "common:core:agent/namespace/center"
	duration      = time.Second * 30
	timeout       = time.Second * 4

	retrievalPath    = "/namespace/retrieval"
	relationPath     = "/namespace/relation"
	requestThingPath = "/namespace/request/thing"
	requestLinkPath  = "/namespace/request/link"
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
	processor template.Agent
}

// init - register an agent constructor
func init() {
	exchange.RegisterConstructor(NamespaceName, func() messaging.Agent {
		agent = newAgent(retrieval.Retriever, request.Requester, template.NewAgent(retrieval.Retriever))
		return agent
	})
}

func newAgent(retriever *retrieval.Interface, requester *request.Interface, processor template.Agent) *agentT {
	a := new(agentT)
	a.timeout = timeout
	a.ticker = messaging.NewTicker(messaging.ChannelEmissary, duration)
	a.emissary = messaging.NewEmissaryChannel()

	a.retriever = retriever
	a.requester = requester
	a.processor = processor
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
	go emissaryAttend(a)
}

func (a *agentT) emissaryFinalize() {
	a.emissary.Close()
	a.ticker.Stop()
}

// Link - chainable exchange
func (a *agentT) Link(next rest.Exchange) rest.Exchange {
	return func(req *http.Request) (resp *http.Response, err error) {
		ctx, cancel := httpx.NewContext(nil, a.timeout)
		defer cancel()
		var buf bytes.Buffer

		switch req.URL.Path {
		case retrievalPath:
			buf, err = a.retrieval(ctx, req)
		case relationPath:
			buf, err = a.relation(ctx, req)
		case requestThingPath:
			buf, err = a.thingRequest(ctx, req)
		case requestLinkPath:
			buf, err = a.linkRequest(ctx, req)
		default:
			return httpx.NewResponse(http.StatusBadRequest, nil, nil), errors.New(fmt.Sprintf("path is invalid [%v]", req.URL.Path))
		}
		if err != nil {
			return httpx.NewResponse(http.StatusInternalServerError, nil, nil), err
		}
		return httpx.NewResponse(http.StatusOK, nil, buf), nil
	}
}
