package dtos

import (
	"time"

	domainlogin "github.com/yelaco/gchess-tui/domains/entities/login"
)

type Login struct {
	Username string
	Password string
}

type User struct {
	Username  string
	Email     string
	Rating    int64
	CreatedAt time.Time
}

func UserEntityToDto(user domainlogin.User) User {
	return User{
		Username:  user.Username,
		Email:     user.Email,
		Rating:    user.Rating,
		CreatedAt: user.CreatedAt,
	}
}
