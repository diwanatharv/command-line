package models

type KeyValueStore struct {
	store map[string]string
}

func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		store: make(map[string]string),
	}
}

type keymethods interface {
	Set(key, value string)
	Get(key string) (string, bool)
	Delete(key string) bool
	ListKeys() []string
	Count() int
}

func (kvs *KeyValueStore) Set(key, value string) {
	kvs.store[key] = value
}

func (kvs *KeyValueStore) Get(key string) (string, bool) {
	value, exists := kvs.store[key]
	return value, exists
}

func (kvs *KeyValueStore) Delete(key string) bool {
	_, exists := kvs.store[key]
	if exists {
		delete(kvs.store, key)
		return true
	}
	return false
}

func (kvs *KeyValueStore) ListKeys() []string {
	keys := make([]string, 0, len(kvs.store))
	for key := range kvs.store {
		keys = append(keys, key)
	}
	return keys
}

func (kvs *KeyValueStore) Count() int {
	return len(kvs.store)
}
