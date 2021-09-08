package server

import (
	"fmt"
	"jwt-auth/util"
	"net"

	"google.golang.org/grpc"
)

var log = util.NewLogger("server")

// Server is a data structure that populates the server with credentials required to run a server.
type Server struct {
	Port    int
	Address string
}

// Creator is interface to that exposes methods associated with server.
type Creator interface {
	Start() error
	GetHttpListener() (net.Listener, error)
	GetGrpcServerInstance() (*grpc.Server, error)
}

// New initializes new server instance.
func New(port int, address string) Creator {
	return &Server{Port: port, Address: address}
}

// GetHttpListener returns a http listener on give port and address.
func (server *Server) GetHttpListener() (net.Listener, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 6000))
	if err != nil {
		log.Error(err)

		return nil, err
	}

	log.Info(fmt.Sprintf("Server listening on port: %d", server.Port))

	return listener, nil
}

// GetGrpcServerInstance returns a grpc server instance.
func (server *Server) GetGrpcServerInstance() (*grpc.Server, error) {
	grpcServer := grpc.NewServer(
	// grpc.UnaryInterceptor(interceptor.NewUnary().Unary()),
	)

	return grpcServer, nil
}

// Start starts the grpc server using the http listener in provided port and address.
func (server *Server) Start() error {
	grpcServer, err := server.GetGrpcServerInstance()

	if err != nil {
		return err
	}

	httpServer, err := server.GetHttpListener()
	if err != nil {
		return err
	}

	grpcServer.Serve(httpServer)

	return nil
}
