package app

import (
	"time"

	"github.com/yelaco/gchess-tui/pkg/ui/domains"
)

// UserProfile struct    stores local user profile
type UserProfile struct {
	UserId    string // unique identifer for each user in system
	PlayerId  string // unique identifier for each match user is still playing in
	Username  string
	Email     string
	Rating    int64
	CreatedAt time.Time
}

// NewUserProfile function    extract user information from user domain
func NewUserProfile(user domains.User) {
	getApp().UserProfile = UserProfile{
		Username:  user.Username,
		Email:     user.Email,
		Rating:    user.Rating,
		CreatedAt: user.CreatedAt,
	}
}

// GetUserProfile function    Return a copy of user profile.
// Any update to user profile should be done via modifying the copied value and SyncUserProfile() method
func GetUserProfile() UserProfile {
	return getApp().UserProfile
}

// SyncProfile function    syncs user profile locally and remotely in database
func SyncUserProfile(userProfile UserProfile) error {
	return nil
}

type Match struct {
	domains.Match
}

func NewMatch(match domains.Match) Match {
	return Match{match}
}

// SyncMatchState function    syncs match state locally and remotely
func SyncMatchState(match domains.Match) error {
	getApp().Match.GameState = match.GameState
	getApp().Match.PlayerState = match.PlayerState
	return nil
}
