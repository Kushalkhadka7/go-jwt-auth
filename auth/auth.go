package auth

import (
	"jwt-auth/model"
	pb "jwt-auth/pb"
	"jwt-auth/user"

	"gorm.io/gorm"
)

type Auth struct {
	conn *gorm.DB
}

type AuthI interface {
	CreateUser(usr *pb.User) (*model.User, error)
	CheckUserExistence(usr *pb.User) (bool, error)
}

func NewAuth(conn *gorm.DB) AuthI {
	return &Auth{
		conn,
	}
}

func (auth *Auth) CreateUser(usr *pb.User) (*model.User, error) {
	user := user.NewUser(auth.conn)

	newUser, err := user.CreateUser(usr)
	if err != nil {
		return nil, err
	}
	if newUser == nil {
		return nil, nil
	}

	return newUser, nil
}

func (auth *Auth) CheckUserExistence(usr *pb.User) (bool, error) {
	user := user.NewUser(auth.conn)

	userExists, err := user.CheckUserExistence(usr)
	if err != nil {
		return false, nil
	}

	if userExists {
		return true, nil
	}

	return false, nil
}
