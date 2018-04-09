package qwik

import (
	proto "github.com/golang/protobuf/proto"
	nats "github.com/nats-io/go-nats"
	"github.com/rs/zerolog/log"
)

var QWIK_PUBLISH_CHANNEL = "qwik_events"

type EventPublisher interface {
	Publish(QwikEvent)
}

func NewNATSEventPublisher(cfg NatsConfig) (EventPublisher, error) {
	client, err := nats.Connect(cfg.ConnectString)

	if err != nil {
		return natsEventPublisher{}, err
	}
	var channel string
	if cfg.Channel == "" {
		channel = QWIK_PUBLISH_CHANNEL
	} else {
		channel = cfg.Channel
	}
	return natsEventPublisher{
		conn:    client,
		channel: channel,
	}, nil
}

type NatsConfig struct {
	ConnectString string
	Channel       string
}

type natsEventPublisher struct {
	conn    *nats.Conn
	channel string
}

func (nats natsEventPublisher) Publish(evt QwikEvent) {
	if nats.conn.IsConnected() {
		data, err := proto.Marshal(&evt)
		if err != nil {
			log.Error().Msgf("error marshaling event: %v", err)
			return
		}

		nats.conn.Publish(nats.channel, data)
	}
}
