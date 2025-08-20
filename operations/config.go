package operations

import (
	"github.com/appellative-ai/core/messaging"
	"github.com/appellative-ai/core/std"
)

func (a *agentT) configure(m *messaging.Message) {
	switch m.ContentType() {
	case messaging.ContentTypeMap:
		//cfg, status := messaging.MapContent(m)
		//if !status.OK() {
		//	messaging.Reply(m, messaging.EmptyMapError(a.Name()), a.Name())
		//	return
		//}
		//a.state = initialize(cfg)
		// Initialize linked collectives

	}
	messaging.Reply(m, std.StatusOK, a.Name())
}
