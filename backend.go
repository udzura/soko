package soko

import (
	"fmt"
	"net/url"
)

type Backend interface {
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

func FindBackend(uri string) (Backend, error) {
	if uri == "" {
		// Defaults to return consul default backend
		return NewConsulBackend("", false)
	}

	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	// TODO: implement AWS and OpenStack backend
	switch u.Scheme {
	case "consul":
		return NewConsulBackend(u.Host, false)
	case "consuls":
		return NewConsulBackend(u.Host, true)
	default:
		return nil, fmt.Errorf("Unsupported schema: %s", uri)
	}
}
