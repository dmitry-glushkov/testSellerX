package model

import "time"

type User struct {
	ID         int
	Username   string
	Created_at time.Time
}
