package app

import (
	"time"

	"github.com/yelaco/gchess-tui/pkg/ui/domains"
)

// UserProfile struct    stores local user profile
type UserProfile struct {
	UserId    string
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

func GetMatch() Match {
	mu.RLock()
	defer mu.RUnlock()
	return getApp().Match
}

func NewMatch(match domains.Match) {
	SyncMatch(match)
}

// SyncMatch function    syncs match state locally and remotely
func SyncMatch(match domains.Match) {
	mu.Lock()
	defer mu.Unlock()
	getApp().Match = Match{match}
}
