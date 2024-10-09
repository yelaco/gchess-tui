package login

import "time"

type LoginDao struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserDao struct {
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Rating    int64     `json:"rating,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
