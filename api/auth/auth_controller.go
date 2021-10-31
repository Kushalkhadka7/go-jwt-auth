package auth

import (
	"context"
	pb "jwt-auth/pb"
)

type Contorller struct {
	Servicer
}

func New(service Servicer) *Contorller {
	return &Contorller{service}
}

// Login is used to login into the app and once the login gets successful it return user along with jwt tokens.
func (us *Contorller) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	res, err := us.Servicer.Login(req.UserName, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Message: "Successfully created user.",
		Status:  200,
		Data: &pb.LoginData{
			Name:         res.Name,
			UserName:     res.Name,
			Role:         "ADMIN",
			Email:        res.Email,
			IsActive:     res.IsActive,
			AccessToken:  res.AccessToken,
			RefreshToken: res.RefreshToken,
		},
	}, nil
}

func (us *Contorller) VerifyUser(ctx context.Context, req *pb.VerifyUserRequest) (*pb.VerifyUserResponse, error) {
	_, err := us.Servicer.VerifyLoggedInUser(req.AccessToken)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
