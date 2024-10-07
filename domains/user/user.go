package user

import "time"

type User struct {
	Username  string
	Email     string
	Rating    int64
	CreatedAt time.Time
}
