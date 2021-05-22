package user

import (
	"fmt"
	pb "jwt-auth/pb"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	user UserI
}

type Servicer interface {
	CreateUser(user *pb.User) (*pb.User, error)
}

func NewService(user UserI) Servicer {
	return &Service{user}
}

// CreateUser checks either the given user exist or not, if not creates new user.
func (s *Service) CreateUser(user *pb.User) (*pb.User, error) {
	userExists, err := s.user.CheckUserExistence(user)
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

	newUser, err := s.user.CreateUser(user)
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
