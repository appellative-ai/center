package operations

import (
	"github.com/appellative-ai/common/messaging"
)

func (a *agentT) configure(m *messaging.Message) {
	if m == nil || m.Name != messaging.ConfigEvent {
		return
	}
	/*
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
		messaging.Reply(m, 0, a.Name())

	*/
}
