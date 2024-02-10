package simple_map

import (
	"reflect"
	"testing"
)

func TestStore_Set_Get_Delete_Save_Load(t *testing.T) {
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

			store.Delete(tt.args.key)
			if _, ok := store.Get(tt.args.key); ok {
				t.Errorf("Delete() = %v, want %v", ok, false)
			}

		})
	}

	// save test
	store.Set("test", "test")
	err := store.Save()
	if err != nil {
		t.Errorf("Save() = %v, want %v", err, nil)
	}

	// load test
	store.items = nil
	store.Load()

	if !reflect.DeepEqual(store.items, map[string]any{"test": "test"}) {
		t.Errorf("Load() = %v, want %v", store.items, map[string]any{"test": "test"})
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
