package app

import (
	"jwt-auth/auth"
	pb "jwt-auth/pb"
	"jwt-auth/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

type App struct {
	db *gorm.DB
}

type AppCreator interface {
	RegisterServers(grpcServer *grpc.Server)
}

func NewAppServer(db *gorm.DB) AppCreator {
	return &App{db: db}
}

func (server *App) RegisterServers(grpcServer *grpc.Server) {

	// Register user server and service to grpc server.
	userModel := user.NewUser(server.db)
	userService := user.NewService(userModel)
	userServer := user.New(userService)
	pb.RegisterUserServiceServer(grpcServer, userServer)

	// Register auth server and auth to grpc server.
	authModel := auth.NewAuth(server.db)
	authService := auth.NewService(authModel)
	authServer := auth.New(authService)
	pb.RegisterAuthServiceServer(grpcServer, authServer)

	reflection.Register(grpcServer)
}
