package model

import "time"

type User struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	UserName     string    `json:"user_name"`
	Password     string    `json:"password"`
	Email        string    `json:"email"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	Role         string    `json:"role"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"modified_at"`
}
