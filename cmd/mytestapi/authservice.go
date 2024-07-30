package mytestapi

import "github.com/jmoiron/sqlx"

// AuthService represents authentication service
type AuthService struct {
	dbClient *sqlx.DB
}

// NewAuthService returns pointer to AuthService
func NewAuthService(db *sqlx.DB) *AuthService {
	return &AuthService{dbClient: db}
}

// IsValidAPIKey validates apiKey
func (a *AuthService) IsValidAPIKey(apiKey string) (bool, error) {
	var authID int
	err := a.dbClient.Get(&authID, "SELECT EXISTS(SELECT id FROM auth WHERE  `api-key` = ?);", apiKey)
	if err != nil {
		return false, err
	}

	return authID != 0, nil
}
