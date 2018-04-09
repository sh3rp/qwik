package qwik

import (
	"encoding/json"
	"os"
)

type QwikConfig struct {
	SrcIP      string               `json:"src_ip"`
	MessageBus QwikMessageBusConfig `json:"message_bus"`
	FilePaths  []string             `json:"paths"`
}

type QwikMessageBusConfig struct {
	ConnectString string `json:"connect_string"`
	Channel       string `json:"channel"`
}

func ReadConfig(filename string) (QwikConfig, error) {
	var config QwikConfig
	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	json.NewDecoder(file).Decode(&config)
	return config, nil
}
