package views

import "time"

var (
	Permission = map[string]int64{
		"user":  1,
		"admin": 2,
	}
)

type CreateUser struct {
	Id         int64
	Name       string
	Email      string
	Passw      string
	Permission struct {
		IsAdmin bool
		ID      int64
		Name    string
	}
	CreatedAt time.Time
	UpdatedAt time.Time
}
