package resource

import "github.com/appellative-ai/core/messaging"

const (
// ServiceKind = "service"
)

type Service struct {
	Representation    func(name string) (messaging.Content, *messaging.Status)
	Context           func(name string) (messaging.Content, *messaging.Status)
	AddRepresentation func(name, author, contentType string, value any) *messaging.Status
	AddContext        func(name, author, contentType string, value any) *messaging.Status
}

// Notification - notification interface
type Notification struct {
	Message func(msg *messaging.Message) bool
	Advise  func(msg *messaging.Message) *messaging.Status
	Trace   func(name, task, observation, action string)
}

// Notifier -
var Notifier = func() *Notification {
	return &Notification{
		Message: func(msg *messaging.Message) bool {
			return true
		},
		Advise: func(msg *messaging.Message) *messaging.Status {
			return messaging.StatusOK()
		},
		Trace: func(name, task, observation, action string) {
		},
	}
}()
