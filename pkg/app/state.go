package app

import (
	"github.com/yelaco/gchess-tui/pkg/ui/domains"
)

func GetUser() domains.User {
	mu.RLock()
	defer mu.RUnlock()
	return getApp().User
}

// SyncProfile function    syncs user profile locally and remotely in database
func SyncUser(user domains.User) {
	mu.Lock()
	defer mu.Unlock()
	getApp().User = user
}

func GetMatch() domains.Match {
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
	getApp().Match = match
}
