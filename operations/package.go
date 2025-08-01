package operations

import "github.com/appellative-ai/core/messaging"

const (
	RegistryHost1Key = "registry-host1"
	RegistryHost2Key = "registry-host2"
)

func Startup(msg *messaging.Message) {
	agent.Message(msg)
	agent.Message(messaging.StartupMessage)
}
