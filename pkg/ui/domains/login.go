package domains

import "time"

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
