//go:generate mockgen -source=./broker.go -destination=../../mocks/broker_infra_mock.go -package=mocks -mock_names=Usecase=MockNrokerUsecase

package broker

import (
	"github.com/nats-io/nats.go"
	"github.com/netorissi/SwordTest/config"
	"github.com/netorissi/SwordTest/shared"
)

type MessageBroker interface {
	Publish(topic string, msg []byte) error
	Subscribe(topic string, f func(m *nats.Msg)) error
}

type impl struct {
	broker *nats.Conn
}

func New() MessageBroker {
	nc, err := nats.Connect(config.Global.MessageBroker.DefaultURL)
	if err != nil {
		shared.Logger.DPanic(err.Error())
	}

	shared.Logger.Info("Message broker connected.")

	return &impl{broker: nc}
}

func (i *impl) Publish(topic string, msg []byte) error {
	return i.broker.Publish(topic, msg)
}

func (i *impl) Subscribe(topic string, f func(m *nats.Msg)) error {
	_, err := i.broker.Subscribe(topic, f)
	return err
}
