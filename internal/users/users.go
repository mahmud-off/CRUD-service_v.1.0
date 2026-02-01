package users

import "time"

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	TimeStamp time.Time
}
