package soko

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pelletier/go-toml"
)

type Config struct {
	Backend string
	SectionConfig
}

type SectionConfig map[string]string

const (
	defaultConfigPath = "/etc/soko.toml"

	tomlTemplate = `[default]
backend = "%s"
`
)

var validKeys = map[string][]string{
	"consul": []string{
		"url",
	},
	"openstack": []string{
		"username",
		"password",
		"tenant_name",
		"auth_url",
		"region",
	},
	"aws": []string{
		"access_key_id",
		"secret_access_key",
		"region",
	},
}

func (c *Config) SetConfigByBackend(data *toml.TomlTree) error {
	switch b := c.Backend; b {
	case "consul", "openstack", "aws":
		cfg := make(SectionConfig, 0)
		keys := validKeys[b]
		for _, key := range keys {
			tomlKey := fmt.Sprintf("%s.%s", b, key)
			if v := data.Get(tomlKey); v != nil {
				cfg[key] = v.(string)
			}
		}
		c.SectionConfig = cfg
		return nil
	default:
		return fmt.Errorf("Invalid backend: %s", c.Backend)
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
	conf.Backend = data.Get("default.backend").(string)
	conf.SetConfigByBackend(data)

	return conf, nil
}

func NewConfig(backend string, args []string) (*Config, error) {
	conf := &Config{}
	conf.Backend = backend
	cfg := make(SectionConfig, 0)
	keys := validKeys[backend]

	for _, arg := range args {
		kv := strings.SplitN(arg, "=", 2)
		if len(kv) != 2 {
			k := kv[0]
			v := kv[1]
			for _, key := range keys {
				if k == key {
					cfg[k] = v
				}
			}
		} else {
			return nil, fmt.Errorf("Invalud config kv pair: %s", arg)
		}
	}
	conf.SectionConfig = cfg
	return conf, nil
}

func WriteToConfig(backend string) error {
	data := fmt.Sprintf(tomlTemplate, backend)
	return ioutil.WriteFile(defaultConfigPath, []byte(data), 0644)
}
