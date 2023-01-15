package model

import "time"

var (
	Permission = map[string]int64{
		"user":  1, // id to level of the user
		"admin": 2, // id to level of the admin
	}
)

type User struct {
	Id         int64
	Name       string
	Email      string
	Passw      string
	Permission LevelUser
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type LevelUser struct {
	IsAdmin bool
	ID      int64
	Name    string
}
