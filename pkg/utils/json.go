package utils

import (
	"encoding/json"
	"os"
)

// WriteJSON writes the data to a file in JSON format.
func WriteJSON(filename string, data map[string]any) error {
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, jsonData, 0644)
}

// ReadJSON reads the data from a file in JSON format.
func ReadJSON(filename string) (map[string]any, error) {
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var data map[string]any
	if err := json.Unmarshal(fileData, &data); err != nil {
		return nil, err
	}
	return data, nil
}
