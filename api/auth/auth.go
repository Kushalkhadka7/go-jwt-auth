package auth

import (
	"jwt-auth/api/user"
	"jwt-auth/model"

	"gorm.io/gorm"
)

type Store struct {
	conn *gorm.DB
}

type Storer interface {
	CheckUserExistence(email string) (bool, error)
	GetUser(email, userName string) (*model.User, error)
}

func NewAuth(conn *gorm.DB) Storer {
	return &Store{
		conn,
	}
}

func (auth *Store) CheckUserExistence(email string) (bool, error) {
	user := user.NewStore(auth.conn)

	userExists, err := user.CheckUserExistence(email)
	if err != nil {
		return false, nil
	}

	if userExists {
		return true, nil
	}

	return false, nil
}

func (auth *Store) GetUser(email, userName string) (*model.User, error) {
	us := user.NewStore(auth.conn)

	response, err := us.GetUser(email, userName)
	if err != nil {
		return nil, nil
	}

	if response == nil {
		return nil, nil
	}

	return response, nil
}
