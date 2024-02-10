package utils

import (
	"os"
	"reflect"
	"testing"
)

func TestWriteJSON_ReadJSON(t *testing.T) {
	t.Parallel()

	filename := "test.json"
	testMap := make(map[string]any)
	testMap["test"] = "test"
	err := WriteJSON(filename, testMap)
	if err != nil {
		t.Errorf("WriteJSON() = %v, want %v", err, nil)
	}
	_, err = os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			t.Errorf("WriteJSON() = %v, want %v", nil, err)
		}
	}

	data, err := ReadJSON(filename)
	if err != nil {
		t.Errorf("ReadJSON() = %v, want %v", err, nil)
	}

	if !reflect.DeepEqual(data, testMap) {
		t.Errorf("ReadJSON() = %v, want %v", data, testMap)
	}
}
