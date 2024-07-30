package middlewares

import (
	"log"
	"mytestapi/cmd/mytestapi"
	"net/http"
)

type AuthMiddleware struct {
	authService *mytestapi.AuthService
}

// NewAuthMiddleware new auth middleware
func NewAuthMiddleware(authService *mytestapi.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

// ServeHTTP auth middleware
func (a *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	apiKey := r.Header.Get("Api-key")
	if apiKey == "" {
		http.Error(w, "No Api Key", http.StatusUnauthorized)
		return
	}

	isValidAPIKey, err := a.authService.IsValidAPIKey(apiKey)
	if err != nil {
		log.Printf("Error validating api key: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if !isValidAPIKey {
		http.Error(w, "Api-key is not valid", http.StatusForbidden)
		return
	}
	next(w, r)
}
