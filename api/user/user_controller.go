package user

import (
	"context"
	pb "jwt-auth/pb"
	"jwt-auth/util"

	"google.golang.org/grpc/codes"
)

// UserController is a interface which exposes the methods that are available for user end points.
type UserController struct {
	Servicer
}

// New initializes new user contorller.
func New(service Servicer) *UserController {
	return &UserController{service}
}

// CreateUser method creates user.
func (us *UserController) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
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
