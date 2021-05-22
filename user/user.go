package user

import (
	"jwt-auth/model"
	pb "jwt-auth/pb"
	"time"

	"gorm.io/gorm"
)

type User struct {
	table string
	db    *gorm.DB
}

type UserI interface {
	CreateUser(user *pb.User) (*model.User, error)
	CheckUserExistence(user *pb.User) (bool, error)
}

func NewUser(db *gorm.DB) UserI {
	return &User{
		table: "user",
		db:    db,
	}
}

func (auth *User) CreateUser(user *pb.User) (*model.User, error) {
	userModel := model.User{}
	newUser := model.User{
		Name:      user.Name,
		Password:  user.Password,
		UserName:  user.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	data := auth.db.Model(&userModel).Create(newUser)
	if data.Error != nil {
		return nil, data.Error
	}

	if data == nil {
		return nil, nil
	}

	return &userModel, nil
}

// CheckUserExistence checks either the given user already exist in db or not.
func (auth *User) CheckUserExistence(user *pb.User) (bool, error) {
	userModel := model.User{}

	data := auth.db.Where("name = ?", user.Name).Find(&userModel)

	if data.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}
