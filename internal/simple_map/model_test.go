package simple_map

import (
	"testing"
)

func TestStore_Set_Get(t *testing.T) {
	t.Parallel()

	store := NewStore()
	type args struct {
		key   string
		value any
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Test String", args: args{key: "test string", value: "test"}},
		{name: "Test Int", args: args{key: "test int", value: 1}},
		{name: "Test Bool", args: args{key: "test bool", value: true}},
		{name: "Test Struct", args: args{key: "test struct", value: struct{}{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store.Set(tt.args.key, tt.args.value)

			if _, ok := store.Get(tt.args.key); !ok {
				t.Errorf("Set() = %v, want %v", ok, true)
			}

		})
	}
}

func TestNewStore(t *testing.T) {
	t.Parallel()
	store := NewStore()
	if store == nil {
		t.Error("NewStore should not return a nil object")
	}

	if store.items == nil {
		t.Error("Store's items area should be a map")
	}

	store2 := NewStore()

	if store != store2 {
		t.Error("NewStore should return a singleton object")
	}

}
