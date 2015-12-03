package core

import (
	"encoding/json"
	"os"
)

type ServerConfiguration struct {
	Host string
	Port string
}

type Configuration struct {
	Server ServerConfiguration
}

func (c Configuration) loadFromFile() error {
	file, _ := os.Open("conf/main.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&c)
	if err != nil {
		return err
	}
	return nil
}
