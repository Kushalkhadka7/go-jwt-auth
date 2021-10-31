package user

import (
	"context"
	pb "jwt-auth/pb"
	"jwt-auth/util"

	"google.golang.org/grpc/codes"
)

// UserController is a interface which exposes the methods that are available for user end points.
type Controller struct {
	svc Servicer
}

// NewController initializes new user contorller.
func NewController(service Servicer) *Controller {
	return &Controller{service}
}

// CreateUser method calls create user service to create new user.
func (uc *Controller) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	res, err := uc.svc.CreateUser(req.User)
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
			Email:    res.Email,
			IsActive: res.IsActive,
		},
	}, nil
}

// CreateUser method creates user.
func (uc *Controller) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserRequest, error) {
	return nil, nil
}

// DeactivateUser method calls deactivate user serivce to deactivate user account.
func (uc *Controller) DeactivateUser(ctx context.Context, req *pb.DeactivateUserRequest) (*pb.DeactivateUserResponse, error) {
	res, err := uc.svc.DeactivateUser(req.Email, req.UserName)
	if err != nil {
		return nil, err
	}

	return &pb.DeactivateUserResponse{
		Message: "Successfully deactivated user.",
		Status:  200,
		Data: &pb.UserData{
			Name:     res.Name,
			UserName: res.Name,
			Role:     pb.UserData_ADMIN,
			Email:    res.Email,
			IsActive: res.IsActive,
		},
	}, nil
}

// DeactivateUser method calls activate user serivce to activate user account.
func (uc *Controller) ActivateUser(ctx context.Context, req *pb.ActivateUserRequest) (*pb.ActivateUserResponse, error) {
	res, err := uc.svc.ActivateUser(req.Email, req.UserName)
	if err != nil {
		return nil, util.NewResponse(codes.Internal, "Internal Server", err).Error()
	}

	return &pb.ActivateUserResponse{
		Message: "Successfully activated user.",
		Status:  200,
		Data: &pb.UserData{
			Name:     res.Name,
			UserName: res.Name,
			IsActive: true,
			Role:     pb.UserData_ADMIN,
		},
	}, nil
}

// DeleteUser method calls delete user serivce to delete user account.
func (uc *Controller) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	res, err := uc.svc.DeleteUser(req.Email, req.UserName)
	if err != nil {
		return nil, util.NewResponse(codes.Internal, "Internal Server", err).Error()
	}

	return &pb.DeleteUserResponse{
		Message: "Successfully deleted user.",
		Status:  200,
		Data: &pb.UserData{
			Name:     res.Name,
			UserName: res.Name,
			IsActive: true,
			Role:     pb.UserData_ADMIN,
		},
	}, nil
}

// DeleteUser method calls delete user serivce to delete user account.
func (uc *Controller) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	res, err := uc.svc.GetUser(req.Email, req.UserName)
	if err != nil {
		return nil, util.NewResponse(codes.Internal, "Internal Server", err).Error()
	}

	return &pb.GetUserResponse{
		Message: "Successfully fetched user details.",
		Status:  200,
		Data: &pb.UserData{
			Name:     res.Name,
			UserName: res.Name,
			IsActive: true,
			Role:     pb.UserData_ADMIN,
		},
	}, nil
}
