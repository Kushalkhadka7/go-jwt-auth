package migration

import "time"

type User struct {
	Id        string    `json:"id" gorm:"primaryKey autoIncrement"`
	Name      string    `json:"name"`
	UserName  string    `json:"user_name" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at" sql:"DEFAULT:'current_timestamp'"`
	UpdatedAt time.Time `json:"modified_at" sql:"DEFAULT:'current_timestamp'"`
}
