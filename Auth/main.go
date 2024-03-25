package main

import (
	"fmt"
	"net/http"
	"strings"
)

var validTokens = map[string]string{
	"valid_token_1": "user1",
	"valid_token_2": "user2",
}

func main() {
	// fmt.Println("Starting authorization service...")

	http.HandleFunc("/authorize", authorizeHandler)
	fmt.Println("Authorization service is listening on port 9002...")
	http.ListenAndServe(":9002", nil)
}

func authorizeHandler(w http.ResponseWriter, r *http.Request) {

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
		return
	}
	token := parts[1]

	username, ok := validTokens[token]
	if !ok {
		http.Error(w, "Invalid token", http.StatusForbidden)
		return
	}

	w.Header().Set("x-current-user", username)
	w.WriteHeader(http.StatusOK)
}
