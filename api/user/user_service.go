package user

import (
	pb "jwt-auth/pb"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	store Storer
}

type Servicer interface {
	CreateUser(user *pb.User) (*pb.UserData, error)
	GetUser(email, userName string) (*pb.UserData, error)
	DeleteUser(email, userName string) (*pb.UserData, error)
	ActivateUser(email, userName string) (*pb.UserData, error)
	DeactivateUser(email, userName string) (*pb.UserData, error)
}

func NewService(store Storer) Servicer {
	return &Service{store}
}

// CreateUser checks either the given user exist or not, if not creates new user.
func (svc *Service) CreateUser(user *pb.User) (*pb.UserData, error) {
	userExists, err := svc.store.CheckUserExistence(user.Email)
	if err != nil {
		return nil, err
	}

	if userExists {
		return nil, err
	}

	// If user doesnot exist generate password hash and create user.
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
		return nil, err
	}

	return &pb.UserData{
		Id:       newUser.Id,
		Name:     newUser.Name,
		UserName: newUser.UserName,
		Email:    newUser.Email,
		IsActive: newUser.IsActive,
	}, nil
}

// Deactivates user deactivates the user account.
func (svc *Service) DeactivateUser(email, userName string) (*pb.UserData, error) {
	userExists, err := svc.store.CheckUserExistence(email)
	if err != nil {
		return nil, err
	}

	if userExists {
		return nil, err
	}

	res, err := svc.store.DeactivateUser(email, userName)
	if err != nil {
		return nil, err
	}

	return &pb.UserData{
		Id:       res.Id,
		Name:     res.Name,
		UserName: res.UserName,
		Email:    res.Email,
		IsActive: res.IsActive,
	}, nil
}

// Activate user activates the user account.
func (svc *Service) ActivateUser(email, userName string) (*pb.UserData, error) {
	userExists, err := svc.store.CheckUserExistence(email)
	if err != nil {
		return nil, err
	}

	if userExists {
		return nil, err
	}

	res, err := svc.store.ActivateUser(email, userName)
	if err != nil {
		return nil, err
	}

	return &pb.UserData{
		Name:  res.Name,
		Email: res.Email,
	}, nil
}

// DeleteUser deletes the user account.
func (svc *Service) DeleteUser(email, userName string) (*pb.UserData, error) {
	userExists, err := svc.store.CheckUserExistence(email)
	if err != nil {
		return nil, err
	}

	if userExists {
		return nil, err
	}

	res, err := svc.store.DeleteUser(email, userName)
	if err != nil {
		return nil, err
	}

	return &pb.UserData{
		Id:       res.Id,
		Name:     res.Name,
		UserName: res.UserName,
		Email:    res.Email,
		IsActive: res.IsActive,
	}, nil
}

// DeleteUser deletes the user account.
func (svc *Service) GetUser(email, userName string) (*pb.UserData, error) {
	res, err := svc.store.GetUser(email, userName)
	if err != nil {
		return nil, err
	}

	return &pb.UserData{
		Id:       res.Id,
		Name:     res.Name,
		UserName: res.UserName,
		Email:    res.Email,
		IsActive: res.IsActive,
	}, nil
}
