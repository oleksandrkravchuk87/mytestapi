package mytestapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Server represents application http server
type Server struct {
	HttpServer     *http.Server
	ProfileService *ProfileService
}

// GetProfile is a profile handler
func (s *Server) GetProfile() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		username := r.URL.Query().Get("username")
		if username == "" {
			userProfiles, err := s.ProfileService.GetProfiles()
			if err != nil {
				http.Error(w, "Error retrieving user profile", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(userProfiles); err != nil {
				http.Error(w, "Error encoding response", http.StatusInternalServerError)
			}
			return
		}

		userProfile, err := s.ProfileService.GetProfileByUsername(username)
		if err != nil {
			if errors.Is(err, ErrNotFound) {
				http.Error(w, fmt.Sprintf("no profile found with user ID %s", username), http.StatusNotFound)
			} else {
				http.Error(w, "Error retrieving user profile", http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(userProfile); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
		}
		return
	})
}
