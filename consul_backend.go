package soko

import (
	"fmt"
	"io/ioutil"

	"github.com/hashicorp/consul/api"
)

type ConsulBackend struct {
	client      *api.KV
	originalURL string
}

const consulTomlTemplate = `[default]
backend = "%s"

[consul]
url = "%s"
`

func NewConsulBackend(hostWithPort string, ssl bool) (*ConsulBackend, error) {
	conf := api.DefaultConfig()
	if hostWithPort == "" {
		hostWithPort = "localhost:8500"
	}

	conf.Address = hostWithPort
	if ssl {
		conf.Scheme = "https"
	}

	client, err := api.NewClient(conf)
	if err != nil {
		return nil, err
	}
	kv := client.KV()

	var url string
	if ssl {
		url = fmt.Sprintf("https://%s", hostWithPort)
	} else {
		url = fmt.Sprintf("http://%s", hostWithPort)
	}

	return &ConsulBackend{
		client:      kv,
		originalURL: url,
	}, nil
}

func (b *ConsulBackend) pathOf(serverID string, key string) string {
	const prefix = "metadata"
	return fmt.Sprintf("%s/%s/%s", prefix, serverID, key)
}

func (b *ConsulBackend) Save() error {
	data := fmt.Sprintf(
		consulTomlTemplate,
		"consul",
		b.originalURL,
	)
	return ioutil.WriteFile(defaultConfigPath, []byte(data), 0644)
}

func (b *ConsulBackend) Get(serverID string, key string) (string, error) {
	p, _, err := b.client.Get(b.pathOf(serverID, key), nil)
	if err != nil {
		return "", err
	}

	if p == nil || len(p.Value) == 0 {
		sayEmpty(key)
		return "", nil
	}

	return string(p.Value), nil
}

func (b *ConsulBackend) Put(serverID string, key string, value string) error {
	data := &api.KVPair{
		Key:   b.pathOf(serverID, key),
		Value: []byte(value),
	}
	_, err := b.client.Put(data, nil)

	return err
}

func (b *ConsulBackend) Delete(serverID string, key string) error {
	_, err := b.client.Delete(b.pathOf(serverID, key), nil)

	return err
}
