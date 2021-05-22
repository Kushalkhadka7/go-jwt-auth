package auth

import (
	"jwt-auth/common"
	pb "jwt-auth/pb"
	"time"
)

type Service struct {
	jwtManager *common.JWTManager
	auth       AuthI
}

type Servicer interface {
	Login(user *pb.User) (map[string]string, error)
	VerifyLoggedInUser(accessToken string) (map[string]string, error)
}

func NewService(auth AuthI) Servicer {
	jwtManager := common.NewJWTManager()

	return &Service{
		jwtManager,
		auth,
	}
}

func (s *Service) Login(usr *pb.User) (map[string]string, error) {
	userExisted, err := s.auth.CheckUserExistence(usr)
	if err != nil {
		return nil, err
	}

	if userExisted {
		tokens, err := s.jwtManager.GenerateTokens("kushal", 1000*time.Second, usr)

		if err != nil {
			return nil, err
		}

		return tokens, nil

	}

	newUser, err := s.auth.CreateUser(usr)
	if err != nil {
		return nil, err
	}

	if newUser == nil {
		return nil, nil
	}

	tokens, err := s.jwtManager.GenerateTokens(newUser.Name, 1000*time.Second, newUser)

	return tokens, nil
}

func (s *Service) VerifyLoggedInUser(accessToken string) (map[string]string, error) {
	res, err := s.jwtManager.VerifyToken(accessToken)
	if err != nil {
		return nil, err
	}

	return res, nil
}
