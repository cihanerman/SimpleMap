package simple_map

import (
	"reflect"
	"testing"
)

func TestNewStoreService(t *testing.T) {
	t.Parallel()

	service := NewStoreService()
	if service == nil {
		t.Error("NewStoreService should not return a nil object")
	}
	if reflect.TypeOf(service) != reflect.TypeOf(&StoreService{}) {
		t.Error("NewStoreService should return a StoreService object")
	}
}

func TestStoreService_Set_Get(t *testing.T) {
	t.Parallel()

	service := NewStoreService()
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
			service.Set(tt.args.key, tt.args.value)

			if _, ok := service.Get(tt.args.key); !ok {
				t.Errorf("Set() = %v, want %v", ok, true)
			}

		})
	}
}
