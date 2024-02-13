package simple_map

type StoreService struct {
	store *Store
}

func NewStoreService() *StoreService {
	return &StoreService{store: NewStore()}
}

func (s *StoreService) Set(key string, value any) {
	s.store.Set(key, value)
}

func (s *StoreService) Get(key string) (any, bool) {
	return s.store.Get(key)
}

func (s *StoreService) Delete(key string) {
	s.store.Delete(key)
}

func (s *StoreService) Save() {
	_ = s.store.Save()
}
