package soko

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pelletier/go-toml"
)

type Config struct {
	URI string

	original *toml.TomlTree
}

type SectionConfig map[string]string

const (
	defaultConfigPath = "/etc/soko.toml"

	tomlTemplate = `[default]
uri = "%s"
`
)

func (c *Config) GetConfigBySection(sectionName string) (SectionConfig, error) {
	switch sectionName {
	case "consul":
		return make(SectionConfig, 0), nil
	case "openstack":
		cfg := make(SectionConfig, 0)
		validKeys := []string{
			"username",
			"password",
			"tenant_name",
			"auth_url",
			"region",
		}
		for _, key := range validKeys {
			tomlKey := fmt.Sprintf("%s.%s", sectionName, key)
			if v := c.original.Get(tomlKey); v != nil {
				cfg[key] = v.(string)
			}
		}
		return cfg, nil
	case "aws":
		return nil, fmt.Errorf("AWS not yet implemented")
	default:
		return nil, fmt.Errorf("Invalid backend: %s", sectionName)
	}
}

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
	conf.original = data

	return conf, nil
}

func WriteToConfig(uri string) error {
	data := fmt.Sprintf(tomlTemplate, uri)
	return ioutil.WriteFile(defaultConfigPath, []byte(data), 0644)
}
