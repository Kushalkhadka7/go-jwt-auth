package user

import (
	"fmt"
	pb "jwt-auth/pb"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	store Storer
}

type Servicer interface {
	DeleteUser(email, userName string) (*pb.User, error)
	CreateUser(user *pb.User) (*pb.User, error)
	ActivateUser(isActive bool) (*pb.User, error)
	DeactivateUser(isActive bool) (*pb.User, error)
}

func NewService(store Storer) Servicer {
	return &Service{store}
}

// CreateUser checks either the given user exist or not, if not creates new user.
func (svc *Service) CreateUser(user *pb.User) (*pb.User, error) {
	fmt.Println("called")
	userExists, err := svc.store.CheckUserExistence(user)
	fmt.Println("called2c")
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

	newUser, err := svc.store.CreateUser(user)
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

func (svc *Service) DeactivateUser(isActive bool) (*pb.User, error) {
	res, err := svc.store.DeactivateUser(isActive)
	fmt.Println("called2cccc")
	if err != nil {
		return nil, err
	}

	return &pb.User{
		Name:  res.Name,
		Email: res.Email,
	}, nil
}

func (svc *Service) ActivateUser(isActive bool) (*pb.User, error) {
	res, err := svc.store.ActivateUser(isActive)
	if err != nil {
		return nil, err
	}

	return &pb.User{
		Name:  res.Name,
		Email: res.Email,
	}, nil
}

func (svc *Service) DeleteUser(email, userName string) (*pb.User, error) {
	res, err := svc.store.DeleteUser(email, userName)
	if err != nil {
		return nil, err
	}

	return &pb.User{
		Name:  res.Name,
		Email: res.Email,
	}, nil
}
