package auth

import (
	"context"
	pb "jwt-auth/pb"
)

type AuthServer struct {
	Servicer
}

func New(service Servicer) *AuthServer {
	return &AuthServer{service}
}

// Login is used to login into the app and once the login gets successful it return user along with jwt tokens.
func (us *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	res, err := us.Servicer.Login(req.User)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		UserName:     req.User.Name,
		Name:         req.User.Name,
		Role:         "ADMIN",
		AccessToken:  res["accessToken"],
		RefreshToken: res["refreshToken"],
	}, nil
}

func (us *AuthServer) VerifyUser(ctx context.Context, req *pb.VerifyUserRequest) (*pb.VerifyUserRequest, error) {
	res, err := us.Servicer.VerifyLoggedInUser(req.AccessToken)
	if err != nil {
		return nil, err
	}

	return &pb.VerifyUserRequest{
		AccessToken: res["accessToken"],
	}, nil
}
