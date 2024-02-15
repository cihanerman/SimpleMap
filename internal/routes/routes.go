package routes

import (
	"encoding/json"
	"fmt"
	"github.com/cihanerman/SimpleMap/internal/simple_map"
	"github.com/cihanerman/SimpleMap/pkg/auth"
	"github.com/go-playground/validator/v10"
	"net/http"
	"os"
)

func NewServer() *http.Server {
	mux := http.NewServeMux()

	// Registering the handlers
	mux.HandleFunc("GET /{$}", auth.BasicAuth(healthCheckHandler))
	mux.HandleFunc("GET /healthcheck", auth.BasicAuth(healthCheckHandler))
	mux.HandleFunc("POST /key", auth.BasicAuth(setKeyHandler))
	mux.HandleFunc("GET /key/{key}", auth.BasicAuth(getKeyHandler))
	mux.HandleFunc("DELETE /key/{key}", auth.BasicAuth(deleteKeyHandler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	return server
}

func deleteKeyHandler(writer http.ResponseWriter, request *http.Request) {
	key := request.PathValue("key")
	if key == "" {
		http.Error(writer, "Key is required", http.StatusBadRequest)
		return
	}

	service := simple_map.NewStoreService()
	service.Delete(key)

	writer.WriteHeader(http.StatusNoContent)
}

func getKeyHandler(writer http.ResponseWriter, request *http.Request) {
	key := request.PathValue("key")
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

func healthCheckHandler(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(http.StatusOK)
	_, err := writer.Write([]byte("I'm still alive"))
	if err != nil {

		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
