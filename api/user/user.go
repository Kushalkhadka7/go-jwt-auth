package user

import (
	"jwt-auth/model"
	pb "jwt-auth/pb"
	"time"

	"gorm.io/gorm"
)

// Store is a data structure to populate with credentails and methods that are required to perform operations against db.
type Store struct {
	table string
	db    *gorm.DB
}

// Storer is a interface that exposses methods that are used to perform operation on user against db.
type Storer interface {
	CreateUser(user *pb.User) (*model.User, error)
	CheckUserExistence(user *pb.User) (bool, error)
}

// NewUserStore initializes new user store with associated table and db connection.
func NewUserStore(db *gorm.DB) Storer {
	return &Store{
		table: "user",
		db:    db,
	}
}

// CreateUser creates and saves user in db.
func (s *Store) CreateUser(user *pb.User) (*model.User, error) {
	userModel := model.User{}

	newUser := model.User{
		Name:      user.Name,
		Password:  user.Password,
		UserName:  user.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	res := s.db.Model(&userModel).Create(newUser)
	if res.Error != nil {
		return nil, res.Error
	}

	if res == nil {
		return nil, nil
	}

	return &userModel, nil
}

// CheckUserExistence checks either the given user already exist in db or not.
func (auth *Store) CheckUserExistence(user *pb.User) (bool, error) {
	userModel := model.User{}

	data := auth.db.Where("name = ?", user.Name).Find(&userModel)

	if data.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}
