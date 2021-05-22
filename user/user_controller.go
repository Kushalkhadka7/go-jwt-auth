package user

import (
	"context"
	pb "jwt-auth/pb"
)

type UserServer struct {
	Servicer
}

func New(service Servicer) *UserServer {
	return &UserServer{service}
}

func (us *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	res, err := us.Servicer.CreateUser(req.User)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		Name:     res.Name,
		Message:  "Successfully created user.",
		UserName: res.Name,
	}, nil
}
