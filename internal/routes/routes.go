package routes

import "net/http"

func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", healthCheckHandler)
	return mux
}

func healthCheckHandler(writer http.ResponseWriter, request *http.Request) {
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
