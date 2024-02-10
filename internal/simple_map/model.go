package simple_map

import (
	"github.com/cihanerman/SimpleMap/pkg/utils"
	"sync"
)

const dataFile = "store.json"

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

func (s *Store) Delete(key string) {
	s.Lock()
	defer s.Unlock()
	delete(s.items, key)
}

func (s *Store) Save() error {
	return utils.WriteJSON(dataFile, s.items)
}

func (s *Store) Load() {
	data, err := utils.ReadJSON(dataFile)
	if err == nil {
		s.items = data
	} else {
		s.items = make(map[string]any)
	}
}

func NewStore() *Store {
	if store == nil {
		lock.Lock()
		defer lock.Unlock()
		if store == nil {
			store = &Store{items: nil}
			store.Load()
		}
	}
	return store
}
