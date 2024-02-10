package routes

import (
	"encoding/json"
	"fmt"
	"github.com/cihanerman/SimpleMap/internal/simple_map"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	//mux.HandleFunc("/", healthCheckHandler)

	// TODO: This will change in go 1.22
	// sample mux.HandleFunc("GET /healthcheck", healthCheckHandler)
	mux.HandleFunc("/healthcheck", healthCheckHandler)
	mux.HandleFunc("/set_key", setKeyHandler)
	mux.HandleFunc("/get_key", getKeyHandler)
	mux.HandleFunc("/delete_key", deleteKeyHandler)
	return mux
}

func deleteKeyHandler(writer http.ResponseWriter, request *http.Request) {
	// TODO: This will delete in go 1.22
	if request.Method != http.MethodGet {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	key := request.URL.Query().Get("key")
	if key == "" {
		http.Error(writer, "Key is required", http.StatusBadRequest)
		return
	}

	service := simple_map.NewStoreService()
	service.Delete(key)

	writer.WriteHeader(http.StatusNoContent)
}

func getKeyHandler(writer http.ResponseWriter, request *http.Request) {
	// TODO: This will delete in go 1.22
	if request.Method != http.MethodGet {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	key := request.URL.Query().Get("key")
	if key == "" {
		http.Error(writer, "Key is required", http.StatusBadRequest)
		return
	}

	service := simple_map.NewStoreService()
	value, ok := service.Get(key)
	if !ok {
		http.Error(writer, "Key not found", http.StatusNotFound)
		return
	}

	response := simple_map.StoreDto{
		Key:   key,
		Value: value.(string),
	}

	writer.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func setKeyHandler(writer http.ResponseWriter, request *http.Request) {
	// TODO: This will delete in go 1.22
	if request.Method != http.MethodPost {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var storeDto simple_map.StoreDto
	err := json.NewDecoder(request.Body).Decode(&storeDto)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// validate request body
	validate := validator.New()
	err = validate.Struct(storeDto)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			http.Error(writer, fmt.Sprintf("Invalid %s value: %s", e.Field(), e.Tag()), http.StatusBadRequest)
			return
		}
	}
	service := simple_map.NewStoreService()
	service.Set(storeDto.Key, storeDto.Value)
	writer.WriteHeader(http.StatusOK)
}

func healthCheckHandler(writer http.ResponseWriter, request *http.Request) {
	// TODO: This will delete in go 1.22
	if request.Method != http.MethodGet {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	_, err := writer.Write([]byte("I'm still alive"))
	if err != nil {

		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
