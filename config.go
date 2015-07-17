package soko

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pelletier/go-toml"
)

type Config struct {
	URI string
}

const (
	defaultConfigPath = "/etc/soko.toml"

	tomlTemplate = `[default]
uri = "%s"
`
)

func DefaultConfig() (*Config, error) {
	if _, err := os.Stat(defaultConfigPath); err != nil {
		// If there is no file, returns empty config
		return &Config{}, nil
	}

	data, err := toml.LoadFile(defaultConfigPath)
	if err != nil {
		return nil, err
	}
	conf := &Config{}
	conf.URI = data.Get("default.uri").(string)

	return conf, nil
}

func WriteToConfig(uri string) error {
	data := fmt.Sprintf(tomlTemplate, uri)
	return ioutil.WriteFile(defaultConfigPath, []byte(data), 0644)
}
