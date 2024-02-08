package simple_map

type StoreDto struct {
	Key   string `json:"key" validate:"required"`
	Value any    `json:"value" validate:"required"`
}
