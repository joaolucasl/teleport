package config

import (
	"github.com/go-yaml/yaml"
	"github.com/pagarme/teleport/database"
	"io/ioutil"
)

type Config struct {
	Databases map[string]database.Database
}

func New() *Config {
	return &Config{}
}

// Open YAML file from path and unmarshal its content
// inside Config object itself
func (c *Config) ReadFromFile(path string) error {
	b, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(b, c); err != nil {
		return err
	}

	return nil
}