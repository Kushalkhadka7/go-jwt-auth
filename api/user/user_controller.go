package user

import (
	"context"
	pb "jwt-auth/pb"
	"jwt-auth/util"

	"google.golang.org/grpc/codes"
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
		return nil, util.NewResponse(codes.Internal, "Internal Server", err).Error()
	}

	return &pb.CreateUserResponse{
		Message: "Successfully created user.",
		Status:  200,
		Data: &pb.UserData{
			Name:     res.Name,
			UserName: res.Name,
			Role:     pb.UserData_ADMIN,
		},
	}, nil
}
