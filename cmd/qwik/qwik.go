package main

import (
	"fmt"
	"time"

	"github.com/sh3rp/qwik"
)

func main() {

	config, err := qwik.ReadConfig("qwik.json")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	publisher, err := qwik.NewNATSEventPublisher(qwik.NatsConfig{
		ConnectString: config.MessageBus.ConnectString,
		Channel:       config.MessageBus.Channel,
	})

	if err != nil {
		fmt.Printf("Error creating message bus client: %v\n", err)
		return
	}

	transformer := qwik.NewEventTransformer(config.SrcIP)

	handler := qwik.NewEventHandler(publisher, transformer)

	watcher, err := qwik.NewFileWatcher(handler.Handle)

	if err != nil {
		fmt.Printf("Error creating watcher: %v\n", err)
		return
	}

	for _, path := range config.FilePaths {
		files := qwik.GetAllFiles(path)
		for _, file := range files {
			fmt.Printf("Registered: %s\n", file)
			watcher.RegisterPath(file)
		}
	}

	for {
		time.Sleep(1000)
	}
}
