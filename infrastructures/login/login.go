package login

import "time"

type UserDao struct {
	Username  string
	Email     string
	Rating    int64
	CreatedAt time.Time
}
