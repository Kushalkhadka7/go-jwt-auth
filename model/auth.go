package model

import "time"

type User struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	UserName  string    `json:"user_name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"modified_at"`
}
