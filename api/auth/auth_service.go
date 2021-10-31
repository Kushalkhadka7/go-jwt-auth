package auth

import (
	"fmt"
	"jwt-auth/common"
	"time"

	pb "jwt-auth/pb"
)

type Service struct {
	jwtManager *common.JWTManager
	auth       Storer
}

type Servicer interface {
	Login(userName, email, password string) (*pb.LoginData, error)
	VerifyLoggedInUser(accessToken string) (map[string]interface{}, error)
}

func NewService(auth Storer) Servicer {
	jwtManager := common.NewJWTManager()

	return &Service{
		jwtManager,
		auth,
	}
}

func (s *Service) Login(userName, email, password string) (*pb.LoginData, error) {
	userExists, err := s.auth.CheckUserExistence(email)
	if err != nil {
		return nil, err
	}

	if userExists {
		return nil, err
	}

	response, err := s.auth.GetUser(email, userName)
	if err != nil {
		return nil, err
	}

	// if response.Password != password {
	// 	return nil, fmt.Errorf("%s", "Password doesnot match")
	// }

	usr := []struct {
		userName string
		email    string
		password string
		role     string
		isActive bool
	}{
		{
			userName: userName,
			email:    email,
			password: password,
			role:     "ADMIN",
			isActive: response.IsActive,
		},
	}

	tokens, err := s.jwtManager.GenerateTokens("kushal", 10000*time.Minute, usr)

	if err != nil {
		return nil, err
	}

	return &pb.LoginData{
		Id:           response.Id,
		Name:         response.Name,
		UserName:     response.UserName,
		Email:        response.Email,
		IsActive:     response.IsActive,
		AccessToken:  tokens["accessToken"],
		RefreshToken: tokens["refreshToken"],
	}, nil

}

func (s *Service) VerifyLoggedInUser(accessToken string) (map[string]interface{}, error) {
	_, err := s.jwtManager.VerifyToken(accessToken)
	if err != nil {
		return nil, err
	}

	res, err := s.jwtManager.ExtractTokenMetadata(accessToken)
	if err != nil {
		fmt.Printf("error occurred%s", err)
	}
	fmt.Printf("%s", res)
	return nil, nil
}
