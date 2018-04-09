package main

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
	nats "github.com/nats-io/go-nats"
	"github.com/sh3rp/qwik"
)

func main() {
	config, err := qwik.ReadConfig("qwik.json")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	client, err := nats.Connect(config.MessageBus.ConnectString)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	client.Subscribe(config.MessageBus.Channel, func(msg *nats.Msg) {
		qwikEvent := &qwik.QwikEvent{}
		err := proto.Unmarshal(msg.Data, qwikEvent)
		if err != nil {
			fmt.Printf("error parsing: %v\n", err)
			return
		}
		fmt.Printf("[%-15s] id=%s type=%d timestamp=%d\n", qwikEvent.SrcIp, qwikEvent.Id, qwikEvent.Type, qwikEvent.Timestamp)

		switch qwikEvent.Event.(type) {
		case *qwik.QwikEvent_Fsevent:
			fsevt := qwikEvent.Event.(*qwik.QwikEvent_Fsevent)
			evt := fsevt.Fsevent
			fmt.Printf("  FS: op=%d path=%s\n", evt.Op, evt.Path)
		}
	})

	for {
		time.Sleep(1000)
	}
}
