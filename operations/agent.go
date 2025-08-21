package operations

import (
	"fmt"
	"github.com/appellative-ai/common/messaging"
	"time"
)

const (
	AgentName = "common:core:agent/operations/collective"
	duration  = time.Second * 30
)

var (
	agent *agentT
)

type agentT struct {
	running  bool
	ticker   *messaging.Ticker
	emissary *messaging.Channel
}

func init() {
	newAgent()
}

func newAgent() *agentT {
	a := new(agentT)
	agent = a

	a.ticker = messaging.NewTicker(messaging.ChannelEmissary, duration)
	a.emissary = messaging.NewEmissaryChannel()

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
	switch m.Name {
	case messaging.ConfigEvent:
		if a.running {
			return
		}
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

func (a *agentT) message(m *messaging.Message) {
	if m == nil {
		return
	}

}
