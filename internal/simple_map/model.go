package simple_map

import "sync"

type Store struct {
	sync.RWMutex
	items map[string]any
}

var lock = &sync.Mutex{}
var store *Store

func (s *Store) Set(key string, value any) {
	s.Lock()
	defer s.Unlock()
	s.items[key] = value
}

func (s *Store) Get(key string) (any, bool) {
	s.RLock()
	defer s.RUnlock()
	value, ok := s.items[key]
	return value, ok
}

func NewStore() *Store {
	if store == nil {
		lock.Lock()
		defer lock.Unlock()
		if store == nil {
			store = &Store{items: make(map[string]any)}
		}
	}
	return store
}
