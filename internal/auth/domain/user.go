package domain

import "time"

type User struct {
	ID            int
	Username      string
	Email         string
	Password      string
	Created_At    time.Time
	Updated_At    time.Time
	Refresh_Token string
}

func NewUser(username, email, password string) *User {
	return &User{
		Username: username,
		Email:    email,
		Password: password,
	}
}
