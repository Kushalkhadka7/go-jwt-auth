package server

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

// Server struct.
type Server struct {
	Port    int
	Address string
}

// Creator is interface to server struct.
type Creator interface {
	GetHttpListener() (net.Listener, error)
	GetGrpcServerInstance() (*grpc.Server, error)
}

// New creates server instance.
func New(port int, address string) Creator {
	return &Server{Port: port, Address: address}
}

// StartHttpServer starts http server listening in given port and address.
func (server *Server) GetHttpListener() (net.Listener, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 6000))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Printf("Server listening on port:%d", server.Port)

	return listener, nil
}

// StartHttpServer starts http server listening in given port and address.
func (server *Server) GetGrpcServerInstance() (*grpc.Server, error) {
	grpcServer := grpc.NewServer()

	return grpcServer, nil
}
