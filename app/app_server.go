package app

import (
	"jwt-auth/api/auth"
	"jwt-auth/api/user"
	pb "jwt-auth/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

// App is a data structure that populates the configurations data needed to run the app.
type App struct {
	db *gorm.DB
}

// AppCreator is a interface that exposes the methods associated with app server registration.
type AppCreator interface {
	RegisterServers(grpcServer *grpc.Server)
}

// NewAppServer initiates a new app instance with database connection.
func NewAppServer(db *gorm.DB) AppCreator {
	return &App{db: db}
}

// RegisterServers registers different grpc services with grpc server and app server..
func (server *App) RegisterServers(grpcServer *grpc.Server) {

	// Register user server and service to grpc server.
	userStore := user.NewStore(server.db)
	userService := user.NewService(userStore)
	userContorller := user.NewController(userService)
	pb.RegisterUserServiceServer(grpcServer, userContorller)

	// Register auth server and auth to grpc server.
	authModel := auth.NewAuth(server.db)
	authService := auth.NewService(authModel)
	authServer := auth.New(authService)
	pb.RegisterAuthServiceServer(grpcServer, authServer)

	reflection.Register(grpcServer)
}
