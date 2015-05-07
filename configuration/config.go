package configuration

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Configuration struct {
	Version    string            `json:"version"`
	PsqlConfig PsqlConfiguration `json:"psql_connection"`
}

func LoadConfiguration(filepath string) *Configuration {
	config := new(Configuration)
	loadConfiguration(filepath, config)
	return config
}

func loadConfiguration(filepath string, configuration interface{}) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Error loading config file at path:%s. Error: %v ", filepath, err.Error())
	}

	if err := parseConfig(data, configuration); err != nil {
		log.Fatalf("Error parsing config file. Error: %v ", err.Error())
	}
}

func parseConfig(data []byte, configuration interface{}) error {
	err := json.Unmarshal(data, configuration)
	if err != nil {
		if configuration != nil {
			configuration = nil
		}
		return err
	}
	return nil
}
