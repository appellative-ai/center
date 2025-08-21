package operations

import (
	"github.com/appellative-ai/core/messaging"
	sqlops "github.com/appellative-ai/postgres/operations"
	"time"
)

// TODO: need a private message queue for internal messaging.
//       Maybe a control channel

func ConfigDatabaseClient(cfg map[string]string) error {
	return sqlops.ConfigClient(cfg)
}

func ConfigDatabaseLogging(log func(start time.Time, duration time.Duration, route string, req any, resp any, timeout time.Duration)) {
	sqlops.ConfigLogging(log)
}

func ConfigDatabaseSource() {
	sqlops.ConfigSourceOverride()
}

func Startup(msg *messaging.Message) {
	agent.Message(msg)
	agent.Message(messaging.StartupMessage)
}
