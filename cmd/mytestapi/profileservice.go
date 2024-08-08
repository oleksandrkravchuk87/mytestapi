package mytestapi

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"mytestapi/cmd/mytestapi/models"

	"github.com/jmoiron/sqlx"
)

const (
	selectUserProfilesQuery = `
  SELECT
	    u.id,
	    u.username,
	    up.first_name,
	    up.last_name,
	    up.city,
	    ud.school
	FROM
	    user u
	JOIN
	    user_profile up ON u.id = up.user_id
	JOIN
	    user_data ud ON u.id = ud.user_id;`

	selectUserProfileByUsernameQuery = `
 SELECT
	    u.id,
	    u.username,
	    up.first_name,
	    up.last_name,
	    up.city,
	    ud.school
	FROM
	    user u
	JOIN
	    user_profile up ON u.id = up.user_id
	JOIN
	    user_data ud ON u.id = ud.user_id
	WHERE
	    u.username = ?;`
)

// ErrNotFound is not found error
var ErrNotFound = errors.New("not found")

// ProfileService is a profile service
type ProfileService struct {
	dbClient *sqlx.DB
}

// NewProfileService returns a pointer to profile service
func NewProfileService(db *sqlx.DB) *ProfileService {
	return &ProfileService{dbClient: db}
}

// GetProfileByUsername returns profile by username
func (p *ProfileService) GetProfileByUsername(ctx context.Context, userID string) (*models.UserProfile, error) {
	var userProfile models.UserProfile
	err := p.dbClient.GetContext(ctx, &userProfile, selectUserProfileByUsernameQuery, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no user found with user_id %s: %w", userID, ErrNotFound)
		}
		return nil, err
	}
	return &userProfile, nil
}

// GetProfiles returns all profiles
func (p *ProfileService) GetProfiles(ctx context.Context) ([]models.UserProfile, error) {
	userProfiles := []models.UserProfile{}
	err := p.dbClient.SelectContext(ctx, &userProfiles, selectUserProfilesQuery)
	if err != nil {
		return nil, err
	}
	return userProfiles, nil
}
