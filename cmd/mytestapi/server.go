package mytestapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"mytestapi/cmd/mytestapi/models"
	"net/http"

	_ "github.com/golang/mock/mockgen/model"
)

//go:generate mockgen -package=mocks -destination=./mocks/profileservice_mock.go mytestapi/cmd/mytestapi IProfileService
type IProfileService interface {
	GetProfileByUsername(userID string) (*models.UserProfile, error)
	GetProfiles() ([]models.UserProfile, error)
}

// Server represents application http server
type Server struct {
	HttpServer     *http.Server
	ProfileService IProfileService
}

// GetProfile is a profile handler
func (s *Server) GetProfile() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
