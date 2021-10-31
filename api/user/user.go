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
	CreateUser(user *pb.User) (*model.User, error)
	CheckUserExistence(email string) (bool, error)
	GetUser(email, userName string) (*model.User, error)
	DeleteUser(email, userName string) (*model.User, error)
	ActivateUser(email, userName string) (*model.User, error)
	DeactivateUser(email, userName string) (*model.User, error)
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

	res := s.db.Model(&userModel).Create(newUser)
	if res.Error != nil {
		return nil, res.Error
	}

	if res == nil {
		return nil, fmt.Errorf("%v", "Cannot create user.")
	}

	return &newUser, nil
}

// CheckUserExistence checks either the given user already exist in db or not.
func (auth *Store) CheckUserExistence(email string) (bool, error) {
	userModel := model.User{}

	data := auth.db.Where("email = ?", email).Find(&userModel)

	if data.RowsAffected == 0 {
		return false, nil
	}

	return true, fmt.Errorf("%s", "User already exists.")
}

func (auth *Store) DeactivateUser(email, userName string) (*model.User, error) {
	userModel := model.User{}

	data := auth.db.Where("email = ?", email).Where("name = ?", userName).Find(&userModel)

	if data.RowsAffected == 0 {
		return nil, nil
	}

	// Update with conditions.
	auth.db.Model(&userModel).Where("email = ?", email).Update("is_active", false)

	result := auth.db.Where("email = ?", email).Find(&userModel)

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return &userModel, nil
}

func (auth *Store) ActivateUser(email, userName string) (*model.User, error) {
	userModel := model.User{}

	data := auth.db.Where("email = ?", email).Where("name = ?", userName).Find(&userModel)

	if data.RowsAffected == 0 {
		return nil, nil
	}

	// Update with conditions.
	auth.db.Model(&userModel).Where("email = ?", email).Update("is_active", true)

	result := auth.db.Where("email = ?", email).Find(&userModel)

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return &userModel, nil
}

func (auth *Store) DeleteUser(email, userName string) (*model.User, error) {
	userModel := model.User{}

	data := auth.db.Where("email = ?", email).Where("name = ?", userName).Find(&userModel)

	if data.RowsAffected == 0 {
		return nil, nil
	}
	// Update with conditions
	auth.db.Model(&userModel).Where("name = ?", userName).Where("email = ?", email).Delete(&userModel)

	return &userModel, nil
}

// CheckUserExistence checks either the given user already exist in db or not.
func (auth *Store) GetUser(email, userName string) (*model.User, error) {
	userModel := model.User{}

	data := auth.db.Where("email = ?", email).Find(&userModel)

	if data.RowsAffected == 0 {
		return nil, nil
	}

	return &userModel, nil
}
