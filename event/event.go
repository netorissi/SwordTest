package event

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/netorissi/SwordTest/infra/broker"
)

const (
	EVENT_NOTIFY_MANAGERS = "event.notify.managers"
)

func New(nc broker.MessageBroker) {

	nc.Subscribe(EVENT_NOTIFY_MANAGERS, func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
}
