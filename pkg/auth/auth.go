package auth

import (
	"encoding/base64"
	"github.com/cihanerman/SimpleMap/pkg/utils"
	"net/http"
	"strings"
)

var users, _ = utils.ReadJSON("users.json")

func BasicAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// If the users.json file is missing, skip the authentication
		if users == nil {
			next.ServeHTTP(w, r)
			return
		}
		// Check the "Authorization" header in the HTTP request
		auth := r.Header.Get("Authorization")
		if auth == "" {
			// If the "Authorization" header is missing, return an authentication error
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Parse the "Authorization" header
		parts := strings.SplitN(auth, " ", 2)
		if len(parts) != 2 || parts[0] != "Basic" {
			http.Error(w, "Invalid Authorization header", http.StatusBadRequest)
			return
		}

		// Decode the Base64 encoded username and password
		decoded, err := base64.StdEncoding.DecodeString(parts[1])
		if err != nil {
			http.Error(w, "Error decoding authorization header", http.StatusBadRequest)
			return
		}

		// Split the username and password
		credentials := strings.SplitN(string(decoded), ":", 2)
		if len(credentials) != 2 {
			http.Error(w, "Invalid authorization header", http.StatusBadRequest)
			return
		}

		// Check the username and password
		if password, ok := users[credentials[0]]; !ok || password != credentials[1] {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If the user authentication is successful, call the next function
		next.ServeHTTP(w, r)
	}
}
