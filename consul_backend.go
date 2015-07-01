package metama

import (
	"github.com/hashicorp/consul/api"
)

type ConsulBackend struct {
	client *api.KV
}

func NewConsulBackend() *ConsulBackend {
	client, _ := api.NewClient(api.DefaultConfig())
	kv := client.KV()

	return &ConsulBackend{
		client: kv,
	}
}
