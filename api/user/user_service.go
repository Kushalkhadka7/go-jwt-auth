package user

import (
	"fmt"
	pb "jwt-auth/pb"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Storer
}

type Servicer interface {
	CreateUser(user *pb.User) (*pb.User, error)
}

func NewService(store Storer) Servicer {
	return &Service{store}
}

// CreateUser checks either the given user exist or not, if not creates new user.
func (svc *Service) CreateUser(user *pb.User) (*pb.User, error) {
	userExists, err := svc.CheckUserExistence(user)
	if err != nil {
		return nil, err
	}

	if userExists {
		return nil, fmt.Errorf("User already exist, try with another username and password.")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	newUser, err := svc.CreateUser(user)
	if err != nil {
		return nil, err
	}

	if newUser == nil {
		return nil, fmt.Errorf("Unable to create user: %s", err)
	}

	return &pb.User{
		Name:     newUser.Name,
		Password: string(hashedPassword),
	}, nil
}
