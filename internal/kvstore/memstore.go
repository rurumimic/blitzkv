package kvstore

type MemStore struct {
	store map[string]string
}

func NewMemStore() *MemStore {
	return &MemStore{
		store: make(map[string]string),
	}
}

func (m *MemStore) Get(key string) (string, error) {
	value, exists := m.store[key]
	if !exists {
		return "", nil
	}
	return value, nil
}

func (m *MemStore) Set(key, value string) error {
	m.store[key] = value
	return nil
}

func (m *MemStore) Delete(key string) error {
	_, exists := m.store[key]
	if !exists {
		return nil
	}
	delete(m.store, key)
	return nil
}

func (m *MemStore) List() ([]string, error) {
	keys := make([]string, 0, len(m.store))
	for key := range m.store {
		keys = append(keys, key)
	}
	return keys, nil
}
