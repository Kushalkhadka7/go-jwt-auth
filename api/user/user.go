package user

import (
	"fmt"
	"jwt-auth/model"
	pb "jwt-auth/pb"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Store is a data structure to populate with credentails and methods that are required to perform operations against db.
type Store struct {
	table string
	db    *gorm.DB
}

// Storer is a interface that exposses methods that are used to perform operation on user against db.
type Storer interface {
	DeleteUser(email, userName string) (*model.User, error)
	DeactivateUser(isActive bool) (*model.User, error)
	ActivateUser(isActive bool) (*model.User, error)
	CreateUser(user *pb.User) (*model.User, error)
	CheckUserExistence(user *pb.User) (bool, error)
}

// NewUserStore initializes new user store with associated table and db connection.
func NewStore(db *gorm.DB) Storer {
	return &Store{
		table: "user",
		db:    db,
	}
}

// CreateUser creates and saves user in db.
func (s *Store) CreateUser(user *pb.User) (*model.User, error) {
	userModel := model.User{}

	newUser := model.User{
		Id:           uuid.NewString(),
		Name:         user.Name,
		Password:     user.Password,
		UserName:     user.Name,
		Email:        user.GetEmail(),
		Role:         "ADMIN",
		AccessToken:  "",
		RefreshToken: "",
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	fmt.Printf(":%v", newUser)

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

	data := auth.db.Where("name = ?", user.Email).Find(&userModel)

	if data.RowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

func (auth *Store) DeactivateUser(isActive bool) (*model.User, error) {
	userModel := model.User{}

	data := auth.db.Where("name = ?", "kushal").Find(&userModel)

	if data.RowsAffected == 0 {
		return nil, nil
	}

	// Update with conditions
	auth.db.Model(&userModel).Where("name = ?", "kushal").Update("is_active", false)

	data1 := auth.db.Where("name = ?", "kushal").Find(&userModel)

	if data1.RowsAffected == 0 {
		return nil, nil
	}

	return &userModel, nil
}

func (auth *Store) ActivateUser(isActive bool) (*model.User, error) {
	userModel := model.User{}

	data := auth.db.Where("name = ?", "kushal").Find(&userModel)

	if data.RowsAffected == 0 {
		return nil, nil
	}

	// Update with conditions
	auth.db.Model(&userModel).Where("name = ?", "kushal").Update("is_active", true)

	data1 := auth.db.Where("name = ?", "kushal").Find(&userModel)

	if data1.RowsAffected == 0 {
		return nil, nil
	}

	return &userModel, nil
}

func (auth *Store) DeleteUser(email, userName string) (*model.User, error) {
	userModel := model.User{}

	data := auth.db.Where("name = ?", "kushal").Find(&userModel)

	if data.RowsAffected == 0 {
		return nil, nil
	}

	// Update with conditions
	auth.db.Model(&userModel).Where("name = ?", "kushal").Where("email = ?", "kushal").Delete(&userModel)

	return &userModel, nil
}
