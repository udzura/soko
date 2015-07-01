package metama

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

func FindBackend(uri string) Backend {
	// TBD parse uri and return proper backend
	return NewConsulBackend()
}
