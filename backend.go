package soko

import (
	"fmt"
	"net/url"
)

type Backend interface {
	// Saves current configuration to a specific file
	Save() error

	// APIs to control backend metadata
	// Gets value from key
	Get(serverID string, key string) (string, error)

	// Put value on the key
	Put(serverID string, key string, value string) error

	// Delete value on the key
	Delete(serverID string, key string) error

	// TODO: implement
	// List(serverID string, prefix string) , returns all of values with serverID
	// Search(key string) , retruns serverID
	// Watch(serverID string, key string) , this blocks until change
	// ...
}

func FindBackend(config *Config) (Backend, error) {
	switch config.Backend {
	case "":
		// Defaults to return consul default backend
		return NewConsulBackend("", false)
	case "consul", "consuls":
		c, err := config.GetConfigBySection("openstack")
		if err != nil {
			return nil, err
		}
		u, err := url.Parse(c["url"])
		if err != nil {
			return nil, err
		}

		ssl := (u.Scheme == "https") || (config.Backend == "consuls")
		return NewConsulBackend(u.Host, ssl)
	case "openstack":
		c, err := config.GetConfigBySection("openstack")
		if err != nil {
			return nil, err
		}
		return NewOpenStackBackend(c)
	case "aws":
		c, err := config.GetConfigBySection("aws")
		if err != nil {
			return nil, err
		}
		return NewAWSBackend(c)
	default:
		return nil, fmt.Errorf("Unsupported backend: %s", config.Backend)
	}
}
