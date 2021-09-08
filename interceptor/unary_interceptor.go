package interceptor

import (
	"context"
	"fmt"
	pb "jwt-auth/pb"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var validaitonRules = []string{"Login", "CreateUser"}

// UnaryInterceptor initializes new auth interceptor.
type UnaryInterceptor struct{}

// NewAuthInterceptor creates new auth interceptor.
func NewUnary() *UnaryInterceptor {
	return &UnaryInterceptor{}
}

// Unary server interceptor.
func (interceptor *UnaryInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		ok := SkipValidation(info.FullMethod)
		if ok {
			return handler(ctx, req)
		}

		fmt.Printf("%s\n", req)
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, err
		}

		fmt.Printf("%s", info.FullMethod)
		values := md["authorization"]
		if len(values) == 0 {
			return nil, err
		}

		accessToken := values[0]
		authReq := &pb.VerifyUserRequest{
			AccessToken: accessToken,
		}

		authCleint, err := createAuthClient()
		if err != nil {
			return nil, err
		}

		res, err := authCleint.VerifyUser(ctx, authReq)
		if err != nil {
			return nil, err

		}

		var result *pb.VerifyUserResponse

		context := context.WithValue(ctx, result, res)
		fmt.Println("i am called1")
		return handler(context, req)
	}

}

// createAuthClient creates new auth client.
func createAuthClient() (pb.AuthServiceClient, error) {
	serverConn, err := grpc.Dial("6000",
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	authCleint := pb.NewAuthServiceClient(serverConn)

	return authCleint, nil
}

func SkipValidation(url string) bool {
	ok := false
	authenticatedList := strings.Split(url, "/")

	for _, v := range validaitonRules {
		for _, vl := range authenticatedList {
			fmt.Println(v, vl)
			if v == vl {
				ok = true

				break
			}
		}
	}

	return ok
}
